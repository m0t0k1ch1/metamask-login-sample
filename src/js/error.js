"use strict";

class AppError {
  constructor(message, code = 0) {
    this.message = message;
    this.code = code;
  }
}

export default AppError;
