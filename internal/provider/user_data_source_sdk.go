// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *UserDataSourceModel) RefreshFromGetResponse(resp *shared.User) {
	if resp.CreatedAt != nil {
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339))
	} else {
		r.CreatedAt = types.StringNull()
	}
	if resp.DelegatedUserID != nil {
		r.DelegatedUserID = types.StringValue(*resp.DelegatedUserID)
	} else {
		r.DelegatedUserID = types.StringNull()
	}
	if resp.DeletedAt != nil {
		r.DeletedAt = types.StringValue(resp.DeletedAt.Format(time.RFC3339))
	} else {
		r.DeletedAt = types.StringNull()
	}
	if resp.Department != nil {
		r.Department = types.StringValue(*resp.Department)
	} else {
		r.Department = types.StringNull()
	}
	r.DepartmentSources = nil
	for _, departmentSourcesItem := range resp.DepartmentSources {
		var departmentSources1 UserAttributeMappingSource
		if departmentSourcesItem.AppID != nil {
			departmentSources1.AppID = types.StringValue(*departmentSourcesItem.AppID)
		} else {
			departmentSources1.AppID = types.StringNull()
		}
		if departmentSourcesItem.AppUserID != nil {
			departmentSources1.AppUserID = types.StringValue(*departmentSourcesItem.AppUserID)
		} else {
			departmentSources1.AppUserID = types.StringNull()
		}
		if departmentSourcesItem.AppUserProfileAttributeKey != nil {
			departmentSources1.AppUserProfileAttributeKey = types.StringValue(*departmentSourcesItem.AppUserProfileAttributeKey)
		} else {
			departmentSources1.AppUserProfileAttributeKey = types.StringNull()
		}
		if departmentSourcesItem.UserAttributeMappingID != nil {
			departmentSources1.UserAttributeMappingID = types.StringValue(*departmentSourcesItem.UserAttributeMappingID)
		} else {
			departmentSources1.UserAttributeMappingID = types.StringNull()
		}
		if departmentSourcesItem.Value != nil {
			departmentSources1.Value = types.StringValue(*departmentSourcesItem.Value)
		} else {
			departmentSources1.Value = types.StringNull()
		}
		r.DepartmentSources = append(r.DepartmentSources, departmentSources1)
	}
	r.DirectoryIds = nil
	for _, v := range resp.DirectoryIds {
		r.DirectoryIds = append(r.DirectoryIds, types.StringValue(v))
	}
	if resp.DirectoryStatus != nil {
		r.DirectoryStatus = types.StringValue(string(*resp.DirectoryStatus))
	} else {
		r.DirectoryStatus = types.StringNull()
	}
	r.DirectoryStatusSources = nil
	for _, directoryStatusSourcesItem := range resp.DirectoryStatusSources {
		var directoryStatusSources1 UserAttributeMappingSource
		if directoryStatusSourcesItem.AppID != nil {
			directoryStatusSources1.AppID = types.StringValue(*directoryStatusSourcesItem.AppID)
		} else {
			directoryStatusSources1.AppID = types.StringNull()
		}
		if directoryStatusSourcesItem.AppUserID != nil {
			directoryStatusSources1.AppUserID = types.StringValue(*directoryStatusSourcesItem.AppUserID)
		} else {
			directoryStatusSources1.AppUserID = types.StringNull()
		}
		if directoryStatusSourcesItem.AppUserProfileAttributeKey != nil {
			directoryStatusSources1.AppUserProfileAttributeKey = types.StringValue(*directoryStatusSourcesItem.AppUserProfileAttributeKey)
		} else {
			directoryStatusSources1.AppUserProfileAttributeKey = types.StringNull()
		}
		if directoryStatusSourcesItem.UserAttributeMappingID != nil {
			directoryStatusSources1.UserAttributeMappingID = types.StringValue(*directoryStatusSourcesItem.UserAttributeMappingID)
		} else {
			directoryStatusSources1.UserAttributeMappingID = types.StringNull()
		}
		if directoryStatusSourcesItem.Value != nil {
			directoryStatusSources1.Value = types.StringValue(*directoryStatusSourcesItem.Value)
		} else {
			directoryStatusSources1.Value = types.StringNull()
		}
		r.DirectoryStatusSources = append(r.DirectoryStatusSources, directoryStatusSources1)
	}
	if resp.DisplayName != nil {
		r.DisplayName = types.StringValue(*resp.DisplayName)
	} else {
		r.DisplayName = types.StringNull()
	}
	if resp.Email != nil {
		r.Email = types.StringValue(*resp.Email)
	} else {
		r.Email = types.StringNull()
	}
	if resp.EmploymentStatus != nil {
		r.EmploymentStatus = types.StringValue(*resp.EmploymentStatus)
	} else {
		r.EmploymentStatus = types.StringNull()
	}
	r.EmploymentStatusSources = nil
	for _, employmentStatusSourcesItem := range resp.EmploymentStatusSources {
		var employmentStatusSources1 UserAttributeMappingSource
		if employmentStatusSourcesItem.AppID != nil {
			employmentStatusSources1.AppID = types.StringValue(*employmentStatusSourcesItem.AppID)
		} else {
			employmentStatusSources1.AppID = types.StringNull()
		}
		if employmentStatusSourcesItem.AppUserID != nil {
			employmentStatusSources1.AppUserID = types.StringValue(*employmentStatusSourcesItem.AppUserID)
		} else {
			employmentStatusSources1.AppUserID = types.StringNull()
		}
		if employmentStatusSourcesItem.AppUserProfileAttributeKey != nil {
			employmentStatusSources1.AppUserProfileAttributeKey = types.StringValue(*employmentStatusSourcesItem.AppUserProfileAttributeKey)
		} else {
			employmentStatusSources1.AppUserProfileAttributeKey = types.StringNull()
		}
		if employmentStatusSourcesItem.UserAttributeMappingID != nil {
			employmentStatusSources1.UserAttributeMappingID = types.StringValue(*employmentStatusSourcesItem.UserAttributeMappingID)
		} else {
			employmentStatusSources1.UserAttributeMappingID = types.StringNull()
		}
		if employmentStatusSourcesItem.Value != nil {
			employmentStatusSources1.Value = types.StringValue(*employmentStatusSourcesItem.Value)
		} else {
			employmentStatusSources1.Value = types.StringNull()
		}
		r.EmploymentStatusSources = append(r.EmploymentStatusSources, employmentStatusSources1)
	}
	if resp.EmploymentType != nil {
		r.EmploymentType = types.StringValue(*resp.EmploymentType)
	} else {
		r.EmploymentType = types.StringNull()
	}
	r.EmploymentTypeSources = nil
	for _, employmentTypeSourcesItem := range resp.EmploymentTypeSources {
		var employmentTypeSources1 UserAttributeMappingSource
		if employmentTypeSourcesItem.AppID != nil {
			employmentTypeSources1.AppID = types.StringValue(*employmentTypeSourcesItem.AppID)
		} else {
			employmentTypeSources1.AppID = types.StringNull()
		}
		if employmentTypeSourcesItem.AppUserID != nil {
			employmentTypeSources1.AppUserID = types.StringValue(*employmentTypeSourcesItem.AppUserID)
		} else {
			employmentTypeSources1.AppUserID = types.StringNull()
		}
		if employmentTypeSourcesItem.AppUserProfileAttributeKey != nil {
			employmentTypeSources1.AppUserProfileAttributeKey = types.StringValue(*employmentTypeSourcesItem.AppUserProfileAttributeKey)
		} else {
			employmentTypeSources1.AppUserProfileAttributeKey = types.StringNull()
		}
		if employmentTypeSourcesItem.UserAttributeMappingID != nil {
			employmentTypeSources1.UserAttributeMappingID = types.StringValue(*employmentTypeSourcesItem.UserAttributeMappingID)
		} else {
			employmentTypeSources1.UserAttributeMappingID = types.StringNull()
		}
		if employmentTypeSourcesItem.Value != nil {
			employmentTypeSources1.Value = types.StringValue(*employmentTypeSourcesItem.Value)
		} else {
			employmentTypeSources1.Value = types.StringNull()
		}
		r.EmploymentTypeSources = append(r.EmploymentTypeSources, employmentTypeSources1)
	}
	if resp.ID != nil {
		r.ID = types.StringValue(*resp.ID)
	} else {
		r.ID = types.StringNull()
	}
	if resp.JobTitle != nil {
		r.JobTitle = types.StringValue(*resp.JobTitle)
	} else {
		r.JobTitle = types.StringNull()
	}
	r.JobTitleSources = nil
	for _, jobTitleSourcesItem := range resp.JobTitleSources {
		var jobTitleSources1 UserAttributeMappingSource
		if jobTitleSourcesItem.AppID != nil {
			jobTitleSources1.AppID = types.StringValue(*jobTitleSourcesItem.AppID)
		} else {
			jobTitleSources1.AppID = types.StringNull()
		}
		if jobTitleSourcesItem.AppUserID != nil {
			jobTitleSources1.AppUserID = types.StringValue(*jobTitleSourcesItem.AppUserID)
		} else {
			jobTitleSources1.AppUserID = types.StringNull()
		}
		if jobTitleSourcesItem.AppUserProfileAttributeKey != nil {
			jobTitleSources1.AppUserProfileAttributeKey = types.StringValue(*jobTitleSourcesItem.AppUserProfileAttributeKey)
		} else {
			jobTitleSources1.AppUserProfileAttributeKey = types.StringNull()
		}
		if jobTitleSourcesItem.UserAttributeMappingID != nil {
			jobTitleSources1.UserAttributeMappingID = types.StringValue(*jobTitleSourcesItem.UserAttributeMappingID)
		} else {
			jobTitleSources1.UserAttributeMappingID = types.StringNull()
		}
		if jobTitleSourcesItem.Value != nil {
			jobTitleSources1.Value = types.StringValue(*jobTitleSourcesItem.Value)
		} else {
			jobTitleSources1.Value = types.StringNull()
		}
		r.JobTitleSources = append(r.JobTitleSources, jobTitleSources1)
	}
	r.ManagerIds = nil
	for _, v := range resp.ManagerIds {
		r.ManagerIds = append(r.ManagerIds, types.StringValue(v))
	}
	r.ManagerSources = nil
	for _, managerSourcesItem := range resp.ManagerSources {
		var managerSources1 UserAttributeMappingSource
		if managerSourcesItem.AppID != nil {
			managerSources1.AppID = types.StringValue(*managerSourcesItem.AppID)
		} else {
			managerSources1.AppID = types.StringNull()
		}
		if managerSourcesItem.AppUserID != nil {
			managerSources1.AppUserID = types.StringValue(*managerSourcesItem.AppUserID)
		} else {
			managerSources1.AppUserID = types.StringNull()
		}
		if managerSourcesItem.AppUserProfileAttributeKey != nil {
			managerSources1.AppUserProfileAttributeKey = types.StringValue(*managerSourcesItem.AppUserProfileAttributeKey)
		} else {
			managerSources1.AppUserProfileAttributeKey = types.StringNull()
		}
		if managerSourcesItem.UserAttributeMappingID != nil {
			managerSources1.UserAttributeMappingID = types.StringValue(*managerSourcesItem.UserAttributeMappingID)
		} else {
			managerSources1.UserAttributeMappingID = types.StringNull()
		}
		if managerSourcesItem.Value != nil {
			managerSources1.Value = types.StringValue(*managerSourcesItem.Value)
		} else {
			managerSources1.Value = types.StringNull()
		}
		r.ManagerSources = append(r.ManagerSources, managerSources1)
	}
	r.RoleIds = nil
	for _, v := range resp.RoleIds {
		r.RoleIds = append(r.RoleIds, types.StringValue(v))
	}
	if resp.Status != nil {
		r.Status = types.StringValue(string(*resp.Status))
	} else {
		r.Status = types.StringNull()
	}
	if resp.UpdatedAt != nil {
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339))
	} else {
		r.UpdatedAt = types.StringNull()
	}
}
