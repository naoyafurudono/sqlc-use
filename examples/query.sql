-- name: GetUser :one
SELECT * FROM user WHERE id = ?;

-- name: ListOrganizationMember :many
SELECT user.* FROM user
INNER JOIN member ON user.id = member.user_id
INNER JOIN organization ON organization.id = member.organization_id
WHERE organization.name = ?;

-- name: AddMember :exec
INSERT INTO member (user_id, organization_id, role) VALUES (?, ?, ?);

-- name: UpdateMemberRole :exec
UPDATE member SET role = ? WHERE user_id = ? AND organization_id = ?;

-- name: RemoveMember :exec
DELETE FROM member WHERE user_id = ? AND organization_id = ?;

-- name: GetAllUsers :many
SELECT * FROM user WHERE active = true
UNION
SELECT * FROM user WHERE last_login > DATE_SUB(NOW(), INTERVAL 30 DAY);
