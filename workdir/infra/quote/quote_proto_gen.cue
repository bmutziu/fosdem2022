package quote

#Language:
	#EN |
	#NL

#EN: 0
#NL: 1

#Language_value: {
	EN: 0
	NL: 1
}

#QuoteRequest: {
	lang?: #Language @protobuf(1,Language)
	num?:  int32     @protobuf(2,int32)
	num?:  >0 & <3
}

#QuoteResponse: {
	quotes?: [...string] @protobuf(1,string)
}
