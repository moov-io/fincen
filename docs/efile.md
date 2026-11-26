---
layout: page
title: BSA E-Filing Information | Moov FinCEN
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Overview

The Bank Secrecy Act (BSA) Electronic Filing (E-Filing) System supports secure electronic transmission of BSA data to the Financial Crimes Enforcement Network (FinCEN). ([Source](https://bsaefiling.fincen.gov/about))

This system supports [individual or batch submissions](https://bsaefiling.fincen.gov/transmission-methods) of the supported report types. Refer to the [BSA costs page](https://bsaefiling.fincen.gov/costs) for more information.

## E-File

The [XML E-File Specifications](https://bsaefiling.fincen.gov/filing-information) can be used to upload the selected report. The Moov FinCEN library can generate the appropriate XML for submission.

Refer to the [FinCEN.gov](https://www.fincen.gov/resources/filing-information) for more information.

## FinCEN documents

FinCEN publishes an XML Schema 2.0 user guide (August 2021) for each report type:

- [FinCEN CTR XML Schema User Guide](https://bsaefiling.fincen.gov/docs/XMLUserGuide_FinCENCTR.pdf) (FinCEN Report 112)
- [FinCEN DOEP XML Schema User Guide](https://bsaefiling.fincen.gov/docs/XMLUserGuide_FinCENDOEP.pdf) (FinCEN Report 110)
- [FinCEN SAR XML Schema User Guide](https://bsaefiling.fincen.gov/docs/XMLUserGuide_FinCENSAR.pdf) (FinCEN Report 111)
- [FinCEN FBAR XML Schema User Guide](https://bsaefiling.fincen.gov/docs/XMLUserGuide_FinCENFBAR.pdf) (FinCEN Report 114)
- [FinCEN Form 8300 XML Schema User Guide](https://bsaefiling.fincen.gov/docs/XMLUserGuide_FinCEN8300.pdf) (FinCEN Form 8300)

Use these documents to validate and test batch files before you submit them:

- [XML Schema Guidance: Schema Validation](https://bsaefiling.fincen.gov/docs/XMLSchemaGuidance_SchemaValidation.pdf)
- [XML Batch Testing Procedures](https://bsaefiling.fincen.gov/docs/XML_Batch_Testing_Procedures.pdf)
- [Secure Direct Transfer Mode (SDTM) Requirements](https://bsaefiling.fincen.gov/docs/SDTMRequirements_2019.pdf)

The official batch schemas are at `https://www.fincen.gov/system/files/schema/base/`: [CTR](https://www.fincen.gov/system/files/schema/base/EFL_CTRXBatchSchema.xsd), [DOEP](https://www.fincen.gov/system/files/schema/base/EFL_DOEPXBatchSchema.xsd), [SAR](https://www.fincen.gov/system/files/schema/base/EFL_SARXBatchSchema.xsd), [FBAR](https://www.fincen.gov/system/files/schema/base/EFL_FBARXBatchSchema.xsd), and [8300](https://www.fincen.gov/system/files/schema/base/EFL_8300XBatchSchema.xsd).
