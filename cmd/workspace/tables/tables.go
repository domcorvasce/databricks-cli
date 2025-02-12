// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tables

import (
	"fmt"

	"github.com/databricks/cli/cmd/root"
	"github.com/databricks/cli/libs/cmdio"
	"github.com/databricks/cli/libs/flags"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/spf13/cobra"
)

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var cmdOverrides []func(*cobra.Command)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tables",
		Short: `A table resides in the third layer of Unity Catalog’s three-level namespace.`,
		Long: `A table resides in the third layer of Unity Catalog’s three-level namespace.
  It contains rows of data. To create a table, users must have CREATE_TABLE and
  USE_SCHEMA permissions on the schema, and they must have the USE_CATALOG
  permission on its parent catalog. To query a table, users must have the SELECT
  permission on the table, and they must have the USE_CATALOG permission on its
  parent catalog and the USE_SCHEMA permission on its parent schema.
  
  A table can be managed or external. From an API perspective, a __VIEW__ is a
  particular kind of table (rather than a managed or external table).`,
		GroupID: "catalog",
		Annotations: map[string]string{
			"package": "catalog",
		},
	}

	// Apply optional overrides to this command.
	for _, fn := range cmdOverrides {
		fn(cmd)
	}

	return cmd
}

// start delete command

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var deleteOverrides []func(
	*cobra.Command,
	*catalog.DeleteTableRequest,
)

func newDelete() *cobra.Command {
	cmd := &cobra.Command{}

	var deleteReq catalog.DeleteTableRequest

	// TODO: short flags

	cmd.Use = "delete FULL_NAME"
	cmd.Short = `Delete a table.`
	cmd.Long = `Delete a table.
  
  Deletes a table from the specified parent catalog and schema. The caller must
  be the owner of the parent catalog, have the **USE_CATALOG** privilege on the
  parent catalog and be the owner of the parent schema, or be the owner of the
  table and have the **USE_CATALOG** privilege on the parent catalog and the
  **USE_SCHEMA** privilege on the parent schema.

  Arguments:
    FULL_NAME: Full name of the table.`

	cmd.Annotations = make(map[string]string)

	cmd.PreRunE = root.MustWorkspaceClient
	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)

		if len(args) == 0 {
			promptSpinner := cmdio.Spinner(ctx)
			promptSpinner <- "No FULL_NAME argument specified. Loading names for Tables drop-down."
			names, err := w.Tables.TableInfoNameToTableIdMap(ctx, catalog.ListTablesRequest{})
			close(promptSpinner)
			if err != nil {
				return fmt.Errorf("failed to load names for Tables drop-down. Please manually specify required arguments. Original error: %w", err)
			}
			id, err := cmdio.Select(ctx, names, "Full name of the table")
			if err != nil {
				return err
			}
			args = append(args, id)
		}
		if len(args) != 1 {
			return fmt.Errorf("expected to have full name of the table")
		}
		deleteReq.FullName = args[0]

		err = w.Tables.Delete(ctx, deleteReq)
		if err != nil {
			return err
		}
		return nil
	}

	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	cmd.ValidArgsFunction = cobra.NoFileCompletions

	// Apply optional overrides to this command.
	for _, fn := range deleteOverrides {
		fn(cmd, &deleteReq)
	}

	return cmd
}

func init() {
	cmdOverrides = append(cmdOverrides, func(cmd *cobra.Command) {
		cmd.AddCommand(newDelete())
	})
}

// start get command

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var getOverrides []func(
	*cobra.Command,
	*catalog.GetTableRequest,
)

func newGet() *cobra.Command {
	cmd := &cobra.Command{}

	var getReq catalog.GetTableRequest

	// TODO: short flags

	cmd.Flags().BoolVar(&getReq.IncludeDeltaMetadata, "include-delta-metadata", getReq.IncludeDeltaMetadata, `Whether delta metadata should be included in the response.`)

	cmd.Use = "get FULL_NAME"
	cmd.Short = `Get a table.`
	cmd.Long = `Get a table.
  
  Gets a table from the metastore for a specific catalog and schema. The caller
  must be a metastore admin, be the owner of the table and have the
  **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
  privilege on the parent schema, or be the owner of the table and have the
  **SELECT** privilege on it as well.

  Arguments:
    FULL_NAME: Full name of the table.`

	cmd.Annotations = make(map[string]string)

	cmd.PreRunE = root.MustWorkspaceClient
	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)

		if len(args) == 0 {
			promptSpinner := cmdio.Spinner(ctx)
			promptSpinner <- "No FULL_NAME argument specified. Loading names for Tables drop-down."
			names, err := w.Tables.TableInfoNameToTableIdMap(ctx, catalog.ListTablesRequest{})
			close(promptSpinner)
			if err != nil {
				return fmt.Errorf("failed to load names for Tables drop-down. Please manually specify required arguments. Original error: %w", err)
			}
			id, err := cmdio.Select(ctx, names, "Full name of the table")
			if err != nil {
				return err
			}
			args = append(args, id)
		}
		if len(args) != 1 {
			return fmt.Errorf("expected to have full name of the table")
		}
		getReq.FullName = args[0]

		response, err := w.Tables.Get(ctx, getReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	}

	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	cmd.ValidArgsFunction = cobra.NoFileCompletions

	// Apply optional overrides to this command.
	for _, fn := range getOverrides {
		fn(cmd, &getReq)
	}

	return cmd
}

func init() {
	cmdOverrides = append(cmdOverrides, func(cmd *cobra.Command) {
		cmd.AddCommand(newGet())
	})
}

// start list command

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var listOverrides []func(
	*cobra.Command,
	*catalog.ListTablesRequest,
)

func newList() *cobra.Command {
	cmd := &cobra.Command{}

	var listReq catalog.ListTablesRequest

	// TODO: short flags

	cmd.Flags().BoolVar(&listReq.IncludeDeltaMetadata, "include-delta-metadata", listReq.IncludeDeltaMetadata, `Whether delta metadata should be included in the response.`)
	cmd.Flags().IntVar(&listReq.MaxResults, "max-results", listReq.MaxResults, `Maximum number of tables to return (page length).`)
	cmd.Flags().StringVar(&listReq.PageToken, "page-token", listReq.PageToken, `Opaque token to send for the next page of results (pagination).`)

	cmd.Use = "list CATALOG_NAME SCHEMA_NAME"
	cmd.Short = `List tables.`
	cmd.Long = `List tables.
  
  Gets an array of all tables for the current metastore under the parent catalog
  and schema. The caller must be a metastore admin or an owner of (or have the
  **SELECT** privilege on) the table. For the latter case, the caller must also
  be the owner or have the **USE_CATALOG** privilege on the parent catalog and
  the **USE_SCHEMA** privilege on the parent schema. There is no guarantee of a
  specific ordering of the elements in the array.

  Arguments:
    CATALOG_NAME: Name of parent catalog for tables of interest.
    SCHEMA_NAME: Parent schema of tables.`

	cmd.Annotations = make(map[string]string)

	cmd.Args = func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(2)
		return check(cmd, args)
	}

	cmd.PreRunE = root.MustWorkspaceClient
	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)

		listReq.CatalogName = args[0]
		listReq.SchemaName = args[1]

		response, err := w.Tables.ListAll(ctx, listReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	}

	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	cmd.ValidArgsFunction = cobra.NoFileCompletions

	// Apply optional overrides to this command.
	for _, fn := range listOverrides {
		fn(cmd, &listReq)
	}

	return cmd
}

func init() {
	cmdOverrides = append(cmdOverrides, func(cmd *cobra.Command) {
		cmd.AddCommand(newList())
	})
}

// start list-summaries command

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var listSummariesOverrides []func(
	*cobra.Command,
	*catalog.ListSummariesRequest,
)

func newListSummaries() *cobra.Command {
	cmd := &cobra.Command{}

	var listSummariesReq catalog.ListSummariesRequest

	// TODO: short flags

	cmd.Flags().IntVar(&listSummariesReq.MaxResults, "max-results", listSummariesReq.MaxResults, `Maximum number of tables to return (page length).`)
	cmd.Flags().StringVar(&listSummariesReq.PageToken, "page-token", listSummariesReq.PageToken, `Opaque token to send for the next page of results (pagination).`)
	cmd.Flags().StringVar(&listSummariesReq.SchemaNamePattern, "schema-name-pattern", listSummariesReq.SchemaNamePattern, `A sql LIKE pattern (% and _) for schema names.`)
	cmd.Flags().StringVar(&listSummariesReq.TableNamePattern, "table-name-pattern", listSummariesReq.TableNamePattern, `A sql LIKE pattern (% and _) for table names.`)

	cmd.Use = "list-summaries CATALOG_NAME"
	cmd.Short = `List table summaries.`
	cmd.Long = `List table summaries.
  
  Gets an array of summaries for tables for a schema and catalog within the
  metastore. The table summaries returned are either:
  
  * summaries for all tables (within the current metastore and parent catalog
  and schema), when the user is a metastore admin, or: * summaries for all
  tables and schemas (within the current metastore and parent catalog) for which
  the user has ownership or the **SELECT** privilege on the table and ownership
  or **USE_SCHEMA** privilege on the schema, provided that the user also has
  ownership or the **USE_CATALOG** privilege on the parent catalog.
  
  There is no guarantee of a specific ordering of the elements in the array.

  Arguments:
    CATALOG_NAME: Name of parent catalog for tables of interest.`

	cmd.Annotations = make(map[string]string)

	cmd.PreRunE = root.MustWorkspaceClient
	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)

		if len(args) == 0 {
			promptSpinner := cmdio.Spinner(ctx)
			promptSpinner <- "No CATALOG_NAME argument specified. Loading names for Tables drop-down."
			names, err := w.Tables.TableInfoNameToTableIdMap(ctx, catalog.ListTablesRequest{})
			close(promptSpinner)
			if err != nil {
				return fmt.Errorf("failed to load names for Tables drop-down. Please manually specify required arguments. Original error: %w", err)
			}
			id, err := cmdio.Select(ctx, names, "Name of parent catalog for tables of interest")
			if err != nil {
				return err
			}
			args = append(args, id)
		}
		if len(args) != 1 {
			return fmt.Errorf("expected to have name of parent catalog for tables of interest")
		}
		listSummariesReq.CatalogName = args[0]

		response, err := w.Tables.ListSummariesAll(ctx, listSummariesReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	}

	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	cmd.ValidArgsFunction = cobra.NoFileCompletions

	// Apply optional overrides to this command.
	for _, fn := range listSummariesOverrides {
		fn(cmd, &listSummariesReq)
	}

	return cmd
}

func init() {
	cmdOverrides = append(cmdOverrides, func(cmd *cobra.Command) {
		cmd.AddCommand(newListSummaries())
	})
}

// start update command

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var updateOverrides []func(
	*cobra.Command,
	*catalog.UpdateTableRequest,
)

func newUpdate() *cobra.Command {
	cmd := &cobra.Command{}

	var updateReq catalog.UpdateTableRequest
	var updateJson flags.JsonFlag

	// TODO: short flags
	cmd.Flags().Var(&updateJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	cmd.Flags().StringVar(&updateReq.Owner, "owner", updateReq.Owner, ``)

	cmd.Use = "update FULL_NAME"
	cmd.Short = `Update a table owner.`
	cmd.Long = `Update a table owner.
  
  Change the owner of the table. The caller must be the owner of the parent
  catalog, have the **USE_CATALOG** privilege on the parent catalog and be the
  owner of the parent schema, or be the owner of the table and have the
  **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
  privilege on the parent schema.

  Arguments:
    FULL_NAME: Full name of the table.`

	// This command is being previewed; hide from help output.
	cmd.Hidden = true

	cmd.Annotations = make(map[string]string)

	cmd.PreRunE = root.MustWorkspaceClient
	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)

		if cmd.Flags().Changed("json") {
			err = updateJson.Unmarshal(&updateReq)
			if err != nil {
				return err
			}
		}
		if len(args) == 0 {
			promptSpinner := cmdio.Spinner(ctx)
			promptSpinner <- "No FULL_NAME argument specified. Loading names for Tables drop-down."
			names, err := w.Tables.TableInfoNameToTableIdMap(ctx, catalog.ListTablesRequest{})
			close(promptSpinner)
			if err != nil {
				return fmt.Errorf("failed to load names for Tables drop-down. Please manually specify required arguments. Original error: %w", err)
			}
			id, err := cmdio.Select(ctx, names, "Full name of the table")
			if err != nil {
				return err
			}
			args = append(args, id)
		}
		if len(args) != 1 {
			return fmt.Errorf("expected to have full name of the table")
		}
		updateReq.FullName = args[0]

		err = w.Tables.Update(ctx, updateReq)
		if err != nil {
			return err
		}
		return nil
	}

	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	cmd.ValidArgsFunction = cobra.NoFileCompletions

	// Apply optional overrides to this command.
	for _, fn := range updateOverrides {
		fn(cmd, &updateReq)
	}

	return cmd
}

func init() {
	cmdOverrides = append(cmdOverrides, func(cmd *cobra.Command) {
		cmd.AddCommand(newUpdate())
	})
}

// end service Tables
