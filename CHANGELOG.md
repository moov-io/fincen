## v0.5.0 (Unreleased)

IMPROVEMENTS

- data: replace the bundled XSD files with the official FinCEN BSA XML 2.0 files. Add `BSA_XML_2.0.xsd`, `FinCENReferenceCodes.xsd`, the IC-ISM schemas, an XML catalog for offline validation, and checksums (#PR)

## v0.4.0 (Released 2026-11-23)

This is the first release from master since v0.3.5. Do not use or pin v0.3.6. The v0.3.6 tag is not on master, and it panics when an `Activity` has no `TotalAmount`. Pin v0.4.0 or later.

IMPROVEMENTS

- fix: `TotalAmount` does not panic when an `Activity` element is empty (CTR, 8300, and SAR) (#111)
- fix: `Marshal` does not panic on types that implement `xml.Marshaler`; add fuzz tests (#110)
- fix: batch attaches activities together correctly (#73)
- fix(batch): keep computed totals when attributes are validated, so a wrong `TotalAmount` or `PartyCount` is rejected (#130)
- fix(batch): write `float64` amounts without scientific notation (#90)
- fix(batch): write the required SARX `ActivityAttachmentCount` and `AttachmentCount` attributes, also when the value is zero (#153)
- fix(batch): emit official `xsi:schemaLocation` values for all five form types (#158)
- fix: return an error instead of a panic from the marshaler
- feat(currency_transaction): accept the cash-out codes for activity detail type codes (#87)
- feat(suspicious_activity): accept EIN, TIN, and FIN `PartyIdentificationTypeCode` values for parties 30 and 34 (#108)
- fix(cmd/fincen): write progress messages to stderr, so stdout has only the report output (#154)
- fix(service): set the correct `Content-Type` on `/reformat` responses (#154)
- docs: fix the HTTP examples, links, and Go library snippets; the example programs report errors (#155)

BUILD

- build: Go 1.26 or later is required
- build: release assets include the CLI binaries (`fincen-cli-linux-amd64`, `fincen-cli-darwin-amd64`, `fincen-cli.exe`) next to the server binaries (`fincen-<os>-amd64`, `fincen.exe`) (#154)
- build: add `docker-compose.yml` for `make test-integration`; fix `go vet` and service tests that did not stop (#154)
- build: push Docker images only from official moov-io releases (#121)
- build: the web UI moved from the Docker image to the docs site
- ci: add scheduled fuzz tests, CodeQL, and zizmor checks
- Update dependencies, including github.com/moov-io/base to v0.68.2

## v0.3.6 (Released 2023-07-11)

IMPROVEMENTS

- fix: batch isn't attaching activites together

BUILD

- fix(deps): update module github.com/moov-io/base to v0.45.0
- fix(deps): update module github.com/stretchr/testify to v1.8.4

## v0.3.5 (Released 2023-02-27)

IMPROVEMENTS

- fix: failing fincen integration with AssetsTable, should be Assets

BUILD

- build: fixup ports in docs and Dockerfiles
- chore(deps): update golang docker tag to v1.20
- fix(deps): update module github.com/moov-io/base to v0.39.0
- fix(deps): update module github.com/stretchr/testify to v1.8.2

## v0.3.4 (Released 2022-11-28)

IMPROVEMENTS

- fix: fincen needs the order of party account association to have party account association type code ahead of the party

## v0.3.3 (Released 2022-11-16)

IMPROVEMENTS

- fix: set/require correct PartyIdentificationTypeCode for SAR

BUILD

- fix(deps): update module github.com/moov-io/base to v0.36.2

## v0.3.2 (Released 2022-11-10)

IMPROVEMENTS

- fix: splitting indicator type

## v0.3.1 (Released 2022-11-09)

IMPROVEMENTS

- fix: fixing xml namespace prefix of EFilingBatchXML

## v0.3.0 (Released 2022-11-09)

IMPROVEMENTS

- feat: fixing fc to fc2 for formTypeCode and adding empty indicator for currency  transactions

## v0.2.9 (Released 2022-11-09)

IMPROVEMENTS

- fix: modified namespace prefix of form type code

BUILD

- fix(deps): removed module github.com/antihax/optional
- fix(deps): removed module golang.org/x/oauth2
- fix(deps): removed module google.golang.org/appengine
- fix(deps): removed module golang.org/x/net

## v0.2.6 (Released 2022-10-28)

IMPROVEMENTS

- fix: use specified xsi:schemaLocation attributes

## v0.2.5 (Released 2022-10-26)

IMPROVEMENTS

- fix: modified unique suquence number checker and samples
- fix: updated logic for generating sequence numbers

BUILD

- fix(deps): update module github.com/moov-io/base to v0.36.1
- fix(deps): update module github.com/stretchr/testify to v1.8.1
- fix(deps): update module golang.org/x/oauth2 to v0.1.0

## v0.2.4 (Released 2022-10-17)

IMPROVEMENTS

- docs: adding simple go usage page
- fix: update Dockerfile with config

BUILD

- fix(deps): update module github.com/moov-io/base to v0.36.0
- fix(deps): update golang.org/x/oauth2 digest to 6fdb5e3

## v0.2.0 (Released 2022-09-26)

IMPROVEMENTS

- Add `webui` interface for parsing and converting files

## v0.1.0 (Released 2022-09-13)

Initial release! Let us know how moov-io/fincen is working for you!
