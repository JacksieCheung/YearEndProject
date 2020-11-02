//存储学生信息数据库的模型
package model

import (
    _"github.com/go-sql-driver/mysql"
)

type Data struct {
    Stunum string `gorm:"type:char(10);not null;index:stunum"`
    Time string `gorm:"type:varchar(20);not null;"`
    Cost string `gorm:"type:varhchar(8);not null;"`
    Restaurant string `gorm:"type:varchar(30);not null;"`
    Place string `gorm:"type:varchar(30);not null;"`
}

type Stunum2018 struct {
    Stunum string `gorm:"type:char(10);not null primary key;"`
}

type Stuinfos struct {
    Stunum string `gorm:"type:char(10);not null;index:stunum"`
    Time string `gorm:"type:varchar(20);not null;"`
    Cost string `gorm:"type:varhchar(8);not null;"`
    Restaurant string `gorm:"type:varchar(30);not null;"`
    Place string `gorm:"type:varchar(30);not null;"`
}

type Data2018 struct {
    Stunum string `gorm:"type:char(10);not null;index:stunum"`
    Time string `gorm:"type:varchar(20);not null;"`
    Cost string `gorm:"type:varhchar(8);not null;"`
    Restaurant string `gorm:"type:varchar(30);not null;"`
    Place string `gorm:"type:varchar(30);not null;"`
}

type Stunum2017 struct {
    Stunum string `gorm:"type:char(10);not null primary key;"`
}

type Stunum2016 struct {
    Stunum string `gorm:"type:char(10);not null primary key;"`
}

type Data2017 struct {
    Stunum string `gorm:"type:char(10);not null;index:stunum"`
    Time string `gorm:"type:varchar(20);not null;"`
    Cost string `gorm:"type:varhchar(8);not null;"`
    Restaurant string `gorm:"type:varchar(30);not null;"`
    Place string `gorm:"type:varchar(30);not null;"`
}
