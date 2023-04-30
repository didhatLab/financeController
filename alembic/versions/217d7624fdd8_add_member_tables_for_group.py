"""add member tables for group

Revision ID: 217d7624fdd8
Revises: e42e5e2a4a73
Create Date: 2023-04-29 19:48:33.836507

"""
from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision = "217d7624fdd8"
down_revision = "e42e5e2a4a73"
branch_labels = None
depends_on = None


def upgrade() -> None:
    op.create_table(
        "group_member",
        sa.Column("user_id", sa.Integer, primary_key=True, nullable=False),
        sa.Column("group_id", sa.Integer, primary_key=True, nullable=False),
    )
    op.create_foreign_key(None, "group_member", "auth", ["user_id"], ["user_id"])
    op.create_foreign_key(None, "group_member", "spend_group", ["group_id"], ["id"])


def downgrade() -> None:
    op.drop_table("group_member")
