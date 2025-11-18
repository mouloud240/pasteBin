package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"
	. "pasteBin/internal/database/models"
	. "pasteBin/internal/models"
	"pasteBin/pkg/exception"
	"pasteBin/pkg/hash"
	"time"

	"gorm.io/gorm"
)
type PastesRepository struct{
	db *gorm.DB
}
func NewPastesRepository(db *gorm.DB) *PastesRepository {
	return &PastesRepository{db: db}
}
func (r *PastesRepository) CreatePaste(ctx context.Context, paste CreatePaste,userId *uint) (*Paste,error){

	
	input:=Paste{Content:paste.Content,Password:paste.Password,MaxViews:paste.MaxViews,UserID:userId,ExpirationDate: &sql.NullTime{Time: *paste.Expires_at,Valid: paste.Expires_at != nil}}
	result:=gorm.WithResult()
	err:=gorm.G[Paste](r.db,result).Create(ctx,&input)
	if err!=nil {
		return nil,err
	}

	return &input,nil;
	}
	func (r *PastesRepository) GetPastes(ctx context.Context, page *int, limit *int)([]Paste,error){

  var pastes []Paste;
	
	offset:=(*page-1)*(*limit)
		result:=r.db.WithContext(ctx).Preload("Author").Where("max_views!=0").Limit(*limit).Offset(offset).Find(&pastes)
		if result.Error!=nil{
			return nil,result.Error
		}
		return pastes,nil;
	}

	func (r *PastesRepository) GetPaste(ctx context.Context, id string,password string)(*Paste,error) {

		
		var paste *Paste
		err := r.db.WithContext(ctx).Session(&gorm.Session{SkipHooks: true}).Where("id=?", id).Preload("Author",nil).First(&paste).Error
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
			
_,err:=gorm.G[Paste](r.db).Where("id=?",paste.ID).Update(ctx,"max_views",gorm.Expr("max_views -?",1))
			if err!=nil{
				return nil,err;
			}
		}
		return paste,nil;
	}
	
	func (r *PastesRepository) DeletePaste(ctx context.Context, id string,userId uint )(error){

		if err:=r.db.WithContext(ctx).Where("UserID=?",userId).Delete(&Paste{},id);err!=nil{
			if err.RowsAffected==0{
				return errors.New("Paste Not found")
			}
			return err.Error
		}
		
		return nil;
	
	}
	func (r *PastesRepository) DeleteExpiredPastes()(error){
		ctx:=context.Background();
	res,err:=gorm.G[Paste](r.db).Where("expiration_date IS NOT NULL AND expiration_date <= ?",sql.NullTime{Time: time.Now(),Valid: true}).Delete(ctx);
	if err!=nil{
		return exception.NewInternalServerError("Failed to delete expired pastes",err)
	}
	log.Printf("Deleted %d expired pastes",res)
	return nil;
	}

