package model

import "github.com/volatiletech/sqlboiler/v4/queries/qm"

var Active = qm.Where("deleted_at IS NULL")
