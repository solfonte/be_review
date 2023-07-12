package utils

type FlagsArray []string

func (i *FlagsArray) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func (i *FlagsArray) String() string {
	return ""
}