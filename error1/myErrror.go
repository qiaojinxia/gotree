package error1

import "errors"

/**
 * Created by @CaomaoBoy on 2021/5/4.
 *  email:<115882934@qq.com>
 */

type MyError error

var RPCERROR = errors.New("Rpc业务异常!")
