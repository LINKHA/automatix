package xerr

// 成功返回
const OK uint32 = 0

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码
const SERVER_COMMON_ERROR uint32 = 100001
const REUQEST_PARAM_ERROR uint32 = 100002
const TOKEN_EXPIRE_ERROR uint32 = 100003
const TOKEN_GENERATE_ERROR uint32 = 100004
const DB_ERROR uint32 = 100005
const DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 100006
const SERVER_CODE_INVALID uint32 = 100007 //ServerCode无效

// 服务器管理
const SERVER_MANAGER_LOGIN_SERVER_QUEUE_MAX uint32 = 001001   //排队人数到达上限
const SERVER_MANAGER_LOGIN_SERVER_QUEUE_ENTER uint32 = 001002 //进入排队队列
