package admin

import (
	"net/http"

	"comma/pkg/library/core/e"
	"comma/pkg/service/admin"

	http2 "comma/pkg/library/net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// SELECT * FROM dict_word c RIGHT JOIN (SELECT
// 	a.id,
// 	substring_index( substring_index( a.word_ids, ',', b.id ), ',',- 1 ) AS word_id
// FROM
// 	( SELECT id, word_ids FROM dict_synonyms LIMIT 10 ) a
// 	JOIN dict_incr_id b ON b.id <= ( length( a.word_ids ) - length( REPLACE ( a.word_ids, ',', '' ))+ 1 )) d ON c.id = d.word_id LIMIT 100

func SynonymList(c *gin.Context) {
	ctx := http2.Context{C: c}
	pager := ctx.GetPager()
	list, _ := admin.NewSynonym().SynonymGetList(pager)
	ctx.Success(gin.H{
		"list":  list,
		"pager": pager.ToMap(),
	}, nil)
}

func SynonymCreate(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.SynonymCreateForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	outData, err := admin.NewSynonym().SynonymCreate(form)
	if err != nil {
		ctx.Error(http.StatusOK, e.Error, err.Error(), nil)
		return
	}
	ctx.Success(gin.H{
		"info": outData,
	}, nil)
}

func SynonymUpdate(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.SynonymUpdateForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	outData, err := admin.NewSynonym().SynonymUpdateById(updateId, form)
	if err != nil {
		ctx.Error(http.StatusOK, e.Error, err.Error(), err)
		return
	}
	ctx.Success(gin.H{
		"info": outData,
	}, nil)
}

func SynonymDelete(c *gin.Context) {
	ctx := http2.Context{C: c}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	ids := make([]int, 0)
	ids = append(ids, updateId)
	form := admin.SynonymUMultipleDeleteForm{
		Ids: ids,
	}
	success := admin.NewSynonym().SynonymDeleteByIds(form)
	if !success {
		ctx.Error(http.StatusOK, e.Error, nil, nil)
		return
	}
	ctx.Success(nil, nil)
}

func SynonymMultipleDelete(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.SynonymUMultipleDeleteForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	success := admin.NewSynonym().SynonymDeleteByIds(form)
	if !success {
		ctx.Error(http.StatusOK, e.Error, nil, nil)
		return
	}
	ctx.Success(nil, nil)
}
