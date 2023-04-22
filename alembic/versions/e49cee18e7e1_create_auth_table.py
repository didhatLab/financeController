"""create auth table

Revision ID: e49cee18e7e1
Revises: f4c9d7b2e101
Create Date: 2023-04-22 12:12:18.650546

"""
from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision = 'e49cee18e7e1'
down_revision = 'f4c9d7b2e101'
branch_labels = None
depends_on = None


def upgrade() -> None:
    op.create_table(
        "auth",
        sa.Column("user_id", sa.Integer, primary_key=True, autoincrement=True),
        sa.Column("username", sa.String, unique=True),
        sa.Column("pass_hash", sa.String)
    )


def downgrade() -> None:
    op.drop_table("auth")
