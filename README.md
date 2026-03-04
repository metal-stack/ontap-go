# ontap-go
Go client to talk to Ontap

## Post-generation fixes

The ONTAP swagger spec has a few property names using hyphens (e.g. `volume-count`) even though the spec states that "REST API properties use underscores instead of hyphens between words" and the actual API returns underscores. The `generate-client` Makefile target includes a sed step to fix these after generation.
