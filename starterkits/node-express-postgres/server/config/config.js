'use strict';

require('dotenv').load();

const REQUIRED_KEYS = [
  'DB_USER',
  'DB_PASSWORD',
  'DB_NAME',
  'DB_HOST',
  'JWT_TOKEN',
  'TOKEN_EXPIRATION_TIME',
];

REQUIRED_KEYS.forEach((key) => {
  if (!(key in process.env)) {
    throw new Error(`Missing required config key: ${key}`);
  }
});

const {
  DB_USER,
  DB_PASSWORD,
  DB_NAME,
  DB_HOST,
  JWT_TOKEN,
  TOKEN_EXPIRATION_TIME,
} = process.env;

module.exports = {
  JWT_TOKEN,
  TOKEN_EXPIRATION_TIME: Number(TOKEN_EXPIRATION_TIME),

  // Sequelize config, sourced based on current NODE_ENV from models/index.js file
  [process.env.NODE_ENV || 'development']: {
    username: DB_USER,
    password: DB_PASSWORD || null,
    database: DB_NAME,
    host: DB_HOST,
    dialect: 'postgres',
  },
};
