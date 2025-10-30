package repository

import (
	"context"
	"errors"
	. "pasteBin/internal/database/models"
	. "pasteBin/internal/initializers"
	. "pasteBin/internal/models"
	"pasteBin/pkg/hash"
	"time"

	"gorm.io/gorm"
)
type PastesRepository struct{
	db *gorm.DB
}
func NewPastesRepository() *PastesRepository {
	return &PastesRepository{db: DB}
}
func (r *PastesRepository) CreatePaste(paste CreatePaste) (*Paste,error){

	
	input:=Paste{Content:paste.Content,Password:paste.Password,MaxViews:paste.MaxViews}
	ctx:=context.Background()
	result:=gorm.WithResult()
	err:=gorm.G[Paste](r.db,result).Create(ctx,&input)
	if err!=nil {
		return nil,err
	}

	return &input,nil;
	}
	func (r *PastesRepository) GetPastes()([]Paste,error){

var		pastes []Paste;
		result:=r.db.Find(&pastes).Where("deleted_at!=null")
		if result.Error!=nil{
			return nil,result.Error
		}
		return pastes,nil;
	}
	func (r *PastesRepository) GetPaste(id string,password string)(*Paste,error) {

		
		var paste *Paste
		err := r.db.Session(&gorm.Session{SkipHooks: true}).Where("id=?", id).First(&paste).Error
		if err!=nil{
			return  nil,err
		}
		if paste==nil{
			
			return nil , errors.New("Paste not found")
			
			
		}
			if paste.Password!=nil{
			if (password==""){
				return nil,errors.New("Paste password protected you need to enter a password to view it!")
			}
		  match,err:=hash.Compare(password,*paste.Password)
			if err!=nil{
				return nil, err;
			}
			if match==nil{
				return nil,errors.New("Incorrect password")
			}
			if *match==false{
				return nil,errors.New("Incorrect password")
			}
		}
		if (paste.ExpirationDate!=nil){
if (!paste.ExpirationDate.Time.After(time.Now())){
			return nil,errors.New("Paste Expired")
			}

		}
		if (paste.MaxViews!=nil){
		if *paste.MaxViews==0{
		return nil,errors.New("Paste destroyed")
		}
			
			ctx:=context.Background()
_,err:=gorm.G[Paste](r.db).Where("id=?",paste.ID).Update(ctx,"max_views",gorm.Expr("max_views -?",1))
			if err!=nil{
				return nil,err;
			}
		}
		return paste,nil;
	}
	
	func (r *PastesRepository) DeletePaste(id string)(error){

		if err:=r.db.Delete(&Paste{},id);err!=nil{
			if err.RowsAffected==0{
				return errors.New("Paste Not found")
			}
			return err.Error
		}
		
		return nil;
	
	}

