package structs

import "strings"

// tagOptions contains a slice of tag options
type tagOptions []string

// Has returns true if the given option is available in tagOptions
func (t tagOptions) Has(opt string) bool {
	for _, tagOpt := range t {
		if tagOpt == opt {
			return true
		}
	}

	return false
}

// ParseTag splits a struct field's tag into its name and a list of options
// which comes after a name. A tag is in the form of: "name,option1,option2".
// The name can be neglectected.
func ParseTag(tag string) (string, tagOptions) {
	// tag is one of followings:
	// ""
	// "name"
	// "name,opt"
	// "name,opt,opt2"
	// ",opt"

	res := strings.Split(tag, ",")
	return res[0], res[1:]
}

type TagOptions map[string]string

/*
	标签解析
	支持
	`name:"opt1,opt2:123,opt3"`
*/
func ParseTag2(tag string) TagOptions {
	tagStrArray := strings.Split(tag, ",")
	tags := map[string]string{}
	for _, tagStr := range tagStrArray {
		opt := strings.Split(tagStr, ":")
		if len(opt) == 1 {
			tags[strings.TrimSpace(opt[0])] = "true"
		} else {
			tags[strings.TrimSpace(opt[0])] = strings.TrimSpace(opt[1])
		}
	}
	return TagOptions(tags)
}
