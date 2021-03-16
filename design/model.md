#

## 术语

name|means
-|-
Namespace|空间
Link|一条短链接
Identity|一个用户在系统中的角色，Admin是超级用户，User为普通用户
Nid|namespace的唯一id
FVid|FoldView的唯一id

## 短路径黑名单

name|why
-|-
/|根路径
/admin/**|超级后台
/dashboard/**|后台

## 访问频率限制

// TODO:

## 数据模型

```plantuml
@startuml

class "Namespace" {
  ID
	CreateAt
	UpdateAt
	Creator

  Name
}

class User {
  ID
	CreateAt
	UpdateAt

  UID
  Username
  Password
  Identity
}

enum Identity {
  ADMIN
  USER
}

class UserNamespace {
  ID
	CreateAt
	UpdateAt
	Creator

  Uid
  Nid
  Role
}

enum Role {
  MANAGER
  EDITOR
  VIEWER
}

class Link {
  ID
	CreateAt
	UpdateAt
	Creator

  Nid
  Short
  Target
  Description
  Enable
}

class Tag {
  ID
	CreateAt
	UpdateAt
	Creator

  Name
  Nid
}

class LinkTag {
  ID
	CreateAt
	UpdateAt
	Creator

  LinkId
  TagId
}

package folder{
  class FolderView {
    ID
    CreateAt
    UpdateAt
    Creator

    Name
    Nid
  }

  class Folder {
    ID
    CreateAt
    UpdateAt
    Creator

    Path
    FVid
  }

  class FolderTag {
    ID
    CreateAt
    UpdateAt
    Creator

    FolderId
    TagId
  }
}

User }-left- Identity
UserNamespace }-up- Role
"Namespace" *-- Link
"Namespace" *-- Tag

User -- Namespace
(User, Namespace) . UserNamespace

Link -right- Tag
(Link, Tag) . LinkTag

Namespace *-right- FolderView
FolderView *-right- Folder
(Tag, Folder) . FolderTag


hide methods
@enduml
```

## 日志

日志直接存储至业务数据库，（后续考虑采用mongo的oplog）  
ROLE和可以看到日志类型的对应关系写死到后端服务。

```plantuml
@startuml

class OpLog {
  ID
	CreateAt
	Creator

  Nid
  OpType
  Content
}

enum OpType {
  APPEND_NS

  APPEND_USER_NS
  REMOVE_USER_NS
  UPDATE_USER_NS_ROLE

  APPEND_LINK
  REMOVE_LINK
  UPDATE_LINK_TARGET
  UPDATE_LINK_TAG
  ENABLE_LINK
  UNABLE_LINK
}

OpLog }-right- OpType

hide methods
@enduml
```

```plantuml
@startuml

class UsageLog {
  ID
	CreateAt
	Creator

  LinkId
  Ip
  UserAgent
}

hide methods
@enduml
```
