#!/usr/bin/env bash
set -e

DB_PATH=${DB_PATH:-/dashboard/data/data.db}

mkdir -p /dashboard/data

echo "🚀 Starting Nezha Dashboard with Traffic Control"
echo "📦 DB: $DB_PATH"

exec /dashboard/nezha-dashboard
