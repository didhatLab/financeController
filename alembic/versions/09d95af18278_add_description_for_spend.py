"""add description for spend

Revision ID: 09d95af18278
Revises: f9ad2d990408
Create Date: 2023-04-26 10:24:04.263337

"""
from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision = "09d95af18278"
down_revision = "f9ad2d990408"
branch_labels = None
depends_on = None


def upgrade() -> None:
    op.add_column("spend", sa.Column("description", sa.Text, server_default=""))


def downgrade() -> None:
    op.drop_column("spend", "description")
