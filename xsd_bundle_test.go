package fincen_test

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"io"
	"io/fs"
	"os"
	"path"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// officialSchemaFiles are the FinCEN BSA XML 2.0 files that must be in data/xsd.
var officialSchemaFiles = []string{
	"BSA_XML_2.0.xsd",
	"FinCENReferenceCodes.xsd",
	"EFL_8300XBatchSchema.xsd",
	"EFL_CTRXBatchSchema.xsd",
	"EFL_DOEPXBatchSchema.xsd",
	"EFL_FBARXBatchSchema.xsd",
	"EFL_SARXBatchSchema.xsd",
	"Schema/ISM/IC-ISM.xsd",
}

func xsdBundle() fs.FS {
	return os.DirFS("data/xsd")
}

func readSHA256SUMS(t *testing.T, fsys fs.FS) map[string]string {
	t.Helper()

	fd, err := fsys.Open("SHA256SUMS")
	require.NoError(t, err)
	defer fd.Close()

	sums := make(map[string]string)
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		require.Len(t, fields, 2, line)
		sums[fields[1]] = fields[0]
	}
	require.NoError(t, scanner.Err())
	return sums
}

func listSchemaFiles(t *testing.T, fsys fs.FS) []string {
	t.Helper()

	var out []string
	err := fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(p, ".xsd") {
			out = append(out, p)
		}
		return nil
	})
	require.NoError(t, err)
	sort.Strings(out)
	return out
}

func TestXSDBundle_OfficialFilesPresent(t *testing.T) {
	fsys := xsdBundle()
	for _, name := range officialSchemaFiles {
		_, err := fs.Stat(fsys, name)
		require.NoError(t, err, name)
	}
}

// The bundled files must stay byte-identical to the files that FinCEN publishes.
// Update data/xsd/SHA256SUMS when you replace a schema file.
func TestXSDBundle_Checksums(t *testing.T) {
	fsys := xsdBundle()
	sums := readSHA256SUMS(t, fsys)

	files := listSchemaFiles(t, fsys)
	require.Len(t, sums, len(files), "each .xsd file must have one entry in SHA256SUMS")

	for _, name := range files {
		want, ok := sums[name]
		require.True(t, ok, "%s is not in SHA256SUMS", name)

		bs, err := fs.ReadFile(fsys, name)
		require.NoError(t, err)

		got := sha256.Sum256(bs)
		require.Equal(t, want, hex.EncodeToString(got[:]), name)
	}
}

type schemaReference struct {
	from     string
	kind     string
	location string
}

func schemaReferences(t *testing.T, fsys fs.FS, name string) []schemaReference {
	t.Helper()

	bs, err := fs.ReadFile(fsys, name)
	require.NoError(t, err)

	var refs []schemaReference
	dec := xml.NewDecoder(bytes.NewReader(bs))
	for {
		tok, err := dec.Token()
		if errors.Is(err, io.EOF) {
			break
		}
		require.NoError(t, err, "%s is not well-formed XML", name)

		start, ok := tok.(xml.StartElement)
		if !ok || start.Name.Space != "http://www.w3.org/2001/XMLSchema" {
			continue
		}
		switch start.Name.Local {
		case "import", "include", "redefine":
			for _, attr := range start.Attr {
				if attr.Name.Local == "schemaLocation" && attr.Name.Space == "" {
					refs = append(refs, schemaReference{from: name, kind: start.Name.Local, location: attr.Value})
				}
			}
		}
	}
	return refs
}

type xmlCatalog struct {
	URIs []struct {
		Name string `xml:"name,attr"`
		URI  string `xml:"uri,attr"`
	} `xml:"uri"`
	Rewrites []struct {
		Start  string `xml:"uriStartString,attr"`
		Prefix string `xml:"rewritePrefix,attr"`
	} `xml:"rewriteURI"`
}

// resolve returns the local path for a remote location, as an XML catalog resolver does.
func (c xmlCatalog) resolve(location string) (string, bool) {
	for _, u := range c.URIs {
		if u.Name == location {
			return path.Clean(u.URI), true
		}
	}
	for _, r := range c.Rewrites {
		if strings.HasPrefix(location, r.Start) {
			return path.Clean(r.Prefix + strings.TrimPrefix(location, r.Start)), true
		}
	}
	return "", false
}

// Each import, include, and redefine must resolve to a bundled file, so validation can run offline.
func TestXSDBundle_ReferencesResolveOffline(t *testing.T) {
	fsys := xsdBundle()

	bs, err := fs.ReadFile(fsys, "catalog.xml")
	require.NoError(t, err)

	var catalog xmlCatalog
	require.NoError(t, xml.Unmarshal(bs, &catalog))
	require.NotEmpty(t, catalog.URIs)

	var checked int
	for _, name := range listSchemaFiles(t, fsys) {
		for _, ref := range schemaReferences(t, fsys, name) {
			var local string
			if strings.Contains(ref.location, "://") {
				var ok bool
				local, ok = catalog.resolve(ref.location)
				require.True(t, ok, "%s: %s %q is not in catalog.xml", ref.from, ref.kind, ref.location)
			} else {
				local = path.Join(path.Dir(ref.from), ref.location)
			}

			_, err := fs.Stat(fsys, local)
			require.NoError(t, err, "%s: %s %q resolves to missing file %q", ref.from, ref.kind, ref.location, local)
			checked++
		}
	}
	require.Positive(t, checked)
}
