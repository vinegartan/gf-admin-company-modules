package co_hook

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
)

type InviteJoinTeamHookFunc func(ctx context.Context, state sys_enum.InviteType, invite *sys_model.InviteRes, teamInfo *co_entity.CompanyTeam) (bool, error)
