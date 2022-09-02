# fincen
Fincen (Financial crimes enforcment network) BSA data transmission methods for the the BSA E-Filing System

A go library for reading and writing Fincen BSA forms. It is capable of generating, validating, and batching submissions.

The project is Fincen which is a financial reporting for the United States. There are many forms that are required in this project but we will start with SAR(Suspicious activity report). I would like a go library that creates the structs to build and validate the SAR form. It also needs to be able to export the file to XML.

In the future we will wrap the service with json/http services so that it can work as a web service and not just a go library.
New project
https://github.com/moov-io/fincen/blob/master/README.md

We will end up with a sub project for each of the form parsers. This is all of the forms.
https://bsaefiling.fincen.treas.gov/FilingInformation.html

The final phase of the project will be a service that collects fillings and then on a scheduled time will batch the fillings and submit them electronically.
https://bsaefiling.fincen.treas.gov/docs/SDTMRequirements.pdf


[About the BSA E-Filing System](https://bsaefiling.fincen.treas.gov/AboutBsa.html)

## Supported Forms

- [ ] FinCEN Currency Transaction Report (FinCEN Report 112)
- [ ] FinCEN Designation of Exempt Person (FinCEN Report 110)
- [ ] FinCEN Suspicious Activity Report (FinCEN Report 111)
- [ ] FinCEN Registration of Money Services Business (FinCEN Report 107)
- [ ] Report of Foreign Bank and Financial Accounts (FinCEN Report 114)
- [ ] Report of Cash Payments Over $10,000 Received in a Trade or Business (FinCEN Form 8300)

## Filing information user guides

The [Filing Information](https://bsaefiling.fincen.treas.gov/FilingInformation.html) website contains PDF's of each of the XML formates. The PDF's guides contain links to XSD's of the formts.

## Become a BSA E-Filer

[Secure Direct Transfer Mode](https://bsaefiling.fincen.treas.gov/SDTMInfo.html)

[Supervisory User Registration](https://bsaefiling1.fincen.treas.gov/AddUser)
