'use strict';

class AppError {
  constructor(message, code = 0) {
    this.message = message;
    this.code    = code;
  }
}

module.exports = AppError;
