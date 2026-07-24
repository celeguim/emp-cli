package kubectl

type Client struct {
	Context    string
	Namespace  string
	Kubeconfig string
}

func New() *Client {
	return &Client{}
}
