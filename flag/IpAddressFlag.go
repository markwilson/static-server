package flag

type IpAddressFlag struct {
	set bool
	value string
}

func (f *IpAddressFlag) Set(ip string) error {
	f.value = ip
	f.set = true

	return nil
}

func (f *IpAddressFlag) String() string {
	return f.value
}
