package util

func SafeChainInt64(f func() *int64) *int64 {
	defer func(){ _ = recover() }()
	if v := f(); v != nil { return v }
	return nil
}

func SafeChainInt(f func() *int) *int {
	defer func(){ _ = recover() }()
	if v := f(); v != nil { return v }
	return nil
}

func SafeChainString(f func() *string) *string {
	defer func(){ _ = recover() }()
	if v := f(); v != nil { return v }
	return nil
}