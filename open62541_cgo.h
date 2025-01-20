#ifndef OPEN62541_CGO_H_
#define OPEN62541_CGO_H_

#include "open62541.h"
#include <stdio.h>
#include <stdlib.h>

typedef struct nodeTree {
  uint32_t level;
  uint32_t index;
  char *nodeID;
  uint32_t nodeLength;

  struct nodeTree *parent;
  struct nodeTree *next;

  struct nodeTree *head;
  struct nodeTree *tail;
} NodeTree;

extern UA_StatusCode UA_VariantType(UA_Variant *variant, UA_UInt32 *ptype);

extern UA_Boolean UA_VariantValueBoolean(UA_Variant *value, int index);

extern UA_SByte UA_VariantValueInt8(UA_Variant *value, int index);

extern UA_Byte UA_VariantValueUint8(UA_Variant *value, int index);

extern UA_Int16 UA_VariantValueInt16(UA_Variant *value, int index);

extern UA_UInt16 UA_VariantValueUint16(UA_Variant *value, int index);

extern UA_Int32 UA_VariantValueInt32(UA_Variant *value, int index);

extern UA_UInt32 UA_VariantValueUint32(UA_Variant *value, int index);

extern UA_Int64 UA_VariantValueInt64(UA_Variant *value, int index);

extern UA_UInt64 UA_VariantValueUint64(UA_Variant *value, int index);

extern UA_Float UA_VariantValueFloat(UA_Variant *value, int index);

extern UA_Double UA_VariantValueDouble(UA_Variant *value, int index);

extern UA_String UA_VariantValueString(UA_Variant *value, int index);

extern UA_DateTime UA_VariantValueDateTime(UA_Variant *value, int index);

extern UA_ByteString UA_VariantValueByteString(UA_Variant *value, int index);
//

extern UA_StatusCode UA_VariantScalarValueBoolean(UA_Variant *variant,
                                                  UA_Boolean value);

extern UA_StatusCode UA_VariantScalarValueInt8(UA_Variant *variant,
                                               UA_SByte value);

extern UA_StatusCode UA_VariantScalarValueUint8(UA_Variant *variant,
                                                UA_Byte value);

extern UA_StatusCode UA_VariantScalarValueInt16(UA_Variant *variant,
                                                UA_Int16 value);

extern UA_StatusCode UA_VariantScalarValueUint16(UA_Variant *variant,
                                                 UA_UInt16 value);

extern UA_StatusCode UA_VariantScalarValueInt32(UA_Variant *variant,
                                                UA_Int32 value);

extern UA_StatusCode UA_VariantScalarValueUint32(UA_Variant *variant,
                                                 UA_UInt32 value);

extern UA_StatusCode UA_VariantScalarValueInt64(UA_Variant *variant,
                                                UA_Int64 value);

extern UA_StatusCode UA_VariantScalarValueUint64(UA_Variant *variant,
                                                 UA_UInt64 value);

extern UA_StatusCode UA_VariantScalarValueFloat(UA_Variant *variant,
                                                UA_Float value);

extern UA_StatusCode UA_VariantScalarValueDouble(UA_Variant *variant,
                                                 UA_Double value);

extern UA_StatusCode UA_VariantScalarValueString(UA_Variant *variant,
                                                 char *value);

extern UA_StatusCode UA_VariantScalarValueDateTime(UA_Variant *variant,
                                                   UA_DateTime value);

extern UA_StatusCode UA_VariantScalarValueByteString(UA_Variant *variant,
                                                     void *value,
                                                     size_t length);

//
typedef struct arrayValue {
  size_t length;
  void *data;
} ArrayValue;

extern UA_StatusCode UA_ArrayValueInit(ArrayValue *value, uint32_t uaType);

extern UA_StatusCode UA_ArrayValueAppendBoolean(ArrayValue *array,
                                                UA_Boolean value);

extern UA_StatusCode UA_ArrayValueAppendInt8(ArrayValue *array, UA_SByte value);

extern UA_StatusCode UA_ArrayValueAppendUint8(ArrayValue *array, UA_Byte value);

extern UA_StatusCode UA_ArrayValueAppendInt16(ArrayValue *array,
                                              UA_Int16 value);

extern UA_StatusCode UA_ArrayValueAppendUint16(ArrayValue *array,
                                               UA_UInt16 value);

extern UA_StatusCode UA_ArrayValueAppendInt32(ArrayValue *array,
                                              UA_Int32 value);

extern UA_StatusCode UA_ArrayValueAppendUint32(ArrayValue *array,
                                               UA_UInt32 value);

extern UA_StatusCode UA_ArrayValueAppendInt64(ArrayValue *array,
                                              UA_Int64 value);

extern UA_StatusCode UA_ArrayValueAppendUint64(ArrayValue *array,
                                               UA_UInt64 value);

extern UA_StatusCode UA_ArrayValueAppendFloat(ArrayValue *array,
                                              UA_Float value);

extern UA_StatusCode UA_ArrayValueAppendDouble(ArrayValue *array,
                                               UA_Double value);

extern UA_StatusCode UA_ArrayValueAppendString(ArrayValue *array, char *value);

extern UA_StatusCode UA_ArrayValueAppendDateTime(ArrayValue *array,
                                                 UA_DateTime value);

extern UA_StatusCode UA_ArrayValueAppendByteString(ArrayValue *array,
                                                   void *value, size_t length);

extern void UA_VariantArrayValue(UA_Variant *variant, ArrayValue *value,
                                 uint32_t uaType);

//
extern UA_StatusCode UA_Browse_nodeTree(UA_Client *client, NodeTree *root);

extern UA_StatusCode UA_VariantValueWrite(UA_Client *client, uint32_t nsIndex,
                                          char *nodeID, UA_Variant *variant);

// node tree init and view functions
extern NodeTree *UA_NodeTree_root_init(void);

extern void UA_NodeTree_clear(NodeTree *nodeTree);

extern NodeTree *UA_NodeTree_next(NodeTree *nodeTree);

extern NodeTree *UA_NodeTree_head(NodeTree *nodeTree);

extern UA_ReadValueId *UA_ReadValueID_alloc(int number);

extern void UA_ReadValueID_free(UA_ReadValueId *readValueId);

extern void UA_ReadValueID_string(UA_ReadValueId *readValueId, int index,
                                  UA_UInt16 nsIndex, char *chars,
                                  UA_UInt32 attributeId);

extern UA_Variant *UA_ReadResponse_variant(UA_ReadResponse *response,
                                           int index);

// logger wrapper functions
typedef void (*UA_Logger_Wrapper_t)(uint32_t level, uint32_t category,
                                    char *msg);

typedef void (*UA_Logger_t)(void *logContext, UA_LogLevel level,
                            UA_LogCategory category, const char *msg,
                            va_list args);

extern void UA_LoggerWrapper(void *callback, UA_LogLevel level,
                             UA_LogCategory category, const char *format,
                             va_list args);

extern void UA_Logger_init(UA_Logger *logger, void *context, void *log,
                           void *clear);

extern void UA_Logger_info(const char *format, ...);

extern void UA_Logger_error(const char *format, ...);

extern void UA_Logger_golang(uint32_t level, uint32_t category, char *msg);

#define UA_LOGGER_INFO(format, ...)                                            \
  UA_Logger_info("[%s:%d] " format, __FILE__, __LINE__, ##__VA_ARGS__)

#define UA_LOGGER_ERROR(format, ...)                                           \
  UA_Logger_error("[%s:%d] " format, __FILE__, __LINE__, ##__VA_ARGS__)

// UA_server add object functions

extern UA_StatusCode UA_ServerAddObject(UA_Server *server, UA_UInt16 index,
                                        char *name);

extern UA_StatusCode
UA_ServerAddVariable(UA_Server *server, UA_UInt16 parentNsIndex,
                     char *parentNodeID, UA_UInt16 aNsIndex, char *aNodeID,
                     char *displayName, UA_Variant *variant);

#endif
