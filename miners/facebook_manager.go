package miners

import (
	"errors"
	"github.com/schmidtjohannes/social-media-data-mining/config"
	"github.com/schmidtjohannes/social-media-data-mining/model"
)

type FacebookManagerInterface interface {
	QueryGroups() (map[string]*model.FacebookGroupResponse, error)
}

type FacebookManager struct {
	Config *config.Configuration
}

func NewFacebookManager(cfg *config.Configuration) (*FacebookManager, error) {
	fbNetwork, ok := cfg.Networks["facebook"]
	if !ok {
		return nil, errors.New("No network named facebook provided")
	}
	if len(fbNetwork.Groups) == 0 {
		return nil, errors.New("No groups provided")
	}
	if fbNetwork.AccessToken == "" {
		return nil, errors.New("No access-token provided")
	}
	return &FacebookManager{Config: cfg}, nil
}

func (fm *FacebookManager) QueryGroups() (map[string]*model.FacebookGroupResponse, error) {
	fgr := make(map[string]*model.FacebookGroupResponse)
	for idx := range fm.Config.Networks["facebook"].Groups {
		fbGroupId := fm.Config.Networks["facebook"].Groups[idx]
		fm := newFacebookMiner(fbGroupId, fm.Config.Networks["facebook"].AccessToken)
		r, err := fm.QueryGroup()
		if err != nil {
			return nil, err
		}
		fgr[fbGroupId] = r
	}
	return fgr, nil
}
