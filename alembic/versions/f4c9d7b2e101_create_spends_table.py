"""create spends table

Revision ID: f4c9d7b2e101
Revises: 
Create Date: 2023-04-16 11:34:14.983057

"""
from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision = "f4c9d7b2e101"
down_revision = None
branch_labels = None
depends_on = None


def upgrade() -> None:
    op.create_table(
        "spend",
        sa.Column("id", sa.Integer, primary_key=True, autoincrement=True),
        sa.Column("name", sa.Integer),
        sa.Column("type", sa.String, default="unknown"),
        sa.Column("user_id", sa.Integer),
        sa.Column("amount", sa.Integer),
        sa.Column("currency", sa.String),
    )


def downgrade() -> None:
    op.drop_table("spend")
