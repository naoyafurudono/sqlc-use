# sqlc-use

sqlc-useはsqlcプラグインで、それぞれのクエリがテーブルに対してどのような操作をするかを表現するJSONを生成します。

## 例

### 入力

```sql
-- name: ListOrganizationMember :many
select user.* from user
inner join member on user.id = member.user_id
inner join organization on organization.id = member.organization_id
where organization.name = ?;

-- name: AddMember :exec
insert into member (user_id, organization_id) (?, ?);

-- name: RemoveMember :exec
delete from member where user_id = ? and organization_id = ?;
```

### 出力

```json
{
  "github.com/sqlc-use/repository.ListOrganizationMember": [
    {
      "operation": "select",
      "table": "user"
    },
    {
      "operation": "select",
      "table": "member"
    },
    {
      "operation": "select",
      "table": "organization"
    }
  ],
  "github.com/sqlc-use/repository.AddMember": [
    {
      "operation": "insert",
      "table": "member"
    }
  ],
  "github.com/sqlc-use/repository.RemoveMember": [
    {
      "operation": "delete",
      "table": "member"
    }
  ]
}
```
