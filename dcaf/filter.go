package main

type Filter func(string) string

func reconstructionFilter(seg Segmenter, des Desegmenter) (Filter) {
	return func(s string) string {return des(seg(s))}
}

func constructFilter(match_str string, join_str string, join_delim string, args []string) (Filter) {
	slice_rules := constructSliceRules(match_str, createCustomTypes(args))
	
	data_joiner := constructJoinRules(join_str)

	if join_str == "" {
		data_joiner = createMatchingJoinRules(slice_rules, join_delim)
	}

	segmenter := getSliceRuleSegmenter(slice_rules)
	desegmenter := getJoinRuleDesegmenter(data_joiner)
	return reconstructionFilter(segmenter, desegmenter)
}