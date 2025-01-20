#include "open62541_cgo.h"

UA_StatusCode UA_VariantType(UA_Variant *variant, UA_UInt32 *ptype) {
  if (variant->type == &UA_TYPES[UA_TYPES_BOOLEAN]) {
    *ptype = UA_TYPES_BOOLEAN;
  } else if (variant->type == &UA_TYPES[UA_TYPES_SBYTE]) {
    *ptype = UA_TYPES_SBYTE;
  } else if (variant->type == &UA_TYPES[UA_TYPES_BYTE]) {
    *ptype = UA_TYPES_BYTE;
  } else if (variant->type == &UA_TYPES[UA_TYPES_INT16]) {
    *ptype = UA_TYPES_INT16;
  } else if (variant->type == &UA_TYPES[UA_TYPES_UINT16]) {
    *ptype = UA_TYPES_UINT16;
  } else if (variant->type == &UA_TYPES[UA_TYPES_INT32]) {
    *ptype = UA_TYPES_INT32;
  } else if (variant->type == &UA_TYPES[UA_TYPES_UINT32]) {
    *ptype = UA_TYPES_UINT32;
  } else if (variant->type == &UA_TYPES[UA_TYPES_INT64]) {
    *ptype = UA_TYPES_INT64;
  } else if (variant->type == &UA_TYPES[UA_TYPES_UINT64]) {
    *ptype = UA_TYPES_UINT64;
  } else if (variant->type == &UA_TYPES[UA_TYPES_FLOAT]) {
    *ptype = UA_TYPES_FLOAT;
  } else if (variant->type == &UA_TYPES[UA_TYPES_DOUBLE]) {
    *ptype = UA_TYPES_DOUBLE;
  } else if (variant->type == &UA_TYPES[UA_TYPES_STRING]) {
    *ptype = UA_TYPES_STRING;
  } else if (variant->type == &UA_TYPES[UA_TYPES_DATETIME]) {
    *ptype = UA_TYPES_DATETIME;
  } else if (variant->type == &UA_TYPES[UA_TYPES_GUID]) {
    *ptype = UA_TYPES_GUID;
  } else if (variant->type == &UA_TYPES[UA_TYPES_BYTESTRING]) {
    *ptype = UA_TYPES_BYTESTRING;
  } else if (variant->type == &UA_TYPES[UA_TYPES_XMLELEMENT]) {
    *ptype = UA_TYPES_XMLELEMENT;
  } else if (variant->type == &UA_TYPES[UA_TYPES_NODEID]) {
    *ptype = UA_TYPES_NODEID;
  } else if (variant->type == &UA_TYPES[UA_TYPES_EXPANDEDNODEID]) {
    *ptype = UA_TYPES_EXPANDEDNODEID;
  } else if (variant->type == &UA_TYPES[UA_TYPES_STATUSCODE]) {
    *ptype = UA_TYPES_STATUSCODE;
  } else if (variant->type == &UA_TYPES[UA_TYPES_QUALIFIEDNAME]) {
    *ptype = UA_TYPES_QUALIFIEDNAME;
  } else if (variant->type == &UA_TYPES[UA_TYPES_LOCALIZEDTEXT]) {
    *ptype = UA_TYPES_LOCALIZEDTEXT;
  } else if (variant->type == &UA_TYPES[UA_TYPES_EXTENSIONOBJECT]) {
    *ptype = UA_TYPES_EXTENSIONOBJECT;
  } else if (variant->type == &UA_TYPES[UA_TYPES_DATAVALUE]) {
    *ptype = UA_TYPES_DATAVALUE;
  } else if (variant->type == &UA_TYPES[UA_TYPES_VARIANT]) {
    *ptype = UA_TYPES_VARIANT;
  } else if (variant->type == &UA_TYPES[UA_TYPES_DIAGNOSTICINFO]) {
    *ptype = UA_TYPES_DIAGNOSTICINFO;
  } else {
    return UA_STATUSCODE_BADTYPEMISMATCH;
  }
  return UA_STATUSCODE_GOOD;
}

UA_Boolean UA_VariantValueBoolean(UA_Variant *value, int index) {
  UA_Boolean *valueData = (UA_Boolean *)value->data;
  return valueData[index];
}

UA_SByte UA_VariantValueInt8(UA_Variant *value, int index) {
  UA_SByte *valueData = (UA_SByte *)value->data;
  return valueData[index];
}

UA_Byte UA_VariantValueUint8(UA_Variant *value, int index) {
  UA_Byte *valueData = (UA_Byte *)value->data;
  return valueData[index];
}

UA_Int16 UA_VariantValueInt16(UA_Variant *value, int index) {
  UA_Int16 *valueData = (UA_Int16 *)value->data;
  return valueData[index];
}

UA_UInt16 UA_VariantValueUint16(UA_Variant *value, int index) {
  UA_UInt16 *valueData = (UA_UInt16 *)value->data;
  return valueData[index];
}

UA_Int32 UA_VariantValueInt32(UA_Variant *value, int index) {
  UA_Int32 *valueData = (UA_Int32 *)value->data;
  return valueData[index];
}

UA_UInt32 UA_VariantValueUint32(UA_Variant *value, int index) {
  UA_UInt32 *valueData = (UA_UInt32 *)value->data;
  return valueData[index];
}

UA_Int64 UA_VariantValueInt64(UA_Variant *value, int index) {
  UA_Int64 *valueData = (UA_Int64 *)value->data;
  return valueData[index];
}

UA_UInt64 UA_VariantValueUint64(UA_Variant *value, int index) {
  UA_UInt64 *valueData = (UA_UInt64 *)value->data;
  return valueData[index];
}

UA_Float UA_VariantValueFloat(UA_Variant *value, int index) {
  UA_Float *valueData = (UA_Float *)value->data;
  return valueData[index];
}

UA_Double UA_VariantValueDouble(UA_Variant *value, int index) {
  UA_Double *valueData = (UA_Double *)value->data;
  return valueData[index];
}

UA_String UA_VariantValueString(UA_Variant *value, int index) {
  UA_String *valueData = (UA_String *)value->data;
  return valueData[index];
}

UA_DateTime UA_VariantValueDateTime(UA_Variant *value, int index) {
  UA_DateTime *valueData = (UA_DateTime *)value->data;
  return valueData[index];
}

UA_ByteString UA_VariantValueByteString(UA_Variant *value, int index) {
  UA_ByteString *valueData = (UA_ByteString *)value->data;
  return valueData[index];
}

UA_StatusCode UA_VariantScalarValueBoolean(UA_Variant *variant,
                                           UA_Boolean value) {
  return UA_Variant_setScalarCopy(variant, &value, &UA_TYPES[UA_TYPES_BOOLEAN]);
}

UA_StatusCode UA_VariantScalarValueInt8(UA_Variant *variant, UA_SByte value) {
  return UA_Variant_setScalarCopy(variant, &value, &UA_TYPES[UA_TYPES_SBYTE]);
}

UA_StatusCode UA_VariantScalarValueUint8(UA_Variant *variant, UA_Byte value) {
  return UA_Variant_setScalarCopy(variant, &value, &UA_TYPES[UA_TYPES_BYTE]);
}

UA_StatusCode UA_VariantScalarValueInt16(UA_Variant *variant, UA_Int16 value) {
  return UA_Variant_setScalarCopy(variant, &value, &UA_TYPES[UA_TYPES_INT16]);
}

UA_StatusCode UA_VariantScalarValueUint16(UA_Variant *variant,
                                          UA_UInt16 value) {
  return UA_Variant_setScalarCopy(variant, &value, &UA_TYPES[UA_TYPES_UINT16]);
}

UA_StatusCode UA_VariantScalarValueInt32(UA_Variant *variant, UA_Int32 value) {
  return UA_Variant_setScalarCopy(variant, &value, &UA_TYPES[UA_TYPES_INT32]);
}

UA_StatusCode UA_VariantScalarValueUint32(UA_Variant *variant,
                                          UA_UInt32 value) {
  return UA_Variant_setScalarCopy(variant, &value, &UA_TYPES[UA_TYPES_UINT32]);
}

UA_StatusCode UA_VariantScalarValueInt64(UA_Variant *variant, UA_Int64 value) {
  return UA_Variant_setScalarCopy(variant, &value, &UA_TYPES[UA_TYPES_INT64]);
}

UA_StatusCode UA_VariantScalarValueUint64(UA_Variant *variant,
                                          UA_UInt64 value) {
  return UA_Variant_setScalarCopy(variant, &value, &UA_TYPES[UA_TYPES_UINT64]);
}

UA_StatusCode UA_VariantScalarValueFloat(UA_Variant *variant, UA_Float value) {
  return UA_Variant_setScalarCopy(variant, &value, &UA_TYPES[UA_TYPES_FLOAT]);
}

UA_StatusCode UA_VariantScalarValueDouble(UA_Variant *variant,
                                          UA_Double value) {
  return UA_Variant_setScalarCopy(variant, &value, &UA_TYPES[UA_TYPES_DOUBLE]);
}

UA_StatusCode UA_VariantScalarValueString(UA_Variant *variant, char *value) {
  UA_String stringValue = UA_STRING(value);
  return UA_Variant_setScalarCopy(variant, &stringValue,
                                  &UA_TYPES[UA_TYPES_STRING]);
}

UA_StatusCode UA_VariantScalarValueDateTime(UA_Variant *variant,
                                            UA_DateTime value) {
  return UA_Variant_setScalarCopy(variant, &value,
                                  &UA_TYPES[UA_TYPES_DATETIME]);
}

UA_StatusCode UA_VariantScalarValueByteString(UA_Variant *variant, void *value,
                                              size_t length) {
  UA_ByteString byteString;
  byteString.data = (UA_Byte *)value;
  byteString.length = length;
  return UA_Variant_setScalarCopy(variant, &byteString,
                                  &UA_TYPES[UA_TYPES_BYTESTRING]);
}

//
UA_StatusCode UA_ArrayValueInit(ArrayValue *value, uint32_t uaType) {
  value->length = 0;
  switch (uaType) {
  case UA_TYPES_BOOLEAN:
    value->data = UA_Array_new(0, &UA_TYPES[UA_TYPES_BOOLEAN]);
    break;
  case UA_TYPES_SBYTE:
    value->data = UA_Array_new(0, &UA_TYPES[UA_TYPES_SBYTE]);
    break;
  case UA_TYPES_BYTE:
    value->data = UA_Array_new(0, &UA_TYPES[UA_TYPES_BYTE]);
    break;
  case UA_TYPES_INT16:
    value->data = UA_Array_new(0, &UA_TYPES[UA_TYPES_INT16]);
    break;
  case UA_TYPES_UINT16:
    value->data = UA_Array_new(0, &UA_TYPES[UA_TYPES_UINT16]);
    break;
  case UA_TYPES_INT32:
    value->data = UA_Array_new(0, &UA_TYPES[UA_TYPES_INT32]);
    break;
  case UA_TYPES_UINT32:
    value->data = UA_Array_new(0, &UA_TYPES[UA_TYPES_UINT32]);
    break;
  case UA_TYPES_INT64:
    value->data = UA_Array_new(0, &UA_TYPES[UA_TYPES_INT64]);
    break;
  case UA_TYPES_UINT64:
    value->data = UA_Array_new(0, &UA_TYPES[UA_TYPES_UINT64]);
    break;
  case UA_TYPES_FLOAT:
    value->data = UA_Array_new(0, &UA_TYPES[UA_TYPES_FLOAT]);
    break;
  case UA_TYPES_DOUBLE:
    value->data = UA_Array_new(0, &UA_TYPES[UA_TYPES_DOUBLE]);
    break;
  case UA_TYPES_STRING:
    value->data = UA_Array_new(0, &UA_TYPES[UA_TYPES_STRING]);
    break;
  case UA_TYPES_DATETIME:
    value->data = UA_Array_new(0, &UA_TYPES[UA_TYPES_DATETIME]);
    break;
  case UA_TYPES_BYTESTRING:
    value->data = UA_Array_new(0, &UA_TYPES[UA_TYPES_BYTESTRING]);
    break;
  default:
    return UA_STATUSCODE_BADTYPEMISMATCH;
  }
  if (value->data == NULL) {
    return UA_STATUSCODE_BADOUTOFMEMORY;
  }
  return UA_STATUSCODE_GOOD;
}

UA_StatusCode UA_ArrayValueAppendBoolean(ArrayValue *array, UA_Boolean value) {
  return UA_Array_appendCopy(&array->data, &array->length, &value,
                             &UA_TYPES[UA_TYPES_BOOLEAN]);
}

UA_StatusCode UA_ArrayValueAppendInt8(ArrayValue *array, UA_SByte value) {
  return UA_Array_appendCopy(&array->data, &array->length, &value,
                             &UA_TYPES[UA_TYPES_SBYTE]);
}

UA_StatusCode UA_ArrayValueAppendUint8(ArrayValue *array, UA_Byte value) {
  return UA_Array_appendCopy(&array->data, &array->length, &value,
                             &UA_TYPES[UA_TYPES_BYTE]);
}

UA_StatusCode UA_ArrayValueAppendInt16(ArrayValue *array, UA_Int16 value) {
  return UA_Array_appendCopy(&array->data, &array->length, &value,
                             &UA_TYPES[UA_TYPES_INT16]);
}

UA_StatusCode UA_ArrayValueAppendUint16(ArrayValue *array, UA_UInt16 value) {
  return UA_Array_appendCopy(&array->data, &array->length, &value,
                             &UA_TYPES[UA_TYPES_UINT16]);
}

UA_StatusCode UA_ArrayValueAppendInt32(ArrayValue *array, UA_Int32 value) {
  return UA_Array_appendCopy(&array->data, &array->length, &value,
                             &UA_TYPES[UA_TYPES_INT32]);
}

UA_StatusCode UA_ArrayValueAppendUint32(ArrayValue *array, UA_UInt32 value) {
  return UA_Array_appendCopy(&array->data, &array->length, &value,
                             &UA_TYPES[UA_TYPES_UINT32]);
}

UA_StatusCode UA_ArrayValueAppendInt64(ArrayValue *array, UA_Int64 value) {
  return UA_Array_appendCopy(&array->data, &array->length, &value,
                             &UA_TYPES[UA_TYPES_INT64]);
}

UA_StatusCode UA_ArrayValueAppendUint64(ArrayValue *array, UA_UInt64 value) {
  return UA_Array_appendCopy(&array->data, &array->length, &value,
                             &UA_TYPES[UA_TYPES_UINT64]);
}

UA_StatusCode UA_ArrayValueAppendFloat(ArrayValue *array, UA_Float value) {
  return UA_Array_appendCopy(&array->data, &array->length, &value,
                             &UA_TYPES[UA_TYPES_FLOAT]);
}

UA_StatusCode UA_ArrayValueAppendDouble(ArrayValue *array, UA_Double value) {
  return UA_Array_appendCopy(&array->data, &array->length, &value,
                             &UA_TYPES[UA_TYPES_DOUBLE]);
}

UA_StatusCode UA_ArrayValueAppendString(ArrayValue *array, char *value) {
  UA_String stringValue = UA_STRING(value);
  return UA_Array_appendCopy(&array->data, &array->length, &stringValue,
                             &UA_TYPES[UA_TYPES_STRING]);
}

UA_StatusCode UA_ArrayValueAppendDateTime(ArrayValue *array,
                                          UA_DateTime value) {
  return UA_Array_appendCopy(&array->data, &array->length, &value,
                             &UA_TYPES[UA_TYPES_DATETIME]);
}

UA_StatusCode UA_ArrayValueAppendByteString(ArrayValue *array, void *value,
                                            size_t length) {
  UA_ByteString byteString;
  byteString.data = (UA_Byte *)value;
  byteString.length = length;

  return UA_Array_appendCopy(&array->data, &array->length, &byteString,
                             &UA_TYPES[UA_TYPES_BYTESTRING]);
}

void UA_VariantArrayValue(UA_Variant *variant, ArrayValue *value,
                          uint32_t uaType) {
  UA_Variant_setArray(variant, value->data, value->length, &UA_TYPES[uaType]);
}

//

NodeTree *ua_NodeTree_init(NodeTree *parent, uint32_t level, uint32_t index,
                           void *nodeID, size_t length) {
  NodeTree *node = (NodeTree *)malloc(sizeof(NodeTree));
  if (node == NULL) {
    return NULL;
  }
  memset(node, 0, sizeof(NodeTree));

  if (length) {
    node->nodeID = (char *)malloc(length + 1);
    if (node->nodeID == NULL) {
      free(node);
      return NULL;
    }
    memset(node->nodeID, '\0', length + 1);
    memcpy(node->nodeID, nodeID, length);
  }

  node->level = level;
  node->index = index;
  node->parent = parent;
  node->nodeLength = length;

  if (parent != NULL) {
    if (parent->head == NULL) {
      parent->head = node;
    } else {
      parent->tail->next = node;
    }
    parent->tail = node;
  }

  return node;
}

NodeTree *UA_NodeTree_root_init() {
  NodeTree *node = (NodeTree *)malloc(sizeof(NodeTree));
  if (node == NULL) {
    return NULL;
  }
  memset(node, 0, sizeof(NodeTree));
  return node;
}

void UA_NodeTree_clear(NodeTree *node) {
  NodeTree *cur = node->head;
  while (cur) {
    NodeTree *next = cur->next;
    UA_NodeTree_clear(cur);
    cur = next;
  }
  if (node->nodeID) {
    memset(node->nodeID, 0, strlen(node->nodeID));
    free(node->nodeID);
  }
  memset(node, 0, sizeof(NodeTree));
  free(node);
}

NodeTree *UA_NodeTree_next(NodeTree *node) { return node->next; }

NodeTree *UA_NodeTree_head(NodeTree *node) { return node->head; }

UA_StatusCode UA_Browse_nodeTreeLevel(UA_Client *client, UA_NodeId nodeId,
                                      NodeTree *parent, uint32_t level) {
  UA_BrowseRequest bReq;
  UA_BrowseRequest_init(&bReq);
  bReq.requestedMaxReferencesPerNode = 0;
  bReq.nodesToBrowse = UA_BrowseDescription_new();
  bReq.nodesToBrowseSize = 1;

  UA_NodeId_copy(&nodeId, &bReq.nodesToBrowse[0].nodeId);
  bReq.nodesToBrowse[0].resultMask = UA_BROWSERESULTMASK_ALL;

  UA_BrowseResponse bResp = UA_Client_Service_browse(client, bReq);
  if (bResp.responseHeader.serviceResult != UA_STATUSCODE_GOOD) {
    UA_BrowseResponse_clear(&bResp);
    return bResp.responseHeader.serviceResult;
  }

  for (int i = 0; i < bResp.resultsSize; i++) {
    for (int j = 0; j < bResp.results[i].referencesSize; j++) {
      NodeTree *node = NULL;

      UA_ReferenceDescription *ref = &(bResp.results[i].references[j]);
      if ((ref->nodeClass == UA_NODECLASS_OBJECT ||
           ref->nodeClass == UA_NODECLASS_VARIABLE ||
           ref->nodeClass == UA_NODECLASS_METHOD)) {
        if (ref->nodeId.nodeId.identifierType == UA_NODEIDTYPE_NUMERIC) {

          node = ua_NodeTree_init(
              parent, level, ref->nodeId.nodeId.namespaceIndex,
              ref->browseName.name.data, ref->browseName.name.length);
          if (node == NULL) {
            return UA_STATUSCODE_BADOUTOFMEMORY;
          }

          UA_StatusCode retval = UA_Browse_nodeTreeLevel(
              client,
              UA_NODEID_NUMERIC(ref->nodeId.nodeId.namespaceIndex,
                                ref->nodeId.nodeId.identifier.numeric),
              node, level + 1);

          if (retval != UA_STATUSCODE_GOOD) {
            return retval;
          }

        } else if (ref->nodeId.nodeId.identifierType == UA_NODEIDTYPE_STRING) {

          node =
              ua_NodeTree_init(parent, level, ref->nodeId.nodeId.namespaceIndex,
                               ref->nodeId.nodeId.identifier.string.data,
                               ref->nodeId.nodeId.identifier.string.length);
          if (node == NULL) {
            return UA_STATUSCODE_BADOUTOFMEMORY;
          }

          UA_StatusCode retval = UA_Browse_nodeTreeLevel(
              client,
              UA_NODEID_STRING(ref->nodeId.nodeId.namespaceIndex, node->nodeID),
              node, level + 1);
          if (retval != UA_STATUSCODE_GOOD) {
            return retval;
          }
        }
      }
    }
  }

  UA_BrowseResponse_clear(&bResp);

  return UA_STATUSCODE_GOOD;
}

UA_StatusCode UA_Browse_nodeTree(UA_Client *client, NodeTree *root) {
  return UA_Browse_nodeTreeLevel(
      client, UA_NODEID_NUMERIC(0, UA_NS0ID_OBJECTSFOLDER), root, 1);
}

UA_StatusCode UA_VariantValueWrite(UA_Client *client, uint32_t nsIndex,
                                   char *nodeID, UA_Variant *variant) {
  UA_WriteValue valueId;
  UA_WriteValue_init(&valueId);
  valueId.nodeId = UA_NODEID_STRING(nsIndex, nodeID);
  valueId.attributeId = UA_ATTRIBUTEID_VALUE;
  valueId.value.value = *variant;
  valueId.value.hasValue = true;

  UA_WriteRequest wReq;
  UA_WriteRequest_init(&wReq);
  wReq.nodesToWrite = &valueId;
  wReq.nodesToWriteSize = 1;

  UA_WriteResponse wResp = UA_Client_Service_write(client, wReq);
  UA_StatusCode retval = wResp.responseHeader.serviceResult;
  UA_WriteResponse_clear(&wResp);

  return retval;
}

UA_ReadValueId *UA_ReadValueID_alloc(int number) {
  UA_ReadValueId *readValueId =
      (UA_ReadValueId *)UA_malloc(sizeof(UA_ReadValueId) * number);
  if (readValueId == NULL) {
    return NULL;
  }
  for (size_t i = 0; i < number; i++) {
    UA_ReadValueId_init(&readValueId[i]);
  }
  return readValueId;
}

void UA_ReadValueID_free(UA_ReadValueId *readValueId) { UA_free(readValueId); }

void UA_ReadValueID_string(UA_ReadValueId *readValueId, int index,
                           UA_UInt16 nsIndex, char *chars,
                           UA_UInt32 attributeId) {
  readValueId[index].nodeId = UA_NODEID_STRING(nsIndex, chars);
  readValueId[index].attributeId = attributeId;
}

UA_Variant *UA_ReadResponse_variant(UA_ReadResponse *response, int index) {
  return &response->results[index].value;
}

void UA_Logger_init(UA_Logger *logger, void *context, void *log, void *clear) {
  logger->log = log;
  logger->context = context;
  logger->clear = clear;
}

void UA_LoggerWrapper(void *callback, UA_LogLevel level,
                      UA_LogCategory category, const char *format,
                      va_list args) {
  char buffer[1024] = {0};
  vsnprintf(buffer, sizeof(buffer) - 1, format, args);
  ((UA_Logger_Wrapper_t)callback)(level, category, buffer);
}

void UA_Logger_info(const char *format, ...) {
  char buffer[1024] = {0};

  va_list args;
  va_start(args, format);
  vsnprintf(buffer, sizeof(buffer) - 1, format, args);
  va_end(args);

  UA_Logger_golang(2, 10000, buffer);
}

void UA_Logger_error(const char *format, ...) {
  char buffer[1024] = {0};

  va_list args;
  va_start(args, format);
  vsnprintf(buffer, sizeof(buffer) - 1, format, args);
  va_end(args);

  UA_Logger_golang(4, 10000, buffer);
}

UA_StatusCode UA_ServerAddObject(UA_Server *server, UA_UInt16 index,
                                 char *name) {
  // 建立folder、object、type等之类的返回节点信息，方便后续使用
  UA_NodeId folderId;

  // 创建默认object节点
  UA_ObjectAttributes folderAttr = UA_ObjectAttributes_default;

  // 设置节点名字
  folderAttr.displayName = UA_LOCALIZEDTEXT("en-US", name);

  // UA_NodeId folderNodeid = UA_NODEID_NUMERIC(spaceIndex, 1);
  UA_NodeId folderNodeid = UA_NODEID_STRING(index, name);

  // 设置父节点，我放在了Objects下面
  UA_NodeId folderParNodeid = UA_NODEID_NUMERIC(0, UA_NS0ID_OBJECTSFOLDER);

  // 设置参考类型，其实就是与Objects的关系
  UA_NodeId folderParReferNodeid = UA_NODEID_NUMERIC(0, UA_NS0ID_ORGANIZES);

  // 设置由命名空间决定的浏览名称
  UA_QualifiedName folderBrowseName = UA_QUALIFIEDNAME(index, name);

  // 设置我们建立的节点的类型
  UA_NodeId folderType = UA_NODEID_NUMERIC(0, UA_NS0ID_FOLDERTYPE);

  // 往server中添加节点
  return UA_Server_addObjectNode(server, folderNodeid, folderParNodeid,
                                 folderParReferNodeid, folderBrowseName,
                                 folderType, folderAttr, NULL, &folderId);
}

UA_StatusCode UA_ServerAddVariable(UA_Server *server, UA_UInt16 parentNsIndex,
                                   char *parentNodeID, UA_UInt16 aNsIndex,
                                   char *aNodeID, char *displayName,
                                   UA_Variant *variant) {
  /*变量节点的属性*/
  UA_VariableAttributes attr = UA_VariableAttributes_default;
  memcpy(&attr.value, variant, sizeof(UA_Variant));

  // 节点在用户接口显示的名字(本地化)
  attr.description = UA_LOCALIZEDTEXT("en-US", displayName);
  // 节点在自身本地化描述
  attr.displayName = UA_LOCALIZEDTEXT("en-US", displayName);
  attr.dataType = variant->type->typeId;
  // 变量节点的访问权限：可读可写
  attr.accessLevel = UA_ACCESSLEVELMASK_READ | UA_ACCESSLEVELMASK_WRITE;

  /*定义1个节点：节点的命名空间 节点的ID标识符*/
  UA_NodeId integerNodeId = UA_NODEID_STRING(aNsIndex, aNodeID);

  // 节点对外的浏览名称(非本地化)
  UA_QualifiedName integerName = UA_QUALIFIEDNAME(aNsIndex, displayName);

  // 定义1个节点
  UA_NodeId parentNodeId = UA_NODEID_STRING(parentNsIndex, parentNodeID);

  // 定义1个节点
  UA_NodeId parentReferenceNodeId = UA_NODEID_NUMERIC(0, UA_NS0ID_ORGANIZES);

  /*************************************************************
  UA_Server_addVariableNode(UA_Server *server,  // 服务
  const UA_NodeId requestedNewNodeId, //请求添加的节点
  const UA_NodeId parentNodeId,      //父节点
  const UA_NodeId referenceTypeId,   //父节点引用类型节点
  const UA_QualifiedName browseName, //节点对外的浏览名称(非本地化)
  const UA_NodeId typeDefinition,	 //引用节点
  const UA_VariableAttributes attr,
  void *nodeContext, UA_NodeId *outNewNodeId)
  **************************************************************/
  UA_NodeId outNodeID;
  return UA_Server_addVariableNode(
      server, integerNodeId, parentNodeId, parentReferenceNodeId, integerName,
      UA_NODEID_NUMERIC(0, UA_NS0ID_BASEDATAVARIABLETYPE), attr, NULL,
      &outNodeID);
}