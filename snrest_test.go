package snrest

import (
	"testing"

	"github.com/mak42/snrest"
)

type TaskSingle struct {
	Result struct {
		SysID  string `json:"sys_id"`
		Number string `json:"number"`
	} `json:"result"`
}
type TaskResult struct {
	Result []struct {
		SysID  string `json:"sys_id"`
		Number string `json:"number"`
	} `json:"result"`
}
type Task struct {
	SysID            string `json:"sys_id"`
	ShortDescription string `json:"short_description"`
}

func TestRestClient(t *testing.T) {
	c := snrest.New("softpointdev2", "xxxx", "xxxxxx")
	/*t.Run("GetKeys", func(t *testing.T) {
		p := snrest.JsonV2Params{Query: "active=true"}
		k, err := c.GetKeys("incident", p)

		if len(k) <= 0 || err != nil {
			t.Fail()
		}
	})*/
	/*t.Run("Get", func(t *testing.T) {
		p := snrest.GetParams{}
		r := TaskSingle{}
		_, err := c.Get("incident", "d71f7935c0a8016700802b64c67c11c6", p, &r)

		if r.Result.Number == "" || err != nil {
			t.Fail()
		}
	})
	t.Run("GetMultiple", func(t *testing.T) {
		p := snrest.GetParams{}
		r := TaskResult{}
		_, err := c.GetMultiple("incident", p, &r)

		if len(r.Result) <= 0 || err != nil {
			t.Fail()
		}
	})*/
	t.Run("Insert", func(t *testing.T) {
		p := snrest.PostParams{}
		r := TaskSingle{}
		data := Task{ShortDescription: "Test123Short"}
		res, err := c.Insert("incident", p, data, &r)
		t.Log(err)
		t.Log(res)

		if r.Result.Number == "" || err != nil {
			t.Fail()
		}
	})
}

/*func TestClient(t *testing.T) {
	res := TaskResult{}
	params := snrest.GetParams{}
	c := snrest.New("softpointdev2", "snow-team@softpoint.at", "!2softpoint")
	c.GetMultiple("incident", params, &res)
	t.Log(res)
	res2 := TaskResult{}
	c.GetMultiple("change_request", params, &res2)
	t.Log(res2)
}*/
