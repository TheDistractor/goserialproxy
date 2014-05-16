goserialproxy
=============

A very quick and dirty serial proxy to combine needed features from tarm/goserial and chimera/rs232


### Overview
Although I don't use windows myself, I am working on a project that should support linux, darwin AND windows. However the serial library we use does not yet support Windows (chimera/rs232). (tarm/goserial) DOES support Windows, but not with the feature set needed. This project implements the needed featureset of the project through an interface in order to support both linux/darwin/windows.

The goal would actually be to either port the refactored windows from goserial to rs232, or to merge the capabilitys of both librarys into a single solution. I dont have time to do this 'just yet' so this proxy serves that temporary purpose.

The proxy's interface will be expanded more when I get time (I dont need some of the interface stubs just yet).

