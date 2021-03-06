--1. 部门信息
drop table if exists fast.safeDepart;
create table fast.safeDepart(
  id                  varchar(255)    default ''     not null,    --系统ID
  serialId            integer         default 0      not null,    --序列号, 标识顺序,仅排序用
  parentPath          varchar(255)    default ''     not null,    --父部门路径
  departName          varchar(255)    default ''     not null,    --部门名称
  departText          varchar(255)    default ''     not null,    --部门文本信息
  departImage         varchar(255)    default ''     not null,    --部门图形信息
  departRemark        varchar(255)    default ''     not null,    --部门备注
  primary key(id)
);


--2. 安全用户
drop table if exists fast.safeUser;
create table fast.safeUser(
  id                  varchar(255)    default ''     not null,    --系统ID
  departId            varchar(255)    default ''     not null,    --部门ID
  name                varchar(255)    default ''     not null,    --用户名称
  password            varchar(255)    default ''     not null,    --登录密码
  employeeId          varchar(255)    default ''     not null,    --员工ID
  nickName            varchar(255)    default ''     not null,    --用户昵称
  firstName           varchar(255)    default ''     not null,    --FirstName
  lastName            varchar(255)    default ''     not null,    --LastName
  mobile              varchar(255)    default ''     not null,    --手机号码
  email               varchar(255)    default ''     not null,    --电子邮箱
  userRemark          varchar(255)    default ''     not null,    --用户备注
  isLock              bool            default false  not null,    --是否锁定
  lockReason          varchar(255)    default ''     not null,    --锁定原因
  isForever           bool            default true   not null,    --是否永久帐号
  accountExpired      varchar(255)    default ''     not null,    --帐号有效期
  primary key(id)
);


--3. 安全会话
drop table if exists fast.safeSession;
create table fast.safeSession(
  id                  varchar(255)    default ''     not null,    --系统ID
  randId              varchar(255)    default ''     not null,    --会话随机ID, 登录会话协商获得
  connTime            varchar(255)    default ''     not null,    --连接时间
  loginUserId         varchar(255)    default ''     not null,    --登录用户ID
  loginUserName       varchar(255)    default ''     not null,    --登录用户名
  clientIp            varchar(255)    default ''     not null,    --客户端IP地址
  clientHost          varchar(255)    default ''     not null,    --客户端主机名
  serverIp            varchar(255)    default ''     not null,    --服务器IP地址
  serverHost          varchar(255)    default ''     not null,    --服务器主机名
  keepAliveTime       varchar(255)    default ''     not null,    --最新会话保活时间
  sessionStatus       varchar(255)    default ''     not null,    --会话状态
  primary key(id)
);


--4. 安全角色
drop table if exists fast.safeRole;
create table fast.safeRole(
  id                  varchar(255)    default ''     not null,    --系统ID
  name                varchar(255)    default ''     not null,    --角色名称
  roleDetail          varchar(255)    default ''     not null,    --角色详细信息
  roleRemark          varchar(255)    default ''     not null,    --角色备注
  isUser              bool            default false  not null,    --用户映射标志,按用户授权使用
  isDepart            bool            default false  not null,    --部门映射标志,按部门授权使用
  defaultSafeObject   varchar(255)    default ''     not null,    --默认加载的安全对象
  primary key(id)
);


--5. 角色分配
drop table if exists fast.roleAlloc;
create table fast.roleAlloc(
  roleId              varchar(255)    default ''     not null,    --角色ID
  userId              varchar(255)    default ''     not null,    --用户ID
  primary key(roleId,userId)
);


--6. 业务子系统定义
drop table if exists fast.subSys;
create table fast.subSys(
  id                varchar(255)    default ''     not null,    --业务子系统ID
  name              varchar(255)    default ''     not null,    --业务子系统名称
  primary key(id)
);


--7. 安全对象组
drop table if exists fast.objectGroup;
create table fast.objectGroup(
  subSysId          varchar(255)    default ''     not null,    --业务子系统ID
  id                varchar(255)    default ''     not null,    --对象组ID
  name              varchar(255)    default ''     not null,    --对象组名称
  primary key(id)
);


--8. 安全对象
drop table if exists fast.safeObject;
create table fast.safeObject(
  id                  varchar(255)   default ''      not null,    --系统ID
  serialId            integer        default 0       not null,    --序列号, 标识顺序,仅排序用
  objectType          varchar(255)   default ''      not null,    --对象类型
  parentPath          varchar(255)   default ''      not null,    --父对象路径
  objectName          varchar(255)   default ''      not null,    --对象名称
  objectText          varchar(255)   default ''      not null,    --对象文本信息
  objectImage         varchar(255)   default ''      not null,    --对象图形信息
  objectRemark        varchar(255)   default ''      not null,    --安全对象备注
  primary key(id)
);


--9. 安全对象授权
drop table if exists fast.objectAuth;
create table fast.objectAuth(
  objectId            varchar(255)   default ''      not null,    --安全对象ID
  roleId              varchar(255)   default ''      not null,    --安全角色ID
  primary key(objectId,roleId)
);


--10. 安全操作组
drop table if exists fast.operateGroup;
create table fast.operateGroup(
  subSysId          varchar(255)    default ''     not null,    --业务子系统ID  
  id                varchar(255)    default ''     not null,    --操作组ID
  name              varchar(255)    default ''     not null,    --操作组名称
  primary key(id)
);


--11. 安全操作
drop table if exists fast.safeOperate;
create table fast.safeOperate(
  id                  varchar(255)   default ''      not null,    --系统ID
  serialId            integer        default 0       not null,    --序列号, 标识顺序,仅排序用
  operateCode         varchar(255)   default ''      not null,    --操作码
  subsys              varchar(255)   default ''      not null,    --子系统
  operateGroup        varchar(255)   default ''      not null,    --操作组
  operateName         varchar(255)   default ''      not null,    --操作名称
  operateDetail       varchar(255)   default ''      not null,    --操作详细信息
  operateRemark       varchar(255)   default ''      not null,    --操作备注
  isAuth              bool           default true    not null,    --需要鉴权标志
  isLog               bool           default true    not null,    --需要记录日志标志
  primary key(id)
);


--12. 安全操作授权
drop table if exists fast.operateAuth;
create table fast.operateAuth(
  operateId           varchar(255)   default ''      not null,    --安全操作ID
  roleId              varchar(255)   default ''      not null,    --安全角色ID
  primary key(operateId,roleId)
);


--13. 系统操作日志表
drop table if exists fast.operateLog;
create table fast.operateLog(
  userName              varchar(255)   default ''    not null,    --用户ID
  subsys                varchar(255)   default ''    not null,    --子系统
  operateGroup          varchar(255)   default ''    not null,    --操作组
  operateName           varchar(255)   default ''    not null,    --操作名称
  operateRet            varchar(255)   default ''    not null,    --操作结果
  operateRetDetail      varchar(255)   default ''    not null,    --结果描述
  operateObj            varchar(255)   default ''    not null,    --操作对象
  operateContent        varchar(255)   default ''    not null,    --操作内容
  userIpAddress         varchar(255)   default ''    not null,    --操作员IP
  userHostName          varchar(255)   default ''    not null,    --操作员机器名
  operateTime           varchar(255)   default ''    not null,    --操作时间
  logType               varchar(255)   default ''    not null,    --日志类型
  serialNum             varchar(255)   default ''    not null,    --流水号
  primary key(serialNum)
);


--14. 系统参数配置表
drop table if exists fast.sysParaConf;
create table fast.sysParaConf(
  catalog            varchar(255)         not null,         --参数目录分组
  paraName           varchar(255)         not null,         --参数名称
  paraValue          varchar(65535)       not null,         --参数值
  paraType           varchar(255)         not null,         --参数类型
  paraRemark         varchar(255)         not null,         --参数说明
  primary key(paraName)
);



