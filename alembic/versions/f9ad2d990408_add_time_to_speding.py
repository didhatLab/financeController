"""add time to speding

Revision ID: f9ad2d990408
Revises: e49cee18e7e1
Create Date: 2023-04-25 20:29:01.430540

"""
import datetime

from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision = "f9ad2d990408"
down_revision = "e49cee18e7e1"
branch_labels = None
depends_on = None


def upgrade() -> None:
    op.add_column(
        "spend",
        sa.Column(
            "time",
            sa.DateTime(),
            server_default=sa.text("NOW()"),
        ),
    )


def downgrade() -> None:
    op.drop_column("spend", "time")
