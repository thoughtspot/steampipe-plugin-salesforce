package salesforce

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func SalesforceObjectPermission(ctx context.Context, dm dynamicMap, config salesforceConfig) *plugin.Table {
	tableName := "ObjectPermissions"

	columns := mergeTableColumns(ctx, config, getCustomCols(dm), []*plugin.Column{
		// Top columns
		{Name: "id", Type: proto.ColumnType_STRING, Description: "The ObjectPermissions ID."},
		{Name: "parent_id", Type: proto.ColumnType_STRING, Description: "The Id of this object's parent PermissionSet."},
		{Name: "sobject_type", Type: proto.ColumnType_STRING, Description: "The object's API name. For example, Merchandise__c."},

		// Other columns
		{Name: "permissions_create", Type: proto.ColumnType_BOOL, Description: "If true, users assigned to the parent PermissionSet can create records for this object. Requires PermissionsRead for the same object to be true."},
		{Name: "permissions_delete", Type: proto.ColumnType_BOOL, Description: "If true, users assigned to the parent PermissionSet can delete records for this object. Requires PermissionsRead and PermissionsEdit for the same object to be true."},
		{Name: "permissions_edit", Type: proto.ColumnType_BOOL, Description: "If true, users assigned to the parent PermissionSet can edit records for this object. Requires PermissionsRead for the same object to be true."},
		{Name: "permissions_read", Type: proto.ColumnType_BOOL, Description: "If true, users assigned to the parent PermissionSet can view records for this object."},
		{Name: "permissions_modify_all_records", Type: proto.ColumnType_BOOL, Description: "If true, users assigned to the parent PermissionSet can edit all records for this object, regardless of sharing settings. Requires PermissionsRead, PermissionsDelete, PermissionsEdit, and PermissionsViewAllRecords for the same object to be true."},
		{Name: "permissions_view_all_records", Type: proto.ColumnType_BOOL, Description: "If true, users assigned to the parent PermissionSet can view all records for this object, regardless of sharing settings. Requires PermissionsRead for the same object to be true."},
	})

	plugin.Logger(ctx).Debug("SalesforceObjectPermission init")

	queryColumnsMap := make(map[string]*plugin.Column)
	for _, column := range columns {
		queryColumnsMap[getSalesforceColumnName(column.Name)] = column
	}

	return &plugin.Table{
		Name:        "salesforce_object_permission",
		Description: "Represents the enabled object permissions for the parent PermissionSet.",
		List: &plugin.ListConfig{
			Hydrate:    listSalesforceObjectsByTable(tableName, dm.salesforceColumns, queryColumnsMap),
			KeyColumns: getKeyColumns(columns),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSalesforceObjectbyID(tableName, queryColumnsMap),
			KeyColumns: plugin.SingleColumn(checkNameScheme(config, dm.cols)),
		},
		Columns: columns,
	}
}
