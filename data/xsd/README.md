# FinCEN BSA XML 2.0 schemas

This directory holds the FinCEN BSA XML Schema 2.0 files. The files are byte-for-byte copies of the files that FinCEN publishes. Do not edit them.

- Source: https://www.fincen.gov/system/files/schema/base/
- Retrieved: 2026-11-24
- Checksums: `SHA256SUMS` (the `go test` suite compares each file to this list)

| File | Version comment in the file |
|------|-----------------------------|
| `EFL_8300XBatchSchema.xsd` | v1.0 (03/15/2019) |
| `EFL_CTRXBatchSchema.xsd` | v1.2 (12/17/2021) |
| `EFL_DOEPXBatchSchema.xsd` | v1.0 (03/14/2019) |
| `EFL_FBARXBatchSchema.xsd` | v1.2 (7/31/2018) |
| `EFL_SARXBatchSchema.xsd` | 1.2 (12/17/2021) |
| `BSA_XML_2.0.xsd` | Last change 12-17-2025 (Real Estate Report elements) |
| `FinCENReferenceCodes.xsd` | version 1.5, last change 02/05/2026 |
| `Schema/ISM/IC-ISM.xsd` and `Schema/ISM*/CVEGenerated/*.xsd` | IC-ISM version 13 |

The five `EFL_*XBatchSchema.xsd` files are the batch schemas, one for each form type. Each batch schema redefines `BSA_XML_2.0.xsd`. `BSA_XML_2.0.xsd` imports `FinCENReferenceCodes.xsd` and `Schema/ISM/IC-ISM.xsd`. The `Schema/` directory keeps the same layout as the FinCEN site, so the relative imports in `IC-ISM.xsd` resolve.

## Offline validation

Some official files refer to remote locations. Some of these locations do not resolve (for example `https://www.fincen.gov/sites/default/files/schema/base/BSA_XML_2.0.xsd` and `https://www.fincen.gov/codes`). `catalog.xml` is an OASIS XML catalog that maps these locations to the local files. With `xmllint` (libxml2):

```
XML_CATALOG_FILES=data/xsd/catalog.xml \
  xmllint --nonet --noout --schema data/xsd/EFL_CTRXBatchSchema.xsd batch.xml
```

The batch schemas declare `vc:minVersion="1.1"`. Processors that support only XSD 1.0, such as libxml2, ignore this attribute and can compile the schemas.

The library does not use these files. `Validate()` and the HTTP `/validator` endpoint check the library's field and attribute rules only. A successful result does not mean that FinCEN would accept the file.

## Update the schemas

1. Download each file from the source URL. Keep the same file names and directory layout.
2. If a file has a new import or include, add the file. If the new location is remote, add an entry to `catalog.xml`.
3. Update `SHA256SUMS`:

   ```
   cd data/xsd && find . -name '*.xsd' | sed 's|^\./||' | LC_ALL=C sort | xargs sha256sum > SHA256SUMS
   ```

4. Update the table in this file, then run `go test ./...`.
