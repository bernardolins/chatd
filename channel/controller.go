package channel

import ()

type Controller struct {
	channelList map[string]*Channel
}

func NewController() *Controller {
	c := new(Controller)
	c.channelList = make(map[string]*Channel)

	return c
}

func (cc *Controller) NewChannel(name string) {
	if _, contains := cc.channelList[name]; contains {
		//error, room already exists
	} else {
	}
	c := New(name)
	cc.channelList[name] = c
}

func (cc *Controller) SelectChannel(name string) *Channel {
	return cc.channelList[name]
}
