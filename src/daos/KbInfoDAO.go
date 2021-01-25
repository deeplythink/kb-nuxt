package daos

import (
	"fmt"
	"gorm.io/gorm"
	"knowledgeBaseNuxt/src/models/DocModel"
)

type KbInfoDAO struct {
	DB *gorm.DB `inject:"-"`
}

func NewKbInfoDao() *KbInfoDAO {
	return &KbInfoDAO{}
}

func (this *KbInfoDAO) GetKbDetail(username string, kbName string) []*DocModel.DocGrpImpl {
	var dgm []*DocModel.DocGrpImpl

	kbID := 120

	this.getKbDetail("test", kbID, 0, &dgm)
	return dgm
}

func (this *KbInfoDAO) getKbDetail(kbName string, kbID, groupID int, result *[]*DocModel.DocGrpImpl) []*DocModel.DocGrpImpl {
	this.DB.Table("doc_grps").Raw(`select group_id,group_name,shorturl from doc_grps 
where kb_id = ? and pid = ? 
order by group_order `, kbID, groupID).Find(&result)
	for _, v := range *result {
		var docs []*DocModel.DocImpl
		fmt.Println(v)
		this.DB.Table("docs").Raw(`select doc_id,doc_title,shorturl from docs 
where kb_id = ? and group_id = ? 
order by  doc_id`, kbID, v.GroupID).Find(&docs)
		fmt.Println(docs)
		for _, doc := range docs {
			doc.DocHref = "/" + kbName + "/" + v.GroupShortUrl + "/" + doc.DocShortUrl

				fmt.Println(doc)
				this.getKbDetail(kbName, kbID, v.GroupID, &doc.Children)

		}

		v.Children = docs




	}
	return *result
}
