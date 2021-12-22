package infra

import (
	"fmt"
	"math"
	"net/url"

	"gitlab.com/odeo/admin-iam/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConfig struct {
	Host  string
	Port  string
	Name  string
	User  string
	Pass  string
	Debug bool
}

type GormClient struct {
	DB   *gorm.DB
	Conf *DbConfig
}

var IamDbSrv = &GormClient{}
var err error

func (g *GormClient) Init(conf *DbConfig) error {
	g.Conf = conf
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		conf.Host,
		conf.User,
		conf.Pass,
		conf.Name,
		conf.Port,
	)
	g.DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		return err
	}

	return nil
}
func (*GormClient) Paginate(query *gorm.DB, dest interface{}, p *utils.PaginationMetadata) (*utils.Pagination, error) {
	query.Count(&p.TotalRecords)
	p.TotalPages = int(math.Ceil(float64(p.TotalRecords) / float64(p.Limit)))
	if err = query.Limit(p.Limit).Offset(p.GetOffset()).Find(dest).Error; err != nil {
		return nil, err
	}

	return &utils.Pagination{
		Records:  dest,
		Metadata: p,
	}, nil
}

func (*GormClient) Count(query *gorm.DB, p *utils.PaginationMetadata) *utils.PaginationMetadata {
	query.Count(&p.TotalRecords)
	p.TotalPages = int(math.Ceil(float64(p.TotalRecords) / float64(p.Limit)))
	return p
}

func (g *GormClient) HasRelation(db *gorm.DB, relationModel interface{}, id string, conds ...interface{}) *gorm.DB {
	var ids []interface{}
	relationDB := g.DB.Model(relationModel)
	if len(conds) > 0 {
		relationDB.Where(conds[0], conds[1:]...)
	}
	relationDB.Pluck("id", &ids)
	return db.Where(id+" IN (?)", ids)
}

func (g *GormClient) Filter(q url.Values, dest interface{}) {
	filter := make(map[string]string)
	for k, v := range q {
		filter[k] = v[0]
	}
	utils.Mapstructure(filter, dest)
}
