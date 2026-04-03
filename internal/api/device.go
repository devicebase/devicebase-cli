package api

import "fmt"

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Bounds struct {
	X1 int `json:"x1"`
	Y1 int `json:"y1"`
	X2 int `json:"x2"`
	Y2 int `json:"y2"`
}

type LaunchAppRequest struct {
	AppName string `json:"app_name"`
}

type InputTextRequest struct {
	Text string `json:"text"`
}

func (c *Client) DeviceInfo(serial string) ([]byte, error) {
	return c.Post(fmt.Sprintf("/v1/deviceinfo/%s", serial), nil)
}

func (c *Client) Tap(serial string, p Point) ([]byte, error) {
	return c.Post(fmt.Sprintf("/v1/tap/%s", serial), p)
}

func (c *Client) DoubleTap(serial string, p Point) ([]byte, error) {
	return c.Post(fmt.Sprintf("/v1/double_tap/%s", serial), p)
}

func (c *Client) LongPress(serial string, p Point) ([]byte, error) {
	return c.Post(fmt.Sprintf("/v1/long_press/%s", serial), p)
}

func (c *Client) Swipe(serial string, b Bounds) ([]byte, error) {
	return c.Post(fmt.Sprintf("/v1/swipe/%s", serial), b)
}

func (c *Client) Back(serial string) ([]byte, error) {
	return c.Post(fmt.Sprintf("/v1/back/%s", serial), nil)
}

func (c *Client) Home(serial string) ([]byte, error) {
	return c.Post(fmt.Sprintf("/v1/home/%s", serial), nil)
}

func (c *Client) LaunchApp(serial string, appName string) ([]byte, error) {
	return c.Post(fmt.Sprintf("/v1/launch_app/%s", serial), LaunchAppRequest{AppName: appName})
}

func (c *Client) InputText(serial string, text string) ([]byte, error) {
	return c.Post(fmt.Sprintf("/v1/input/%s", serial), InputTextRequest{Text: text})
}

func (c *Client) ClearText(serial string) ([]byte, error) {
	return c.Post(fmt.Sprintf("/v1/clear_text/%s", serial), nil)
}

func (c *Client) CurrentApp(serial string) ([]byte, error) {
	return c.Post(fmt.Sprintf("/v1/current_app/%s", serial), nil)
}

func (c *Client) DumpHierarchy(serial string) ([]byte, error) {
	return c.Post(fmt.Sprintf("/v1/dump_hierarchy/%s", serial), nil)
}

func (c *Client) Screenshot(serial string) ([]byte, error) {
	return c.Post(fmt.Sprintf("/v1/screen/%s", serial), nil)
}
