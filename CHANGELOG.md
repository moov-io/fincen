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
