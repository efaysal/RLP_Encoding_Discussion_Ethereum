# RLP_Encoding_Discussion_Ethereum



RLP Encoding, specially when no-ASCII are used, let's say an Ethereum application used the following string, "cafés", the RLP of this simple string will depend how it was edited first time: 

cafés has this number of bytes: 7

Its hex is
63 61 66 65 CC 81 73

cafés has this number of bytes: 6

Its hex is
63 61 66 C3 A9 73 

The coding is in https://github.com/efaysal/RLP_Encoding_Discussion_Ethereum. 

I dont know if this can be exploited in any Ethereum application, 
however using only ASCII characters seems to be the most reasonable way.
