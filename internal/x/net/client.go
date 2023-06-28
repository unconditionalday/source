package netx

type Client interface {
	Download(src string) ([]byte, error)
}
