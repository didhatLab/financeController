"""create objects for groups

Revision ID: e42e5e2a4a73
Revises: 09d95af18278
Create Date: 2023-04-29 19:19:43.708715

"""
import sqlalchemy as sa
from alembic import op

# revision identifiers, used by Alembic.
revision = "e42e5e2a4a73"
down_revision = "09d95af18278"
branch_labels = None
depends_on = None


def upgrade() -> None:
    op.create_table(
        "spend_group",
        sa.Column("id", sa.Integer, primary_key=True, autoincrement=True),
        sa.Column("name", sa.String, nullable=False),
        sa.Column("description", sa.String)
    ),

    op.add_column("spend", sa.Column("group_id", sa.Integer))
    op.create_foreign_key(
        None,
        "spend",
        "spend_group",
        ["group_id"],
        ["id"],
        ondelete="CASCADE"
    )


def downgrade() -> None:
    op.drop_column("spend", "group_id")
    op.drop_table("spend_group")
