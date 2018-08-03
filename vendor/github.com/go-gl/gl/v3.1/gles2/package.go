// Copyright (c) 2010 Khronos Group.
// This material may be distributed subject to the terms and conditions
// set forth in the Open Publication License, v 1.0, 8 June 1999.
// http://opencontent.org/openpub/.
//
// Copyright (c) 1991-2006 Silicon Graphics, Inc.
// This document is licensed under the SGI Free Software B License.
// For details, see http://oss.sgi.com/projects/FreeB.

// Package gles2 implements Go bindings to OpenGL.
//
// This package was automatically generated using Glow:
//  http://github.com/go-gl/glow
//
package gles2

// #cgo darwin        LDFLAGS: -framework OpenGL
// #cgo linux freebsd LDFLAGS: -lGL
// #cgo windows       LDFLAGS: -lopengl32
// #if defined(_WIN32) && !defined(APIENTRY) && !defined(__CYGWIN__) && !defined(__SCITECH_SNAP__)
// #ifndef WIN32_LEAN_AND_MEAN
// #define WIN32_LEAN_AND_MEAN 1
// #endif
// #include <windows.h>
// #endif
// #ifndef APIENTRY
// #define APIENTRY
// #endif
// #ifndef APIENTRYP
// #define APIENTRYP APIENTRY *
// #endif
// #ifndef GLAPI
// #define GLAPI extern
// #endif
// #include <KHR/khrplatform.h>
// typedef unsigned int GLenum;
// typedef unsigned char GLboolean;
// typedef unsigned int GLbitfield;
// typedef int GLint;
// typedef unsigned int GLuint;
// typedef int GLsizei;
// typedef double GLdouble;
// typedef void *GLeglClientBufferEXT;
// typedef void *GLeglImageOES;
// typedef char GLchar;
// typedef uintptr_t GLsync;
// typedef void (APIENTRY *GLDEBUGPROC)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,const void *userParam);
// typedef void (APIENTRY *GLDEBUGPROCKHR)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,const void *userParam);
// typedef khronos_uint8_t GLubyte;
// typedef khronos_float_t GLfloat;
// typedef khronos_float_t GLclampf;
// typedef khronos_int64_t GLint64;
// typedef khronos_uint64_t GLuint64;
// typedef khronos_int64_t GLint64EXT;
// typedef khronos_uint64_t GLuint64EXT;
// typedef khronos_intptr_t GLintptr;
// typedef khronos_ssize_t GLsizeiptr;
// typedef void (APIENTRY *GLVULKANPROCNV)(void);
// extern void glowDebugCallback_gles231(GLenum source, GLenum type, GLuint id, GLenum severity, GLsizei length, const GLchar* message, const void* userParam);
// static void APIENTRY glowCDebugCallback(GLenum source, GLenum type, GLuint id, GLenum severity, GLsizei length, const GLchar* message, const void* userParam) {
//   glowDebugCallback_gles231(source, type, id, severity, length, message, userParam);
// }
// typedef GLboolean  (APIENTRYP GPACQUIREKEYEDMUTEXWIN32EXT)(GLuint  memory, GLuint64  key, GLuint  timeout);
// typedef void  (APIENTRYP GPACTIVEPROGRAMEXT)(GLuint  program);
// typedef void  (APIENTRYP GPACTIVESHADERPROGRAM)(GLuint  pipeline, GLuint  program);
// typedef void  (APIENTRYP GPACTIVESHADERPROGRAMEXT)(GLuint  pipeline, GLuint  program);
// typedef void  (APIENTRYP GPACTIVETEXTURE)(GLenum  texture);
// typedef void  (APIENTRYP GPALPHAFUNCQCOM)(GLenum  func, GLclampf  ref);
// typedef void  (APIENTRYP GPAPPLYFRAMEBUFFERATTACHMENTCMAAINTEL)();
// typedef void  (APIENTRYP GPATTACHSHADER)(GLuint  program, GLuint  shader);
// typedef void  (APIENTRYP GPBEGINCONDITIONALRENDERNV)(GLuint  id, GLenum  mode);
// typedef void  (APIENTRYP GPBEGINPERFMONITORAMD)(GLuint  monitor);
// typedef void  (APIENTRYP GPBEGINPERFQUERYINTEL)(GLuint  queryHandle);
// typedef void  (APIENTRYP GPBEGINQUERY)(GLenum  target, GLuint  id);
// typedef void  (APIENTRYP GPBEGINQUERYEXT)(GLenum  target, GLuint  id);
// typedef void  (APIENTRYP GPBEGINTRANSFORMFEEDBACK)(GLenum  primitiveMode);
// typedef void  (APIENTRYP GPBINDATTRIBLOCATION)(GLuint  program, GLuint  index, const GLchar * name);
// typedef void  (APIENTRYP GPBINDBUFFER)(GLenum  target, GLuint  buffer);
// typedef void  (APIENTRYP GPBINDBUFFERBASE)(GLenum  target, GLuint  index, GLuint  buffer);
// typedef void  (APIENTRYP GPBINDBUFFERRANGE)(GLenum  target, GLuint  index, GLuint  buffer, GLintptr  offset, GLsizeiptr  size);
// typedef void  (APIENTRYP GPBINDFRAGDATALOCATIONEXT)(GLuint  program, GLuint  color, const GLchar * name);
// typedef void  (APIENTRYP GPBINDFRAGDATALOCATIONINDEXEDEXT)(GLuint  program, GLuint  colorNumber, GLuint  index, const GLchar * name);
// typedef void  (APIENTRYP GPBINDFRAMEBUFFER)(GLenum  target, GLuint  framebuffer);
// typedef void  (APIENTRYP GPBINDIMAGETEXTURE)(GLuint  unit, GLuint  texture, GLint  level, GLboolean  layered, GLint  layer, GLenum  access, GLenum  format);
// typedef void  (APIENTRYP GPBINDPROGRAMPIPELINE)(GLuint  pipeline);
// typedef void  (APIENTRYP GPBINDPROGRAMPIPELINEEXT)(GLuint  pipeline);
// typedef void  (APIENTRYP GPBINDRENDERBUFFER)(GLenum  target, GLuint  renderbuffer);
// typedef void  (APIENTRYP GPBINDSAMPLER)(GLuint  unit, GLuint  sampler);
// typedef void  (APIENTRYP GPBINDTEXTURE)(GLenum  target, GLuint  texture);
// typedef void  (APIENTRYP GPBINDTRANSFORMFEEDBACK)(GLenum  target, GLuint  id);
// typedef void  (APIENTRYP GPBINDVERTEXARRAY)(GLuint  array);
// typedef void  (APIENTRYP GPBINDVERTEXARRAYOES)(GLuint  array);
// typedef void  (APIENTRYP GPBINDVERTEXBUFFER)(GLuint  bindingindex, GLuint  buffer, GLintptr  offset, GLsizei  stride);
// typedef void  (APIENTRYP GPBLENDBARRIERKHR)();
// typedef void  (APIENTRYP GPBLENDBARRIERNV)();
// typedef void  (APIENTRYP GPBLENDCOLOR)(GLfloat  red, GLfloat  green, GLfloat  blue, GLfloat  alpha);
// typedef void  (APIENTRYP GPBLENDEQUATION)(GLenum  mode);
// typedef void  (APIENTRYP GPBLENDEQUATIONEXT)(GLenum  mode);
// typedef void  (APIENTRYP GPBLENDEQUATIONSEPARATE)(GLenum  modeRGB, GLenum  modeAlpha);
// typedef void  (APIENTRYP GPBLENDEQUATIONSEPARATEIEXT)(GLuint  buf, GLenum  modeRGB, GLenum  modeAlpha);
// typedef void  (APIENTRYP GPBLENDEQUATIONSEPARATEIOES)(GLuint  buf, GLenum  modeRGB, GLenum  modeAlpha);
// typedef void  (APIENTRYP GPBLENDEQUATIONIEXT)(GLuint  buf, GLenum  mode);
// typedef void  (APIENTRYP GPBLENDEQUATIONIOES)(GLuint  buf, GLenum  mode);
// typedef void  (APIENTRYP GPBLENDFUNC)(GLenum  sfactor, GLenum  dfactor);
// typedef void  (APIENTRYP GPBLENDFUNCSEPARATE)(GLenum  sfactorRGB, GLenum  dfactorRGB, GLenum  sfactorAlpha, GLenum  dfactorAlpha);
// typedef void  (APIENTRYP GPBLENDFUNCSEPARATEIEXT)(GLuint  buf, GLenum  srcRGB, GLenum  dstRGB, GLenum  srcAlpha, GLenum  dstAlpha);
// typedef void  (APIENTRYP GPBLENDFUNCSEPARATEIOES)(GLuint  buf, GLenum  srcRGB, GLenum  dstRGB, GLenum  srcAlpha, GLenum  dstAlpha);
// typedef void  (APIENTRYP GPBLENDFUNCIEXT)(GLuint  buf, GLenum  src, GLenum  dst);
// typedef void  (APIENTRYP GPBLENDFUNCIOES)(GLuint  buf, GLenum  src, GLenum  dst);
// typedef void  (APIENTRYP GPBLENDPARAMETERINV)(GLenum  pname, GLint  value);
// typedef void  (APIENTRYP GPBLITFRAMEBUFFER)(GLint  srcX0, GLint  srcY0, GLint  srcX1, GLint  srcY1, GLint  dstX0, GLint  dstY0, GLint  dstX1, GLint  dstY1, GLbitfield  mask, GLenum  filter);
// typedef void  (APIENTRYP GPBLITFRAMEBUFFERANGLE)(GLint  srcX0, GLint  srcY0, GLint  srcX1, GLint  srcY1, GLint  dstX0, GLint  dstY0, GLint  dstX1, GLint  dstY1, GLbitfield  mask, GLenum  filter);
// typedef void  (APIENTRYP GPBLITFRAMEBUFFERNV)(GLint  srcX0, GLint  srcY0, GLint  srcX1, GLint  srcY1, GLint  dstX0, GLint  dstY0, GLint  dstX1, GLint  dstY1, GLbitfield  mask, GLenum  filter);
// typedef void  (APIENTRYP GPBUFFERDATA)(GLenum  target, GLsizeiptr  size, const void * data, GLenum  usage);
// typedef void  (APIENTRYP GPBUFFERSTORAGEEXT)(GLenum  target, GLsizeiptr  size, const void * data, GLbitfield  flags);
// typedef void  (APIENTRYP GPBUFFERSTORAGEEXTERNALEXT)(GLenum  target, GLintptr  offset, GLsizeiptr  size, GLeglClientBufferEXT  clientBuffer, GLbitfield  flags);
// typedef void  (APIENTRYP GPBUFFERSTORAGEMEMEXT)(GLenum  target, GLsizeiptr  size, GLuint  memory, GLuint64  offset);
// typedef void  (APIENTRYP GPBUFFERSUBDATA)(GLenum  target, GLintptr  offset, GLsizeiptr  size, const void * data);
// typedef GLenum  (APIENTRYP GPCHECKFRAMEBUFFERSTATUS)(GLenum  target);
// typedef void  (APIENTRYP GPCLEAR)(GLbitfield  mask);
// typedef void  (APIENTRYP GPCLEARBUFFERFI)(GLenum  buffer, GLint  drawbuffer, GLfloat  depth, GLint  stencil);
// typedef void  (APIENTRYP GPCLEARBUFFERFV)(GLenum  buffer, GLint  drawbuffer, const GLfloat * value);
// typedef void  (APIENTRYP GPCLEARBUFFERIV)(GLenum  buffer, GLint  drawbuffer, const GLint * value);
// typedef void  (APIENTRYP GPCLEARBUFFERUIV)(GLenum  buffer, GLint  drawbuffer, const GLuint * value);
// typedef void  (APIENTRYP GPCLEARCOLOR)(GLfloat  red, GLfloat  green, GLfloat  blue, GLfloat  alpha);
// typedef void  (APIENTRYP GPCLEARDEPTHF)(GLfloat  d);
// typedef void  (APIENTRYP GPCLEARPIXELLOCALSTORAGEUIEXT)(GLsizei  offset, GLsizei  n, const GLuint * values);
// typedef void  (APIENTRYP GPCLEARSTENCIL)(GLint  s);
// typedef void  (APIENTRYP GPCLEARTEXIMAGEEXT)(GLuint  texture, GLint  level, GLenum  format, GLenum  type, const void * data);
// typedef void  (APIENTRYP GPCLEARTEXSUBIMAGEEXT)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * data);
// typedef GLenum  (APIENTRYP GPCLIENTWAITSYNC)(GLsync  sync, GLbitfield  flags, GLuint64  timeout);
// typedef GLenum  (APIENTRYP GPCLIENTWAITSYNCAPPLE)(GLsync  sync, GLbitfield  flags, GLuint64  timeout);
// typedef void  (APIENTRYP GPCLIPCONTROLEXT)(GLenum  origin, GLenum  depth);
// typedef void  (APIENTRYP GPCOLORMASK)(GLboolean  red, GLboolean  green, GLboolean  blue, GLboolean  alpha);
// typedef void  (APIENTRYP GPCOLORMASKIEXT)(GLuint  index, GLboolean  r, GLboolean  g, GLboolean  b, GLboolean  a);
// typedef void  (APIENTRYP GPCOLORMASKIOES)(GLuint  index, GLboolean  r, GLboolean  g, GLboolean  b, GLboolean  a);
// typedef void  (APIENTRYP GPCOMPILESHADER)(GLuint  shader);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXIMAGE2D)(GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLsizei  height, GLint  border, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXIMAGE3D)(GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLint  border, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXIMAGE3DOES)(GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLint  border, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXSUBIMAGE2D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXSUBIMAGE3D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXSUBIMAGE3DOES)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCONSERVATIVERASTERPARAMETERINV)(GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPCOPYBUFFERSUBDATA)(GLenum  readTarget, GLenum  writeTarget, GLintptr  readOffset, GLintptr  writeOffset, GLsizeiptr  size);
// typedef void  (APIENTRYP GPCOPYBUFFERSUBDATANV)(GLenum  readTarget, GLenum  writeTarget, GLintptr  readOffset, GLintptr  writeOffset, GLsizeiptr  size);
// typedef void  (APIENTRYP GPCOPYIMAGESUBDATAEXT)(GLuint  srcName, GLenum  srcTarget, GLint  srcLevel, GLint  srcX, GLint  srcY, GLint  srcZ, GLuint  dstName, GLenum  dstTarget, GLint  dstLevel, GLint  dstX, GLint  dstY, GLint  dstZ, GLsizei  srcWidth, GLsizei  srcHeight, GLsizei  srcDepth);
// typedef void  (APIENTRYP GPCOPYIMAGESUBDATAOES)(GLuint  srcName, GLenum  srcTarget, GLint  srcLevel, GLint  srcX, GLint  srcY, GLint  srcZ, GLuint  dstName, GLenum  dstTarget, GLint  dstLevel, GLint  dstX, GLint  dstY, GLint  dstZ, GLsizei  srcWidth, GLsizei  srcHeight, GLsizei  srcDepth);
// typedef void  (APIENTRYP GPCOPYPATHNV)(GLuint  resultPath, GLuint  srcPath);
// typedef void  (APIENTRYP GPCOPYTEXIMAGE2D)(GLenum  target, GLint  level, GLenum  internalformat, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLint  border);
// typedef void  (APIENTRYP GPCOPYTEXSUBIMAGE2D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPCOPYTEXSUBIMAGE3D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPCOPYTEXSUBIMAGE3DOES)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPCOPYTEXTURELEVELSAPPLE)(GLuint  destinationTexture, GLuint  sourceTexture, GLint  sourceBaseLevel, GLsizei  sourceLevelCount);
// typedef void  (APIENTRYP GPCOVERFILLPATHINSTANCEDNV)(GLsizei  numPaths, GLenum  pathNameType, const void * paths, GLuint  pathBase, GLenum  coverMode, GLenum  transformType, const GLfloat * transformValues);
// typedef void  (APIENTRYP GPCOVERFILLPATHNV)(GLuint  path, GLenum  coverMode);
// typedef void  (APIENTRYP GPCOVERSTROKEPATHINSTANCEDNV)(GLsizei  numPaths, GLenum  pathNameType, const void * paths, GLuint  pathBase, GLenum  coverMode, GLenum  transformType, const GLfloat * transformValues);
// typedef void  (APIENTRYP GPCOVERSTROKEPATHNV)(GLuint  path, GLenum  coverMode);
// typedef void  (APIENTRYP GPCOVERAGEMASKNV)(GLboolean  mask);
// typedef void  (APIENTRYP GPCOVERAGEMODULATIONNV)(GLenum  components);
// typedef void  (APIENTRYP GPCOVERAGEMODULATIONTABLENV)(GLsizei  n, const GLfloat * v);
// typedef void  (APIENTRYP GPCOVERAGEOPERATIONNV)(GLenum  operation);
// typedef void  (APIENTRYP GPCREATEMEMORYOBJECTSEXT)(GLsizei  n, GLuint * memoryObjects);
// typedef void  (APIENTRYP GPCREATEPERFQUERYINTEL)(GLuint  queryId, GLuint * queryHandle);
// typedef GLuint  (APIENTRYP GPCREATEPROGRAM)();
// typedef GLuint  (APIENTRYP GPCREATESHADER)(GLenum  type);
// typedef GLuint  (APIENTRYP GPCREATESHADERPROGRAMEXT)(GLenum  type, const GLchar * string);
// typedef GLuint  (APIENTRYP GPCREATESHADERPROGRAMV)(GLenum  type, GLsizei  count, const GLchar *const* strings);
// typedef GLuint  (APIENTRYP GPCREATESHADERPROGRAMVEXT)(GLenum  type, GLsizei  count, const GLchar ** strings);
// typedef void  (APIENTRYP GPCULLFACE)(GLenum  mode);
// typedef void  (APIENTRYP GPDEBUGMESSAGECALLBACK)(GLDEBUGPROC  callback, const void * userParam);
// typedef void  (APIENTRYP GPDEBUGMESSAGECALLBACKKHR)(GLDEBUGPROCKHR  callback, const void * userParam);
// typedef void  (APIENTRYP GPDEBUGMESSAGECONTROL)(GLenum  source, GLenum  type, GLenum  severity, GLsizei  count, const GLuint * ids, GLboolean  enabled);
// typedef void  (APIENTRYP GPDEBUGMESSAGECONTROLKHR)(GLenum  source, GLenum  type, GLenum  severity, GLsizei  count, const GLuint * ids, GLboolean  enabled);
// typedef void  (APIENTRYP GPDEBUGMESSAGEINSERT)(GLenum  source, GLenum  type, GLuint  id, GLenum  severity, GLsizei  length, const GLchar * buf);
// typedef void  (APIENTRYP GPDEBUGMESSAGEINSERTKHR)(GLenum  source, GLenum  type, GLuint  id, GLenum  severity, GLsizei  length, const GLchar * buf);
// typedef void  (APIENTRYP GPDELETEBUFFERS)(GLsizei  n, const GLuint * buffers);
// typedef void  (APIENTRYP GPDELETEFENCESNV)(GLsizei  n, const GLuint * fences);
// typedef void  (APIENTRYP GPDELETEFRAMEBUFFERS)(GLsizei  n, const GLuint * framebuffers);
// typedef void  (APIENTRYP GPDELETEMEMORYOBJECTSEXT)(GLsizei  n, const GLuint * memoryObjects);
// typedef void  (APIENTRYP GPDELETEPATHSNV)(GLuint  path, GLsizei  range);
// typedef void  (APIENTRYP GPDELETEPERFMONITORSAMD)(GLsizei  n, GLuint * monitors);
// typedef void  (APIENTRYP GPDELETEPERFQUERYINTEL)(GLuint  queryHandle);
// typedef void  (APIENTRYP GPDELETEPROGRAM)(GLuint  program);
// typedef void  (APIENTRYP GPDELETEPROGRAMPIPELINES)(GLsizei  n, const GLuint * pipelines);
// typedef void  (APIENTRYP GPDELETEPROGRAMPIPELINESEXT)(GLsizei  n, const GLuint * pipelines);
// typedef void  (APIENTRYP GPDELETEQUERIES)(GLsizei  n, const GLuint * ids);
// typedef void  (APIENTRYP GPDELETEQUERIESEXT)(GLsizei  n, const GLuint * ids);
// typedef void  (APIENTRYP GPDELETERENDERBUFFERS)(GLsizei  n, const GLuint * renderbuffers);
// typedef void  (APIENTRYP GPDELETESAMPLERS)(GLsizei  count, const GLuint * samplers);
// typedef void  (APIENTRYP GPDELETESEMAPHORESEXT)(GLsizei  n, const GLuint * semaphores);
// typedef void  (APIENTRYP GPDELETESHADER)(GLuint  shader);
// typedef void  (APIENTRYP GPDELETESYNC)(GLsync  sync);
// typedef void  (APIENTRYP GPDELETESYNCAPPLE)(GLsync  sync);
// typedef void  (APIENTRYP GPDELETETEXTURES)(GLsizei  n, const GLuint * textures);
// typedef void  (APIENTRYP GPDELETETRANSFORMFEEDBACKS)(GLsizei  n, const GLuint * ids);
// typedef void  (APIENTRYP GPDELETEVERTEXARRAYS)(GLsizei  n, const GLuint * arrays);
// typedef void  (APIENTRYP GPDELETEVERTEXARRAYSOES)(GLsizei  n, const GLuint * arrays);
// typedef void  (APIENTRYP GPDEPTHFUNC)(GLenum  func);
// typedef void  (APIENTRYP GPDEPTHMASK)(GLboolean  flag);
// typedef void  (APIENTRYP GPDEPTHRANGEARRAYFVNV)(GLuint  first, GLsizei  count, const GLfloat * v);
// typedef void  (APIENTRYP GPDEPTHRANGEARRAYFVOES)(GLuint  first, GLsizei  count, const GLfloat * v);
// typedef void  (APIENTRYP GPDEPTHRANGEINDEXEDFNV)(GLuint  index, GLfloat  n, GLfloat  f);
// typedef void  (APIENTRYP GPDEPTHRANGEINDEXEDFOES)(GLuint  index, GLfloat  n, GLfloat  f);
// typedef void  (APIENTRYP GPDEPTHRANGEF)(GLfloat  n, GLfloat  f);
// typedef void  (APIENTRYP GPDETACHSHADER)(GLuint  program, GLuint  shader);
// typedef void  (APIENTRYP GPDISABLE)(GLenum  cap);
// typedef void  (APIENTRYP GPDISABLEDRIVERCONTROLQCOM)(GLuint  driverControl);
// typedef void  (APIENTRYP GPDISABLEVERTEXATTRIBARRAY)(GLuint  index);
// typedef void  (APIENTRYP GPDISABLEIEXT)(GLenum  target, GLuint  index);
// typedef void  (APIENTRYP GPDISABLEINV)(GLenum  target, GLuint  index);
// typedef void  (APIENTRYP GPDISABLEIOES)(GLenum  target, GLuint  index);
// typedef void  (APIENTRYP GPDISCARDFRAMEBUFFEREXT)(GLenum  target, GLsizei  numAttachments, const GLenum * attachments);
// typedef void  (APIENTRYP GPDISPATCHCOMPUTE)(GLuint  num_groups_x, GLuint  num_groups_y, GLuint  num_groups_z);
// typedef void  (APIENTRYP GPDISPATCHCOMPUTEINDIRECT)(GLintptr  indirect);
// typedef void  (APIENTRYP GPDRAWARRAYS)(GLenum  mode, GLint  first, GLsizei  count);
// typedef void  (APIENTRYP GPDRAWARRAYSINDIRECT)(GLenum  mode, const void * indirect);
// typedef void  (APIENTRYP GPDRAWARRAYSINSTANCED)(GLenum  mode, GLint  first, GLsizei  count, GLsizei  instancecount);
// typedef void  (APIENTRYP GPDRAWARRAYSINSTANCEDANGLE)(GLenum  mode, GLint  first, GLsizei  count, GLsizei  primcount);
// typedef void  (APIENTRYP GPDRAWARRAYSINSTANCEDBASEINSTANCEEXT)(GLenum  mode, GLint  first, GLsizei  count, GLsizei  instancecount, GLuint  baseinstance);
// typedef void  (APIENTRYP GPDRAWARRAYSINSTANCEDEXT)(GLenum  mode, GLint  start, GLsizei  count, GLsizei  primcount);
// typedef void  (APIENTRYP GPDRAWARRAYSINSTANCEDNV)(GLenum  mode, GLint  first, GLsizei  count, GLsizei  primcount);
// typedef void  (APIENTRYP GPDRAWBUFFERS)(GLsizei  n, const GLenum * bufs);
// typedef void  (APIENTRYP GPDRAWBUFFERSEXT)(GLsizei  n, const GLenum * bufs);
// typedef void  (APIENTRYP GPDRAWBUFFERSINDEXEDEXT)(GLint  n, const GLenum * location, const GLint * indices);
// typedef void  (APIENTRYP GPDRAWBUFFERSNV)(GLsizei  n, const GLenum * bufs);
// typedef void  (APIENTRYP GPDRAWELEMENTS)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices);
// typedef void  (APIENTRYP GPDRAWELEMENTSBASEVERTEXEXT)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLint  basevertex);
// typedef void  (APIENTRYP GPDRAWELEMENTSBASEVERTEXOES)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLint  basevertex);
// typedef void  (APIENTRYP GPDRAWELEMENTSINDIRECT)(GLenum  mode, GLenum  type, const void * indirect);
// typedef void  (APIENTRYP GPDRAWELEMENTSINSTANCED)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount);
// typedef void  (APIENTRYP GPDRAWELEMENTSINSTANCEDANGLE)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  primcount);
// typedef void  (APIENTRYP GPDRAWELEMENTSINSTANCEDBASEINSTANCEEXT)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLuint  baseinstance);
// typedef void  (APIENTRYP GPDRAWELEMENTSINSTANCEDBASEVERTEXBASEINSTANCEEXT)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLint  basevertex, GLuint  baseinstance);
// typedef void  (APIENTRYP GPDRAWELEMENTSINSTANCEDBASEVERTEXEXT)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLint  basevertex);
// typedef void  (APIENTRYP GPDRAWELEMENTSINSTANCEDBASEVERTEXOES)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLint  basevertex);
// typedef void  (APIENTRYP GPDRAWELEMENTSINSTANCEDEXT)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  primcount);
// typedef void  (APIENTRYP GPDRAWELEMENTSINSTANCEDNV)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  primcount);
// typedef void  (APIENTRYP GPDRAWRANGEELEMENTS)(GLenum  mode, GLuint  start, GLuint  end, GLsizei  count, GLenum  type, const void * indices);
// typedef void  (APIENTRYP GPDRAWRANGEELEMENTSBASEVERTEXEXT)(GLenum  mode, GLuint  start, GLuint  end, GLsizei  count, GLenum  type, const void * indices, GLint  basevertex);
// typedef void  (APIENTRYP GPDRAWRANGEELEMENTSBASEVERTEXOES)(GLenum  mode, GLuint  start, GLuint  end, GLsizei  count, GLenum  type, const void * indices, GLint  basevertex);
// typedef void  (APIENTRYP GPDRAWTRANSFORMFEEDBACKEXT)(GLenum  mode, GLuint  id);
// typedef void  (APIENTRYP GPDRAWTRANSFORMFEEDBACKINSTANCEDEXT)(GLenum  mode, GLuint  id, GLsizei  instancecount);
// typedef void  (APIENTRYP GPDRAWVKIMAGENV)(GLuint64  vkImage, GLuint  sampler, GLfloat  x0, GLfloat  y0, GLfloat  x1, GLfloat  y1, GLfloat  z, GLfloat  s0, GLfloat  t0, GLfloat  s1, GLfloat  t1);
// typedef void  (APIENTRYP GPEGLIMAGETARGETRENDERBUFFERSTORAGEOES)(GLenum  target, GLeglImageOES  image);
// typedef void  (APIENTRYP GPEGLIMAGETARGETTEXSTORAGEEXT)(GLenum  target, GLeglImageOES  image, const GLint*  attrib_list);
// typedef void  (APIENTRYP GPEGLIMAGETARGETTEXTURE2DOES)(GLenum  target, GLeglImageOES  image);
// typedef void  (APIENTRYP GPEGLIMAGETARGETTEXTURESTORAGEEXT)(GLuint  texture, GLeglImageOES  image, const GLint*  attrib_list);
// typedef void  (APIENTRYP GPENABLE)(GLenum  cap);
// typedef void  (APIENTRYP GPENABLEDRIVERCONTROLQCOM)(GLuint  driverControl);
// typedef void  (APIENTRYP GPENABLEVERTEXATTRIBARRAY)(GLuint  index);
// typedef void  (APIENTRYP GPENABLEIEXT)(GLenum  target, GLuint  index);
// typedef void  (APIENTRYP GPENABLEINV)(GLenum  target, GLuint  index);
// typedef void  (APIENTRYP GPENABLEIOES)(GLenum  target, GLuint  index);
// typedef void  (APIENTRYP GPENDCONDITIONALRENDERNV)();
// typedef void  (APIENTRYP GPENDPERFMONITORAMD)(GLuint  monitor);
// typedef void  (APIENTRYP GPENDPERFQUERYINTEL)(GLuint  queryHandle);
// typedef void  (APIENTRYP GPENDQUERY)(GLenum  target);
// typedef void  (APIENTRYP GPENDQUERYEXT)(GLenum  target);
// typedef void  (APIENTRYP GPENDTILINGQCOM)(GLbitfield  preserveMask);
// typedef void  (APIENTRYP GPENDTRANSFORMFEEDBACK)();
// typedef void  (APIENTRYP GPEXTGETBUFFERPOINTERVQCOM)(GLenum  target, void ** params);
// typedef void  (APIENTRYP GPEXTGETBUFFERSQCOM)(GLuint * buffers, GLint  maxBuffers, GLint * numBuffers);
// typedef void  (APIENTRYP GPEXTGETFRAMEBUFFERSQCOM)(GLuint * framebuffers, GLint  maxFramebuffers, GLint * numFramebuffers);
// typedef void  (APIENTRYP GPEXTGETPROGRAMBINARYSOURCEQCOM)(GLuint  program, GLenum  shadertype, GLchar * source, GLint * length);
// typedef void  (APIENTRYP GPEXTGETPROGRAMSQCOM)(GLuint * programs, GLint  maxPrograms, GLint * numPrograms);
// typedef void  (APIENTRYP GPEXTGETRENDERBUFFERSQCOM)(GLuint * renderbuffers, GLint  maxRenderbuffers, GLint * numRenderbuffers);
// typedef void  (APIENTRYP GPEXTGETSHADERSQCOM)(GLuint * shaders, GLint  maxShaders, GLint * numShaders);
// typedef void  (APIENTRYP GPEXTGETTEXLEVELPARAMETERIVQCOM)(GLuint  texture, GLenum  face, GLint  level, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPEXTGETTEXSUBIMAGEQCOM)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, void * texels);
// typedef void  (APIENTRYP GPEXTGETTEXTURESQCOM)(GLuint * textures, GLint  maxTextures, GLint * numTextures);
// typedef GLboolean  (APIENTRYP GPEXTISPROGRAMBINARYQCOM)(GLuint  program);
// typedef void  (APIENTRYP GPEXTTEXOBJECTSTATEOVERRIDEIQCOM)(GLenum  target, GLenum  pname, GLint  param);
// typedef GLsync  (APIENTRYP GPFENCESYNC)(GLenum  condition, GLbitfield  flags);
// typedef GLsync  (APIENTRYP GPFENCESYNCAPPLE)(GLenum  condition, GLbitfield  flags);
// typedef void  (APIENTRYP GPFINISH)();
// typedef void  (APIENTRYP GPFINISHFENCENV)(GLuint  fence);
// typedef void  (APIENTRYP GPFLUSH)();
// typedef void  (APIENTRYP GPFLUSHMAPPEDBUFFERRANGE)(GLenum  target, GLintptr  offset, GLsizeiptr  length);
// typedef void  (APIENTRYP GPFLUSHMAPPEDBUFFERRANGEEXT)(GLenum  target, GLintptr  offset, GLsizeiptr  length);
// typedef void  (APIENTRYP GPFRAGMENTCOVERAGECOLORNV)(GLuint  color);
// typedef void  (APIENTRYP GPFRAMEBUFFERFETCHBARRIEREXT)();
// typedef void  (APIENTRYP GPFRAMEBUFFERFETCHBARRIERQCOM)();
// typedef void  (APIENTRYP GPFRAMEBUFFERFOVEATIONCONFIGQCOM)(GLuint  framebuffer, GLuint  numLayers, GLuint  focalPointsPerLayer, GLuint  requestedFeatures, GLuint * providedFeatures);
// typedef void  (APIENTRYP GPFRAMEBUFFERFOVEATIONPARAMETERSQCOM)(GLuint  framebuffer, GLuint  layer, GLuint  focalPoint, GLfloat  focalX, GLfloat  focalY, GLfloat  gainX, GLfloat  gainY, GLfloat  foveaArea);
// typedef void  (APIENTRYP GPFRAMEBUFFERPARAMETERI)(GLenum  target, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPFRAMEBUFFERPIXELLOCALSTORAGESIZEEXT)(GLuint  target, GLsizei  size);
// typedef void  (APIENTRYP GPFRAMEBUFFERRENDERBUFFER)(GLenum  target, GLenum  attachment, GLenum  renderbuffertarget, GLuint  renderbuffer);
// typedef void  (APIENTRYP GPFRAMEBUFFERSAMPLELOCATIONSFVNV)(GLenum  target, GLuint  start, GLsizei  count, const GLfloat * v);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURE2D)(GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURE2DDOWNSAMPLEIMG)(GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level, GLint  xscale, GLint  yscale);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURE2DMULTISAMPLEEXT)(GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level, GLsizei  samples);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURE2DMULTISAMPLEIMG)(GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level, GLsizei  samples);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURE3DOES)(GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level, GLint  zoffset);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTUREEXT)(GLenum  target, GLenum  attachment, GLuint  texture, GLint  level);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURELAYER)(GLenum  target, GLenum  attachment, GLuint  texture, GLint  level, GLint  layer);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURELAYERDOWNSAMPLEIMG)(GLenum  target, GLenum  attachment, GLuint  texture, GLint  level, GLint  layer, GLint  xscale, GLint  yscale);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTUREMULTISAMPLEMULTIVIEWOVR)(GLenum  target, GLenum  attachment, GLuint  texture, GLint  level, GLsizei  samples, GLint  baseViewIndex, GLsizei  numViews);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTUREMULTIVIEWOVR)(GLenum  target, GLenum  attachment, GLuint  texture, GLint  level, GLint  baseViewIndex, GLsizei  numViews);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTUREOES)(GLenum  target, GLenum  attachment, GLuint  texture, GLint  level);
// typedef void  (APIENTRYP GPFRONTFACE)(GLenum  mode);
// typedef void  (APIENTRYP GPGENBUFFERS)(GLsizei  n, GLuint * buffers);
// typedef void  (APIENTRYP GPGENFENCESNV)(GLsizei  n, GLuint * fences);
// typedef void  (APIENTRYP GPGENFRAMEBUFFERS)(GLsizei  n, GLuint * framebuffers);
// typedef GLuint  (APIENTRYP GPGENPATHSNV)(GLsizei  range);
// typedef void  (APIENTRYP GPGENPERFMONITORSAMD)(GLsizei  n, GLuint * monitors);
// typedef void  (APIENTRYP GPGENPROGRAMPIPELINES)(GLsizei  n, GLuint * pipelines);
// typedef void  (APIENTRYP GPGENPROGRAMPIPELINESEXT)(GLsizei  n, GLuint * pipelines);
// typedef void  (APIENTRYP GPGENQUERIES)(GLsizei  n, GLuint * ids);
// typedef void  (APIENTRYP GPGENQUERIESEXT)(GLsizei  n, GLuint * ids);
// typedef void  (APIENTRYP GPGENRENDERBUFFERS)(GLsizei  n, GLuint * renderbuffers);
// typedef void  (APIENTRYP GPGENSAMPLERS)(GLsizei  count, GLuint * samplers);
// typedef void  (APIENTRYP GPGENSEMAPHORESEXT)(GLsizei  n, GLuint * semaphores);
// typedef void  (APIENTRYP GPGENTEXTURES)(GLsizei  n, GLuint * textures);
// typedef void  (APIENTRYP GPGENTRANSFORMFEEDBACKS)(GLsizei  n, GLuint * ids);
// typedef void  (APIENTRYP GPGENVERTEXARRAYS)(GLsizei  n, GLuint * arrays);
// typedef void  (APIENTRYP GPGENVERTEXARRAYSOES)(GLsizei  n, GLuint * arrays);
// typedef void  (APIENTRYP GPGENERATEMIPMAP)(GLenum  target);
// typedef void  (APIENTRYP GPGETACTIVEATTRIB)(GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLint * size, GLenum * type, GLchar * name);
// typedef void  (APIENTRYP GPGETACTIVEUNIFORM)(GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLint * size, GLenum * type, GLchar * name);
// typedef void  (APIENTRYP GPGETACTIVEUNIFORMBLOCKNAME)(GLuint  program, GLuint  uniformBlockIndex, GLsizei  bufSize, GLsizei * length, GLchar * uniformBlockName);
// typedef void  (APIENTRYP GPGETACTIVEUNIFORMBLOCKIV)(GLuint  program, GLuint  uniformBlockIndex, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETACTIVEUNIFORMSIV)(GLuint  program, GLsizei  uniformCount, const GLuint * uniformIndices, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETATTACHEDSHADERS)(GLuint  program, GLsizei  maxCount, GLsizei * count, GLuint * shaders);
// typedef GLint  (APIENTRYP GPGETATTRIBLOCATION)(GLuint  program, const GLchar * name);
// typedef void  (APIENTRYP GPGETBOOLEANI_V)(GLenum  target, GLuint  index, GLboolean * data);
// typedef void  (APIENTRYP GPGETBOOLEANV)(GLenum  pname, GLboolean * data);
// typedef void  (APIENTRYP GPGETBUFFERPARAMETERI64V)(GLenum  target, GLenum  pname, GLint64 * params);
// typedef void  (APIENTRYP GPGETBUFFERPARAMETERIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETBUFFERPOINTERV)(GLenum  target, GLenum  pname, void ** params);
// typedef void  (APIENTRYP GPGETBUFFERPOINTERVOES)(GLenum  target, GLenum  pname, void ** params);
// typedef void  (APIENTRYP GPGETCOVERAGEMODULATIONTABLENV)(GLsizei  bufsize, GLfloat * v);
// typedef GLuint  (APIENTRYP GPGETDEBUGMESSAGELOG)(GLuint  count, GLsizei  bufSize, GLenum * sources, GLenum * types, GLuint * ids, GLenum * severities, GLsizei * lengths, GLchar * messageLog);
// typedef GLuint  (APIENTRYP GPGETDEBUGMESSAGELOGKHR)(GLuint  count, GLsizei  bufSize, GLenum * sources, GLenum * types, GLuint * ids, GLenum * severities, GLsizei * lengths, GLchar * messageLog);
// typedef void  (APIENTRYP GPGETDRIVERCONTROLSTRINGQCOM)(GLuint  driverControl, GLsizei  bufSize, GLsizei * length, GLchar * driverControlString);
// typedef void  (APIENTRYP GPGETDRIVERCONTROLSQCOM)(GLint * num, GLsizei  size, GLuint * driverControls);
// typedef GLenum  (APIENTRYP GPGETERROR)();
// typedef void  (APIENTRYP GPGETFENCEIVNV)(GLuint  fence, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETFIRSTPERFQUERYIDINTEL)(GLuint * queryId);
// typedef void  (APIENTRYP GPGETFLOATI_VNV)(GLenum  target, GLuint  index, GLfloat * data);
// typedef void  (APIENTRYP GPGETFLOATI_VOES)(GLenum  target, GLuint  index, GLfloat * data);
// typedef void  (APIENTRYP GPGETFLOATV)(GLenum  pname, GLfloat * data);
// typedef GLint  (APIENTRYP GPGETFRAGDATAINDEXEXT)(GLuint  program, const GLchar * name);
// typedef GLint  (APIENTRYP GPGETFRAGDATALOCATION)(GLuint  program, const GLchar * name);
// typedef void  (APIENTRYP GPGETFRAMEBUFFERATTACHMENTPARAMETERIV)(GLenum  target, GLenum  attachment, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETFRAMEBUFFERPARAMETERIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef GLsizei  (APIENTRYP GPGETFRAMEBUFFERPIXELLOCALSTORAGESIZEEXT)(GLuint  target);
// typedef GLenum  (APIENTRYP GPGETGRAPHICSRESETSTATUS)();
// typedef GLenum  (APIENTRYP GPGETGRAPHICSRESETSTATUSEXT)();
// typedef GLenum  (APIENTRYP GPGETGRAPHICSRESETSTATUSKHR)();
// typedef GLuint64  (APIENTRYP GPGETIMAGEHANDLENV)(GLuint  texture, GLint  level, GLboolean  layered, GLint  layer, GLenum  format);
// typedef void  (APIENTRYP GPGETINTEGER64I_V)(GLenum  target, GLuint  index, GLint64 * data);
// typedef void  (APIENTRYP GPGETINTEGER64V)(GLenum  pname, GLint64 * data);
// typedef void  (APIENTRYP GPGETINTEGER64VAPPLE)(GLenum  pname, GLint64 * params);
// typedef void  (APIENTRYP GPGETINTEGERI_V)(GLenum  target, GLuint  index, GLint * data);
// typedef void  (APIENTRYP GPGETINTEGERI_VEXT)(GLenum  target, GLuint  index, GLint * data);
// typedef void  (APIENTRYP GPGETINTEGERV)(GLenum  pname, GLint * data);
// typedef void  (APIENTRYP GPGETINTERNALFORMATSAMPLEIVNV)(GLenum  target, GLenum  internalformat, GLsizei  samples, GLenum  pname, GLsizei  bufSize, GLint * params);
// typedef void  (APIENTRYP GPGETINTERNALFORMATIV)(GLenum  target, GLenum  internalformat, GLenum  pname, GLsizei  bufSize, GLint * params);
// typedef void  (APIENTRYP GPGETMEMORYOBJECTPARAMETERIVEXT)(GLuint  memoryObject, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETMULTISAMPLEFV)(GLenum  pname, GLuint  index, GLfloat * val);
// typedef void  (APIENTRYP GPGETNEXTPERFQUERYIDINTEL)(GLuint  queryId, GLuint * nextQueryId);
// typedef void  (APIENTRYP GPGETOBJECTLABEL)(GLenum  identifier, GLuint  name, GLsizei  bufSize, GLsizei * length, GLchar * label);
// typedef void  (APIENTRYP GPGETOBJECTLABELEXT)(GLenum  type, GLuint  object, GLsizei  bufSize, GLsizei * length, GLchar * label);
// typedef void  (APIENTRYP GPGETOBJECTLABELKHR)(GLenum  identifier, GLuint  name, GLsizei  bufSize, GLsizei * length, GLchar * label);
// typedef void  (APIENTRYP GPGETOBJECTPTRLABEL)(const void * ptr, GLsizei  bufSize, GLsizei * length, GLchar * label);
// typedef void  (APIENTRYP GPGETOBJECTPTRLABELKHR)(const void * ptr, GLsizei  bufSize, GLsizei * length, GLchar * label);
// typedef void  (APIENTRYP GPGETPATHCOMMANDSNV)(GLuint  path, GLubyte * commands);
// typedef void  (APIENTRYP GPGETPATHCOORDSNV)(GLuint  path, GLfloat * coords);
// typedef void  (APIENTRYP GPGETPATHDASHARRAYNV)(GLuint  path, GLfloat * dashArray);
// typedef GLfloat  (APIENTRYP GPGETPATHLENGTHNV)(GLuint  path, GLsizei  startSegment, GLsizei  numSegments);
// typedef void  (APIENTRYP GPGETPATHMETRICRANGENV)(GLbitfield  metricQueryMask, GLuint  firstPathName, GLsizei  numPaths, GLsizei  stride, GLfloat * metrics);
// typedef void  (APIENTRYP GPGETPATHMETRICSNV)(GLbitfield  metricQueryMask, GLsizei  numPaths, GLenum  pathNameType, const void * paths, GLuint  pathBase, GLsizei  stride, GLfloat * metrics);
// typedef void  (APIENTRYP GPGETPATHPARAMETERFVNV)(GLuint  path, GLenum  pname, GLfloat * value);
// typedef void  (APIENTRYP GPGETPATHPARAMETERIVNV)(GLuint  path, GLenum  pname, GLint * value);
// typedef void  (APIENTRYP GPGETPATHSPACINGNV)(GLenum  pathListMode, GLsizei  numPaths, GLenum  pathNameType, const void * paths, GLuint  pathBase, GLfloat  advanceScale, GLfloat  kerningScale, GLenum  transformType, GLfloat * returnedSpacing);
// typedef void  (APIENTRYP GPGETPERFCOUNTERINFOINTEL)(GLuint  queryId, GLuint  counterId, GLuint  counterNameLength, GLchar * counterName, GLuint  counterDescLength, GLchar * counterDesc, GLuint * counterOffset, GLuint * counterDataSize, GLuint * counterTypeEnum, GLuint * counterDataTypeEnum, GLuint64 * rawCounterMaxValue);
// typedef void  (APIENTRYP GPGETPERFMONITORCOUNTERDATAAMD)(GLuint  monitor, GLenum  pname, GLsizei  dataSize, GLuint * data, GLint * bytesWritten);
// typedef void  (APIENTRYP GPGETPERFMONITORCOUNTERINFOAMD)(GLuint  group, GLuint  counter, GLenum  pname, void * data);
// typedef void  (APIENTRYP GPGETPERFMONITORCOUNTERSTRINGAMD)(GLuint  group, GLuint  counter, GLsizei  bufSize, GLsizei * length, GLchar * counterString);
// typedef void  (APIENTRYP GPGETPERFMONITORCOUNTERSAMD)(GLuint  group, GLint * numCounters, GLint * maxActiveCounters, GLsizei  counterSize, GLuint * counters);
// typedef void  (APIENTRYP GPGETPERFMONITORGROUPSTRINGAMD)(GLuint  group, GLsizei  bufSize, GLsizei * length, GLchar * groupString);
// typedef void  (APIENTRYP GPGETPERFMONITORGROUPSAMD)(GLint * numGroups, GLsizei  groupsSize, GLuint * groups);
// typedef void  (APIENTRYP GPGETPERFQUERYDATAINTEL)(GLuint  queryHandle, GLuint  flags, GLsizei  dataSize, void * data, GLuint * bytesWritten);
// typedef void  (APIENTRYP GPGETPERFQUERYIDBYNAMEINTEL)(GLchar * queryName, GLuint * queryId);
// typedef void  (APIENTRYP GPGETPERFQUERYINFOINTEL)(GLuint  queryId, GLuint  queryNameLength, GLchar * queryName, GLuint * dataSize, GLuint * noCounters, GLuint * noInstances, GLuint * capsMask);
// typedef void  (APIENTRYP GPGETPOINTERV)(GLenum  pname, void ** params);
// typedef void  (APIENTRYP GPGETPOINTERVKHR)(GLenum  pname, void ** params);
// typedef void  (APIENTRYP GPGETPROGRAMBINARY)(GLuint  program, GLsizei  bufSize, GLsizei * length, GLenum * binaryFormat, void * binary);
// typedef void  (APIENTRYP GPGETPROGRAMBINARYOES)(GLuint  program, GLsizei  bufSize, GLsizei * length, GLenum * binaryFormat, void * binary);
// typedef void  (APIENTRYP GPGETPROGRAMINFOLOG)(GLuint  program, GLsizei  bufSize, GLsizei * length, GLchar * infoLog);
// typedef void  (APIENTRYP GPGETPROGRAMINTERFACEIV)(GLuint  program, GLenum  programInterface, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETPROGRAMPIPELINEINFOLOG)(GLuint  pipeline, GLsizei  bufSize, GLsizei * length, GLchar * infoLog);
// typedef void  (APIENTRYP GPGETPROGRAMPIPELINEINFOLOGEXT)(GLuint  pipeline, GLsizei  bufSize, GLsizei * length, GLchar * infoLog);
// typedef void  (APIENTRYP GPGETPROGRAMPIPELINEIV)(GLuint  pipeline, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETPROGRAMPIPELINEIVEXT)(GLuint  pipeline, GLenum  pname, GLint * params);
// typedef GLuint  (APIENTRYP GPGETPROGRAMRESOURCEINDEX)(GLuint  program, GLenum  programInterface, const GLchar * name);
// typedef GLint  (APIENTRYP GPGETPROGRAMRESOURCELOCATION)(GLuint  program, GLenum  programInterface, const GLchar * name);
// typedef GLint  (APIENTRYP GPGETPROGRAMRESOURCELOCATIONINDEXEXT)(GLuint  program, GLenum  programInterface, const GLchar * name);
// typedef void  (APIENTRYP GPGETPROGRAMRESOURCENAME)(GLuint  program, GLenum  programInterface, GLuint  index, GLsizei  bufSize, GLsizei * length, GLchar * name);
// typedef void  (APIENTRYP GPGETPROGRAMRESOURCEFVNV)(GLuint  program, GLenum  programInterface, GLuint  index, GLsizei  propCount, const GLenum * props, GLsizei  bufSize, GLsizei * length, GLfloat * params);
// typedef void  (APIENTRYP GPGETPROGRAMRESOURCEIV)(GLuint  program, GLenum  programInterface, GLuint  index, GLsizei  propCount, const GLenum * props, GLsizei  bufSize, GLsizei * length, GLint * params);
// typedef void  (APIENTRYP GPGETPROGRAMIV)(GLuint  program, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETQUERYOBJECTI64VEXT)(GLuint  id, GLenum  pname, GLint64 * params);
// typedef void  (APIENTRYP GPGETQUERYOBJECTIVEXT)(GLuint  id, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETQUERYOBJECTUI64VEXT)(GLuint  id, GLenum  pname, GLuint64 * params);
// typedef void  (APIENTRYP GPGETQUERYOBJECTUIV)(GLuint  id, GLenum  pname, GLuint * params);
// typedef void  (APIENTRYP GPGETQUERYOBJECTUIVEXT)(GLuint  id, GLenum  pname, GLuint * params);
// typedef void  (APIENTRYP GPGETQUERYIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETQUERYIVEXT)(GLenum  target, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETRENDERBUFFERPARAMETERIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETSAMPLERPARAMETERIIVEXT)(GLuint  sampler, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETSAMPLERPARAMETERIIVOES)(GLuint  sampler, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETSAMPLERPARAMETERIUIVEXT)(GLuint  sampler, GLenum  pname, GLuint * params);
// typedef void  (APIENTRYP GPGETSAMPLERPARAMETERIUIVOES)(GLuint  sampler, GLenum  pname, GLuint * params);
// typedef void  (APIENTRYP GPGETSAMPLERPARAMETERFV)(GLuint  sampler, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETSAMPLERPARAMETERIV)(GLuint  sampler, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETSEMAPHOREPARAMETERUI64VEXT)(GLuint  semaphore, GLenum  pname, GLuint64 * params);
// typedef void  (APIENTRYP GPGETSHADERINFOLOG)(GLuint  shader, GLsizei  bufSize, GLsizei * length, GLchar * infoLog);
// typedef void  (APIENTRYP GPGETSHADERPRECISIONFORMAT)(GLenum  shadertype, GLenum  precisiontype, GLint * range, GLint * precision);
// typedef void  (APIENTRYP GPGETSHADERSOURCE)(GLuint  shader, GLsizei  bufSize, GLsizei * length, GLchar * source);
// typedef void  (APIENTRYP GPGETSHADERIV)(GLuint  shader, GLenum  pname, GLint * params);
// typedef const GLubyte * (APIENTRYP GPGETSTRING)(GLenum  name);
// typedef const GLubyte * (APIENTRYP GPGETSTRINGI)(GLenum  name, GLuint  index);
// typedef void  (APIENTRYP GPGETSYNCIV)(GLsync  sync, GLenum  pname, GLsizei  bufSize, GLsizei * length, GLint * values);
// typedef void  (APIENTRYP GPGETSYNCIVAPPLE)(GLsync  sync, GLenum  pname, GLsizei  bufSize, GLsizei * length, GLint * values);
// typedef void  (APIENTRYP GPGETTEXLEVELPARAMETERFV)(GLenum  target, GLint  level, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETTEXLEVELPARAMETERIV)(GLenum  target, GLint  level, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETTEXPARAMETERIIVEXT)(GLenum  target, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETTEXPARAMETERIIVOES)(GLenum  target, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETTEXPARAMETERIUIVEXT)(GLenum  target, GLenum  pname, GLuint * params);
// typedef void  (APIENTRYP GPGETTEXPARAMETERIUIVOES)(GLenum  target, GLenum  pname, GLuint * params);
// typedef void  (APIENTRYP GPGETTEXPARAMETERFV)(GLenum  target, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETTEXPARAMETERIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef GLuint64  (APIENTRYP GPGETTEXTUREHANDLEIMG)(GLuint  texture);
// typedef GLuint64  (APIENTRYP GPGETTEXTUREHANDLENV)(GLuint  texture);
// typedef GLuint64  (APIENTRYP GPGETTEXTURESAMPLERHANDLEIMG)(GLuint  texture, GLuint  sampler);
// typedef GLuint64  (APIENTRYP GPGETTEXTURESAMPLERHANDLENV)(GLuint  texture, GLuint  sampler);
// typedef void  (APIENTRYP GPGETTRANSFORMFEEDBACKVARYING)(GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLsizei * size, GLenum * type, GLchar * name);
// typedef void  (APIENTRYP GPGETTRANSLATEDSHADERSOURCEANGLE)(GLuint  shader, GLsizei  bufsize, GLsizei * length, GLchar * source);
// typedef GLuint  (APIENTRYP GPGETUNIFORMBLOCKINDEX)(GLuint  program, const GLchar * uniformBlockName);
// typedef void  (APIENTRYP GPGETUNIFORMINDICES)(GLuint  program, GLsizei  uniformCount, const GLchar *const* uniformNames, GLuint * uniformIndices);
// typedef GLint  (APIENTRYP GPGETUNIFORMLOCATION)(GLuint  program, const GLchar * name);
// typedef void  (APIENTRYP GPGETUNIFORMFV)(GLuint  program, GLint  location, GLfloat * params);
// typedef void  (APIENTRYP GPGETUNIFORMI64VNV)(GLuint  program, GLint  location, GLint64EXT * params);
// typedef void  (APIENTRYP GPGETUNIFORMIV)(GLuint  program, GLint  location, GLint * params);
// typedef void  (APIENTRYP GPGETUNIFORMUIV)(GLuint  program, GLint  location, GLuint * params);
// typedef void  (APIENTRYP GPGETUNSIGNEDBYTEI_VEXT)(GLenum  target, GLuint  index, GLubyte * data);
// typedef void  (APIENTRYP GPGETUNSIGNEDBYTEVEXT)(GLenum  pname, GLubyte * data);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBIIV)(GLuint  index, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBIUIV)(GLuint  index, GLenum  pname, GLuint * params);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBPOINTERV)(GLuint  index, GLenum  pname, void ** pointer);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBFV)(GLuint  index, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBIV)(GLuint  index, GLenum  pname, GLint * params);
// typedef GLVULKANPROCNV  (APIENTRYP GPGETVKPROCADDRNV)(const GLchar * name);
// typedef void  (APIENTRYP GPGETNUNIFORMFV)(GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params);
// typedef void  (APIENTRYP GPGETNUNIFORMFVEXT)(GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params);
// typedef void  (APIENTRYP GPGETNUNIFORMFVKHR)(GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params);
// typedef void  (APIENTRYP GPGETNUNIFORMIV)(GLuint  program, GLint  location, GLsizei  bufSize, GLint * params);
// typedef void  (APIENTRYP GPGETNUNIFORMIVEXT)(GLuint  program, GLint  location, GLsizei  bufSize, GLint * params);
// typedef void  (APIENTRYP GPGETNUNIFORMIVKHR)(GLuint  program, GLint  location, GLsizei  bufSize, GLint * params);
// typedef void  (APIENTRYP GPGETNUNIFORMUIV)(GLuint  program, GLint  location, GLsizei  bufSize, GLuint * params);
// typedef void  (APIENTRYP GPGETNUNIFORMUIVKHR)(GLuint  program, GLint  location, GLsizei  bufSize, GLuint * params);
// typedef void  (APIENTRYP GPHINT)(GLenum  target, GLenum  mode);
// typedef void  (APIENTRYP GPIMPORTMEMORYFDEXT)(GLuint  memory, GLuint64  size, GLenum  handleType, GLint  fd);
// typedef void  (APIENTRYP GPIMPORTMEMORYWIN32HANDLEEXT)(GLuint  memory, GLuint64  size, GLenum  handleType, void * handle);
// typedef void  (APIENTRYP GPIMPORTMEMORYWIN32NAMEEXT)(GLuint  memory, GLuint64  size, GLenum  handleType, const void * name);
// typedef void  (APIENTRYP GPIMPORTSEMAPHOREFDEXT)(GLuint  semaphore, GLenum  handleType, GLint  fd);
// typedef void  (APIENTRYP GPIMPORTSEMAPHOREWIN32HANDLEEXT)(GLuint  semaphore, GLenum  handleType, void * handle);
// typedef void  (APIENTRYP GPIMPORTSEMAPHOREWIN32NAMEEXT)(GLuint  semaphore, GLenum  handleType, const void * name);
// typedef void  (APIENTRYP GPINSERTEVENTMARKEREXT)(GLsizei  length, const GLchar * marker);
// typedef void  (APIENTRYP GPINTERPOLATEPATHSNV)(GLuint  resultPath, GLuint  pathA, GLuint  pathB, GLfloat  weight);
// typedef void  (APIENTRYP GPINVALIDATEFRAMEBUFFER)(GLenum  target, GLsizei  numAttachments, const GLenum * attachments);
// typedef void  (APIENTRYP GPINVALIDATESUBFRAMEBUFFER)(GLenum  target, GLsizei  numAttachments, const GLenum * attachments, GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef GLboolean  (APIENTRYP GPISBUFFER)(GLuint  buffer);
// typedef GLboolean  (APIENTRYP GPISENABLED)(GLenum  cap);
// typedef GLboolean  (APIENTRYP GPISENABLEDIEXT)(GLenum  target, GLuint  index);
// typedef GLboolean  (APIENTRYP GPISENABLEDINV)(GLenum  target, GLuint  index);
// typedef GLboolean  (APIENTRYP GPISENABLEDIOES)(GLenum  target, GLuint  index);
// typedef GLboolean  (APIENTRYP GPISFENCENV)(GLuint  fence);
// typedef GLboolean  (APIENTRYP GPISFRAMEBUFFER)(GLuint  framebuffer);
// typedef GLboolean  (APIENTRYP GPISIMAGEHANDLERESIDENTNV)(GLuint64  handle);
// typedef GLboolean  (APIENTRYP GPISMEMORYOBJECTEXT)(GLuint  memoryObject);
// typedef GLboolean  (APIENTRYP GPISPATHNV)(GLuint  path);
// typedef GLboolean  (APIENTRYP GPISPOINTINFILLPATHNV)(GLuint  path, GLuint  mask, GLfloat  x, GLfloat  y);
// typedef GLboolean  (APIENTRYP GPISPOINTINSTROKEPATHNV)(GLuint  path, GLfloat  x, GLfloat  y);
// typedef GLboolean  (APIENTRYP GPISPROGRAM)(GLuint  program);
// typedef GLboolean  (APIENTRYP GPISPROGRAMPIPELINE)(GLuint  pipeline);
// typedef GLboolean  (APIENTRYP GPISPROGRAMPIPELINEEXT)(GLuint  pipeline);
// typedef GLboolean  (APIENTRYP GPISQUERY)(GLuint  id);
// typedef GLboolean  (APIENTRYP GPISQUERYEXT)(GLuint  id);
// typedef GLboolean  (APIENTRYP GPISRENDERBUFFER)(GLuint  renderbuffer);
// typedef GLboolean  (APIENTRYP GPISSAMPLER)(GLuint  sampler);
// typedef GLboolean  (APIENTRYP GPISSEMAPHOREEXT)(GLuint  semaphore);
// typedef GLboolean  (APIENTRYP GPISSHADER)(GLuint  shader);
// typedef GLboolean  (APIENTRYP GPISSYNC)(GLsync  sync);
// typedef GLboolean  (APIENTRYP GPISSYNCAPPLE)(GLsync  sync);
// typedef GLboolean  (APIENTRYP GPISTEXTURE)(GLuint  texture);
// typedef GLboolean  (APIENTRYP GPISTEXTUREHANDLERESIDENTNV)(GLuint64  handle);
// typedef GLboolean  (APIENTRYP GPISTRANSFORMFEEDBACK)(GLuint  id);
// typedef GLboolean  (APIENTRYP GPISVERTEXARRAY)(GLuint  array);
// typedef GLboolean  (APIENTRYP GPISVERTEXARRAYOES)(GLuint  array);
// typedef void  (APIENTRYP GPLABELOBJECTEXT)(GLenum  type, GLuint  object, GLsizei  length, const GLchar * label);
// typedef void  (APIENTRYP GPLINEWIDTH)(GLfloat  width);
// typedef void  (APIENTRYP GPLINKPROGRAM)(GLuint  program);
// typedef void  (APIENTRYP GPMAKEIMAGEHANDLENONRESIDENTNV)(GLuint64  handle);
// typedef void  (APIENTRYP GPMAKEIMAGEHANDLERESIDENTNV)(GLuint64  handle, GLenum  access);
// typedef void  (APIENTRYP GPMAKETEXTUREHANDLENONRESIDENTNV)(GLuint64  handle);
// typedef void  (APIENTRYP GPMAKETEXTUREHANDLERESIDENTNV)(GLuint64  handle);
// typedef void * (APIENTRYP GPMAPBUFFEROES)(GLenum  target, GLenum  access);
// typedef void * (APIENTRYP GPMAPBUFFERRANGE)(GLenum  target, GLintptr  offset, GLsizeiptr  length, GLbitfield  access);
// typedef void * (APIENTRYP GPMAPBUFFERRANGEEXT)(GLenum  target, GLintptr  offset, GLsizeiptr  length, GLbitfield  access);
// typedef void  (APIENTRYP GPMATRIXFRUSTUMEXT)(GLenum  mode, GLdouble  left, GLdouble  right, GLdouble  bottom, GLdouble  top, GLdouble  zNear, GLdouble  zFar);
// typedef void  (APIENTRYP GPMATRIXLOAD3X2FNV)(GLenum  matrixMode, const GLfloat * m);
// typedef void  (APIENTRYP GPMATRIXLOAD3X3FNV)(GLenum  matrixMode, const GLfloat * m);
// typedef void  (APIENTRYP GPMATRIXLOADIDENTITYEXT)(GLenum  mode);
// typedef void  (APIENTRYP GPMATRIXLOADTRANSPOSE3X3FNV)(GLenum  matrixMode, const GLfloat * m);
// typedef void  (APIENTRYP GPMATRIXLOADTRANSPOSEDEXT)(GLenum  mode, const GLdouble * m);
// typedef void  (APIENTRYP GPMATRIXLOADTRANSPOSEFEXT)(GLenum  mode, const GLfloat * m);
// typedef void  (APIENTRYP GPMATRIXLOADDEXT)(GLenum  mode, const GLdouble * m);
// typedef void  (APIENTRYP GPMATRIXLOADFEXT)(GLenum  mode, const GLfloat * m);
// typedef void  (APIENTRYP GPMATRIXMULT3X2FNV)(GLenum  matrixMode, const GLfloat * m);
// typedef void  (APIENTRYP GPMATRIXMULT3X3FNV)(GLenum  matrixMode, const GLfloat * m);
// typedef void  (APIENTRYP GPMATRIXMULTTRANSPOSE3X3FNV)(GLenum  matrixMode, const GLfloat * m);
// typedef void  (APIENTRYP GPMATRIXMULTTRANSPOSEDEXT)(GLenum  mode, const GLdouble * m);
// typedef void  (APIENTRYP GPMATRIXMULTTRANSPOSEFEXT)(GLenum  mode, const GLfloat * m);
// typedef void  (APIENTRYP GPMATRIXMULTDEXT)(GLenum  mode, const GLdouble * m);
// typedef void  (APIENTRYP GPMATRIXMULTFEXT)(GLenum  mode, const GLfloat * m);
// typedef void  (APIENTRYP GPMATRIXORTHOEXT)(GLenum  mode, GLdouble  left, GLdouble  right, GLdouble  bottom, GLdouble  top, GLdouble  zNear, GLdouble  zFar);
// typedef void  (APIENTRYP GPMATRIXPOPEXT)(GLenum  mode);
// typedef void  (APIENTRYP GPMATRIXPUSHEXT)(GLenum  mode);
// typedef void  (APIENTRYP GPMATRIXROTATEDEXT)(GLenum  mode, GLdouble  angle, GLdouble  x, GLdouble  y, GLdouble  z);
// typedef void  (APIENTRYP GPMATRIXROTATEFEXT)(GLenum  mode, GLfloat  angle, GLfloat  x, GLfloat  y, GLfloat  z);
// typedef void  (APIENTRYP GPMATRIXSCALEDEXT)(GLenum  mode, GLdouble  x, GLdouble  y, GLdouble  z);
// typedef void  (APIENTRYP GPMATRIXSCALEFEXT)(GLenum  mode, GLfloat  x, GLfloat  y, GLfloat  z);
// typedef void  (APIENTRYP GPMATRIXTRANSLATEDEXT)(GLenum  mode, GLdouble  x, GLdouble  y, GLdouble  z);
// typedef void  (APIENTRYP GPMATRIXTRANSLATEFEXT)(GLenum  mode, GLfloat  x, GLfloat  y, GLfloat  z);
// typedef void  (APIENTRYP GPMAXSHADERCOMPILERTHREADSKHR)(GLuint  count);
// typedef void  (APIENTRYP GPMEMORYBARRIER)(GLbitfield  barriers);
// typedef void  (APIENTRYP GPMEMORYBARRIERBYREGION)(GLbitfield  barriers);
// typedef void  (APIENTRYP GPMEMORYOBJECTPARAMETERIVEXT)(GLuint  memoryObject, GLenum  pname, const GLint * params);
// typedef void  (APIENTRYP GPMINSAMPLESHADINGOES)(GLfloat  value);
// typedef void  (APIENTRYP GPMULTIDRAWARRAYSEXT)(GLenum  mode, const GLint * first, const GLsizei * count, GLsizei  primcount);
// typedef void  (APIENTRYP GPMULTIDRAWARRAYSINDIRECTEXT)(GLenum  mode, const void * indirect, GLsizei  drawcount, GLsizei  stride);
// typedef void  (APIENTRYP GPMULTIDRAWELEMENTSBASEVERTEXEXT)(GLenum  mode, const GLsizei * count, GLenum  type, const void *const* indices, GLsizei  primcount, const GLint * basevertex);
// typedef void  (APIENTRYP GPMULTIDRAWELEMENTSEXT)(GLenum  mode, const GLsizei * count, GLenum  type, const void *const* indices, GLsizei  primcount);
// typedef void  (APIENTRYP GPMULTIDRAWELEMENTSINDIRECTEXT)(GLenum  mode, GLenum  type, const void * indirect, GLsizei  drawcount, GLsizei  stride);
// typedef void  (APIENTRYP GPNAMEDBUFFERSTORAGEEXTERNALEXT)(GLuint  buffer, GLintptr  offset, GLsizeiptr  size, GLeglClientBufferEXT  clientBuffer, GLbitfield  flags);
// typedef void  (APIENTRYP GPNAMEDBUFFERSTORAGEMEMEXT)(GLuint  buffer, GLsizeiptr  size, GLuint  memory, GLuint64  offset);
// typedef void  (APIENTRYP GPNAMEDFRAMEBUFFERSAMPLELOCATIONSFVNV)(GLuint  framebuffer, GLuint  start, GLsizei  count, const GLfloat * v);
// typedef void  (APIENTRYP GPOBJECTLABEL)(GLenum  identifier, GLuint  name, GLsizei  length, const GLchar * label);
// typedef void  (APIENTRYP GPOBJECTLABELKHR)(GLenum  identifier, GLuint  name, GLsizei  length, const GLchar * label);
// typedef void  (APIENTRYP GPOBJECTPTRLABEL)(const void * ptr, GLsizei  length, const GLchar * label);
// typedef void  (APIENTRYP GPOBJECTPTRLABELKHR)(const void * ptr, GLsizei  length, const GLchar * label);
// typedef void  (APIENTRYP GPPATCHPARAMETERIEXT)(GLenum  pname, GLint  value);
// typedef void  (APIENTRYP GPPATCHPARAMETERIOES)(GLenum  pname, GLint  value);
// typedef void  (APIENTRYP GPPATHCOMMANDSNV)(GLuint  path, GLsizei  numCommands, const GLubyte * commands, GLsizei  numCoords, GLenum  coordType, const void * coords);
// typedef void  (APIENTRYP GPPATHCOORDSNV)(GLuint  path, GLsizei  numCoords, GLenum  coordType, const void * coords);
// typedef void  (APIENTRYP GPPATHCOVERDEPTHFUNCNV)(GLenum  func);
// typedef void  (APIENTRYP GPPATHDASHARRAYNV)(GLuint  path, GLsizei  dashCount, const GLfloat * dashArray);
// typedef GLenum  (APIENTRYP GPPATHGLYPHINDEXARRAYNV)(GLuint  firstPathName, GLenum  fontTarget, const void * fontName, GLbitfield  fontStyle, GLuint  firstGlyphIndex, GLsizei  numGlyphs, GLuint  pathParameterTemplate, GLfloat  emScale);
// typedef GLenum  (APIENTRYP GPPATHGLYPHINDEXRANGENV)(GLenum  fontTarget, const void * fontName, GLbitfield  fontStyle, GLuint  pathParameterTemplate, GLfloat  emScale, GLuint * baseAndCount);
// typedef void  (APIENTRYP GPPATHGLYPHRANGENV)(GLuint  firstPathName, GLenum  fontTarget, const void * fontName, GLbitfield  fontStyle, GLuint  firstGlyph, GLsizei  numGlyphs, GLenum  handleMissingGlyphs, GLuint  pathParameterTemplate, GLfloat  emScale);
// typedef void  (APIENTRYP GPPATHGLYPHSNV)(GLuint  firstPathName, GLenum  fontTarget, const void * fontName, GLbitfield  fontStyle, GLsizei  numGlyphs, GLenum  type, const void * charcodes, GLenum  handleMissingGlyphs, GLuint  pathParameterTemplate, GLfloat  emScale);
// typedef GLenum  (APIENTRYP GPPATHMEMORYGLYPHINDEXARRAYNV)(GLuint  firstPathName, GLenum  fontTarget, GLsizeiptr  fontSize, const void * fontData, GLsizei  faceIndex, GLuint  firstGlyphIndex, GLsizei  numGlyphs, GLuint  pathParameterTemplate, GLfloat  emScale);
// typedef void  (APIENTRYP GPPATHPARAMETERFNV)(GLuint  path, GLenum  pname, GLfloat  value);
// typedef void  (APIENTRYP GPPATHPARAMETERFVNV)(GLuint  path, GLenum  pname, const GLfloat * value);
// typedef void  (APIENTRYP GPPATHPARAMETERINV)(GLuint  path, GLenum  pname, GLint  value);
// typedef void  (APIENTRYP GPPATHPARAMETERIVNV)(GLuint  path, GLenum  pname, const GLint * value);
// typedef void  (APIENTRYP GPPATHSTENCILDEPTHOFFSETNV)(GLfloat  factor, GLfloat  units);
// typedef void  (APIENTRYP GPPATHSTENCILFUNCNV)(GLenum  func, GLint  ref, GLuint  mask);
// typedef void  (APIENTRYP GPPATHSTRINGNV)(GLuint  path, GLenum  format, GLsizei  length, const void * pathString);
// typedef void  (APIENTRYP GPPATHSUBCOMMANDSNV)(GLuint  path, GLsizei  commandStart, GLsizei  commandsToDelete, GLsizei  numCommands, const GLubyte * commands, GLsizei  numCoords, GLenum  coordType, const void * coords);
// typedef void  (APIENTRYP GPPATHSUBCOORDSNV)(GLuint  path, GLsizei  coordStart, GLsizei  numCoords, GLenum  coordType, const void * coords);
// typedef void  (APIENTRYP GPPAUSETRANSFORMFEEDBACK)();
// typedef void  (APIENTRYP GPPIXELSTOREI)(GLenum  pname, GLint  param);
// typedef GLboolean  (APIENTRYP GPPOINTALONGPATHNV)(GLuint  path, GLsizei  startSegment, GLsizei  numSegments, GLfloat  distance, GLfloat * x, GLfloat * y, GLfloat * tangentX, GLfloat * tangentY);
// typedef void  (APIENTRYP GPPOLYGONMODENV)(GLenum  face, GLenum  mode);
// typedef void  (APIENTRYP GPPOLYGONOFFSET)(GLfloat  factor, GLfloat  units);
// typedef void  (APIENTRYP GPPOLYGONOFFSETCLAMPEXT)(GLfloat  factor, GLfloat  units, GLfloat  clamp);
// typedef void  (APIENTRYP GPPOPDEBUGGROUP)();
// typedef void  (APIENTRYP GPPOPDEBUGGROUPKHR)();
// typedef void  (APIENTRYP GPPOPGROUPMARKEREXT)();
// typedef void  (APIENTRYP GPPRIMITIVEBOUNDINGBOXEXT)(GLfloat  minX, GLfloat  minY, GLfloat  minZ, GLfloat  minW, GLfloat  maxX, GLfloat  maxY, GLfloat  maxZ, GLfloat  maxW);
// typedef void  (APIENTRYP GPPRIMITIVEBOUNDINGBOXOES)(GLfloat  minX, GLfloat  minY, GLfloat  minZ, GLfloat  minW, GLfloat  maxX, GLfloat  maxY, GLfloat  maxZ, GLfloat  maxW);
// typedef void  (APIENTRYP GPPROGRAMBINARY)(GLuint  program, GLenum  binaryFormat, const void * binary, GLsizei  length);
// typedef void  (APIENTRYP GPPROGRAMBINARYOES)(GLuint  program, GLenum  binaryFormat, const void * binary, GLint  length);
// typedef void  (APIENTRYP GPPROGRAMPARAMETERI)(GLuint  program, GLenum  pname, GLint  value);
// typedef void  (APIENTRYP GPPROGRAMPARAMETERIEXT)(GLuint  program, GLenum  pname, GLint  value);
// typedef void  (APIENTRYP GPPROGRAMPATHFRAGMENTINPUTGENNV)(GLuint  program, GLint  location, GLenum  genMode, GLint  components, const GLfloat * coeffs);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1F)(GLuint  program, GLint  location, GLfloat  v0);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1FEXT)(GLuint  program, GLint  location, GLfloat  v0);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1FV)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1FVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1I)(GLuint  program, GLint  location, GLint  v0);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1I64NV)(GLuint  program, GLint  location, GLint64EXT  x);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1I64VNV)(GLuint  program, GLint  location, GLsizei  count, const GLint64EXT * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1IEXT)(GLuint  program, GLint  location, GLint  v0);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1IV)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1IVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1UI)(GLuint  program, GLint  location, GLuint  v0);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1UI64NV)(GLuint  program, GLint  location, GLuint64EXT  x);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1UI64VNV)(GLuint  program, GLint  location, GLsizei  count, const GLuint64EXT * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1UIEXT)(GLuint  program, GLint  location, GLuint  v0);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1UIV)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1UIVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2F)(GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2FEXT)(GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2FV)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2FVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2I)(GLuint  program, GLint  location, GLint  v0, GLint  v1);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2I64NV)(GLuint  program, GLint  location, GLint64EXT  x, GLint64EXT  y);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2I64VNV)(GLuint  program, GLint  location, GLsizei  count, const GLint64EXT * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2IEXT)(GLuint  program, GLint  location, GLint  v0, GLint  v1);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2IV)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2IVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2UI)(GLuint  program, GLint  location, GLuint  v0, GLuint  v1);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2UI64NV)(GLuint  program, GLint  location, GLuint64EXT  x, GLuint64EXT  y);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2UI64VNV)(GLuint  program, GLint  location, GLsizei  count, const GLuint64EXT * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2UIEXT)(GLuint  program, GLint  location, GLuint  v0, GLuint  v1);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2UIV)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2UIVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3F)(GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3FEXT)(GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3FV)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3FVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3I)(GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3I64NV)(GLuint  program, GLint  location, GLint64EXT  x, GLint64EXT  y, GLint64EXT  z);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3I64VNV)(GLuint  program, GLint  location, GLsizei  count, const GLint64EXT * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3IEXT)(GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3IV)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3IVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3UI)(GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3UI64NV)(GLuint  program, GLint  location, GLuint64EXT  x, GLuint64EXT  y, GLuint64EXT  z);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3UI64VNV)(GLuint  program, GLint  location, GLsizei  count, const GLuint64EXT * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3UIEXT)(GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3UIV)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3UIVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4F)(GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2, GLfloat  v3);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4FEXT)(GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2, GLfloat  v3);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4FV)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4FVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4I)(GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2, GLint  v3);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4I64NV)(GLuint  program, GLint  location, GLint64EXT  x, GLint64EXT  y, GLint64EXT  z, GLint64EXT  w);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4I64VNV)(GLuint  program, GLint  location, GLsizei  count, const GLint64EXT * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4IEXT)(GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2, GLint  v3);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4IV)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4IVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4UI)(GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2, GLuint  v3);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4UI64NV)(GLuint  program, GLint  location, GLuint64EXT  x, GLuint64EXT  y, GLuint64EXT  z, GLuint64EXT  w);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4UI64VNV)(GLuint  program, GLint  location, GLsizei  count, const GLuint64EXT * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4UIEXT)(GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2, GLuint  v3);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4UIV)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4UIVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMHANDLEUI64IMG)(GLuint  program, GLint  location, GLuint64  value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMHANDLEUI64NV)(GLuint  program, GLint  location, GLuint64  value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMHANDLEUI64VIMG)(GLuint  program, GLint  location, GLsizei  count, const GLuint64 * values);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMHANDLEUI64VNV)(GLuint  program, GLint  location, GLsizei  count, const GLuint64 * values);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2FVEXT)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2X3FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2X3FVEXT)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2X4FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2X4FVEXT)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3FVEXT)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3X2FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3X2FVEXT)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3X4FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3X4FVEXT)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4FVEXT)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4X2FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4X2FVEXT)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4X3FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4X3FVEXT)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPUSHDEBUGGROUP)(GLenum  source, GLuint  id, GLsizei  length, const GLchar * message);
// typedef void  (APIENTRYP GPPUSHDEBUGGROUPKHR)(GLenum  source, GLuint  id, GLsizei  length, const GLchar * message);
// typedef void  (APIENTRYP GPPUSHGROUPMARKEREXT)(GLsizei  length, const GLchar * marker);
// typedef void  (APIENTRYP GPQUERYCOUNTEREXT)(GLuint  id, GLenum  target);
// typedef void  (APIENTRYP GPRASTERSAMPLESEXT)(GLuint  samples, GLboolean  fixedsamplelocations);
// typedef void  (APIENTRYP GPREADBUFFER)(GLenum  src);
// typedef void  (APIENTRYP GPREADBUFFERINDEXEDEXT)(GLenum  src, GLint  index);
// typedef void  (APIENTRYP GPREADBUFFERNV)(GLenum  mode);
// typedef void  (APIENTRYP GPREADPIXELS)(GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, void * pixels);
// typedef void  (APIENTRYP GPREADNPIXELS)(GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data);
// typedef void  (APIENTRYP GPREADNPIXELSEXT)(GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data);
// typedef void  (APIENTRYP GPREADNPIXELSKHR)(GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data);
// typedef GLboolean  (APIENTRYP GPRELEASEKEYEDMUTEXWIN32EXT)(GLuint  memory, GLuint64  key);
// typedef void  (APIENTRYP GPRELEASESHADERCOMPILER)();
// typedef void  (APIENTRYP GPRENDERBUFFERSTORAGE)(GLenum  target, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPRENDERBUFFERSTORAGEMULTISAMPLE)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPRENDERBUFFERSTORAGEMULTISAMPLEANGLE)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPRENDERBUFFERSTORAGEMULTISAMPLEAPPLE)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPRENDERBUFFERSTORAGEMULTISAMPLEEXT)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPRENDERBUFFERSTORAGEMULTISAMPLEIMG)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPRENDERBUFFERSTORAGEMULTISAMPLENV)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPRESOLVEDEPTHVALUESNV)();
// typedef void  (APIENTRYP GPRESOLVEMULTISAMPLEFRAMEBUFFERAPPLE)();
// typedef void  (APIENTRYP GPRESUMETRANSFORMFEEDBACK)();
// typedef void  (APIENTRYP GPSAMPLECOVERAGE)(GLfloat  value, GLboolean  invert);
// typedef void  (APIENTRYP GPSAMPLEMASKI)(GLuint  maskNumber, GLbitfield  mask);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERIIVEXT)(GLuint  sampler, GLenum  pname, const GLint * param);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERIIVOES)(GLuint  sampler, GLenum  pname, const GLint * param);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERIUIVEXT)(GLuint  sampler, GLenum  pname, const GLuint * param);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERIUIVOES)(GLuint  sampler, GLenum  pname, const GLuint * param);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERF)(GLuint  sampler, GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERFV)(GLuint  sampler, GLenum  pname, const GLfloat * param);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERI)(GLuint  sampler, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERIV)(GLuint  sampler, GLenum  pname, const GLint * param);
// typedef void  (APIENTRYP GPSCISSOR)(GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPSCISSORARRAYVNV)(GLuint  first, GLsizei  count, const GLint * v);
// typedef void  (APIENTRYP GPSCISSORARRAYVOES)(GLuint  first, GLsizei  count, const GLint * v);
// typedef void  (APIENTRYP GPSCISSORINDEXEDNV)(GLuint  index, GLint  left, GLint  bottom, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPSCISSORINDEXEDOES)(GLuint  index, GLint  left, GLint  bottom, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPSCISSORINDEXEDVNV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPSCISSORINDEXEDVOES)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPSELECTPERFMONITORCOUNTERSAMD)(GLuint  monitor, GLboolean  enable, GLuint  group, GLint  numCounters, GLuint * counterList);
// typedef void  (APIENTRYP GPSEMAPHOREPARAMETERUI64VEXT)(GLuint  semaphore, GLenum  pname, const GLuint64 * params);
// typedef void  (APIENTRYP GPSETFENCENV)(GLuint  fence, GLenum  condition);
// typedef void  (APIENTRYP GPSHADERBINARY)(GLsizei  count, const GLuint * shaders, GLenum  binaryformat, const void * binary, GLsizei  length);
// typedef void  (APIENTRYP GPSHADERSOURCE)(GLuint  shader, GLsizei  count, const GLchar *const* string, const GLint * length);
// typedef void  (APIENTRYP GPSIGNALSEMAPHOREEXT)(GLuint  semaphore, GLuint  numBufferBarriers, const GLuint * buffers, GLuint  numTextureBarriers, const GLuint * textures, const GLenum * dstLayouts);
// typedef void  (APIENTRYP GPSIGNALVKFENCENV)(GLuint64  vkFence);
// typedef void  (APIENTRYP GPSIGNALVKSEMAPHORENV)(GLuint64  vkSemaphore);
// typedef void  (APIENTRYP GPSTARTTILINGQCOM)(GLuint  x, GLuint  y, GLuint  width, GLuint  height, GLbitfield  preserveMask);
// typedef void  (APIENTRYP GPSTENCILFILLPATHINSTANCEDNV)(GLsizei  numPaths, GLenum  pathNameType, const void * paths, GLuint  pathBase, GLenum  fillMode, GLuint  mask, GLenum  transformType, const GLfloat * transformValues);
// typedef void  (APIENTRYP GPSTENCILFILLPATHNV)(GLuint  path, GLenum  fillMode, GLuint  mask);
// typedef void  (APIENTRYP GPSTENCILFUNC)(GLenum  func, GLint  ref, GLuint  mask);
// typedef void  (APIENTRYP GPSTENCILFUNCSEPARATE)(GLenum  face, GLenum  func, GLint  ref, GLuint  mask);
// typedef void  (APIENTRYP GPSTENCILMASK)(GLuint  mask);
// typedef void  (APIENTRYP GPSTENCILMASKSEPARATE)(GLenum  face, GLuint  mask);
// typedef void  (APIENTRYP GPSTENCILOP)(GLenum  fail, GLenum  zfail, GLenum  zpass);
// typedef void  (APIENTRYP GPSTENCILOPSEPARATE)(GLenum  face, GLenum  sfail, GLenum  dpfail, GLenum  dppass);
// typedef void  (APIENTRYP GPSTENCILSTROKEPATHINSTANCEDNV)(GLsizei  numPaths, GLenum  pathNameType, const void * paths, GLuint  pathBase, GLint  reference, GLuint  mask, GLenum  transformType, const GLfloat * transformValues);
// typedef void  (APIENTRYP GPSTENCILSTROKEPATHNV)(GLuint  path, GLint  reference, GLuint  mask);
// typedef void  (APIENTRYP GPSTENCILTHENCOVERFILLPATHINSTANCEDNV)(GLsizei  numPaths, GLenum  pathNameType, const void * paths, GLuint  pathBase, GLenum  fillMode, GLuint  mask, GLenum  coverMode, GLenum  transformType, const GLfloat * transformValues);
// typedef void  (APIENTRYP GPSTENCILTHENCOVERFILLPATHNV)(GLuint  path, GLenum  fillMode, GLuint  mask, GLenum  coverMode);
// typedef void  (APIENTRYP GPSTENCILTHENCOVERSTROKEPATHINSTANCEDNV)(GLsizei  numPaths, GLenum  pathNameType, const void * paths, GLuint  pathBase, GLint  reference, GLuint  mask, GLenum  coverMode, GLenum  transformType, const GLfloat * transformValues);
// typedef void  (APIENTRYP GPSTENCILTHENCOVERSTROKEPATHNV)(GLuint  path, GLint  reference, GLuint  mask, GLenum  coverMode);
// typedef void  (APIENTRYP GPSUBPIXELPRECISIONBIASNV)(GLuint  xbits, GLuint  ybits);
// typedef GLboolean  (APIENTRYP GPTESTFENCENV)(GLuint  fence);
// typedef void  (APIENTRYP GPTEXBUFFEREXT)(GLenum  target, GLenum  internalformat, GLuint  buffer);
// typedef void  (APIENTRYP GPTEXBUFFEROES)(GLenum  target, GLenum  internalformat, GLuint  buffer);
// typedef void  (APIENTRYP GPTEXBUFFERRANGEEXT)(GLenum  target, GLenum  internalformat, GLuint  buffer, GLintptr  offset, GLsizeiptr  size);
// typedef void  (APIENTRYP GPTEXBUFFERRANGEOES)(GLenum  target, GLenum  internalformat, GLuint  buffer, GLintptr  offset, GLsizeiptr  size);
// typedef void  (APIENTRYP GPTEXIMAGE2D)(GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLsizei  height, GLint  border, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXIMAGE3D)(GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLint  border, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXIMAGE3DOES)(GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLint  border, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXPAGECOMMITMENTEXT)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  commit);
// typedef void  (APIENTRYP GPTEXPARAMETERIIVEXT)(GLenum  target, GLenum  pname, const GLint * params);
// typedef void  (APIENTRYP GPTEXPARAMETERIIVOES)(GLenum  target, GLenum  pname, const GLint * params);
// typedef void  (APIENTRYP GPTEXPARAMETERIUIVEXT)(GLenum  target, GLenum  pname, const GLuint * params);
// typedef void  (APIENTRYP GPTEXPARAMETERIUIVOES)(GLenum  target, GLenum  pname, const GLuint * params);
// typedef void  (APIENTRYP GPTEXPARAMETERF)(GLenum  target, GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPTEXPARAMETERFV)(GLenum  target, GLenum  pname, const GLfloat * params);
// typedef void  (APIENTRYP GPTEXPARAMETERI)(GLenum  target, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPTEXPARAMETERIV)(GLenum  target, GLenum  pname, const GLint * params);
// typedef void  (APIENTRYP GPTEXSTORAGE1DEXT)(GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width);
// typedef void  (APIENTRYP GPTEXSTORAGE2D)(GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPTEXSTORAGE2DEXT)(GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPTEXSTORAGE2DMULTISAMPLE)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLboolean  fixedsamplelocations);
// typedef void  (APIENTRYP GPTEXSTORAGE3D)(GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth);
// typedef void  (APIENTRYP GPTEXSTORAGE3DEXT)(GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth);
// typedef void  (APIENTRYP GPTEXSTORAGE3DMULTISAMPLEOES)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedsamplelocations);
// typedef void  (APIENTRYP GPTEXSTORAGEMEM1DEXT)(GLenum  target, GLsizei  levels, GLenum  internalFormat, GLsizei  width, GLuint  memory, GLuint64  offset);
// typedef void  (APIENTRYP GPTEXSTORAGEMEM2DEXT)(GLenum  target, GLsizei  levels, GLenum  internalFormat, GLsizei  width, GLsizei  height, GLuint  memory, GLuint64  offset);
// typedef void  (APIENTRYP GPTEXSTORAGEMEM2DMULTISAMPLEEXT)(GLenum  target, GLsizei  samples, GLenum  internalFormat, GLsizei  width, GLsizei  height, GLboolean  fixedSampleLocations, GLuint  memory, GLuint64  offset);
// typedef void  (APIENTRYP GPTEXSTORAGEMEM3DEXT)(GLenum  target, GLsizei  levels, GLenum  internalFormat, GLsizei  width, GLsizei  height, GLsizei  depth, GLuint  memory, GLuint64  offset);
// typedef void  (APIENTRYP GPTEXSTORAGEMEM3DMULTISAMPLEEXT)(GLenum  target, GLsizei  samples, GLenum  internalFormat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedSampleLocations, GLuint  memory, GLuint64  offset);
// typedef void  (APIENTRYP GPTEXSUBIMAGE2D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXSUBIMAGE3D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXSUBIMAGE3DOES)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXTUREFOVEATIONPARAMETERSQCOM)(GLuint  texture, GLuint  layer, GLuint  focalPoint, GLfloat  focalX, GLfloat  focalY, GLfloat  gainX, GLfloat  gainY, GLfloat  foveaArea);
// typedef void  (APIENTRYP GPTEXTURESTORAGE1DEXT)(GLuint  texture, GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width);
// typedef void  (APIENTRYP GPTEXTURESTORAGE2DEXT)(GLuint  texture, GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPTEXTURESTORAGE3DEXT)(GLuint  texture, GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth);
// typedef void  (APIENTRYP GPTEXTURESTORAGEMEM1DEXT)(GLuint  texture, GLsizei  levels, GLenum  internalFormat, GLsizei  width, GLuint  memory, GLuint64  offset);
// typedef void  (APIENTRYP GPTEXTURESTORAGEMEM2DEXT)(GLuint  texture, GLsizei  levels, GLenum  internalFormat, GLsizei  width, GLsizei  height, GLuint  memory, GLuint64  offset);
// typedef void  (APIENTRYP GPTEXTURESTORAGEMEM2DMULTISAMPLEEXT)(GLuint  texture, GLsizei  samples, GLenum  internalFormat, GLsizei  width, GLsizei  height, GLboolean  fixedSampleLocations, GLuint  memory, GLuint64  offset);
// typedef void  (APIENTRYP GPTEXTURESTORAGEMEM3DEXT)(GLuint  texture, GLsizei  levels, GLenum  internalFormat, GLsizei  width, GLsizei  height, GLsizei  depth, GLuint  memory, GLuint64  offset);
// typedef void  (APIENTRYP GPTEXTURESTORAGEMEM3DMULTISAMPLEEXT)(GLuint  texture, GLsizei  samples, GLenum  internalFormat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedSampleLocations, GLuint  memory, GLuint64  offset);
// typedef void  (APIENTRYP GPTEXTUREVIEWEXT)(GLuint  texture, GLenum  target, GLuint  origtexture, GLenum  internalformat, GLuint  minlevel, GLuint  numlevels, GLuint  minlayer, GLuint  numlayers);
// typedef void  (APIENTRYP GPTEXTUREVIEWOES)(GLuint  texture, GLenum  target, GLuint  origtexture, GLenum  internalformat, GLuint  minlevel, GLuint  numlevels, GLuint  minlayer, GLuint  numlayers);
// typedef void  (APIENTRYP GPTRANSFORMFEEDBACKVARYINGS)(GLuint  program, GLsizei  count, const GLchar *const* varyings, GLenum  bufferMode);
// typedef void  (APIENTRYP GPTRANSFORMPATHNV)(GLuint  resultPath, GLuint  srcPath, GLenum  transformType, const GLfloat * transformValues);
// typedef void  (APIENTRYP GPUNIFORM1F)(GLint  location, GLfloat  v0);
// typedef void  (APIENTRYP GPUNIFORM1FV)(GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORM1I)(GLint  location, GLint  v0);
// typedef void  (APIENTRYP GPUNIFORM1I64NV)(GLint  location, GLint64EXT  x);
// typedef void  (APIENTRYP GPUNIFORM1I64VNV)(GLint  location, GLsizei  count, const GLint64EXT * value);
// typedef void  (APIENTRYP GPUNIFORM1IV)(GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPUNIFORM1UI)(GLint  location, GLuint  v0);
// typedef void  (APIENTRYP GPUNIFORM1UI64NV)(GLint  location, GLuint64EXT  x);
// typedef void  (APIENTRYP GPUNIFORM1UI64VNV)(GLint  location, GLsizei  count, const GLuint64EXT * value);
// typedef void  (APIENTRYP GPUNIFORM1UIV)(GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPUNIFORM2F)(GLint  location, GLfloat  v0, GLfloat  v1);
// typedef void  (APIENTRYP GPUNIFORM2FV)(GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORM2I)(GLint  location, GLint  v0, GLint  v1);
// typedef void  (APIENTRYP GPUNIFORM2I64NV)(GLint  location, GLint64EXT  x, GLint64EXT  y);
// typedef void  (APIENTRYP GPUNIFORM2I64VNV)(GLint  location, GLsizei  count, const GLint64EXT * value);
// typedef void  (APIENTRYP GPUNIFORM2IV)(GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPUNIFORM2UI)(GLint  location, GLuint  v0, GLuint  v1);
// typedef void  (APIENTRYP GPUNIFORM2UI64NV)(GLint  location, GLuint64EXT  x, GLuint64EXT  y);
// typedef void  (APIENTRYP GPUNIFORM2UI64VNV)(GLint  location, GLsizei  count, const GLuint64EXT * value);
// typedef void  (APIENTRYP GPUNIFORM2UIV)(GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPUNIFORM3F)(GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2);
// typedef void  (APIENTRYP GPUNIFORM3FV)(GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORM3I)(GLint  location, GLint  v0, GLint  v1, GLint  v2);
// typedef void  (APIENTRYP GPUNIFORM3I64NV)(GLint  location, GLint64EXT  x, GLint64EXT  y, GLint64EXT  z);
// typedef void  (APIENTRYP GPUNIFORM3I64VNV)(GLint  location, GLsizei  count, const GLint64EXT * value);
// typedef void  (APIENTRYP GPUNIFORM3IV)(GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPUNIFORM3UI)(GLint  location, GLuint  v0, GLuint  v1, GLuint  v2);
// typedef void  (APIENTRYP GPUNIFORM3UI64NV)(GLint  location, GLuint64EXT  x, GLuint64EXT  y, GLuint64EXT  z);
// typedef void  (APIENTRYP GPUNIFORM3UI64VNV)(GLint  location, GLsizei  count, const GLuint64EXT * value);
// typedef void  (APIENTRYP GPUNIFORM3UIV)(GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPUNIFORM4F)(GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2, GLfloat  v3);
// typedef void  (APIENTRYP GPUNIFORM4FV)(GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORM4I)(GLint  location, GLint  v0, GLint  v1, GLint  v2, GLint  v3);
// typedef void  (APIENTRYP GPUNIFORM4I64NV)(GLint  location, GLint64EXT  x, GLint64EXT  y, GLint64EXT  z, GLint64EXT  w);
// typedef void  (APIENTRYP GPUNIFORM4I64VNV)(GLint  location, GLsizei  count, const GLint64EXT * value);
// typedef void  (APIENTRYP GPUNIFORM4IV)(GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPUNIFORM4UI)(GLint  location, GLuint  v0, GLuint  v1, GLuint  v2, GLuint  v3);
// typedef void  (APIENTRYP GPUNIFORM4UI64NV)(GLint  location, GLuint64EXT  x, GLuint64EXT  y, GLuint64EXT  z, GLuint64EXT  w);
// typedef void  (APIENTRYP GPUNIFORM4UI64VNV)(GLint  location, GLsizei  count, const GLuint64EXT * value);
// typedef void  (APIENTRYP GPUNIFORM4UIV)(GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPUNIFORMBLOCKBINDING)(GLuint  program, GLuint  uniformBlockIndex, GLuint  uniformBlockBinding);
// typedef void  (APIENTRYP GPUNIFORMHANDLEUI64IMG)(GLint  location, GLuint64  value);
// typedef void  (APIENTRYP GPUNIFORMHANDLEUI64NV)(GLint  location, GLuint64  value);
// typedef void  (APIENTRYP GPUNIFORMHANDLEUI64VIMG)(GLint  location, GLsizei  count, const GLuint64 * value);
// typedef void  (APIENTRYP GPUNIFORMHANDLEUI64VNV)(GLint  location, GLsizei  count, const GLuint64 * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2X3FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2X3FVNV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2X4FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2X4FVNV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3X2FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3X2FVNV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3X4FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3X4FVNV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4X2FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4X2FVNV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4X3FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4X3FVNV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef GLboolean  (APIENTRYP GPUNMAPBUFFER)(GLenum  target);
// typedef GLboolean  (APIENTRYP GPUNMAPBUFFEROES)(GLenum  target);
// typedef void  (APIENTRYP GPUSEPROGRAM)(GLuint  program);
// typedef void  (APIENTRYP GPUSEPROGRAMSTAGES)(GLuint  pipeline, GLbitfield  stages, GLuint  program);
// typedef void  (APIENTRYP GPUSEPROGRAMSTAGESEXT)(GLuint  pipeline, GLbitfield  stages, GLuint  program);
// typedef void  (APIENTRYP GPUSESHADERPROGRAMEXT)(GLenum  type, GLuint  program);
// typedef void  (APIENTRYP GPVALIDATEPROGRAM)(GLuint  program);
// typedef void  (APIENTRYP GPVALIDATEPROGRAMPIPELINE)(GLuint  pipeline);
// typedef void  (APIENTRYP GPVALIDATEPROGRAMPIPELINEEXT)(GLuint  pipeline);
// typedef void  (APIENTRYP GPVERTEXATTRIB1F)(GLuint  index, GLfloat  x);
// typedef void  (APIENTRYP GPVERTEXATTRIB1FV)(GLuint  index, const GLfloat * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB2F)(GLuint  index, GLfloat  x, GLfloat  y);
// typedef void  (APIENTRYP GPVERTEXATTRIB2FV)(GLuint  index, const GLfloat * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB3F)(GLuint  index, GLfloat  x, GLfloat  y, GLfloat  z);
// typedef void  (APIENTRYP GPVERTEXATTRIB3FV)(GLuint  index, const GLfloat * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4F)(GLuint  index, GLfloat  x, GLfloat  y, GLfloat  z, GLfloat  w);
// typedef void  (APIENTRYP GPVERTEXATTRIB4FV)(GLuint  index, const GLfloat * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBBINDING)(GLuint  attribindex, GLuint  bindingindex);
// typedef void  (APIENTRYP GPVERTEXATTRIBDIVISOR)(GLuint  index, GLuint  divisor);
// typedef void  (APIENTRYP GPVERTEXATTRIBDIVISORANGLE)(GLuint  index, GLuint  divisor);
// typedef void  (APIENTRYP GPVERTEXATTRIBDIVISOREXT)(GLuint  index, GLuint  divisor);
// typedef void  (APIENTRYP GPVERTEXATTRIBDIVISORNV)(GLuint  index, GLuint  divisor);
// typedef void  (APIENTRYP GPVERTEXATTRIBFORMAT)(GLuint  attribindex, GLint  size, GLenum  type, GLboolean  normalized, GLuint  relativeoffset);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4I)(GLuint  index, GLint  x, GLint  y, GLint  z, GLint  w);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4IV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4UI)(GLuint  index, GLuint  x, GLuint  y, GLuint  z, GLuint  w);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4UIV)(GLuint  index, const GLuint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBIFORMAT)(GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset);
// typedef void  (APIENTRYP GPVERTEXATTRIBIPOINTER)(GLuint  index, GLint  size, GLenum  type, GLsizei  stride, const void * pointer);
// typedef void  (APIENTRYP GPVERTEXATTRIBPOINTER)(GLuint  index, GLint  size, GLenum  type, GLboolean  normalized, GLsizei  stride, const void * pointer);
// typedef void  (APIENTRYP GPVERTEXBINDINGDIVISOR)(GLuint  bindingindex, GLuint  divisor);
// typedef void  (APIENTRYP GPVIEWPORT)(GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPVIEWPORTARRAYVNV)(GLuint  first, GLsizei  count, const GLfloat * v);
// typedef void  (APIENTRYP GPVIEWPORTARRAYVOES)(GLuint  first, GLsizei  count, const GLfloat * v);
// typedef void  (APIENTRYP GPVIEWPORTINDEXEDFNV)(GLuint  index, GLfloat  x, GLfloat  y, GLfloat  w, GLfloat  h);
// typedef void  (APIENTRYP GPVIEWPORTINDEXEDFOES)(GLuint  index, GLfloat  x, GLfloat  y, GLfloat  w, GLfloat  h);
// typedef void  (APIENTRYP GPVIEWPORTINDEXEDFVNV)(GLuint  index, const GLfloat * v);
// typedef void  (APIENTRYP GPVIEWPORTINDEXEDFVOES)(GLuint  index, const GLfloat * v);
// typedef void  (APIENTRYP GPVIEWPORTPOSITIONWSCALENV)(GLuint  index, GLfloat  xcoeff, GLfloat  ycoeff);
// typedef void  (APIENTRYP GPVIEWPORTSWIZZLENV)(GLuint  index, GLenum  swizzlex, GLenum  swizzley, GLenum  swizzlez, GLenum  swizzlew);
// typedef void  (APIENTRYP GPWAITSEMAPHOREEXT)(GLuint  semaphore, GLuint  numBufferBarriers, const GLuint * buffers, GLuint  numTextureBarriers, const GLuint * textures, const GLenum * srcLayouts);
// typedef void  (APIENTRYP GPWAITSYNC)(GLsync  sync, GLbitfield  flags, GLuint64  timeout);
// typedef void  (APIENTRYP GPWAITSYNCAPPLE)(GLsync  sync, GLbitfield  flags, GLuint64  timeout);
// typedef void  (APIENTRYP GPWAITVKSEMAPHORENV)(GLuint64  vkSemaphore);
// typedef void  (APIENTRYP GPWEIGHTPATHSNV)(GLuint  resultPath, GLsizei  numPaths, const GLuint * paths, const GLfloat * weights);
// typedef void  (APIENTRYP GPWINDOWRECTANGLESEXT)(GLenum  mode, GLsizei  count, const GLint * box);
// static GLboolean  glowAcquireKeyedMutexWin32EXT(GPACQUIREKEYEDMUTEXWIN32EXT fnptr, GLuint  memory, GLuint64  key, GLuint  timeout) {
//   return (*fnptr)(memory, key, timeout);
// }
// static void  glowActiveProgramEXT(GPACTIVEPROGRAMEXT fnptr, GLuint  program) {
//   (*fnptr)(program);
// }
// static void  glowActiveShaderProgram(GPACTIVESHADERPROGRAM fnptr, GLuint  pipeline, GLuint  program) {
//   (*fnptr)(pipeline, program);
// }
// static void  glowActiveShaderProgramEXT(GPACTIVESHADERPROGRAMEXT fnptr, GLuint  pipeline, GLuint  program) {
//   (*fnptr)(pipeline, program);
// }
// static void  glowActiveTexture(GPACTIVETEXTURE fnptr, GLenum  texture) {
//   (*fnptr)(texture);
// }
// static void  glowAlphaFuncQCOM(GPALPHAFUNCQCOM fnptr, GLenum  func, GLclampf  ref) {
//   (*fnptr)(func, ref);
// }
// static void  glowApplyFramebufferAttachmentCMAAINTEL(GPAPPLYFRAMEBUFFERATTACHMENTCMAAINTEL fnptr) {
//   (*fnptr)();
// }
// static void  glowAttachShader(GPATTACHSHADER fnptr, GLuint  program, GLuint  shader) {
//   (*fnptr)(program, shader);
// }
// static void  glowBeginConditionalRenderNV(GPBEGINCONDITIONALRENDERNV fnptr, GLuint  id, GLenum  mode) {
//   (*fnptr)(id, mode);
// }
// static void  glowBeginPerfMonitorAMD(GPBEGINPERFMONITORAMD fnptr, GLuint  monitor) {
//   (*fnptr)(monitor);
// }
// static void  glowBeginPerfQueryINTEL(GPBEGINPERFQUERYINTEL fnptr, GLuint  queryHandle) {
//   (*fnptr)(queryHandle);
// }
// static void  glowBeginQuery(GPBEGINQUERY fnptr, GLenum  target, GLuint  id) {
//   (*fnptr)(target, id);
// }
// static void  glowBeginQueryEXT(GPBEGINQUERYEXT fnptr, GLenum  target, GLuint  id) {
//   (*fnptr)(target, id);
// }
// static void  glowBeginTransformFeedback(GPBEGINTRANSFORMFEEDBACK fnptr, GLenum  primitiveMode) {
//   (*fnptr)(primitiveMode);
// }
// static void  glowBindAttribLocation(GPBINDATTRIBLOCATION fnptr, GLuint  program, GLuint  index, const GLchar * name) {
//   (*fnptr)(program, index, name);
// }
// static void  glowBindBuffer(GPBINDBUFFER fnptr, GLenum  target, GLuint  buffer) {
//   (*fnptr)(target, buffer);
// }
// static void  glowBindBufferBase(GPBINDBUFFERBASE fnptr, GLenum  target, GLuint  index, GLuint  buffer) {
//   (*fnptr)(target, index, buffer);
// }
// static void  glowBindBufferRange(GPBINDBUFFERRANGE fnptr, GLenum  target, GLuint  index, GLuint  buffer, GLintptr  offset, GLsizeiptr  size) {
//   (*fnptr)(target, index, buffer, offset, size);
// }
// static void  glowBindFragDataLocationEXT(GPBINDFRAGDATALOCATIONEXT fnptr, GLuint  program, GLuint  color, const GLchar * name) {
//   (*fnptr)(program, color, name);
// }
// static void  glowBindFragDataLocationIndexedEXT(GPBINDFRAGDATALOCATIONINDEXEDEXT fnptr, GLuint  program, GLuint  colorNumber, GLuint  index, const GLchar * name) {
//   (*fnptr)(program, colorNumber, index, name);
// }
// static void  glowBindFramebuffer(GPBINDFRAMEBUFFER fnptr, GLenum  target, GLuint  framebuffer) {
//   (*fnptr)(target, framebuffer);
// }
// static void  glowBindImageTexture(GPBINDIMAGETEXTURE fnptr, GLuint  unit, GLuint  texture, GLint  level, GLboolean  layered, GLint  layer, GLenum  access, GLenum  format) {
//   (*fnptr)(unit, texture, level, layered, layer, access, format);
// }
// static void  glowBindProgramPipeline(GPBINDPROGRAMPIPELINE fnptr, GLuint  pipeline) {
//   (*fnptr)(pipeline);
// }
// static void  glowBindProgramPipelineEXT(GPBINDPROGRAMPIPELINEEXT fnptr, GLuint  pipeline) {
//   (*fnptr)(pipeline);
// }
// static void  glowBindRenderbuffer(GPBINDRENDERBUFFER fnptr, GLenum  target, GLuint  renderbuffer) {
//   (*fnptr)(target, renderbuffer);
// }
// static void  glowBindSampler(GPBINDSAMPLER fnptr, GLuint  unit, GLuint  sampler) {
//   (*fnptr)(unit, sampler);
// }
// static void  glowBindTexture(GPBINDTEXTURE fnptr, GLenum  target, GLuint  texture) {
//   (*fnptr)(target, texture);
// }
// static void  glowBindTransformFeedback(GPBINDTRANSFORMFEEDBACK fnptr, GLenum  target, GLuint  id) {
//   (*fnptr)(target, id);
// }
// static void  glowBindVertexArray(GPBINDVERTEXARRAY fnptr, GLuint  array) {
//   (*fnptr)(array);
// }
// static void  glowBindVertexArrayOES(GPBINDVERTEXARRAYOES fnptr, GLuint  array) {
//   (*fnptr)(array);
// }
// static void  glowBindVertexBuffer(GPBINDVERTEXBUFFER fnptr, GLuint  bindingindex, GLuint  buffer, GLintptr  offset, GLsizei  stride) {
//   (*fnptr)(bindingindex, buffer, offset, stride);
// }
// static void  glowBlendBarrierKHR(GPBLENDBARRIERKHR fnptr) {
//   (*fnptr)();
// }
// static void  glowBlendBarrierNV(GPBLENDBARRIERNV fnptr) {
//   (*fnptr)();
// }
// static void  glowBlendColor(GPBLENDCOLOR fnptr, GLfloat  red, GLfloat  green, GLfloat  blue, GLfloat  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowBlendEquation(GPBLENDEQUATION fnptr, GLenum  mode) {
//   (*fnptr)(mode);
// }
// static void  glowBlendEquationEXT(GPBLENDEQUATIONEXT fnptr, GLenum  mode) {
//   (*fnptr)(mode);
// }
// static void  glowBlendEquationSeparate(GPBLENDEQUATIONSEPARATE fnptr, GLenum  modeRGB, GLenum  modeAlpha) {
//   (*fnptr)(modeRGB, modeAlpha);
// }
// static void  glowBlendEquationSeparateiEXT(GPBLENDEQUATIONSEPARATEIEXT fnptr, GLuint  buf, GLenum  modeRGB, GLenum  modeAlpha) {
//   (*fnptr)(buf, modeRGB, modeAlpha);
// }
// static void  glowBlendEquationSeparateiOES(GPBLENDEQUATIONSEPARATEIOES fnptr, GLuint  buf, GLenum  modeRGB, GLenum  modeAlpha) {
//   (*fnptr)(buf, modeRGB, modeAlpha);
// }
// static void  glowBlendEquationiEXT(GPBLENDEQUATIONIEXT fnptr, GLuint  buf, GLenum  mode) {
//   (*fnptr)(buf, mode);
// }
// static void  glowBlendEquationiOES(GPBLENDEQUATIONIOES fnptr, GLuint  buf, GLenum  mode) {
//   (*fnptr)(buf, mode);
// }
// static void  glowBlendFunc(GPBLENDFUNC fnptr, GLenum  sfactor, GLenum  dfactor) {
//   (*fnptr)(sfactor, dfactor);
// }
// static void  glowBlendFuncSeparate(GPBLENDFUNCSEPARATE fnptr, GLenum  sfactorRGB, GLenum  dfactorRGB, GLenum  sfactorAlpha, GLenum  dfactorAlpha) {
//   (*fnptr)(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha);
// }
// static void  glowBlendFuncSeparateiEXT(GPBLENDFUNCSEPARATEIEXT fnptr, GLuint  buf, GLenum  srcRGB, GLenum  dstRGB, GLenum  srcAlpha, GLenum  dstAlpha) {
//   (*fnptr)(buf, srcRGB, dstRGB, srcAlpha, dstAlpha);
// }
// static void  glowBlendFuncSeparateiOES(GPBLENDFUNCSEPARATEIOES fnptr, GLuint  buf, GLenum  srcRGB, GLenum  dstRGB, GLenum  srcAlpha, GLenum  dstAlpha) {
//   (*fnptr)(buf, srcRGB, dstRGB, srcAlpha, dstAlpha);
// }
// static void  glowBlendFunciEXT(GPBLENDFUNCIEXT fnptr, GLuint  buf, GLenum  src, GLenum  dst) {
//   (*fnptr)(buf, src, dst);
// }
// static void  glowBlendFunciOES(GPBLENDFUNCIOES fnptr, GLuint  buf, GLenum  src, GLenum  dst) {
//   (*fnptr)(buf, src, dst);
// }
// static void  glowBlendParameteriNV(GPBLENDPARAMETERINV fnptr, GLenum  pname, GLint  value) {
//   (*fnptr)(pname, value);
// }
// static void  glowBlitFramebuffer(GPBLITFRAMEBUFFER fnptr, GLint  srcX0, GLint  srcY0, GLint  srcX1, GLint  srcY1, GLint  dstX0, GLint  dstY0, GLint  dstX1, GLint  dstY1, GLbitfield  mask, GLenum  filter) {
//   (*fnptr)(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
// }
// static void  glowBlitFramebufferANGLE(GPBLITFRAMEBUFFERANGLE fnptr, GLint  srcX0, GLint  srcY0, GLint  srcX1, GLint  srcY1, GLint  dstX0, GLint  dstY0, GLint  dstX1, GLint  dstY1, GLbitfield  mask, GLenum  filter) {
//   (*fnptr)(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
// }
// static void  glowBlitFramebufferNV(GPBLITFRAMEBUFFERNV fnptr, GLint  srcX0, GLint  srcY0, GLint  srcX1, GLint  srcY1, GLint  dstX0, GLint  dstY0, GLint  dstX1, GLint  dstY1, GLbitfield  mask, GLenum  filter) {
//   (*fnptr)(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
// }
// static void  glowBufferData(GPBUFFERDATA fnptr, GLenum  target, GLsizeiptr  size, const void * data, GLenum  usage) {
//   (*fnptr)(target, size, data, usage);
// }
// static void  glowBufferStorageEXT(GPBUFFERSTORAGEEXT fnptr, GLenum  target, GLsizeiptr  size, const void * data, GLbitfield  flags) {
//   (*fnptr)(target, size, data, flags);
// }
// static void  glowBufferStorageExternalEXT(GPBUFFERSTORAGEEXTERNALEXT fnptr, GLenum  target, GLintptr  offset, GLsizeiptr  size, GLeglClientBufferEXT  clientBuffer, GLbitfield  flags) {
//   (*fnptr)(target, offset, size, clientBuffer, flags);
// }
// static void  glowBufferStorageMemEXT(GPBUFFERSTORAGEMEMEXT fnptr, GLenum  target, GLsizeiptr  size, GLuint  memory, GLuint64  offset) {
//   (*fnptr)(target, size, memory, offset);
// }
// static void  glowBufferSubData(GPBUFFERSUBDATA fnptr, GLenum  target, GLintptr  offset, GLsizeiptr  size, const void * data) {
//   (*fnptr)(target, offset, size, data);
// }
// static GLenum  glowCheckFramebufferStatus(GPCHECKFRAMEBUFFERSTATUS fnptr, GLenum  target) {
//   return (*fnptr)(target);
// }
// static void  glowClear(GPCLEAR fnptr, GLbitfield  mask) {
//   (*fnptr)(mask);
// }
// static void  glowClearBufferfi(GPCLEARBUFFERFI fnptr, GLenum  buffer, GLint  drawbuffer, GLfloat  depth, GLint  stencil) {
//   (*fnptr)(buffer, drawbuffer, depth, stencil);
// }
// static void  glowClearBufferfv(GPCLEARBUFFERFV fnptr, GLenum  buffer, GLint  drawbuffer, const GLfloat * value) {
//   (*fnptr)(buffer, drawbuffer, value);
// }
// static void  glowClearBufferiv(GPCLEARBUFFERIV fnptr, GLenum  buffer, GLint  drawbuffer, const GLint * value) {
//   (*fnptr)(buffer, drawbuffer, value);
// }
// static void  glowClearBufferuiv(GPCLEARBUFFERUIV fnptr, GLenum  buffer, GLint  drawbuffer, const GLuint * value) {
//   (*fnptr)(buffer, drawbuffer, value);
// }
// static void  glowClearColor(GPCLEARCOLOR fnptr, GLfloat  red, GLfloat  green, GLfloat  blue, GLfloat  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowClearDepthf(GPCLEARDEPTHF fnptr, GLfloat  d) {
//   (*fnptr)(d);
// }
// static void  glowClearPixelLocalStorageuiEXT(GPCLEARPIXELLOCALSTORAGEUIEXT fnptr, GLsizei  offset, GLsizei  n, const GLuint * values) {
//   (*fnptr)(offset, n, values);
// }
// static void  glowClearStencil(GPCLEARSTENCIL fnptr, GLint  s) {
//   (*fnptr)(s);
// }
// static void  glowClearTexImageEXT(GPCLEARTEXIMAGEEXT fnptr, GLuint  texture, GLint  level, GLenum  format, GLenum  type, const void * data) {
//   (*fnptr)(texture, level, format, type, data);
// }
// static void  glowClearTexSubImageEXT(GPCLEARTEXSUBIMAGEEXT fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * data) {
//   (*fnptr)(texture, level, xoffset, yoffset, zoffset, width, height, depth, format, type, data);
// }
// static GLenum  glowClientWaitSync(GPCLIENTWAITSYNC fnptr, GLsync  sync, GLbitfield  flags, GLuint64  timeout) {
//   return (*fnptr)(sync, flags, timeout);
// }
// static GLenum  glowClientWaitSyncAPPLE(GPCLIENTWAITSYNCAPPLE fnptr, GLsync  sync, GLbitfield  flags, GLuint64  timeout) {
//   return (*fnptr)(sync, flags, timeout);
// }
// static void  glowClipControlEXT(GPCLIPCONTROLEXT fnptr, GLenum  origin, GLenum  depth) {
//   (*fnptr)(origin, depth);
// }
// static void  glowColorMask(GPCOLORMASK fnptr, GLboolean  red, GLboolean  green, GLboolean  blue, GLboolean  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowColorMaskiEXT(GPCOLORMASKIEXT fnptr, GLuint  index, GLboolean  r, GLboolean  g, GLboolean  b, GLboolean  a) {
//   (*fnptr)(index, r, g, b, a);
// }
// static void  glowColorMaskiOES(GPCOLORMASKIOES fnptr, GLuint  index, GLboolean  r, GLboolean  g, GLboolean  b, GLboolean  a) {
//   (*fnptr)(index, r, g, b, a);
// }
// static void  glowCompileShader(GPCOMPILESHADER fnptr, GLuint  shader) {
//   (*fnptr)(shader);
// }
// static void  glowCompressedTexImage2D(GPCOMPRESSEDTEXIMAGE2D fnptr, GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLsizei  height, GLint  border, GLsizei  imageSize, const void * data) {
//   (*fnptr)(target, level, internalformat, width, height, border, imageSize, data);
// }
// static void  glowCompressedTexImage3D(GPCOMPRESSEDTEXIMAGE3D fnptr, GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLint  border, GLsizei  imageSize, const void * data) {
//   (*fnptr)(target, level, internalformat, width, height, depth, border, imageSize, data);
// }
// static void  glowCompressedTexImage3DOES(GPCOMPRESSEDTEXIMAGE3DOES fnptr, GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLint  border, GLsizei  imageSize, const void * data) {
//   (*fnptr)(target, level, internalformat, width, height, depth, border, imageSize, data);
// }
// static void  glowCompressedTexSubImage2D(GPCOMPRESSEDTEXSUBIMAGE2D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLsizei  imageSize, const void * data) {
//   (*fnptr)(target, level, xoffset, yoffset, width, height, format, imageSize, data);
// }
// static void  glowCompressedTexSubImage3D(GPCOMPRESSEDTEXSUBIMAGE3D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLsizei  imageSize, const void * data) {
//   (*fnptr)(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data);
// }
// static void  glowCompressedTexSubImage3DOES(GPCOMPRESSEDTEXSUBIMAGE3DOES fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLsizei  imageSize, const void * data) {
//   (*fnptr)(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data);
// }
// static void  glowConservativeRasterParameteriNV(GPCONSERVATIVERASTERPARAMETERINV fnptr, GLenum  pname, GLint  param) {
//   (*fnptr)(pname, param);
// }
// static void  glowCopyBufferSubData(GPCOPYBUFFERSUBDATA fnptr, GLenum  readTarget, GLenum  writeTarget, GLintptr  readOffset, GLintptr  writeOffset, GLsizeiptr  size) {
//   (*fnptr)(readTarget, writeTarget, readOffset, writeOffset, size);
// }
// static void  glowCopyBufferSubDataNV(GPCOPYBUFFERSUBDATANV fnptr, GLenum  readTarget, GLenum  writeTarget, GLintptr  readOffset, GLintptr  writeOffset, GLsizeiptr  size) {
//   (*fnptr)(readTarget, writeTarget, readOffset, writeOffset, size);
// }
// static void  glowCopyImageSubDataEXT(GPCOPYIMAGESUBDATAEXT fnptr, GLuint  srcName, GLenum  srcTarget, GLint  srcLevel, GLint  srcX, GLint  srcY, GLint  srcZ, GLuint  dstName, GLenum  dstTarget, GLint  dstLevel, GLint  dstX, GLint  dstY, GLint  dstZ, GLsizei  srcWidth, GLsizei  srcHeight, GLsizei  srcDepth) {
//   (*fnptr)(srcName, srcTarget, srcLevel, srcX, srcY, srcZ, dstName, dstTarget, dstLevel, dstX, dstY, dstZ, srcWidth, srcHeight, srcDepth);
// }
// static void  glowCopyImageSubDataOES(GPCOPYIMAGESUBDATAOES fnptr, GLuint  srcName, GLenum  srcTarget, GLint  srcLevel, GLint  srcX, GLint  srcY, GLint  srcZ, GLuint  dstName, GLenum  dstTarget, GLint  dstLevel, GLint  dstX, GLint  dstY, GLint  dstZ, GLsizei  srcWidth, GLsizei  srcHeight, GLsizei  srcDepth) {
//   (*fnptr)(srcName, srcTarget, srcLevel, srcX, srcY, srcZ, dstName, dstTarget, dstLevel, dstX, dstY, dstZ, srcWidth, srcHeight, srcDepth);
// }
// static void  glowCopyPathNV(GPCOPYPATHNV fnptr, GLuint  resultPath, GLuint  srcPath) {
//   (*fnptr)(resultPath, srcPath);
// }
// static void  glowCopyTexImage2D(GPCOPYTEXIMAGE2D fnptr, GLenum  target, GLint  level, GLenum  internalformat, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLint  border) {
//   (*fnptr)(target, level, internalformat, x, y, width, height, border);
// }
// static void  glowCopyTexSubImage2D(GPCOPYTEXSUBIMAGE2D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, level, xoffset, yoffset, x, y, width, height);
// }
// static void  glowCopyTexSubImage3D(GPCOPYTEXSUBIMAGE3D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, level, xoffset, yoffset, zoffset, x, y, width, height);
// }
// static void  glowCopyTexSubImage3DOES(GPCOPYTEXSUBIMAGE3DOES fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, level, xoffset, yoffset, zoffset, x, y, width, height);
// }
// static void  glowCopyTextureLevelsAPPLE(GPCOPYTEXTURELEVELSAPPLE fnptr, GLuint  destinationTexture, GLuint  sourceTexture, GLint  sourceBaseLevel, GLsizei  sourceLevelCount) {
//   (*fnptr)(destinationTexture, sourceTexture, sourceBaseLevel, sourceLevelCount);
// }
// static void  glowCoverFillPathInstancedNV(GPCOVERFILLPATHINSTANCEDNV fnptr, GLsizei  numPaths, GLenum  pathNameType, const void * paths, GLuint  pathBase, GLenum  coverMode, GLenum  transformType, const GLfloat * transformValues) {
//   (*fnptr)(numPaths, pathNameType, paths, pathBase, coverMode, transformType, transformValues);
// }
// static void  glowCoverFillPathNV(GPCOVERFILLPATHNV fnptr, GLuint  path, GLenum  coverMode) {
//   (*fnptr)(path, coverMode);
// }
// static void  glowCoverStrokePathInstancedNV(GPCOVERSTROKEPATHINSTANCEDNV fnptr, GLsizei  numPaths, GLenum  pathNameType, const void * paths, GLuint  pathBase, GLenum  coverMode, GLenum  transformType, const GLfloat * transformValues) {
//   (*fnptr)(numPaths, pathNameType, paths, pathBase, coverMode, transformType, transformValues);
// }
// static void  glowCoverStrokePathNV(GPCOVERSTROKEPATHNV fnptr, GLuint  path, GLenum  coverMode) {
//   (*fnptr)(path, coverMode);
// }
// static void  glowCoverageMaskNV(GPCOVERAGEMASKNV fnptr, GLboolean  mask) {
//   (*fnptr)(mask);
// }
// static void  glowCoverageModulationNV(GPCOVERAGEMODULATIONNV fnptr, GLenum  components) {
//   (*fnptr)(components);
// }
// static void  glowCoverageModulationTableNV(GPCOVERAGEMODULATIONTABLENV fnptr, GLsizei  n, const GLfloat * v) {
//   (*fnptr)(n, v);
// }
// static void  glowCoverageOperationNV(GPCOVERAGEOPERATIONNV fnptr, GLenum  operation) {
//   (*fnptr)(operation);
// }
// static void  glowCreateMemoryObjectsEXT(GPCREATEMEMORYOBJECTSEXT fnptr, GLsizei  n, GLuint * memoryObjects) {
//   (*fnptr)(n, memoryObjects);
// }
// static void  glowCreatePerfQueryINTEL(GPCREATEPERFQUERYINTEL fnptr, GLuint  queryId, GLuint * queryHandle) {
//   (*fnptr)(queryId, queryHandle);
// }
// static GLuint  glowCreateProgram(GPCREATEPROGRAM fnptr) {
//   return (*fnptr)();
// }
// static GLuint  glowCreateShader(GPCREATESHADER fnptr, GLenum  type) {
//   return (*fnptr)(type);
// }
// static GLuint  glowCreateShaderProgramEXT(GPCREATESHADERPROGRAMEXT fnptr, GLenum  type, const GLchar * string) {
//   return (*fnptr)(type, string);
// }
// static GLuint  glowCreateShaderProgramv(GPCREATESHADERPROGRAMV fnptr, GLenum  type, GLsizei  count, const GLchar *const* strings) {
//   return (*fnptr)(type, count, strings);
// }
// static GLuint  glowCreateShaderProgramvEXT(GPCREATESHADERPROGRAMVEXT fnptr, GLenum  type, GLsizei  count, const GLchar ** strings) {
//   return (*fnptr)(type, count, strings);
// }
// static void  glowCullFace(GPCULLFACE fnptr, GLenum  mode) {
//   (*fnptr)(mode);
// }
// static void  glowDebugMessageCallback(GPDEBUGMESSAGECALLBACK fnptr, GLDEBUGPROC  callback, const void * userParam) {
//   (*fnptr)(glowCDebugCallback, userParam);
// }
// static void  glowDebugMessageCallbackKHR(GPDEBUGMESSAGECALLBACKKHR fnptr, GLDEBUGPROCKHR  callback, const void * userParam) {
//   (*fnptr)(glowCDebugCallback, userParam);
// }
// static void  glowDebugMessageControl(GPDEBUGMESSAGECONTROL fnptr, GLenum  source, GLenum  type, GLenum  severity, GLsizei  count, const GLuint * ids, GLboolean  enabled) {
//   (*fnptr)(source, type, severity, count, ids, enabled);
// }
// static void  glowDebugMessageControlKHR(GPDEBUGMESSAGECONTROLKHR fnptr, GLenum  source, GLenum  type, GLenum  severity, GLsizei  count, const GLuint * ids, GLboolean  enabled) {
//   (*fnptr)(source, type, severity, count, ids, enabled);
// }
// static void  glowDebugMessageInsert(GPDEBUGMESSAGEINSERT fnptr, GLenum  source, GLenum  type, GLuint  id, GLenum  severity, GLsizei  length, const GLchar * buf) {
//   (*fnptr)(source, type, id, severity, length, buf);
// }
// static void  glowDebugMessageInsertKHR(GPDEBUGMESSAGEINSERTKHR fnptr, GLenum  source, GLenum  type, GLuint  id, GLenum  severity, GLsizei  length, const GLchar * buf) {
//   (*fnptr)(source, type, id, severity, length, buf);
// }
// static void  glowDeleteBuffers(GPDELETEBUFFERS fnptr, GLsizei  n, const GLuint * buffers) {
//   (*fnptr)(n, buffers);
// }
// static void  glowDeleteFencesNV(GPDELETEFENCESNV fnptr, GLsizei  n, const GLuint * fences) {
//   (*fnptr)(n, fences);
// }
// static void  glowDeleteFramebuffers(GPDELETEFRAMEBUFFERS fnptr, GLsizei  n, const GLuint * framebuffers) {
//   (*fnptr)(n, framebuffers);
// }
// static void  glowDeleteMemoryObjectsEXT(GPDELETEMEMORYOBJECTSEXT fnptr, GLsizei  n, const GLuint * memoryObjects) {
//   (*fnptr)(n, memoryObjects);
// }
// static void  glowDeletePathsNV(GPDELETEPATHSNV fnptr, GLuint  path, GLsizei  range) {
//   (*fnptr)(path, range);
// }
// static void  glowDeletePerfMonitorsAMD(GPDELETEPERFMONITORSAMD fnptr, GLsizei  n, GLuint * monitors) {
//   (*fnptr)(n, monitors);
// }
// static void  glowDeletePerfQueryINTEL(GPDELETEPERFQUERYINTEL fnptr, GLuint  queryHandle) {
//   (*fnptr)(queryHandle);
// }
// static void  glowDeleteProgram(GPDELETEPROGRAM fnptr, GLuint  program) {
//   (*fnptr)(program);
// }
// static void  glowDeleteProgramPipelines(GPDELETEPROGRAMPIPELINES fnptr, GLsizei  n, const GLuint * pipelines) {
//   (*fnptr)(n, pipelines);
// }
// static void  glowDeleteProgramPipelinesEXT(GPDELETEPROGRAMPIPELINESEXT fnptr, GLsizei  n, const GLuint * pipelines) {
//   (*fnptr)(n, pipelines);
// }
// static void  glowDeleteQueries(GPDELETEQUERIES fnptr, GLsizei  n, const GLuint * ids) {
//   (*fnptr)(n, ids);
// }
// static void  glowDeleteQueriesEXT(GPDELETEQUERIESEXT fnptr, GLsizei  n, const GLuint * ids) {
//   (*fnptr)(n, ids);
// }
// static void  glowDeleteRenderbuffers(GPDELETERENDERBUFFERS fnptr, GLsizei  n, const GLuint * renderbuffers) {
//   (*fnptr)(n, renderbuffers);
// }
// static void  glowDeleteSamplers(GPDELETESAMPLERS fnptr, GLsizei  count, const GLuint * samplers) {
//   (*fnptr)(count, samplers);
// }
// static void  glowDeleteSemaphoresEXT(GPDELETESEMAPHORESEXT fnptr, GLsizei  n, const GLuint * semaphores) {
//   (*fnptr)(n, semaphores);
// }
// static void  glowDeleteShader(GPDELETESHADER fnptr, GLuint  shader) {
//   (*fnptr)(shader);
// }
// static void  glowDeleteSync(GPDELETESYNC fnptr, GLsync  sync) {
//   (*fnptr)(sync);
// }
// static void  glowDeleteSyncAPPLE(GPDELETESYNCAPPLE fnptr, GLsync  sync) {
//   (*fnptr)(sync);
// }
// static void  glowDeleteTextures(GPDELETETEXTURES fnptr, GLsizei  n, const GLuint * textures) {
//   (*fnptr)(n, textures);
// }
// static void  glowDeleteTransformFeedbacks(GPDELETETRANSFORMFEEDBACKS fnptr, GLsizei  n, const GLuint * ids) {
//   (*fnptr)(n, ids);
// }
// static void  glowDeleteVertexArrays(GPDELETEVERTEXARRAYS fnptr, GLsizei  n, const GLuint * arrays) {
//   (*fnptr)(n, arrays);
// }
// static void  glowDeleteVertexArraysOES(GPDELETEVERTEXARRAYSOES fnptr, GLsizei  n, const GLuint * arrays) {
//   (*fnptr)(n, arrays);
// }
// static void  glowDepthFunc(GPDEPTHFUNC fnptr, GLenum  func) {
//   (*fnptr)(func);
// }
// static void  glowDepthMask(GPDEPTHMASK fnptr, GLboolean  flag) {
//   (*fnptr)(flag);
// }
// static void  glowDepthRangeArrayfvNV(GPDEPTHRANGEARRAYFVNV fnptr, GLuint  first, GLsizei  count, const GLfloat * v) {
//   (*fnptr)(first, count, v);
// }
// static void  glowDepthRangeArrayfvOES(GPDEPTHRANGEARRAYFVOES fnptr, GLuint  first, GLsizei  count, const GLfloat * v) {
//   (*fnptr)(first, count, v);
// }
// static void  glowDepthRangeIndexedfNV(GPDEPTHRANGEINDEXEDFNV fnptr, GLuint  index, GLfloat  n, GLfloat  f) {
//   (*fnptr)(index, n, f);
// }
// static void  glowDepthRangeIndexedfOES(GPDEPTHRANGEINDEXEDFOES fnptr, GLuint  index, GLfloat  n, GLfloat  f) {
//   (*fnptr)(index, n, f);
// }
// static void  glowDepthRangef(GPDEPTHRANGEF fnptr, GLfloat  n, GLfloat  f) {
//   (*fnptr)(n, f);
// }
// static void  glowDetachShader(GPDETACHSHADER fnptr, GLuint  program, GLuint  shader) {
//   (*fnptr)(program, shader);
// }
// static void  glowDisable(GPDISABLE fnptr, GLenum  cap) {
//   (*fnptr)(cap);
// }
// static void  glowDisableDriverControlQCOM(GPDISABLEDRIVERCONTROLQCOM fnptr, GLuint  driverControl) {
//   (*fnptr)(driverControl);
// }
// static void  glowDisableVertexAttribArray(GPDISABLEVERTEXATTRIBARRAY fnptr, GLuint  index) {
//   (*fnptr)(index);
// }
// static void  glowDisableiEXT(GPDISABLEIEXT fnptr, GLenum  target, GLuint  index) {
//   (*fnptr)(target, index);
// }
// static void  glowDisableiNV(GPDISABLEINV fnptr, GLenum  target, GLuint  index) {
//   (*fnptr)(target, index);
// }
// static void  glowDisableiOES(GPDISABLEIOES fnptr, GLenum  target, GLuint  index) {
//   (*fnptr)(target, index);
// }
// static void  glowDiscardFramebufferEXT(GPDISCARDFRAMEBUFFEREXT fnptr, GLenum  target, GLsizei  numAttachments, const GLenum * attachments) {
//   (*fnptr)(target, numAttachments, attachments);
// }
// static void  glowDispatchCompute(GPDISPATCHCOMPUTE fnptr, GLuint  num_groups_x, GLuint  num_groups_y, GLuint  num_groups_z) {
//   (*fnptr)(num_groups_x, num_groups_y, num_groups_z);
// }
// static void  glowDispatchComputeIndirect(GPDISPATCHCOMPUTEINDIRECT fnptr, GLintptr  indirect) {
//   (*fnptr)(indirect);
// }
// static void  glowDrawArrays(GPDRAWARRAYS fnptr, GLenum  mode, GLint  first, GLsizei  count) {
//   (*fnptr)(mode, first, count);
// }
// static void  glowDrawArraysIndirect(GPDRAWARRAYSINDIRECT fnptr, GLenum  mode, const void * indirect) {
//   (*fnptr)(mode, indirect);
// }
// static void  glowDrawArraysInstanced(GPDRAWARRAYSINSTANCED fnptr, GLenum  mode, GLint  first, GLsizei  count, GLsizei  instancecount) {
//   (*fnptr)(mode, first, count, instancecount);
// }
// static void  glowDrawArraysInstancedANGLE(GPDRAWARRAYSINSTANCEDANGLE fnptr, GLenum  mode, GLint  first, GLsizei  count, GLsizei  primcount) {
//   (*fnptr)(mode, first, count, primcount);
// }
// static void  glowDrawArraysInstancedBaseInstanceEXT(GPDRAWARRAYSINSTANCEDBASEINSTANCEEXT fnptr, GLenum  mode, GLint  first, GLsizei  count, GLsizei  instancecount, GLuint  baseinstance) {
//   (*fnptr)(mode, first, count, instancecount, baseinstance);
// }
// static void  glowDrawArraysInstancedEXT(GPDRAWARRAYSINSTANCEDEXT fnptr, GLenum  mode, GLint  start, GLsizei  count, GLsizei  primcount) {
//   (*fnptr)(mode, start, count, primcount);
// }
// static void  glowDrawArraysInstancedNV(GPDRAWARRAYSINSTANCEDNV fnptr, GLenum  mode, GLint  first, GLsizei  count, GLsizei  primcount) {
//   (*fnptr)(mode, first, count, primcount);
// }
// static void  glowDrawBuffers(GPDRAWBUFFERS fnptr, GLsizei  n, const GLenum * bufs) {
//   (*fnptr)(n, bufs);
// }
// static void  glowDrawBuffersEXT(GPDRAWBUFFERSEXT fnptr, GLsizei  n, const GLenum * bufs) {
//   (*fnptr)(n, bufs);
// }
// static void  glowDrawBuffersIndexedEXT(GPDRAWBUFFERSINDEXEDEXT fnptr, GLint  n, const GLenum * location, const GLint * indices) {
//   (*fnptr)(n, location, indices);
// }
// static void  glowDrawBuffersNV(GPDRAWBUFFERSNV fnptr, GLsizei  n, const GLenum * bufs) {
//   (*fnptr)(n, bufs);
// }
// static void  glowDrawElements(GPDRAWELEMENTS fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices) {
//   (*fnptr)(mode, count, type, indices);
// }
// static void  glowDrawElementsBaseVertexEXT(GPDRAWELEMENTSBASEVERTEXEXT fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLint  basevertex) {
//   (*fnptr)(mode, count, type, indices, basevertex);
// }
// static void  glowDrawElementsBaseVertexOES(GPDRAWELEMENTSBASEVERTEXOES fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLint  basevertex) {
//   (*fnptr)(mode, count, type, indices, basevertex);
// }
// static void  glowDrawElementsIndirect(GPDRAWELEMENTSINDIRECT fnptr, GLenum  mode, GLenum  type, const void * indirect) {
//   (*fnptr)(mode, type, indirect);
// }
// static void  glowDrawElementsInstanced(GPDRAWELEMENTSINSTANCED fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount) {
//   (*fnptr)(mode, count, type, indices, instancecount);
// }
// static void  glowDrawElementsInstancedANGLE(GPDRAWELEMENTSINSTANCEDANGLE fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  primcount) {
//   (*fnptr)(mode, count, type, indices, primcount);
// }
// static void  glowDrawElementsInstancedBaseInstanceEXT(GPDRAWELEMENTSINSTANCEDBASEINSTANCEEXT fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLuint  baseinstance) {
//   (*fnptr)(mode, count, type, indices, instancecount, baseinstance);
// }
// static void  glowDrawElementsInstancedBaseVertexBaseInstanceEXT(GPDRAWELEMENTSINSTANCEDBASEVERTEXBASEINSTANCEEXT fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLint  basevertex, GLuint  baseinstance) {
//   (*fnptr)(mode, count, type, indices, instancecount, basevertex, baseinstance);
// }
// static void  glowDrawElementsInstancedBaseVertexEXT(GPDRAWELEMENTSINSTANCEDBASEVERTEXEXT fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLint  basevertex) {
//   (*fnptr)(mode, count, type, indices, instancecount, basevertex);
// }
// static void  glowDrawElementsInstancedBaseVertexOES(GPDRAWELEMENTSINSTANCEDBASEVERTEXOES fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLint  basevertex) {
//   (*fnptr)(mode, count, type, indices, instancecount, basevertex);
// }
// static void  glowDrawElementsInstancedEXT(GPDRAWELEMENTSINSTANCEDEXT fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  primcount) {
//   (*fnptr)(mode, count, type, indices, primcount);
// }
// static void  glowDrawElementsInstancedNV(GPDRAWELEMENTSINSTANCEDNV fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  primcount) {
//   (*fnptr)(mode, count, type, indices, primcount);
// }
// static void  glowDrawRangeElements(GPDRAWRANGEELEMENTS fnptr, GLenum  mode, GLuint  start, GLuint  end, GLsizei  count, GLenum  type, const void * indices) {
//   (*fnptr)(mode, start, end, count, type, indices);
// }
// static void  glowDrawRangeElementsBaseVertexEXT(GPDRAWRANGEELEMENTSBASEVERTEXEXT fnptr, GLenum  mode, GLuint  start, GLuint  end, GLsizei  count, GLenum  type, const void * indices, GLint  basevertex) {
//   (*fnptr)(mode, start, end, count, type, indices, basevertex);
// }
// static void  glowDrawRangeElementsBaseVertexOES(GPDRAWRANGEELEMENTSBASEVERTEXOES fnptr, GLenum  mode, GLuint  start, GLuint  end, GLsizei  count, GLenum  type, const void * indices, GLint  basevertex) {
//   (*fnptr)(mode, start, end, count, type, indices, basevertex);
// }
// static void  glowDrawTransformFeedbackEXT(GPDRAWTRANSFORMFEEDBACKEXT fnptr, GLenum  mode, GLuint  id) {
//   (*fnptr)(mode, id);
// }
// static void  glowDrawTransformFeedbackInstancedEXT(GPDRAWTRANSFORMFEEDBACKINSTANCEDEXT fnptr, GLenum  mode, GLuint  id, GLsizei  instancecount) {
//   (*fnptr)(mode, id, instancecount);
// }
// static void  glowDrawVkImageNV(GPDRAWVKIMAGENV fnptr, GLuint64  vkImage, GLuint  sampler, GLfloat  x0, GLfloat  y0, GLfloat  x1, GLfloat  y1, GLfloat  z, GLfloat  s0, GLfloat  t0, GLfloat  s1, GLfloat  t1) {
//   (*fnptr)(vkImage, sampler, x0, y0, x1, y1, z, s0, t0, s1, t1);
// }
// static void  glowEGLImageTargetRenderbufferStorageOES(GPEGLIMAGETARGETRENDERBUFFERSTORAGEOES fnptr, GLenum  target, GLeglImageOES  image) {
//   (*fnptr)(target, image);
// }
// static void  glowEGLImageTargetTexStorageEXT(GPEGLIMAGETARGETTEXSTORAGEEXT fnptr, GLenum  target, GLeglImageOES  image, const GLint*  attrib_list) {
//   (*fnptr)(target, image, attrib_list);
// }
// static void  glowEGLImageTargetTexture2DOES(GPEGLIMAGETARGETTEXTURE2DOES fnptr, GLenum  target, GLeglImageOES  image) {
//   (*fnptr)(target, image);
// }
// static void  glowEGLImageTargetTextureStorageEXT(GPEGLIMAGETARGETTEXTURESTORAGEEXT fnptr, GLuint  texture, GLeglImageOES  image, const GLint*  attrib_list) {
//   (*fnptr)(texture, image, attrib_list);
// }
// static void  glowEnable(GPENABLE fnptr, GLenum  cap) {
//   (*fnptr)(cap);
// }
// static void  glowEnableDriverControlQCOM(GPENABLEDRIVERCONTROLQCOM fnptr, GLuint  driverControl) {
//   (*fnptr)(driverControl);
// }
// static void  glowEnableVertexAttribArray(GPENABLEVERTEXATTRIBARRAY fnptr, GLuint  index) {
//   (*fnptr)(index);
// }
// static void  glowEnableiEXT(GPENABLEIEXT fnptr, GLenum  target, GLuint  index) {
//   (*fnptr)(target, index);
// }
// static void  glowEnableiNV(GPENABLEINV fnptr, GLenum  target, GLuint  index) {
//   (*fnptr)(target, index);
// }
// static void  glowEnableiOES(GPENABLEIOES fnptr, GLenum  target, GLuint  index) {
//   (*fnptr)(target, index);
// }
// static void  glowEndConditionalRenderNV(GPENDCONDITIONALRENDERNV fnptr) {
//   (*fnptr)();
// }
// static void  glowEndPerfMonitorAMD(GPENDPERFMONITORAMD fnptr, GLuint  monitor) {
//   (*fnptr)(monitor);
// }
// static void  glowEndPerfQueryINTEL(GPENDPERFQUERYINTEL fnptr, GLuint  queryHandle) {
//   (*fnptr)(queryHandle);
// }
// static void  glowEndQuery(GPENDQUERY fnptr, GLenum  target) {
//   (*fnptr)(target);
// }
// static void  glowEndQueryEXT(GPENDQUERYEXT fnptr, GLenum  target) {
//   (*fnptr)(target);
// }
// static void  glowEndTilingQCOM(GPENDTILINGQCOM fnptr, GLbitfield  preserveMask) {
//   (*fnptr)(preserveMask);
// }
// static void  glowEndTransformFeedback(GPENDTRANSFORMFEEDBACK fnptr) {
//   (*fnptr)();
// }
// static void  glowExtGetBufferPointervQCOM(GPEXTGETBUFFERPOINTERVQCOM fnptr, GLenum  target, void ** params) {
//   (*fnptr)(target, params);
// }
// static void  glowExtGetBuffersQCOM(GPEXTGETBUFFERSQCOM fnptr, GLuint * buffers, GLint  maxBuffers, GLint * numBuffers) {
//   (*fnptr)(buffers, maxBuffers, numBuffers);
// }
// static void  glowExtGetFramebuffersQCOM(GPEXTGETFRAMEBUFFERSQCOM fnptr, GLuint * framebuffers, GLint  maxFramebuffers, GLint * numFramebuffers) {
//   (*fnptr)(framebuffers, maxFramebuffers, numFramebuffers);
// }
// static void  glowExtGetProgramBinarySourceQCOM(GPEXTGETPROGRAMBINARYSOURCEQCOM fnptr, GLuint  program, GLenum  shadertype, GLchar * source, GLint * length) {
//   (*fnptr)(program, shadertype, source, length);
// }
// static void  glowExtGetProgramsQCOM(GPEXTGETPROGRAMSQCOM fnptr, GLuint * programs, GLint  maxPrograms, GLint * numPrograms) {
//   (*fnptr)(programs, maxPrograms, numPrograms);
// }
// static void  glowExtGetRenderbuffersQCOM(GPEXTGETRENDERBUFFERSQCOM fnptr, GLuint * renderbuffers, GLint  maxRenderbuffers, GLint * numRenderbuffers) {
//   (*fnptr)(renderbuffers, maxRenderbuffers, numRenderbuffers);
// }
// static void  glowExtGetShadersQCOM(GPEXTGETSHADERSQCOM fnptr, GLuint * shaders, GLint  maxShaders, GLint * numShaders) {
//   (*fnptr)(shaders, maxShaders, numShaders);
// }
// static void  glowExtGetTexLevelParameterivQCOM(GPEXTGETTEXLEVELPARAMETERIVQCOM fnptr, GLuint  texture, GLenum  face, GLint  level, GLenum  pname, GLint * params) {
//   (*fnptr)(texture, face, level, pname, params);
// }
// static void  glowExtGetTexSubImageQCOM(GPEXTGETTEXSUBIMAGEQCOM fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, void * texels) {
//   (*fnptr)(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, texels);
// }
// static void  glowExtGetTexturesQCOM(GPEXTGETTEXTURESQCOM fnptr, GLuint * textures, GLint  maxTextures, GLint * numTextures) {
//   (*fnptr)(textures, maxTextures, numTextures);
// }
// static GLboolean  glowExtIsProgramBinaryQCOM(GPEXTISPROGRAMBINARYQCOM fnptr, GLuint  program) {
//   return (*fnptr)(program);
// }
// static void  glowExtTexObjectStateOverrideiQCOM(GPEXTTEXOBJECTSTATEOVERRIDEIQCOM fnptr, GLenum  target, GLenum  pname, GLint  param) {
//   (*fnptr)(target, pname, param);
// }
// static GLsync  glowFenceSync(GPFENCESYNC fnptr, GLenum  condition, GLbitfield  flags) {
//   return (*fnptr)(condition, flags);
// }
// static GLsync  glowFenceSyncAPPLE(GPFENCESYNCAPPLE fnptr, GLenum  condition, GLbitfield  flags) {
//   return (*fnptr)(condition, flags);
// }
// static void  glowFinish(GPFINISH fnptr) {
//   (*fnptr)();
// }
// static void  glowFinishFenceNV(GPFINISHFENCENV fnptr, GLuint  fence) {
//   (*fnptr)(fence);
// }
// static void  glowFlush(GPFLUSH fnptr) {
//   (*fnptr)();
// }
// static void  glowFlushMappedBufferRange(GPFLUSHMAPPEDBUFFERRANGE fnptr, GLenum  target, GLintptr  offset, GLsizeiptr  length) {
//   (*fnptr)(target, offset, length);
// }
// static void  glowFlushMappedBufferRangeEXT(GPFLUSHMAPPEDBUFFERRANGEEXT fnptr, GLenum  target, GLintptr  offset, GLsizeiptr  length) {
//   (*fnptr)(target, offset, length);
// }
// static void  glowFragmentCoverageColorNV(GPFRAGMENTCOVERAGECOLORNV fnptr, GLuint  color) {
//   (*fnptr)(color);
// }
// static void  glowFramebufferFetchBarrierEXT(GPFRAMEBUFFERFETCHBARRIEREXT fnptr) {
//   (*fnptr)();
// }
// static void  glowFramebufferFetchBarrierQCOM(GPFRAMEBUFFERFETCHBARRIERQCOM fnptr) {
//   (*fnptr)();
// }
// static void  glowFramebufferFoveationConfigQCOM(GPFRAMEBUFFERFOVEATIONCONFIGQCOM fnptr, GLuint  framebuffer, GLuint  numLayers, GLuint  focalPointsPerLayer, GLuint  requestedFeatures, GLuint * providedFeatures) {
//   (*fnptr)(framebuffer, numLayers, focalPointsPerLayer, requestedFeatures, providedFeatures);
// }
// static void  glowFramebufferFoveationParametersQCOM(GPFRAMEBUFFERFOVEATIONPARAMETERSQCOM fnptr, GLuint  framebuffer, GLuint  layer, GLuint  focalPoint, GLfloat  focalX, GLfloat  focalY, GLfloat  gainX, GLfloat  gainY, GLfloat  foveaArea) {
//   (*fnptr)(framebuffer, layer, focalPoint, focalX, focalY, gainX, gainY, foveaArea);
// }
// static void  glowFramebufferParameteri(GPFRAMEBUFFERPARAMETERI fnptr, GLenum  target, GLenum  pname, GLint  param) {
//   (*fnptr)(target, pname, param);
// }
// static void  glowFramebufferPixelLocalStorageSizeEXT(GPFRAMEBUFFERPIXELLOCALSTORAGESIZEEXT fnptr, GLuint  target, GLsizei  size) {
//   (*fnptr)(target, size);
// }
// static void  glowFramebufferRenderbuffer(GPFRAMEBUFFERRENDERBUFFER fnptr, GLenum  target, GLenum  attachment, GLenum  renderbuffertarget, GLuint  renderbuffer) {
//   (*fnptr)(target, attachment, renderbuffertarget, renderbuffer);
// }
// static void  glowFramebufferSampleLocationsfvNV(GPFRAMEBUFFERSAMPLELOCATIONSFVNV fnptr, GLenum  target, GLuint  start, GLsizei  count, const GLfloat * v) {
//   (*fnptr)(target, start, count, v);
// }
// static void  glowFramebufferTexture2D(GPFRAMEBUFFERTEXTURE2D fnptr, GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level) {
//   (*fnptr)(target, attachment, textarget, texture, level);
// }
// static void  glowFramebufferTexture2DDownsampleIMG(GPFRAMEBUFFERTEXTURE2DDOWNSAMPLEIMG fnptr, GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level, GLint  xscale, GLint  yscale) {
//   (*fnptr)(target, attachment, textarget, texture, level, xscale, yscale);
// }
// static void  glowFramebufferTexture2DMultisampleEXT(GPFRAMEBUFFERTEXTURE2DMULTISAMPLEEXT fnptr, GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level, GLsizei  samples) {
//   (*fnptr)(target, attachment, textarget, texture, level, samples);
// }
// static void  glowFramebufferTexture2DMultisampleIMG(GPFRAMEBUFFERTEXTURE2DMULTISAMPLEIMG fnptr, GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level, GLsizei  samples) {
//   (*fnptr)(target, attachment, textarget, texture, level, samples);
// }
// static void  glowFramebufferTexture3DOES(GPFRAMEBUFFERTEXTURE3DOES fnptr, GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level, GLint  zoffset) {
//   (*fnptr)(target, attachment, textarget, texture, level, zoffset);
// }
// static void  glowFramebufferTextureEXT(GPFRAMEBUFFERTEXTUREEXT fnptr, GLenum  target, GLenum  attachment, GLuint  texture, GLint  level) {
//   (*fnptr)(target, attachment, texture, level);
// }
// static void  glowFramebufferTextureLayer(GPFRAMEBUFFERTEXTURELAYER fnptr, GLenum  target, GLenum  attachment, GLuint  texture, GLint  level, GLint  layer) {
//   (*fnptr)(target, attachment, texture, level, layer);
// }
// static void  glowFramebufferTextureLayerDownsampleIMG(GPFRAMEBUFFERTEXTURELAYERDOWNSAMPLEIMG fnptr, GLenum  target, GLenum  attachment, GLuint  texture, GLint  level, GLint  layer, GLint  xscale, GLint  yscale) {
//   (*fnptr)(target, attachment, texture, level, layer, xscale, yscale);
// }
// static void  glowFramebufferTextureMultisampleMultiviewOVR(GPFRAMEBUFFERTEXTUREMULTISAMPLEMULTIVIEWOVR fnptr, GLenum  target, GLenum  attachment, GLuint  texture, GLint  level, GLsizei  samples, GLint  baseViewIndex, GLsizei  numViews) {
//   (*fnptr)(target, attachment, texture, level, samples, baseViewIndex, numViews);
// }
// static void  glowFramebufferTextureMultiviewOVR(GPFRAMEBUFFERTEXTUREMULTIVIEWOVR fnptr, GLenum  target, GLenum  attachment, GLuint  texture, GLint  level, GLint  baseViewIndex, GLsizei  numViews) {
//   (*fnptr)(target, attachment, texture, level, baseViewIndex, numViews);
// }
// static void  glowFramebufferTextureOES(GPFRAMEBUFFERTEXTUREOES fnptr, GLenum  target, GLenum  attachment, GLuint  texture, GLint  level) {
//   (*fnptr)(target, attachment, texture, level);
// }
// static void  glowFrontFace(GPFRONTFACE fnptr, GLenum  mode) {
//   (*fnptr)(mode);
// }
// static void  glowGenBuffers(GPGENBUFFERS fnptr, GLsizei  n, GLuint * buffers) {
//   (*fnptr)(n, buffers);
// }
// static void  glowGenFencesNV(GPGENFENCESNV fnptr, GLsizei  n, GLuint * fences) {
//   (*fnptr)(n, fences);
// }
// static void  glowGenFramebuffers(GPGENFRAMEBUFFERS fnptr, GLsizei  n, GLuint * framebuffers) {
//   (*fnptr)(n, framebuffers);
// }
// static GLuint  glowGenPathsNV(GPGENPATHSNV fnptr, GLsizei  range) {
//   return (*fnptr)(range);
// }
// static void  glowGenPerfMonitorsAMD(GPGENPERFMONITORSAMD fnptr, GLsizei  n, GLuint * monitors) {
//   (*fnptr)(n, monitors);
// }
// static void  glowGenProgramPipelines(GPGENPROGRAMPIPELINES fnptr, GLsizei  n, GLuint * pipelines) {
//   (*fnptr)(n, pipelines);
// }
// static void  glowGenProgramPipelinesEXT(GPGENPROGRAMPIPELINESEXT fnptr, GLsizei  n, GLuint * pipelines) {
//   (*fnptr)(n, pipelines);
// }
// static void  glowGenQueries(GPGENQUERIES fnptr, GLsizei  n, GLuint * ids) {
//   (*fnptr)(n, ids);
// }
// static void  glowGenQueriesEXT(GPGENQUERIESEXT fnptr, GLsizei  n, GLuint * ids) {
//   (*fnptr)(n, ids);
// }
// static void  glowGenRenderbuffers(GPGENRENDERBUFFERS fnptr, GLsizei  n, GLuint * renderbuffers) {
//   (*fnptr)(n, renderbuffers);
// }
// static void  glowGenSamplers(GPGENSAMPLERS fnptr, GLsizei  count, GLuint * samplers) {
//   (*fnptr)(count, samplers);
// }
// static void  glowGenSemaphoresEXT(GPGENSEMAPHORESEXT fnptr, GLsizei  n, GLuint * semaphores) {
//   (*fnptr)(n, semaphores);
// }
// static void  glowGenTextures(GPGENTEXTURES fnptr, GLsizei  n, GLuint * textures) {
//   (*fnptr)(n, textures);
// }
// static void  glowGenTransformFeedbacks(GPGENTRANSFORMFEEDBACKS fnptr, GLsizei  n, GLuint * ids) {
//   (*fnptr)(n, ids);
// }
// static void  glowGenVertexArrays(GPGENVERTEXARRAYS fnptr, GLsizei  n, GLuint * arrays) {
//   (*fnptr)(n, arrays);
// }
// static void  glowGenVertexArraysOES(GPGENVERTEXARRAYSOES fnptr, GLsizei  n, GLuint * arrays) {
//   (*fnptr)(n, arrays);
// }
// static void  glowGenerateMipmap(GPGENERATEMIPMAP fnptr, GLenum  target) {
//   (*fnptr)(target);
// }
// static void  glowGetActiveAttrib(GPGETACTIVEATTRIB fnptr, GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLint * size, GLenum * type, GLchar * name) {
//   (*fnptr)(program, index, bufSize, length, size, type, name);
// }
// static void  glowGetActiveUniform(GPGETACTIVEUNIFORM fnptr, GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLint * size, GLenum * type, GLchar * name) {
//   (*fnptr)(program, index, bufSize, length, size, type, name);
// }
// static void  glowGetActiveUniformBlockName(GPGETACTIVEUNIFORMBLOCKNAME fnptr, GLuint  program, GLuint  uniformBlockIndex, GLsizei  bufSize, GLsizei * length, GLchar * uniformBlockName) {
//   (*fnptr)(program, uniformBlockIndex, bufSize, length, uniformBlockName);
// }
// static void  glowGetActiveUniformBlockiv(GPGETACTIVEUNIFORMBLOCKIV fnptr, GLuint  program, GLuint  uniformBlockIndex, GLenum  pname, GLint * params) {
//   (*fnptr)(program, uniformBlockIndex, pname, params);
// }
// static void  glowGetActiveUniformsiv(GPGETACTIVEUNIFORMSIV fnptr, GLuint  program, GLsizei  uniformCount, const GLuint * uniformIndices, GLenum  pname, GLint * params) {
//   (*fnptr)(program, uniformCount, uniformIndices, pname, params);
// }
// static void  glowGetAttachedShaders(GPGETATTACHEDSHADERS fnptr, GLuint  program, GLsizei  maxCount, GLsizei * count, GLuint * shaders) {
//   (*fnptr)(program, maxCount, count, shaders);
// }
// static GLint  glowGetAttribLocation(GPGETATTRIBLOCATION fnptr, GLuint  program, const GLchar * name) {
//   return (*fnptr)(program, name);
// }
// static void  glowGetBooleani_v(GPGETBOOLEANI_V fnptr, GLenum  target, GLuint  index, GLboolean * data) {
//   (*fnptr)(target, index, data);
// }
// static void  glowGetBooleanv(GPGETBOOLEANV fnptr, GLenum  pname, GLboolean * data) {
//   (*fnptr)(pname, data);
// }
// static void  glowGetBufferParameteri64v(GPGETBUFFERPARAMETERI64V fnptr, GLenum  target, GLenum  pname, GLint64 * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetBufferParameteriv(GPGETBUFFERPARAMETERIV fnptr, GLenum  target, GLenum  pname, GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetBufferPointerv(GPGETBUFFERPOINTERV fnptr, GLenum  target, GLenum  pname, void ** params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetBufferPointervOES(GPGETBUFFERPOINTERVOES fnptr, GLenum  target, GLenum  pname, void ** params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetCoverageModulationTableNV(GPGETCOVERAGEMODULATIONTABLENV fnptr, GLsizei  bufsize, GLfloat * v) {
//   (*fnptr)(bufsize, v);
// }
// static GLuint  glowGetDebugMessageLog(GPGETDEBUGMESSAGELOG fnptr, GLuint  count, GLsizei  bufSize, GLenum * sources, GLenum * types, GLuint * ids, GLenum * severities, GLsizei * lengths, GLchar * messageLog) {
//   return (*fnptr)(count, bufSize, sources, types, ids, severities, lengths, messageLog);
// }
// static GLuint  glowGetDebugMessageLogKHR(GPGETDEBUGMESSAGELOGKHR fnptr, GLuint  count, GLsizei  bufSize, GLenum * sources, GLenum * types, GLuint * ids, GLenum * severities, GLsizei * lengths, GLchar * messageLog) {
//   return (*fnptr)(count, bufSize, sources, types, ids, severities, lengths, messageLog);
// }
// static void  glowGetDriverControlStringQCOM(GPGETDRIVERCONTROLSTRINGQCOM fnptr, GLuint  driverControl, GLsizei  bufSize, GLsizei * length, GLchar * driverControlString) {
//   (*fnptr)(driverControl, bufSize, length, driverControlString);
// }
// static void  glowGetDriverControlsQCOM(GPGETDRIVERCONTROLSQCOM fnptr, GLint * num, GLsizei  size, GLuint * driverControls) {
//   (*fnptr)(num, size, driverControls);
// }
// static GLenum  glowGetError(GPGETERROR fnptr) {
//   return (*fnptr)();
// }
// static void  glowGetFenceivNV(GPGETFENCEIVNV fnptr, GLuint  fence, GLenum  pname, GLint * params) {
//   (*fnptr)(fence, pname, params);
// }
// static void  glowGetFirstPerfQueryIdINTEL(GPGETFIRSTPERFQUERYIDINTEL fnptr, GLuint * queryId) {
//   (*fnptr)(queryId);
// }
// static void  glowGetFloati_vNV(GPGETFLOATI_VNV fnptr, GLenum  target, GLuint  index, GLfloat * data) {
//   (*fnptr)(target, index, data);
// }
// static void  glowGetFloati_vOES(GPGETFLOATI_VOES fnptr, GLenum  target, GLuint  index, GLfloat * data) {
//   (*fnptr)(target, index, data);
// }
// static void  glowGetFloatv(GPGETFLOATV fnptr, GLenum  pname, GLfloat * data) {
//   (*fnptr)(pname, data);
// }
// static GLint  glowGetFragDataIndexEXT(GPGETFRAGDATAINDEXEXT fnptr, GLuint  program, const GLchar * name) {
//   return (*fnptr)(program, name);
// }
// static GLint  glowGetFragDataLocation(GPGETFRAGDATALOCATION fnptr, GLuint  program, const GLchar * name) {
//   return (*fnptr)(program, name);
// }
// static void  glowGetFramebufferAttachmentParameteriv(GPGETFRAMEBUFFERATTACHMENTPARAMETERIV fnptr, GLenum  target, GLenum  attachment, GLenum  pname, GLint * params) {
//   (*fnptr)(target, attachment, pname, params);
// }
// static void  glowGetFramebufferParameteriv(GPGETFRAMEBUFFERPARAMETERIV fnptr, GLenum  target, GLenum  pname, GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static GLsizei  glowGetFramebufferPixelLocalStorageSizeEXT(GPGETFRAMEBUFFERPIXELLOCALSTORAGESIZEEXT fnptr, GLuint  target) {
//   return (*fnptr)(target);
// }
// static GLenum  glowGetGraphicsResetStatus(GPGETGRAPHICSRESETSTATUS fnptr) {
//   return (*fnptr)();
// }
// static GLenum  glowGetGraphicsResetStatusEXT(GPGETGRAPHICSRESETSTATUSEXT fnptr) {
//   return (*fnptr)();
// }
// static GLenum  glowGetGraphicsResetStatusKHR(GPGETGRAPHICSRESETSTATUSKHR fnptr) {
//   return (*fnptr)();
// }
// static GLuint64  glowGetImageHandleNV(GPGETIMAGEHANDLENV fnptr, GLuint  texture, GLint  level, GLboolean  layered, GLint  layer, GLenum  format) {
//   return (*fnptr)(texture, level, layered, layer, format);
// }
// static void  glowGetInteger64i_v(GPGETINTEGER64I_V fnptr, GLenum  target, GLuint  index, GLint64 * data) {
//   (*fnptr)(target, index, data);
// }
// static void  glowGetInteger64v(GPGETINTEGER64V fnptr, GLenum  pname, GLint64 * data) {
//   (*fnptr)(pname, data);
// }
// static void  glowGetInteger64vAPPLE(GPGETINTEGER64VAPPLE fnptr, GLenum  pname, GLint64 * params) {
//   (*fnptr)(pname, params);
// }
// static void  glowGetIntegeri_v(GPGETINTEGERI_V fnptr, GLenum  target, GLuint  index, GLint * data) {
//   (*fnptr)(target, index, data);
// }
// static void  glowGetIntegeri_vEXT(GPGETINTEGERI_VEXT fnptr, GLenum  target, GLuint  index, GLint * data) {
//   (*fnptr)(target, index, data);
// }
// static void  glowGetIntegerv(GPGETINTEGERV fnptr, GLenum  pname, GLint * data) {
//   (*fnptr)(pname, data);
// }
// static void  glowGetInternalformatSampleivNV(GPGETINTERNALFORMATSAMPLEIVNV fnptr, GLenum  target, GLenum  internalformat, GLsizei  samples, GLenum  pname, GLsizei  bufSize, GLint * params) {
//   (*fnptr)(target, internalformat, samples, pname, bufSize, params);
// }
// static void  glowGetInternalformativ(GPGETINTERNALFORMATIV fnptr, GLenum  target, GLenum  internalformat, GLenum  pname, GLsizei  bufSize, GLint * params) {
//   (*fnptr)(target, internalformat, pname, bufSize, params);
// }
// static void  glowGetMemoryObjectParameterivEXT(GPGETMEMORYOBJECTPARAMETERIVEXT fnptr, GLuint  memoryObject, GLenum  pname, GLint * params) {
//   (*fnptr)(memoryObject, pname, params);
// }
// static void  glowGetMultisamplefv(GPGETMULTISAMPLEFV fnptr, GLenum  pname, GLuint  index, GLfloat * val) {
//   (*fnptr)(pname, index, val);
// }
// static void  glowGetNextPerfQueryIdINTEL(GPGETNEXTPERFQUERYIDINTEL fnptr, GLuint  queryId, GLuint * nextQueryId) {
//   (*fnptr)(queryId, nextQueryId);
// }
// static void  glowGetObjectLabel(GPGETOBJECTLABEL fnptr, GLenum  identifier, GLuint  name, GLsizei  bufSize, GLsizei * length, GLchar * label) {
//   (*fnptr)(identifier, name, bufSize, length, label);
// }
// static void  glowGetObjectLabelEXT(GPGETOBJECTLABELEXT fnptr, GLenum  type, GLuint  object, GLsizei  bufSize, GLsizei * length, GLchar * label) {
//   (*fnptr)(type, object, bufSize, length, label);
// }
// static void  glowGetObjectLabelKHR(GPGETOBJECTLABELKHR fnptr, GLenum  identifier, GLuint  name, GLsizei  bufSize, GLsizei * length, GLchar * label) {
//   (*fnptr)(identifier, name, bufSize, length, label);
// }
// static void  glowGetObjectPtrLabel(GPGETOBJECTPTRLABEL fnptr, const void * ptr, GLsizei  bufSize, GLsizei * length, GLchar * label) {
//   (*fnptr)(ptr, bufSize, length, label);
// }
// static void  glowGetObjectPtrLabelKHR(GPGETOBJECTPTRLABELKHR fnptr, const void * ptr, GLsizei  bufSize, GLsizei * length, GLchar * label) {
//   (*fnptr)(ptr, bufSize, length, label);
// }
// static void  glowGetPathCommandsNV(GPGETPATHCOMMANDSNV fnptr, GLuint  path, GLubyte * commands) {
//   (*fnptr)(path, commands);
// }
// static void  glowGetPathCoordsNV(GPGETPATHCOORDSNV fnptr, GLuint  path, GLfloat * coords) {
//   (*fnptr)(path, coords);
// }
// static void  glowGetPathDashArrayNV(GPGETPATHDASHARRAYNV fnptr, GLuint  path, GLfloat * dashArray) {
//   (*fnptr)(path, dashArray);
// }
// static GLfloat  glowGetPathLengthNV(GPGETPATHLENGTHNV fnptr, GLuint  path, GLsizei  startSegment, GLsizei  numSegments) {
//   return (*fnptr)(path, startSegment, numSegments);
// }
// static void  glowGetPathMetricRangeNV(GPGETPATHMETRICRANGENV fnptr, GLbitfield  metricQueryMask, GLuint  firstPathName, GLsizei  numPaths, GLsizei  stride, GLfloat * metrics) {
//   (*fnptr)(metricQueryMask, firstPathName, numPaths, stride, metrics);
// }
// static void  glowGetPathMetricsNV(GPGETPATHMETRICSNV fnptr, GLbitfield  metricQueryMask, GLsizei  numPaths, GLenum  pathNameType, const void * paths, GLuint  pathBase, GLsizei  stride, GLfloat * metrics) {
//   (*fnptr)(metricQueryMask, numPaths, pathNameType, paths, pathBase, stride, metrics);
// }
// static void  glowGetPathParameterfvNV(GPGETPATHPARAMETERFVNV fnptr, GLuint  path, GLenum  pname, GLfloat * value) {
//   (*fnptr)(path, pname, value);
// }
// static void  glowGetPathParameterivNV(GPGETPATHPARAMETERIVNV fnptr, GLuint  path, GLenum  pname, GLint * value) {
//   (*fnptr)(path, pname, value);
// }
// static void  glowGetPathSpacingNV(GPGETPATHSPACINGNV fnptr, GLenum  pathListMode, GLsizei  numPaths, GLenum  pathNameType, const void * paths, GLuint  pathBase, GLfloat  advanceScale, GLfloat  kerningScale, GLenum  transformType, GLfloat * returnedSpacing) {
//   (*fnptr)(pathListMode, numPaths, pathNameType, paths, pathBase, advanceScale, kerningScale, transformType, returnedSpacing);
// }
// static void  glowGetPerfCounterInfoINTEL(GPGETPERFCOUNTERINFOINTEL fnptr, GLuint  queryId, GLuint  counterId, GLuint  counterNameLength, GLchar * counterName, GLuint  counterDescLength, GLchar * counterDesc, GLuint * counterOffset, GLuint * counterDataSize, GLuint * counterTypeEnum, GLuint * counterDataTypeEnum, GLuint64 * rawCounterMaxValue) {
//   (*fnptr)(queryId, counterId, counterNameLength, counterName, counterDescLength, counterDesc, counterOffset, counterDataSize, counterTypeEnum, counterDataTypeEnum, rawCounterMaxValue);
// }
// static void  glowGetPerfMonitorCounterDataAMD(GPGETPERFMONITORCOUNTERDATAAMD fnptr, GLuint  monitor, GLenum  pname, GLsizei  dataSize, GLuint * data, GLint * bytesWritten) {
//   (*fnptr)(monitor, pname, dataSize, data, bytesWritten);
// }
// static void  glowGetPerfMonitorCounterInfoAMD(GPGETPERFMONITORCOUNTERINFOAMD fnptr, GLuint  group, GLuint  counter, GLenum  pname, void * data) {
//   (*fnptr)(group, counter, pname, data);
// }
// static void  glowGetPerfMonitorCounterStringAMD(GPGETPERFMONITORCOUNTERSTRINGAMD fnptr, GLuint  group, GLuint  counter, GLsizei  bufSize, GLsizei * length, GLchar * counterString) {
//   (*fnptr)(group, counter, bufSize, length, counterString);
// }
// static void  glowGetPerfMonitorCountersAMD(GPGETPERFMONITORCOUNTERSAMD fnptr, GLuint  group, GLint * numCounters, GLint * maxActiveCounters, GLsizei  counterSize, GLuint * counters) {
//   (*fnptr)(group, numCounters, maxActiveCounters, counterSize, counters);
// }
// static void  glowGetPerfMonitorGroupStringAMD(GPGETPERFMONITORGROUPSTRINGAMD fnptr, GLuint  group, GLsizei  bufSize, GLsizei * length, GLchar * groupString) {
//   (*fnptr)(group, bufSize, length, groupString);
// }
// static void  glowGetPerfMonitorGroupsAMD(GPGETPERFMONITORGROUPSAMD fnptr, GLint * numGroups, GLsizei  groupsSize, GLuint * groups) {
//   (*fnptr)(numGroups, groupsSize, groups);
// }
// static void  glowGetPerfQueryDataINTEL(GPGETPERFQUERYDATAINTEL fnptr, GLuint  queryHandle, GLuint  flags, GLsizei  dataSize, void * data, GLuint * bytesWritten) {
//   (*fnptr)(queryHandle, flags, dataSize, data, bytesWritten);
// }
// static void  glowGetPerfQueryIdByNameINTEL(GPGETPERFQUERYIDBYNAMEINTEL fnptr, GLchar * queryName, GLuint * queryId) {
//   (*fnptr)(queryName, queryId);
// }
// static void  glowGetPerfQueryInfoINTEL(GPGETPERFQUERYINFOINTEL fnptr, GLuint  queryId, GLuint  queryNameLength, GLchar * queryName, GLuint * dataSize, GLuint * noCounters, GLuint * noInstances, GLuint * capsMask) {
//   (*fnptr)(queryId, queryNameLength, queryName, dataSize, noCounters, noInstances, capsMask);
// }
// static void  glowGetPointerv(GPGETPOINTERV fnptr, GLenum  pname, void ** params) {
//   (*fnptr)(pname, params);
// }
// static void  glowGetPointervKHR(GPGETPOINTERVKHR fnptr, GLenum  pname, void ** params) {
//   (*fnptr)(pname, params);
// }
// static void  glowGetProgramBinary(GPGETPROGRAMBINARY fnptr, GLuint  program, GLsizei  bufSize, GLsizei * length, GLenum * binaryFormat, void * binary) {
//   (*fnptr)(program, bufSize, length, binaryFormat, binary);
// }
// static void  glowGetProgramBinaryOES(GPGETPROGRAMBINARYOES fnptr, GLuint  program, GLsizei  bufSize, GLsizei * length, GLenum * binaryFormat, void * binary) {
//   (*fnptr)(program, bufSize, length, binaryFormat, binary);
// }
// static void  glowGetProgramInfoLog(GPGETPROGRAMINFOLOG fnptr, GLuint  program, GLsizei  bufSize, GLsizei * length, GLchar * infoLog) {
//   (*fnptr)(program, bufSize, length, infoLog);
// }
// static void  glowGetProgramInterfaceiv(GPGETPROGRAMINTERFACEIV fnptr, GLuint  program, GLenum  programInterface, GLenum  pname, GLint * params) {
//   (*fnptr)(program, programInterface, pname, params);
// }
// static void  glowGetProgramPipelineInfoLog(GPGETPROGRAMPIPELINEINFOLOG fnptr, GLuint  pipeline, GLsizei  bufSize, GLsizei * length, GLchar * infoLog) {
//   (*fnptr)(pipeline, bufSize, length, infoLog);
// }
// static void  glowGetProgramPipelineInfoLogEXT(GPGETPROGRAMPIPELINEINFOLOGEXT fnptr, GLuint  pipeline, GLsizei  bufSize, GLsizei * length, GLchar * infoLog) {
//   (*fnptr)(pipeline, bufSize, length, infoLog);
// }
// static void  glowGetProgramPipelineiv(GPGETPROGRAMPIPELINEIV fnptr, GLuint  pipeline, GLenum  pname, GLint * params) {
//   (*fnptr)(pipeline, pname, params);
// }
// static void  glowGetProgramPipelineivEXT(GPGETPROGRAMPIPELINEIVEXT fnptr, GLuint  pipeline, GLenum  pname, GLint * params) {
//   (*fnptr)(pipeline, pname, params);
// }
// static GLuint  glowGetProgramResourceIndex(GPGETPROGRAMRESOURCEINDEX fnptr, GLuint  program, GLenum  programInterface, const GLchar * name) {
//   return (*fnptr)(program, programInterface, name);
// }
// static GLint  glowGetProgramResourceLocation(GPGETPROGRAMRESOURCELOCATION fnptr, GLuint  program, GLenum  programInterface, const GLchar * name) {
//   return (*fnptr)(program, programInterface, name);
// }
// static GLint  glowGetProgramResourceLocationIndexEXT(GPGETPROGRAMRESOURCELOCATIONINDEXEXT fnptr, GLuint  program, GLenum  programInterface, const GLchar * name) {
//   return (*fnptr)(program, programInterface, name);
// }
// static void  glowGetProgramResourceName(GPGETPROGRAMRESOURCENAME fnptr, GLuint  program, GLenum  programInterface, GLuint  index, GLsizei  bufSize, GLsizei * length, GLchar * name) {
//   (*fnptr)(program, programInterface, index, bufSize, length, name);
// }
// static void  glowGetProgramResourcefvNV(GPGETPROGRAMRESOURCEFVNV fnptr, GLuint  program, GLenum  programInterface, GLuint  index, GLsizei  propCount, const GLenum * props, GLsizei  bufSize, GLsizei * length, GLfloat * params) {
//   (*fnptr)(program, programInterface, index, propCount, props, bufSize, length, params);
// }
// static void  glowGetProgramResourceiv(GPGETPROGRAMRESOURCEIV fnptr, GLuint  program, GLenum  programInterface, GLuint  index, GLsizei  propCount, const GLenum * props, GLsizei  bufSize, GLsizei * length, GLint * params) {
//   (*fnptr)(program, programInterface, index, propCount, props, bufSize, length, params);
// }
// static void  glowGetProgramiv(GPGETPROGRAMIV fnptr, GLuint  program, GLenum  pname, GLint * params) {
//   (*fnptr)(program, pname, params);
// }
// static void  glowGetQueryObjecti64vEXT(GPGETQUERYOBJECTI64VEXT fnptr, GLuint  id, GLenum  pname, GLint64 * params) {
//   (*fnptr)(id, pname, params);
// }
// static void  glowGetQueryObjectivEXT(GPGETQUERYOBJECTIVEXT fnptr, GLuint  id, GLenum  pname, GLint * params) {
//   (*fnptr)(id, pname, params);
// }
// static void  glowGetQueryObjectui64vEXT(GPGETQUERYOBJECTUI64VEXT fnptr, GLuint  id, GLenum  pname, GLuint64 * params) {
//   (*fnptr)(id, pname, params);
// }
// static void  glowGetQueryObjectuiv(GPGETQUERYOBJECTUIV fnptr, GLuint  id, GLenum  pname, GLuint * params) {
//   (*fnptr)(id, pname, params);
// }
// static void  glowGetQueryObjectuivEXT(GPGETQUERYOBJECTUIVEXT fnptr, GLuint  id, GLenum  pname, GLuint * params) {
//   (*fnptr)(id, pname, params);
// }
// static void  glowGetQueryiv(GPGETQUERYIV fnptr, GLenum  target, GLenum  pname, GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetQueryivEXT(GPGETQUERYIVEXT fnptr, GLenum  target, GLenum  pname, GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetRenderbufferParameteriv(GPGETRENDERBUFFERPARAMETERIV fnptr, GLenum  target, GLenum  pname, GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetSamplerParameterIivEXT(GPGETSAMPLERPARAMETERIIVEXT fnptr, GLuint  sampler, GLenum  pname, GLint * params) {
//   (*fnptr)(sampler, pname, params);
// }
// static void  glowGetSamplerParameterIivOES(GPGETSAMPLERPARAMETERIIVOES fnptr, GLuint  sampler, GLenum  pname, GLint * params) {
//   (*fnptr)(sampler, pname, params);
// }
// static void  glowGetSamplerParameterIuivEXT(GPGETSAMPLERPARAMETERIUIVEXT fnptr, GLuint  sampler, GLenum  pname, GLuint * params) {
//   (*fnptr)(sampler, pname, params);
// }
// static void  glowGetSamplerParameterIuivOES(GPGETSAMPLERPARAMETERIUIVOES fnptr, GLuint  sampler, GLenum  pname, GLuint * params) {
//   (*fnptr)(sampler, pname, params);
// }
// static void  glowGetSamplerParameterfv(GPGETSAMPLERPARAMETERFV fnptr, GLuint  sampler, GLenum  pname, GLfloat * params) {
//   (*fnptr)(sampler, pname, params);
// }
// static void  glowGetSamplerParameteriv(GPGETSAMPLERPARAMETERIV fnptr, GLuint  sampler, GLenum  pname, GLint * params) {
//   (*fnptr)(sampler, pname, params);
// }
// static void  glowGetSemaphoreParameterui64vEXT(GPGETSEMAPHOREPARAMETERUI64VEXT fnptr, GLuint  semaphore, GLenum  pname, GLuint64 * params) {
//   (*fnptr)(semaphore, pname, params);
// }
// static void  glowGetShaderInfoLog(GPGETSHADERINFOLOG fnptr, GLuint  shader, GLsizei  bufSize, GLsizei * length, GLchar * infoLog) {
//   (*fnptr)(shader, bufSize, length, infoLog);
// }
// static void  glowGetShaderPrecisionFormat(GPGETSHADERPRECISIONFORMAT fnptr, GLenum  shadertype, GLenum  precisiontype, GLint * range, GLint * precision) {
//   (*fnptr)(shadertype, precisiontype, range, precision);
// }
// static void  glowGetShaderSource(GPGETSHADERSOURCE fnptr, GLuint  shader, GLsizei  bufSize, GLsizei * length, GLchar * source) {
//   (*fnptr)(shader, bufSize, length, source);
// }
// static void  glowGetShaderiv(GPGETSHADERIV fnptr, GLuint  shader, GLenum  pname, GLint * params) {
//   (*fnptr)(shader, pname, params);
// }
// static const GLubyte * glowGetString(GPGETSTRING fnptr, GLenum  name) {
//   return (*fnptr)(name);
// }
// static const GLubyte * glowGetStringi(GPGETSTRINGI fnptr, GLenum  name, GLuint  index) {
//   return (*fnptr)(name, index);
// }
// static void  glowGetSynciv(GPGETSYNCIV fnptr, GLsync  sync, GLenum  pname, GLsizei  bufSize, GLsizei * length, GLint * values) {
//   (*fnptr)(sync, pname, bufSize, length, values);
// }
// static void  glowGetSyncivAPPLE(GPGETSYNCIVAPPLE fnptr, GLsync  sync, GLenum  pname, GLsizei  bufSize, GLsizei * length, GLint * values) {
//   (*fnptr)(sync, pname, bufSize, length, values);
// }
// static void  glowGetTexLevelParameterfv(GPGETTEXLEVELPARAMETERFV fnptr, GLenum  target, GLint  level, GLenum  pname, GLfloat * params) {
//   (*fnptr)(target, level, pname, params);
// }
// static void  glowGetTexLevelParameteriv(GPGETTEXLEVELPARAMETERIV fnptr, GLenum  target, GLint  level, GLenum  pname, GLint * params) {
//   (*fnptr)(target, level, pname, params);
// }
// static void  glowGetTexParameterIivEXT(GPGETTEXPARAMETERIIVEXT fnptr, GLenum  target, GLenum  pname, GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetTexParameterIivOES(GPGETTEXPARAMETERIIVOES fnptr, GLenum  target, GLenum  pname, GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetTexParameterIuivEXT(GPGETTEXPARAMETERIUIVEXT fnptr, GLenum  target, GLenum  pname, GLuint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetTexParameterIuivOES(GPGETTEXPARAMETERIUIVOES fnptr, GLenum  target, GLenum  pname, GLuint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetTexParameterfv(GPGETTEXPARAMETERFV fnptr, GLenum  target, GLenum  pname, GLfloat * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetTexParameteriv(GPGETTEXPARAMETERIV fnptr, GLenum  target, GLenum  pname, GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static GLuint64  glowGetTextureHandleIMG(GPGETTEXTUREHANDLEIMG fnptr, GLuint  texture) {
//   return (*fnptr)(texture);
// }
// static GLuint64  glowGetTextureHandleNV(GPGETTEXTUREHANDLENV fnptr, GLuint  texture) {
//   return (*fnptr)(texture);
// }
// static GLuint64  glowGetTextureSamplerHandleIMG(GPGETTEXTURESAMPLERHANDLEIMG fnptr, GLuint  texture, GLuint  sampler) {
//   return (*fnptr)(texture, sampler);
// }
// static GLuint64  glowGetTextureSamplerHandleNV(GPGETTEXTURESAMPLERHANDLENV fnptr, GLuint  texture, GLuint  sampler) {
//   return (*fnptr)(texture, sampler);
// }
// static void  glowGetTransformFeedbackVarying(GPGETTRANSFORMFEEDBACKVARYING fnptr, GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLsizei * size, GLenum * type, GLchar * name) {
//   (*fnptr)(program, index, bufSize, length, size, type, name);
// }
// static void  glowGetTranslatedShaderSourceANGLE(GPGETTRANSLATEDSHADERSOURCEANGLE fnptr, GLuint  shader, GLsizei  bufsize, GLsizei * length, GLchar * source) {
//   (*fnptr)(shader, bufsize, length, source);
// }
// static GLuint  glowGetUniformBlockIndex(GPGETUNIFORMBLOCKINDEX fnptr, GLuint  program, const GLchar * uniformBlockName) {
//   return (*fnptr)(program, uniformBlockName);
// }
// static void  glowGetUniformIndices(GPGETUNIFORMINDICES fnptr, GLuint  program, GLsizei  uniformCount, const GLchar *const* uniformNames, GLuint * uniformIndices) {
//   (*fnptr)(program, uniformCount, uniformNames, uniformIndices);
// }
// static GLint  glowGetUniformLocation(GPGETUNIFORMLOCATION fnptr, GLuint  program, const GLchar * name) {
//   return (*fnptr)(program, name);
// }
// static void  glowGetUniformfv(GPGETUNIFORMFV fnptr, GLuint  program, GLint  location, GLfloat * params) {
//   (*fnptr)(program, location, params);
// }
// static void  glowGetUniformi64vNV(GPGETUNIFORMI64VNV fnptr, GLuint  program, GLint  location, GLint64EXT * params) {
//   (*fnptr)(program, location, params);
// }
// static void  glowGetUniformiv(GPGETUNIFORMIV fnptr, GLuint  program, GLint  location, GLint * params) {
//   (*fnptr)(program, location, params);
// }
// static void  glowGetUniformuiv(GPGETUNIFORMUIV fnptr, GLuint  program, GLint  location, GLuint * params) {
//   (*fnptr)(program, location, params);
// }
// static void  glowGetUnsignedBytei_vEXT(GPGETUNSIGNEDBYTEI_VEXT fnptr, GLenum  target, GLuint  index, GLubyte * data) {
//   (*fnptr)(target, index, data);
// }
// static void  glowGetUnsignedBytevEXT(GPGETUNSIGNEDBYTEVEXT fnptr, GLenum  pname, GLubyte * data) {
//   (*fnptr)(pname, data);
// }
// static void  glowGetVertexAttribIiv(GPGETVERTEXATTRIBIIV fnptr, GLuint  index, GLenum  pname, GLint * params) {
//   (*fnptr)(index, pname, params);
// }
// static void  glowGetVertexAttribIuiv(GPGETVERTEXATTRIBIUIV fnptr, GLuint  index, GLenum  pname, GLuint * params) {
//   (*fnptr)(index, pname, params);
// }
// static void  glowGetVertexAttribPointerv(GPGETVERTEXATTRIBPOINTERV fnptr, GLuint  index, GLenum  pname, void ** pointer) {
//   (*fnptr)(index, pname, pointer);
// }
// static void  glowGetVertexAttribfv(GPGETVERTEXATTRIBFV fnptr, GLuint  index, GLenum  pname, GLfloat * params) {
//   (*fnptr)(index, pname, params);
// }
// static void  glowGetVertexAttribiv(GPGETVERTEXATTRIBIV fnptr, GLuint  index, GLenum  pname, GLint * params) {
//   (*fnptr)(index, pname, params);
// }
// static GLVULKANPROCNV  glowGetVkProcAddrNV(GPGETVKPROCADDRNV fnptr, const GLchar * name) {
//   return (*fnptr)(name);
// }
// static void  glowGetnUniformfv(GPGETNUNIFORMFV fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformfvEXT(GPGETNUNIFORMFVEXT fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformfvKHR(GPGETNUNIFORMFVKHR fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformiv(GPGETNUNIFORMIV fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLint * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformivEXT(GPGETNUNIFORMIVEXT fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLint * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformivKHR(GPGETNUNIFORMIVKHR fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLint * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformuiv(GPGETNUNIFORMUIV fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLuint * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformuivKHR(GPGETNUNIFORMUIVKHR fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLuint * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowHint(GPHINT fnptr, GLenum  target, GLenum  mode) {
//   (*fnptr)(target, mode);
// }
// static void  glowImportMemoryFdEXT(GPIMPORTMEMORYFDEXT fnptr, GLuint  memory, GLuint64  size, GLenum  handleType, GLint  fd) {
//   (*fnptr)(memory, size, handleType, fd);
// }
// static void  glowImportMemoryWin32HandleEXT(GPIMPORTMEMORYWIN32HANDLEEXT fnptr, GLuint  memory, GLuint64  size, GLenum  handleType, void * handle) {
//   (*fnptr)(memory, size, handleType, handle);
// }
// static void  glowImportMemoryWin32NameEXT(GPIMPORTMEMORYWIN32NAMEEXT fnptr, GLuint  memory, GLuint64  size, GLenum  handleType, const void * name) {
//   (*fnptr)(memory, size, handleType, name);
// }
// static void  glowImportSemaphoreFdEXT(GPIMPORTSEMAPHOREFDEXT fnptr, GLuint  semaphore, GLenum  handleType, GLint  fd) {
//   (*fnptr)(semaphore, handleType, fd);
// }
// static void  glowImportSemaphoreWin32HandleEXT(GPIMPORTSEMAPHOREWIN32HANDLEEXT fnptr, GLuint  semaphore, GLenum  handleType, void * handle) {
//   (*fnptr)(semaphore, handleType, handle);
// }
// static void  glowImportSemaphoreWin32NameEXT(GPIMPORTSEMAPHOREWIN32NAMEEXT fnptr, GLuint  semaphore, GLenum  handleType, const void * name) {
//   (*fnptr)(semaphore, handleType, name);
// }
// static void  glowInsertEventMarkerEXT(GPINSERTEVENTMARKEREXT fnptr, GLsizei  length, const GLchar * marker) {
//   (*fnptr)(length, marker);
// }
// static void  glowInterpolatePathsNV(GPINTERPOLATEPATHSNV fnptr, GLuint  resultPath, GLuint  pathA, GLuint  pathB, GLfloat  weight) {
//   (*fnptr)(resultPath, pathA, pathB, weight);
// }
// static void  glowInvalidateFramebuffer(GPINVALIDATEFRAMEBUFFER fnptr, GLenum  target, GLsizei  numAttachments, const GLenum * attachments) {
//   (*fnptr)(target, numAttachments, attachments);
// }
// static void  glowInvalidateSubFramebuffer(GPINVALIDATESUBFRAMEBUFFER fnptr, GLenum  target, GLsizei  numAttachments, const GLenum * attachments, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, numAttachments, attachments, x, y, width, height);
// }
// static GLboolean  glowIsBuffer(GPISBUFFER fnptr, GLuint  buffer) {
//   return (*fnptr)(buffer);
// }
// static GLboolean  glowIsEnabled(GPISENABLED fnptr, GLenum  cap) {
//   return (*fnptr)(cap);
// }
// static GLboolean  glowIsEnablediEXT(GPISENABLEDIEXT fnptr, GLenum  target, GLuint  index) {
//   return (*fnptr)(target, index);
// }
// static GLboolean  glowIsEnablediNV(GPISENABLEDINV fnptr, GLenum  target, GLuint  index) {
//   return (*fnptr)(target, index);
// }
// static GLboolean  glowIsEnablediOES(GPISENABLEDIOES fnptr, GLenum  target, GLuint  index) {
//   return (*fnptr)(target, index);
// }
// static GLboolean  glowIsFenceNV(GPISFENCENV fnptr, GLuint  fence) {
//   return (*fnptr)(fence);
// }
// static GLboolean  glowIsFramebuffer(GPISFRAMEBUFFER fnptr, GLuint  framebuffer) {
//   return (*fnptr)(framebuffer);
// }
// static GLboolean  glowIsImageHandleResidentNV(GPISIMAGEHANDLERESIDENTNV fnptr, GLuint64  handle) {
//   return (*fnptr)(handle);
// }
// static GLboolean  glowIsMemoryObjectEXT(GPISMEMORYOBJECTEXT fnptr, GLuint  memoryObject) {
//   return (*fnptr)(memoryObject);
// }
// static GLboolean  glowIsPathNV(GPISPATHNV fnptr, GLuint  path) {
//   return (*fnptr)(path);
// }
// static GLboolean  glowIsPointInFillPathNV(GPISPOINTINFILLPATHNV fnptr, GLuint  path, GLuint  mask, GLfloat  x, GLfloat  y) {
//   return (*fnptr)(path, mask, x, y);
// }
// static GLboolean  glowIsPointInStrokePathNV(GPISPOINTINSTROKEPATHNV fnptr, GLuint  path, GLfloat  x, GLfloat  y) {
//   return (*fnptr)(path, x, y);
// }
// static GLboolean  glowIsProgram(GPISPROGRAM fnptr, GLuint  program) {
//   return (*fnptr)(program);
// }
// static GLboolean  glowIsProgramPipeline(GPISPROGRAMPIPELINE fnptr, GLuint  pipeline) {
//   return (*fnptr)(pipeline);
// }
// static GLboolean  glowIsProgramPipelineEXT(GPISPROGRAMPIPELINEEXT fnptr, GLuint  pipeline) {
//   return (*fnptr)(pipeline);
// }
// static GLboolean  glowIsQuery(GPISQUERY fnptr, GLuint  id) {
//   return (*fnptr)(id);
// }
// static GLboolean  glowIsQueryEXT(GPISQUERYEXT fnptr, GLuint  id) {
//   return (*fnptr)(id);
// }
// static GLboolean  glowIsRenderbuffer(GPISRENDERBUFFER fnptr, GLuint  renderbuffer) {
//   return (*fnptr)(renderbuffer);
// }
// static GLboolean  glowIsSampler(GPISSAMPLER fnptr, GLuint  sampler) {
//   return (*fnptr)(sampler);
// }
// static GLboolean  glowIsSemaphoreEXT(GPISSEMAPHOREEXT fnptr, GLuint  semaphore) {
//   return (*fnptr)(semaphore);
// }
// static GLboolean  glowIsShader(GPISSHADER fnptr, GLuint  shader) {
//   return (*fnptr)(shader);
// }
// static GLboolean  glowIsSync(GPISSYNC fnptr, GLsync  sync) {
//   return (*fnptr)(sync);
// }
// static GLboolean  glowIsSyncAPPLE(GPISSYNCAPPLE fnptr, GLsync  sync) {
//   return (*fnptr)(sync);
// }
// static GLboolean  glowIsTexture(GPISTEXTURE fnptr, GLuint  texture) {
//   return (*fnptr)(texture);
// }
// static GLboolean  glowIsTextureHandleResidentNV(GPISTEXTUREHANDLERESIDENTNV fnptr, GLuint64  handle) {
//   return (*fnptr)(handle);
// }
// static GLboolean  glowIsTransformFeedback(GPISTRANSFORMFEEDBACK fnptr, GLuint  id) {
//   return (*fnptr)(id);
// }
// static GLboolean  glowIsVertexArray(GPISVERTEXARRAY fnptr, GLuint  array) {
//   return (*fnptr)(array);
// }
// static GLboolean  glowIsVertexArrayOES(GPISVERTEXARRAYOES fnptr, GLuint  array) {
//   return (*fnptr)(array);
// }
// static void  glowLabelObjectEXT(GPLABELOBJECTEXT fnptr, GLenum  type, GLuint  object, GLsizei  length, const GLchar * label) {
//   (*fnptr)(type, object, length, label);
// }
// static void  glowLineWidth(GPLINEWIDTH fnptr, GLfloat  width) {
//   (*fnptr)(width);
// }
// static void  glowLinkProgram(GPLINKPROGRAM fnptr, GLuint  program) {
//   (*fnptr)(program);
// }
// static void  glowMakeImageHandleNonResidentNV(GPMAKEIMAGEHANDLENONRESIDENTNV fnptr, GLuint64  handle) {
//   (*fnptr)(handle);
// }
// static void  glowMakeImageHandleResidentNV(GPMAKEIMAGEHANDLERESIDENTNV fnptr, GLuint64  handle, GLenum  access) {
//   (*fnptr)(handle, access);
// }
// static void  glowMakeTextureHandleNonResidentNV(GPMAKETEXTUREHANDLENONRESIDENTNV fnptr, GLuint64  handle) {
//   (*fnptr)(handle);
// }
// static void  glowMakeTextureHandleResidentNV(GPMAKETEXTUREHANDLERESIDENTNV fnptr, GLuint64  handle) {
//   (*fnptr)(handle);
// }
// static void * glowMapBufferOES(GPMAPBUFFEROES fnptr, GLenum  target, GLenum  access) {
//   return (*fnptr)(target, access);
// }
// static void * glowMapBufferRange(GPMAPBUFFERRANGE fnptr, GLenum  target, GLintptr  offset, GLsizeiptr  length, GLbitfield  access) {
//   return (*fnptr)(target, offset, length, access);
// }
// static void * glowMapBufferRangeEXT(GPMAPBUFFERRANGEEXT fnptr, GLenum  target, GLintptr  offset, GLsizeiptr  length, GLbitfield  access) {
//   return (*fnptr)(target, offset, length, access);
// }
// static void  glowMatrixFrustumEXT(GPMATRIXFRUSTUMEXT fnptr, GLenum  mode, GLdouble  left, GLdouble  right, GLdouble  bottom, GLdouble  top, GLdouble  zNear, GLdouble  zFar) {
//   (*fnptr)(mode, left, right, bottom, top, zNear, zFar);
// }
// static void  glowMatrixLoad3x2fNV(GPMATRIXLOAD3X2FNV fnptr, GLenum  matrixMode, const GLfloat * m) {
//   (*fnptr)(matrixMode, m);
// }
// static void  glowMatrixLoad3x3fNV(GPMATRIXLOAD3X3FNV fnptr, GLenum  matrixMode, const GLfloat * m) {
//   (*fnptr)(matrixMode, m);
// }
// static void  glowMatrixLoadIdentityEXT(GPMATRIXLOADIDENTITYEXT fnptr, GLenum  mode) {
//   (*fnptr)(mode);
// }
// static void  glowMatrixLoadTranspose3x3fNV(GPMATRIXLOADTRANSPOSE3X3FNV fnptr, GLenum  matrixMode, const GLfloat * m) {
//   (*fnptr)(matrixMode, m);
// }
// static void  glowMatrixLoadTransposedEXT(GPMATRIXLOADTRANSPOSEDEXT fnptr, GLenum  mode, const GLdouble * m) {
//   (*fnptr)(mode, m);
// }
// static void  glowMatrixLoadTransposefEXT(GPMATRIXLOADTRANSPOSEFEXT fnptr, GLenum  mode, const GLfloat * m) {
//   (*fnptr)(mode, m);
// }
// static void  glowMatrixLoaddEXT(GPMATRIXLOADDEXT fnptr, GLenum  mode, const GLdouble * m) {
//   (*fnptr)(mode, m);
// }
// static void  glowMatrixLoadfEXT(GPMATRIXLOADFEXT fnptr, GLenum  mode, const GLfloat * m) {
//   (*fnptr)(mode, m);
// }
// static void  glowMatrixMult3x2fNV(GPMATRIXMULT3X2FNV fnptr, GLenum  matrixMode, const GLfloat * m) {
//   (*fnptr)(matrixMode, m);
// }
// static void  glowMatrixMult3x3fNV(GPMATRIXMULT3X3FNV fnptr, GLenum  matrixMode, const GLfloat * m) {
//   (*fnptr)(matrixMode, m);
// }
// static void  glowMatrixMultTranspose3x3fNV(GPMATRIXMULTTRANSPOSE3X3FNV fnptr, GLenum  matrixMode, const GLfloat * m) {
//   (*fnptr)(matrixMode, m);
// }
// static void  glowMatrixMultTransposedEXT(GPMATRIXMULTTRANSPOSEDEXT fnptr, GLenum  mode, const GLdouble * m) {
//   (*fnptr)(mode, m);
// }
// static void  glowMatrixMultTransposefEXT(GPMATRIXMULTTRANSPOSEFEXT fnptr, GLenum  mode, const GLfloat * m) {
//   (*fnptr)(mode, m);
// }
// static void  glowMatrixMultdEXT(GPMATRIXMULTDEXT fnptr, GLenum  mode, const GLdouble * m) {
//   (*fnptr)(mode, m);
// }
// static void  glowMatrixMultfEXT(GPMATRIXMULTFEXT fnptr, GLenum  mode, const GLfloat * m) {
//   (*fnptr)(mode, m);
// }
// static void  glowMatrixOrthoEXT(GPMATRIXORTHOEXT fnptr, GLenum  mode, GLdouble  left, GLdouble  right, GLdouble  bottom, GLdouble  top, GLdouble  zNear, GLdouble  zFar) {
//   (*fnptr)(mode, left, right, bottom, top, zNear, zFar);
// }
// static void  glowMatrixPopEXT(GPMATRIXPOPEXT fnptr, GLenum  mode) {
//   (*fnptr)(mode);
// }
// static void  glowMatrixPushEXT(GPMATRIXPUSHEXT fnptr, GLenum  mode) {
//   (*fnptr)(mode);
// }
// static void  glowMatrixRotatedEXT(GPMATRIXROTATEDEXT fnptr, GLenum  mode, GLdouble  angle, GLdouble  x, GLdouble  y, GLdouble  z) {
//   (*fnptr)(mode, angle, x, y, z);
// }
// static void  glowMatrixRotatefEXT(GPMATRIXROTATEFEXT fnptr, GLenum  mode, GLfloat  angle, GLfloat  x, GLfloat  y, GLfloat  z) {
//   (*fnptr)(mode, angle, x, y, z);
// }
// static void  glowMatrixScaledEXT(GPMATRIXSCALEDEXT fnptr, GLenum  mode, GLdouble  x, GLdouble  y, GLdouble  z) {
//   (*fnptr)(mode, x, y, z);
// }
// static void  glowMatrixScalefEXT(GPMATRIXSCALEFEXT fnptr, GLenum  mode, GLfloat  x, GLfloat  y, GLfloat  z) {
//   (*fnptr)(mode, x, y, z);
// }
// static void  glowMatrixTranslatedEXT(GPMATRIXTRANSLATEDEXT fnptr, GLenum  mode, GLdouble  x, GLdouble  y, GLdouble  z) {
//   (*fnptr)(mode, x, y, z);
// }
// static void  glowMatrixTranslatefEXT(GPMATRIXTRANSLATEFEXT fnptr, GLenum  mode, GLfloat  x, GLfloat  y, GLfloat  z) {
//   (*fnptr)(mode, x, y, z);
// }
// static void  glowMaxShaderCompilerThreadsKHR(GPMAXSHADERCOMPILERTHREADSKHR fnptr, GLuint  count) {
//   (*fnptr)(count);
// }
// static void  glowMemoryBarrier(GPMEMORYBARRIER fnptr, GLbitfield  barriers) {
//   (*fnptr)(barriers);
// }
// static void  glowMemoryBarrierByRegion(GPMEMORYBARRIERBYREGION fnptr, GLbitfield  barriers) {
//   (*fnptr)(barriers);
// }
// static void  glowMemoryObjectParameterivEXT(GPMEMORYOBJECTPARAMETERIVEXT fnptr, GLuint  memoryObject, GLenum  pname, const GLint * params) {
//   (*fnptr)(memoryObject, pname, params);
// }
// static void  glowMinSampleShadingOES(GPMINSAMPLESHADINGOES fnptr, GLfloat  value) {
//   (*fnptr)(value);
// }
// static void  glowMultiDrawArraysEXT(GPMULTIDRAWARRAYSEXT fnptr, GLenum  mode, const GLint * first, const GLsizei * count, GLsizei  primcount) {
//   (*fnptr)(mode, first, count, primcount);
// }
// static void  glowMultiDrawArraysIndirectEXT(GPMULTIDRAWARRAYSINDIRECTEXT fnptr, GLenum  mode, const void * indirect, GLsizei  drawcount, GLsizei  stride) {
//   (*fnptr)(mode, indirect, drawcount, stride);
// }
// static void  glowMultiDrawElementsBaseVertexEXT(GPMULTIDRAWELEMENTSBASEVERTEXEXT fnptr, GLenum  mode, const GLsizei * count, GLenum  type, const void *const* indices, GLsizei  primcount, const GLint * basevertex) {
//   (*fnptr)(mode, count, type, indices, primcount, basevertex);
// }
// static void  glowMultiDrawElementsEXT(GPMULTIDRAWELEMENTSEXT fnptr, GLenum  mode, const GLsizei * count, GLenum  type, const void *const* indices, GLsizei  primcount) {
//   (*fnptr)(mode, count, type, indices, primcount);
// }
// static void  glowMultiDrawElementsIndirectEXT(GPMULTIDRAWELEMENTSINDIRECTEXT fnptr, GLenum  mode, GLenum  type, const void * indirect, GLsizei  drawcount, GLsizei  stride) {
//   (*fnptr)(mode, type, indirect, drawcount, stride);
// }
// static void  glowNamedBufferStorageExternalEXT(GPNAMEDBUFFERSTORAGEEXTERNALEXT fnptr, GLuint  buffer, GLintptr  offset, GLsizeiptr  size, GLeglClientBufferEXT  clientBuffer, GLbitfield  flags) {
//   (*fnptr)(buffer, offset, size, clientBuffer, flags);
// }
// static void  glowNamedBufferStorageMemEXT(GPNAMEDBUFFERSTORAGEMEMEXT fnptr, GLuint  buffer, GLsizeiptr  size, GLuint  memory, GLuint64  offset) {
//   (*fnptr)(buffer, size, memory, offset);
// }
// static void  glowNamedFramebufferSampleLocationsfvNV(GPNAMEDFRAMEBUFFERSAMPLELOCATIONSFVNV fnptr, GLuint  framebuffer, GLuint  start, GLsizei  count, const GLfloat * v) {
//   (*fnptr)(framebuffer, start, count, v);
// }
// static void  glowObjectLabel(GPOBJECTLABEL fnptr, GLenum  identifier, GLuint  name, GLsizei  length, const GLchar * label) {
//   (*fnptr)(identifier, name, length, label);
// }
// static void  glowObjectLabelKHR(GPOBJECTLABELKHR fnptr, GLenum  identifier, GLuint  name, GLsizei  length, const GLchar * label) {
//   (*fnptr)(identifier, name, length, label);
// }
// static void  glowObjectPtrLabel(GPOBJECTPTRLABEL fnptr, const void * ptr, GLsizei  length, const GLchar * label) {
//   (*fnptr)(ptr, length, label);
// }
// static void  glowObjectPtrLabelKHR(GPOBJECTPTRLABELKHR fnptr, const void * ptr, GLsizei  length, const GLchar * label) {
//   (*fnptr)(ptr, length, label);
// }
// static void  glowPatchParameteriEXT(GPPATCHPARAMETERIEXT fnptr, GLenum  pname, GLint  value) {
//   (*fnptr)(pname, value);
// }
// static void  glowPatchParameteriOES(GPPATCHPARAMETERIOES fnptr, GLenum  pname, GLint  value) {
//   (*fnptr)(pname, value);
// }
// static void  glowPathCommandsNV(GPPATHCOMMANDSNV fnptr, GLuint  path, GLsizei  numCommands, const GLubyte * commands, GLsizei  numCoords, GLenum  coordType, const void * coords) {
//   (*fnptr)(path, numCommands, commands, numCoords, coordType, coords);
// }
// static void  glowPathCoordsNV(GPPATHCOORDSNV fnptr, GLuint  path, GLsizei  numCoords, GLenum  coordType, const void * coords) {
//   (*fnptr)(path, numCoords, coordType, coords);
// }
// static void  glowPathCoverDepthFuncNV(GPPATHCOVERDEPTHFUNCNV fnptr, GLenum  func) {
//   (*fnptr)(func);
// }
// static void  glowPathDashArrayNV(GPPATHDASHARRAYNV fnptr, GLuint  path, GLsizei  dashCount, const GLfloat * dashArray) {
//   (*fnptr)(path, dashCount, dashArray);
// }
// static GLenum  glowPathGlyphIndexArrayNV(GPPATHGLYPHINDEXARRAYNV fnptr, GLuint  firstPathName, GLenum  fontTarget, const void * fontName, GLbitfield  fontStyle, GLuint  firstGlyphIndex, GLsizei  numGlyphs, GLuint  pathParameterTemplate, GLfloat  emScale) {
//   return (*fnptr)(firstPathName, fontTarget, fontName, fontStyle, firstGlyphIndex, numGlyphs, pathParameterTemplate, emScale);
// }
// static GLenum  glowPathGlyphIndexRangeNV(GPPATHGLYPHINDEXRANGENV fnptr, GLenum  fontTarget, const void * fontName, GLbitfield  fontStyle, GLuint  pathParameterTemplate, GLfloat  emScale, GLuint * baseAndCount) {
//   return (*fnptr)(fontTarget, fontName, fontStyle, pathParameterTemplate, emScale, baseAndCount);
// }
// static void  glowPathGlyphRangeNV(GPPATHGLYPHRANGENV fnptr, GLuint  firstPathName, GLenum  fontTarget, const void * fontName, GLbitfield  fontStyle, GLuint  firstGlyph, GLsizei  numGlyphs, GLenum  handleMissingGlyphs, GLuint  pathParameterTemplate, GLfloat  emScale) {
//   (*fnptr)(firstPathName, fontTarget, fontName, fontStyle, firstGlyph, numGlyphs, handleMissingGlyphs, pathParameterTemplate, emScale);
// }
// static void  glowPathGlyphsNV(GPPATHGLYPHSNV fnptr, GLuint  firstPathName, GLenum  fontTarget, const void * fontName, GLbitfield  fontStyle, GLsizei  numGlyphs, GLenum  type, const void * charcodes, GLenum  handleMissingGlyphs, GLuint  pathParameterTemplate, GLfloat  emScale) {
//   (*fnptr)(firstPathName, fontTarget, fontName, fontStyle, numGlyphs, type, charcodes, handleMissingGlyphs, pathParameterTemplate, emScale);
// }
// static GLenum  glowPathMemoryGlyphIndexArrayNV(GPPATHMEMORYGLYPHINDEXARRAYNV fnptr, GLuint  firstPathName, GLenum  fontTarget, GLsizeiptr  fontSize, const void * fontData, GLsizei  faceIndex, GLuint  firstGlyphIndex, GLsizei  numGlyphs, GLuint  pathParameterTemplate, GLfloat  emScale) {
//   return (*fnptr)(firstPathName, fontTarget, fontSize, fontData, faceIndex, firstGlyphIndex, numGlyphs, pathParameterTemplate, emScale);
// }
// static void  glowPathParameterfNV(GPPATHPARAMETERFNV fnptr, GLuint  path, GLenum  pname, GLfloat  value) {
//   (*fnptr)(path, pname, value);
// }
// static void  glowPathParameterfvNV(GPPATHPARAMETERFVNV fnptr, GLuint  path, GLenum  pname, const GLfloat * value) {
//   (*fnptr)(path, pname, value);
// }
// static void  glowPathParameteriNV(GPPATHPARAMETERINV fnptr, GLuint  path, GLenum  pname, GLint  value) {
//   (*fnptr)(path, pname, value);
// }
// static void  glowPathParameterivNV(GPPATHPARAMETERIVNV fnptr, GLuint  path, GLenum  pname, const GLint * value) {
//   (*fnptr)(path, pname, value);
// }
// static void  glowPathStencilDepthOffsetNV(GPPATHSTENCILDEPTHOFFSETNV fnptr, GLfloat  factor, GLfloat  units) {
//   (*fnptr)(factor, units);
// }
// static void  glowPathStencilFuncNV(GPPATHSTENCILFUNCNV fnptr, GLenum  func, GLint  ref, GLuint  mask) {
//   (*fnptr)(func, ref, mask);
// }
// static void  glowPathStringNV(GPPATHSTRINGNV fnptr, GLuint  path, GLenum  format, GLsizei  length, const void * pathString) {
//   (*fnptr)(path, format, length, pathString);
// }
// static void  glowPathSubCommandsNV(GPPATHSUBCOMMANDSNV fnptr, GLuint  path, GLsizei  commandStart, GLsizei  commandsToDelete, GLsizei  numCommands, const GLubyte * commands, GLsizei  numCoords, GLenum  coordType, const void * coords) {
//   (*fnptr)(path, commandStart, commandsToDelete, numCommands, commands, numCoords, coordType, coords);
// }
// static void  glowPathSubCoordsNV(GPPATHSUBCOORDSNV fnptr, GLuint  path, GLsizei  coordStart, GLsizei  numCoords, GLenum  coordType, const void * coords) {
//   (*fnptr)(path, coordStart, numCoords, coordType, coords);
// }
// static void  glowPauseTransformFeedback(GPPAUSETRANSFORMFEEDBACK fnptr) {
//   (*fnptr)();
// }
// static void  glowPixelStorei(GPPIXELSTOREI fnptr, GLenum  pname, GLint  param) {
//   (*fnptr)(pname, param);
// }
// static GLboolean  glowPointAlongPathNV(GPPOINTALONGPATHNV fnptr, GLuint  path, GLsizei  startSegment, GLsizei  numSegments, GLfloat  distance, GLfloat * x, GLfloat * y, GLfloat * tangentX, GLfloat * tangentY) {
//   return (*fnptr)(path, startSegment, numSegments, distance, x, y, tangentX, tangentY);
// }
// static void  glowPolygonModeNV(GPPOLYGONMODENV fnptr, GLenum  face, GLenum  mode) {
//   (*fnptr)(face, mode);
// }
// static void  glowPolygonOffset(GPPOLYGONOFFSET fnptr, GLfloat  factor, GLfloat  units) {
//   (*fnptr)(factor, units);
// }
// static void  glowPolygonOffsetClampEXT(GPPOLYGONOFFSETCLAMPEXT fnptr, GLfloat  factor, GLfloat  units, GLfloat  clamp) {
//   (*fnptr)(factor, units, clamp);
// }
// static void  glowPopDebugGroup(GPPOPDEBUGGROUP fnptr) {
//   (*fnptr)();
// }
// static void  glowPopDebugGroupKHR(GPPOPDEBUGGROUPKHR fnptr) {
//   (*fnptr)();
// }
// static void  glowPopGroupMarkerEXT(GPPOPGROUPMARKEREXT fnptr) {
//   (*fnptr)();
// }
// static void  glowPrimitiveBoundingBoxEXT(GPPRIMITIVEBOUNDINGBOXEXT fnptr, GLfloat  minX, GLfloat  minY, GLfloat  minZ, GLfloat  minW, GLfloat  maxX, GLfloat  maxY, GLfloat  maxZ, GLfloat  maxW) {
//   (*fnptr)(minX, minY, minZ, minW, maxX, maxY, maxZ, maxW);
// }
// static void  glowPrimitiveBoundingBoxOES(GPPRIMITIVEBOUNDINGBOXOES fnptr, GLfloat  minX, GLfloat  minY, GLfloat  minZ, GLfloat  minW, GLfloat  maxX, GLfloat  maxY, GLfloat  maxZ, GLfloat  maxW) {
//   (*fnptr)(minX, minY, minZ, minW, maxX, maxY, maxZ, maxW);
// }
// static void  glowProgramBinary(GPPROGRAMBINARY fnptr, GLuint  program, GLenum  binaryFormat, const void * binary, GLsizei  length) {
//   (*fnptr)(program, binaryFormat, binary, length);
// }
// static void  glowProgramBinaryOES(GPPROGRAMBINARYOES fnptr, GLuint  program, GLenum  binaryFormat, const void * binary, GLint  length) {
//   (*fnptr)(program, binaryFormat, binary, length);
// }
// static void  glowProgramParameteri(GPPROGRAMPARAMETERI fnptr, GLuint  program, GLenum  pname, GLint  value) {
//   (*fnptr)(program, pname, value);
// }
// static void  glowProgramParameteriEXT(GPPROGRAMPARAMETERIEXT fnptr, GLuint  program, GLenum  pname, GLint  value) {
//   (*fnptr)(program, pname, value);
// }
// static void  glowProgramPathFragmentInputGenNV(GPPROGRAMPATHFRAGMENTINPUTGENNV fnptr, GLuint  program, GLint  location, GLenum  genMode, GLint  components, const GLfloat * coeffs) {
//   (*fnptr)(program, location, genMode, components, coeffs);
// }
// static void  glowProgramUniform1f(GPPROGRAMUNIFORM1F fnptr, GLuint  program, GLint  location, GLfloat  v0) {
//   (*fnptr)(program, location, v0);
// }
// static void  glowProgramUniform1fEXT(GPPROGRAMUNIFORM1FEXT fnptr, GLuint  program, GLint  location, GLfloat  v0) {
//   (*fnptr)(program, location, v0);
// }
// static void  glowProgramUniform1fv(GPPROGRAMUNIFORM1FV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform1fvEXT(GPPROGRAMUNIFORM1FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform1i(GPPROGRAMUNIFORM1I fnptr, GLuint  program, GLint  location, GLint  v0) {
//   (*fnptr)(program, location, v0);
// }
// static void  glowProgramUniform1i64NV(GPPROGRAMUNIFORM1I64NV fnptr, GLuint  program, GLint  location, GLint64EXT  x) {
//   (*fnptr)(program, location, x);
// }
// static void  glowProgramUniform1i64vNV(GPPROGRAMUNIFORM1I64VNV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint64EXT * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform1iEXT(GPPROGRAMUNIFORM1IEXT fnptr, GLuint  program, GLint  location, GLint  v0) {
//   (*fnptr)(program, location, v0);
// }
// static void  glowProgramUniform1iv(GPPROGRAMUNIFORM1IV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform1ivEXT(GPPROGRAMUNIFORM1IVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform1ui(GPPROGRAMUNIFORM1UI fnptr, GLuint  program, GLint  location, GLuint  v0) {
//   (*fnptr)(program, location, v0);
// }
// static void  glowProgramUniform1ui64NV(GPPROGRAMUNIFORM1UI64NV fnptr, GLuint  program, GLint  location, GLuint64EXT  x) {
//   (*fnptr)(program, location, x);
// }
// static void  glowProgramUniform1ui64vNV(GPPROGRAMUNIFORM1UI64VNV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint64EXT * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform1uiEXT(GPPROGRAMUNIFORM1UIEXT fnptr, GLuint  program, GLint  location, GLuint  v0) {
//   (*fnptr)(program, location, v0);
// }
// static void  glowProgramUniform1uiv(GPPROGRAMUNIFORM1UIV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform1uivEXT(GPPROGRAMUNIFORM1UIVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform2f(GPPROGRAMUNIFORM2F fnptr, GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1) {
//   (*fnptr)(program, location, v0, v1);
// }
// static void  glowProgramUniform2fEXT(GPPROGRAMUNIFORM2FEXT fnptr, GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1) {
//   (*fnptr)(program, location, v0, v1);
// }
// static void  glowProgramUniform2fv(GPPROGRAMUNIFORM2FV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform2fvEXT(GPPROGRAMUNIFORM2FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform2i(GPPROGRAMUNIFORM2I fnptr, GLuint  program, GLint  location, GLint  v0, GLint  v1) {
//   (*fnptr)(program, location, v0, v1);
// }
// static void  glowProgramUniform2i64NV(GPPROGRAMUNIFORM2I64NV fnptr, GLuint  program, GLint  location, GLint64EXT  x, GLint64EXT  y) {
//   (*fnptr)(program, location, x, y);
// }
// static void  glowProgramUniform2i64vNV(GPPROGRAMUNIFORM2I64VNV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint64EXT * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform2iEXT(GPPROGRAMUNIFORM2IEXT fnptr, GLuint  program, GLint  location, GLint  v0, GLint  v1) {
//   (*fnptr)(program, location, v0, v1);
// }
// static void  glowProgramUniform2iv(GPPROGRAMUNIFORM2IV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform2ivEXT(GPPROGRAMUNIFORM2IVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform2ui(GPPROGRAMUNIFORM2UI fnptr, GLuint  program, GLint  location, GLuint  v0, GLuint  v1) {
//   (*fnptr)(program, location, v0, v1);
// }
// static void  glowProgramUniform2ui64NV(GPPROGRAMUNIFORM2UI64NV fnptr, GLuint  program, GLint  location, GLuint64EXT  x, GLuint64EXT  y) {
//   (*fnptr)(program, location, x, y);
// }
// static void  glowProgramUniform2ui64vNV(GPPROGRAMUNIFORM2UI64VNV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint64EXT * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform2uiEXT(GPPROGRAMUNIFORM2UIEXT fnptr, GLuint  program, GLint  location, GLuint  v0, GLuint  v1) {
//   (*fnptr)(program, location, v0, v1);
// }
// static void  glowProgramUniform2uiv(GPPROGRAMUNIFORM2UIV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform2uivEXT(GPPROGRAMUNIFORM2UIVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform3f(GPPROGRAMUNIFORM3F fnptr, GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2) {
//   (*fnptr)(program, location, v0, v1, v2);
// }
// static void  glowProgramUniform3fEXT(GPPROGRAMUNIFORM3FEXT fnptr, GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2) {
//   (*fnptr)(program, location, v0, v1, v2);
// }
// static void  glowProgramUniform3fv(GPPROGRAMUNIFORM3FV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform3fvEXT(GPPROGRAMUNIFORM3FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform3i(GPPROGRAMUNIFORM3I fnptr, GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2) {
//   (*fnptr)(program, location, v0, v1, v2);
// }
// static void  glowProgramUniform3i64NV(GPPROGRAMUNIFORM3I64NV fnptr, GLuint  program, GLint  location, GLint64EXT  x, GLint64EXT  y, GLint64EXT  z) {
//   (*fnptr)(program, location, x, y, z);
// }
// static void  glowProgramUniform3i64vNV(GPPROGRAMUNIFORM3I64VNV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint64EXT * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform3iEXT(GPPROGRAMUNIFORM3IEXT fnptr, GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2) {
//   (*fnptr)(program, location, v0, v1, v2);
// }
// static void  glowProgramUniform3iv(GPPROGRAMUNIFORM3IV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform3ivEXT(GPPROGRAMUNIFORM3IVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform3ui(GPPROGRAMUNIFORM3UI fnptr, GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2) {
//   (*fnptr)(program, location, v0, v1, v2);
// }
// static void  glowProgramUniform3ui64NV(GPPROGRAMUNIFORM3UI64NV fnptr, GLuint  program, GLint  location, GLuint64EXT  x, GLuint64EXT  y, GLuint64EXT  z) {
//   (*fnptr)(program, location, x, y, z);
// }
// static void  glowProgramUniform3ui64vNV(GPPROGRAMUNIFORM3UI64VNV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint64EXT * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform3uiEXT(GPPROGRAMUNIFORM3UIEXT fnptr, GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2) {
//   (*fnptr)(program, location, v0, v1, v2);
// }
// static void  glowProgramUniform3uiv(GPPROGRAMUNIFORM3UIV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform3uivEXT(GPPROGRAMUNIFORM3UIVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform4f(GPPROGRAMUNIFORM4F fnptr, GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2, GLfloat  v3) {
//   (*fnptr)(program, location, v0, v1, v2, v3);
// }
// static void  glowProgramUniform4fEXT(GPPROGRAMUNIFORM4FEXT fnptr, GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2, GLfloat  v3) {
//   (*fnptr)(program, location, v0, v1, v2, v3);
// }
// static void  glowProgramUniform4fv(GPPROGRAMUNIFORM4FV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform4fvEXT(GPPROGRAMUNIFORM4FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform4i(GPPROGRAMUNIFORM4I fnptr, GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2, GLint  v3) {
//   (*fnptr)(program, location, v0, v1, v2, v3);
// }
// static void  glowProgramUniform4i64NV(GPPROGRAMUNIFORM4I64NV fnptr, GLuint  program, GLint  location, GLint64EXT  x, GLint64EXT  y, GLint64EXT  z, GLint64EXT  w) {
//   (*fnptr)(program, location, x, y, z, w);
// }
// static void  glowProgramUniform4i64vNV(GPPROGRAMUNIFORM4I64VNV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint64EXT * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform4iEXT(GPPROGRAMUNIFORM4IEXT fnptr, GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2, GLint  v3) {
//   (*fnptr)(program, location, v0, v1, v2, v3);
// }
// static void  glowProgramUniform4iv(GPPROGRAMUNIFORM4IV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform4ivEXT(GPPROGRAMUNIFORM4IVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform4ui(GPPROGRAMUNIFORM4UI fnptr, GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2, GLuint  v3) {
//   (*fnptr)(program, location, v0, v1, v2, v3);
// }
// static void  glowProgramUniform4ui64NV(GPPROGRAMUNIFORM4UI64NV fnptr, GLuint  program, GLint  location, GLuint64EXT  x, GLuint64EXT  y, GLuint64EXT  z, GLuint64EXT  w) {
//   (*fnptr)(program, location, x, y, z, w);
// }
// static void  glowProgramUniform4ui64vNV(GPPROGRAMUNIFORM4UI64VNV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint64EXT * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform4uiEXT(GPPROGRAMUNIFORM4UIEXT fnptr, GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2, GLuint  v3) {
//   (*fnptr)(program, location, v0, v1, v2, v3);
// }
// static void  glowProgramUniform4uiv(GPPROGRAMUNIFORM4UIV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform4uivEXT(GPPROGRAMUNIFORM4UIVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniformHandleui64IMG(GPPROGRAMUNIFORMHANDLEUI64IMG fnptr, GLuint  program, GLint  location, GLuint64  value) {
//   (*fnptr)(program, location, value);
// }
// static void  glowProgramUniformHandleui64NV(GPPROGRAMUNIFORMHANDLEUI64NV fnptr, GLuint  program, GLint  location, GLuint64  value) {
//   (*fnptr)(program, location, value);
// }
// static void  glowProgramUniformHandleui64vIMG(GPPROGRAMUNIFORMHANDLEUI64VIMG fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint64 * values) {
//   (*fnptr)(program, location, count, values);
// }
// static void  glowProgramUniformHandleui64vNV(GPPROGRAMUNIFORMHANDLEUI64VNV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint64 * values) {
//   (*fnptr)(program, location, count, values);
// }
// static void  glowProgramUniformMatrix2fv(GPPROGRAMUNIFORMMATRIX2FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix2fvEXT(GPPROGRAMUNIFORMMATRIX2FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix2x3fv(GPPROGRAMUNIFORMMATRIX2X3FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix2x3fvEXT(GPPROGRAMUNIFORMMATRIX2X3FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix2x4fv(GPPROGRAMUNIFORMMATRIX2X4FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix2x4fvEXT(GPPROGRAMUNIFORMMATRIX2X4FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix3fv(GPPROGRAMUNIFORMMATRIX3FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix3fvEXT(GPPROGRAMUNIFORMMATRIX3FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix3x2fv(GPPROGRAMUNIFORMMATRIX3X2FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix3x2fvEXT(GPPROGRAMUNIFORMMATRIX3X2FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix3x4fv(GPPROGRAMUNIFORMMATRIX3X4FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix3x4fvEXT(GPPROGRAMUNIFORMMATRIX3X4FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix4fv(GPPROGRAMUNIFORMMATRIX4FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix4fvEXT(GPPROGRAMUNIFORMMATRIX4FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix4x2fv(GPPROGRAMUNIFORMMATRIX4X2FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix4x2fvEXT(GPPROGRAMUNIFORMMATRIX4X2FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix4x3fv(GPPROGRAMUNIFORMMATRIX4X3FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix4x3fvEXT(GPPROGRAMUNIFORMMATRIX4X3FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowPushDebugGroup(GPPUSHDEBUGGROUP fnptr, GLenum  source, GLuint  id, GLsizei  length, const GLchar * message) {
//   (*fnptr)(source, id, length, message);
// }
// static void  glowPushDebugGroupKHR(GPPUSHDEBUGGROUPKHR fnptr, GLenum  source, GLuint  id, GLsizei  length, const GLchar * message) {
//   (*fnptr)(source, id, length, message);
// }
// static void  glowPushGroupMarkerEXT(GPPUSHGROUPMARKEREXT fnptr, GLsizei  length, const GLchar * marker) {
//   (*fnptr)(length, marker);
// }
// static void  glowQueryCounterEXT(GPQUERYCOUNTEREXT fnptr, GLuint  id, GLenum  target) {
//   (*fnptr)(id, target);
// }
// static void  glowRasterSamplesEXT(GPRASTERSAMPLESEXT fnptr, GLuint  samples, GLboolean  fixedsamplelocations) {
//   (*fnptr)(samples, fixedsamplelocations);
// }
// static void  glowReadBuffer(GPREADBUFFER fnptr, GLenum  src) {
//   (*fnptr)(src);
// }
// static void  glowReadBufferIndexedEXT(GPREADBUFFERINDEXEDEXT fnptr, GLenum  src, GLint  index) {
//   (*fnptr)(src, index);
// }
// static void  glowReadBufferNV(GPREADBUFFERNV fnptr, GLenum  mode) {
//   (*fnptr)(mode);
// }
// static void  glowReadPixels(GPREADPIXELS fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, void * pixels) {
//   (*fnptr)(x, y, width, height, format, type, pixels);
// }
// static void  glowReadnPixels(GPREADNPIXELS fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data) {
//   (*fnptr)(x, y, width, height, format, type, bufSize, data);
// }
// static void  glowReadnPixelsEXT(GPREADNPIXELSEXT fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data) {
//   (*fnptr)(x, y, width, height, format, type, bufSize, data);
// }
// static void  glowReadnPixelsKHR(GPREADNPIXELSKHR fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data) {
//   (*fnptr)(x, y, width, height, format, type, bufSize, data);
// }
// static GLboolean  glowReleaseKeyedMutexWin32EXT(GPRELEASEKEYEDMUTEXWIN32EXT fnptr, GLuint  memory, GLuint64  key) {
//   return (*fnptr)(memory, key);
// }
// static void  glowReleaseShaderCompiler(GPRELEASESHADERCOMPILER fnptr) {
//   (*fnptr)();
// }
// static void  glowRenderbufferStorage(GPRENDERBUFFERSTORAGE fnptr, GLenum  target, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, internalformat, width, height);
// }
// static void  glowRenderbufferStorageMultisample(GPRENDERBUFFERSTORAGEMULTISAMPLE fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, samples, internalformat, width, height);
// }
// static void  glowRenderbufferStorageMultisampleANGLE(GPRENDERBUFFERSTORAGEMULTISAMPLEANGLE fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, samples, internalformat, width, height);
// }
// static void  glowRenderbufferStorageMultisampleAPPLE(GPRENDERBUFFERSTORAGEMULTISAMPLEAPPLE fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, samples, internalformat, width, height);
// }
// static void  glowRenderbufferStorageMultisampleEXT(GPRENDERBUFFERSTORAGEMULTISAMPLEEXT fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, samples, internalformat, width, height);
// }
// static void  glowRenderbufferStorageMultisampleIMG(GPRENDERBUFFERSTORAGEMULTISAMPLEIMG fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, samples, internalformat, width, height);
// }
// static void  glowRenderbufferStorageMultisampleNV(GPRENDERBUFFERSTORAGEMULTISAMPLENV fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, samples, internalformat, width, height);
// }
// static void  glowResolveDepthValuesNV(GPRESOLVEDEPTHVALUESNV fnptr) {
//   (*fnptr)();
// }
// static void  glowResolveMultisampleFramebufferAPPLE(GPRESOLVEMULTISAMPLEFRAMEBUFFERAPPLE fnptr) {
//   (*fnptr)();
// }
// static void  glowResumeTransformFeedback(GPRESUMETRANSFORMFEEDBACK fnptr) {
//   (*fnptr)();
// }
// static void  glowSampleCoverage(GPSAMPLECOVERAGE fnptr, GLfloat  value, GLboolean  invert) {
//   (*fnptr)(value, invert);
// }
// static void  glowSampleMaski(GPSAMPLEMASKI fnptr, GLuint  maskNumber, GLbitfield  mask) {
//   (*fnptr)(maskNumber, mask);
// }
// static void  glowSamplerParameterIivEXT(GPSAMPLERPARAMETERIIVEXT fnptr, GLuint  sampler, GLenum  pname, const GLint * param) {
//   (*fnptr)(sampler, pname, param);
// }
// static void  glowSamplerParameterIivOES(GPSAMPLERPARAMETERIIVOES fnptr, GLuint  sampler, GLenum  pname, const GLint * param) {
//   (*fnptr)(sampler, pname, param);
// }
// static void  glowSamplerParameterIuivEXT(GPSAMPLERPARAMETERIUIVEXT fnptr, GLuint  sampler, GLenum  pname, const GLuint * param) {
//   (*fnptr)(sampler, pname, param);
// }
// static void  glowSamplerParameterIuivOES(GPSAMPLERPARAMETERIUIVOES fnptr, GLuint  sampler, GLenum  pname, const GLuint * param) {
//   (*fnptr)(sampler, pname, param);
// }
// static void  glowSamplerParameterf(GPSAMPLERPARAMETERF fnptr, GLuint  sampler, GLenum  pname, GLfloat  param) {
//   (*fnptr)(sampler, pname, param);
// }
// static void  glowSamplerParameterfv(GPSAMPLERPARAMETERFV fnptr, GLuint  sampler, GLenum  pname, const GLfloat * param) {
//   (*fnptr)(sampler, pname, param);
// }
// static void  glowSamplerParameteri(GPSAMPLERPARAMETERI fnptr, GLuint  sampler, GLenum  pname, GLint  param) {
//   (*fnptr)(sampler, pname, param);
// }
// static void  glowSamplerParameteriv(GPSAMPLERPARAMETERIV fnptr, GLuint  sampler, GLenum  pname, const GLint * param) {
//   (*fnptr)(sampler, pname, param);
// }
// static void  glowScissor(GPSCISSOR fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(x, y, width, height);
// }
// static void  glowScissorArrayvNV(GPSCISSORARRAYVNV fnptr, GLuint  first, GLsizei  count, const GLint * v) {
//   (*fnptr)(first, count, v);
// }
// static void  glowScissorArrayvOES(GPSCISSORARRAYVOES fnptr, GLuint  first, GLsizei  count, const GLint * v) {
//   (*fnptr)(first, count, v);
// }
// static void  glowScissorIndexedNV(GPSCISSORINDEXEDNV fnptr, GLuint  index, GLint  left, GLint  bottom, GLsizei  width, GLsizei  height) {
//   (*fnptr)(index, left, bottom, width, height);
// }
// static void  glowScissorIndexedOES(GPSCISSORINDEXEDOES fnptr, GLuint  index, GLint  left, GLint  bottom, GLsizei  width, GLsizei  height) {
//   (*fnptr)(index, left, bottom, width, height);
// }
// static void  glowScissorIndexedvNV(GPSCISSORINDEXEDVNV fnptr, GLuint  index, const GLint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowScissorIndexedvOES(GPSCISSORINDEXEDVOES fnptr, GLuint  index, const GLint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowSelectPerfMonitorCountersAMD(GPSELECTPERFMONITORCOUNTERSAMD fnptr, GLuint  monitor, GLboolean  enable, GLuint  group, GLint  numCounters, GLuint * counterList) {
//   (*fnptr)(monitor, enable, group, numCounters, counterList);
// }
// static void  glowSemaphoreParameterui64vEXT(GPSEMAPHOREPARAMETERUI64VEXT fnptr, GLuint  semaphore, GLenum  pname, const GLuint64 * params) {
//   (*fnptr)(semaphore, pname, params);
// }
// static void  glowSetFenceNV(GPSETFENCENV fnptr, GLuint  fence, GLenum  condition) {
//   (*fnptr)(fence, condition);
// }
// static void  glowShaderBinary(GPSHADERBINARY fnptr, GLsizei  count, const GLuint * shaders, GLenum  binaryformat, const void * binary, GLsizei  length) {
//   (*fnptr)(count, shaders, binaryformat, binary, length);
// }
// static void  glowShaderSource(GPSHADERSOURCE fnptr, GLuint  shader, GLsizei  count, const GLchar *const* string, const GLint * length) {
//   (*fnptr)(shader, count, string, length);
// }
// static void  glowSignalSemaphoreEXT(GPSIGNALSEMAPHOREEXT fnptr, GLuint  semaphore, GLuint  numBufferBarriers, const GLuint * buffers, GLuint  numTextureBarriers, const GLuint * textures, const GLenum * dstLayouts) {
//   (*fnptr)(semaphore, numBufferBarriers, buffers, numTextureBarriers, textures, dstLayouts);
// }
// static void  glowSignalVkFenceNV(GPSIGNALVKFENCENV fnptr, GLuint64  vkFence) {
//   (*fnptr)(vkFence);
// }
// static void  glowSignalVkSemaphoreNV(GPSIGNALVKSEMAPHORENV fnptr, GLuint64  vkSemaphore) {
//   (*fnptr)(vkSemaphore);
// }
// static void  glowStartTilingQCOM(GPSTARTTILINGQCOM fnptr, GLuint  x, GLuint  y, GLuint  width, GLuint  height, GLbitfield  preserveMask) {
//   (*fnptr)(x, y, width, height, preserveMask);
// }
// static void  glowStencilFillPathInstancedNV(GPSTENCILFILLPATHINSTANCEDNV fnptr, GLsizei  numPaths, GLenum  pathNameType, const void * paths, GLuint  pathBase, GLenum  fillMode, GLuint  mask, GLenum  transformType, const GLfloat * transformValues) {
//   (*fnptr)(numPaths, pathNameType, paths, pathBase, fillMode, mask, transformType, transformValues);
// }
// static void  glowStencilFillPathNV(GPSTENCILFILLPATHNV fnptr, GLuint  path, GLenum  fillMode, GLuint  mask) {
//   (*fnptr)(path, fillMode, mask);
// }
// static void  glowStencilFunc(GPSTENCILFUNC fnptr, GLenum  func, GLint  ref, GLuint  mask) {
//   (*fnptr)(func, ref, mask);
// }
// static void  glowStencilFuncSeparate(GPSTENCILFUNCSEPARATE fnptr, GLenum  face, GLenum  func, GLint  ref, GLuint  mask) {
//   (*fnptr)(face, func, ref, mask);
// }
// static void  glowStencilMask(GPSTENCILMASK fnptr, GLuint  mask) {
//   (*fnptr)(mask);
// }
// static void  glowStencilMaskSeparate(GPSTENCILMASKSEPARATE fnptr, GLenum  face, GLuint  mask) {
//   (*fnptr)(face, mask);
// }
// static void  glowStencilOp(GPSTENCILOP fnptr, GLenum  fail, GLenum  zfail, GLenum  zpass) {
//   (*fnptr)(fail, zfail, zpass);
// }
// static void  glowStencilOpSeparate(GPSTENCILOPSEPARATE fnptr, GLenum  face, GLenum  sfail, GLenum  dpfail, GLenum  dppass) {
//   (*fnptr)(face, sfail, dpfail, dppass);
// }
// static void  glowStencilStrokePathInstancedNV(GPSTENCILSTROKEPATHINSTANCEDNV fnptr, GLsizei  numPaths, GLenum  pathNameType, const void * paths, GLuint  pathBase, GLint  reference, GLuint  mask, GLenum  transformType, const GLfloat * transformValues) {
//   (*fnptr)(numPaths, pathNameType, paths, pathBase, reference, mask, transformType, transformValues);
// }
// static void  glowStencilStrokePathNV(GPSTENCILSTROKEPATHNV fnptr, GLuint  path, GLint  reference, GLuint  mask) {
//   (*fnptr)(path, reference, mask);
// }
// static void  glowStencilThenCoverFillPathInstancedNV(GPSTENCILTHENCOVERFILLPATHINSTANCEDNV fnptr, GLsizei  numPaths, GLenum  pathNameType, const void * paths, GLuint  pathBase, GLenum  fillMode, GLuint  mask, GLenum  coverMode, GLenum  transformType, const GLfloat * transformValues) {
//   (*fnptr)(numPaths, pathNameType, paths, pathBase, fillMode, mask, coverMode, transformType, transformValues);
// }
// static void  glowStencilThenCoverFillPathNV(GPSTENCILTHENCOVERFILLPATHNV fnptr, GLuint  path, GLenum  fillMode, GLuint  mask, GLenum  coverMode) {
//   (*fnptr)(path, fillMode, mask, coverMode);
// }
// static void  glowStencilThenCoverStrokePathInstancedNV(GPSTENCILTHENCOVERSTROKEPATHINSTANCEDNV fnptr, GLsizei  numPaths, GLenum  pathNameType, const void * paths, GLuint  pathBase, GLint  reference, GLuint  mask, GLenum  coverMode, GLenum  transformType, const GLfloat * transformValues) {
//   (*fnptr)(numPaths, pathNameType, paths, pathBase, reference, mask, coverMode, transformType, transformValues);
// }
// static void  glowStencilThenCoverStrokePathNV(GPSTENCILTHENCOVERSTROKEPATHNV fnptr, GLuint  path, GLint  reference, GLuint  mask, GLenum  coverMode) {
//   (*fnptr)(path, reference, mask, coverMode);
// }
// static void  glowSubpixelPrecisionBiasNV(GPSUBPIXELPRECISIONBIASNV fnptr, GLuint  xbits, GLuint  ybits) {
//   (*fnptr)(xbits, ybits);
// }
// static GLboolean  glowTestFenceNV(GPTESTFENCENV fnptr, GLuint  fence) {
//   return (*fnptr)(fence);
// }
// static void  glowTexBufferEXT(GPTEXBUFFEREXT fnptr, GLenum  target, GLenum  internalformat, GLuint  buffer) {
//   (*fnptr)(target, internalformat, buffer);
// }
// static void  glowTexBufferOES(GPTEXBUFFEROES fnptr, GLenum  target, GLenum  internalformat, GLuint  buffer) {
//   (*fnptr)(target, internalformat, buffer);
// }
// static void  glowTexBufferRangeEXT(GPTEXBUFFERRANGEEXT fnptr, GLenum  target, GLenum  internalformat, GLuint  buffer, GLintptr  offset, GLsizeiptr  size) {
//   (*fnptr)(target, internalformat, buffer, offset, size);
// }
// static void  glowTexBufferRangeOES(GPTEXBUFFERRANGEOES fnptr, GLenum  target, GLenum  internalformat, GLuint  buffer, GLintptr  offset, GLsizeiptr  size) {
//   (*fnptr)(target, internalformat, buffer, offset, size);
// }
// static void  glowTexImage2D(GPTEXIMAGE2D fnptr, GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLsizei  height, GLint  border, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(target, level, internalformat, width, height, border, format, type, pixels);
// }
// static void  glowTexImage3D(GPTEXIMAGE3D fnptr, GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLint  border, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(target, level, internalformat, width, height, depth, border, format, type, pixels);
// }
// static void  glowTexImage3DOES(GPTEXIMAGE3DOES fnptr, GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLint  border, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(target, level, internalformat, width, height, depth, border, format, type, pixels);
// }
// static void  glowTexPageCommitmentEXT(GPTEXPAGECOMMITMENTEXT fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  commit) {
//   (*fnptr)(target, level, xoffset, yoffset, zoffset, width, height, depth, commit);
// }
// static void  glowTexParameterIivEXT(GPTEXPARAMETERIIVEXT fnptr, GLenum  target, GLenum  pname, const GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowTexParameterIivOES(GPTEXPARAMETERIIVOES fnptr, GLenum  target, GLenum  pname, const GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowTexParameterIuivEXT(GPTEXPARAMETERIUIVEXT fnptr, GLenum  target, GLenum  pname, const GLuint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowTexParameterIuivOES(GPTEXPARAMETERIUIVOES fnptr, GLenum  target, GLenum  pname, const GLuint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowTexParameterf(GPTEXPARAMETERF fnptr, GLenum  target, GLenum  pname, GLfloat  param) {
//   (*fnptr)(target, pname, param);
// }
// static void  glowTexParameterfv(GPTEXPARAMETERFV fnptr, GLenum  target, GLenum  pname, const GLfloat * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowTexParameteri(GPTEXPARAMETERI fnptr, GLenum  target, GLenum  pname, GLint  param) {
//   (*fnptr)(target, pname, param);
// }
// static void  glowTexParameteriv(GPTEXPARAMETERIV fnptr, GLenum  target, GLenum  pname, const GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowTexStorage1DEXT(GPTEXSTORAGE1DEXT fnptr, GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width) {
//   (*fnptr)(target, levels, internalformat, width);
// }
// static void  glowTexStorage2D(GPTEXSTORAGE2D fnptr, GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, levels, internalformat, width, height);
// }
// static void  glowTexStorage2DEXT(GPTEXSTORAGE2DEXT fnptr, GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, levels, internalformat, width, height);
// }
// static void  glowTexStorage2DMultisample(GPTEXSTORAGE2DMULTISAMPLE fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLboolean  fixedsamplelocations) {
//   (*fnptr)(target, samples, internalformat, width, height, fixedsamplelocations);
// }
// static void  glowTexStorage3D(GPTEXSTORAGE3D fnptr, GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth) {
//   (*fnptr)(target, levels, internalformat, width, height, depth);
// }
// static void  glowTexStorage3DEXT(GPTEXSTORAGE3DEXT fnptr, GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth) {
//   (*fnptr)(target, levels, internalformat, width, height, depth);
// }
// static void  glowTexStorage3DMultisampleOES(GPTEXSTORAGE3DMULTISAMPLEOES fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedsamplelocations) {
//   (*fnptr)(target, samples, internalformat, width, height, depth, fixedsamplelocations);
// }
// static void  glowTexStorageMem1DEXT(GPTEXSTORAGEMEM1DEXT fnptr, GLenum  target, GLsizei  levels, GLenum  internalFormat, GLsizei  width, GLuint  memory, GLuint64  offset) {
//   (*fnptr)(target, levels, internalFormat, width, memory, offset);
// }
// static void  glowTexStorageMem2DEXT(GPTEXSTORAGEMEM2DEXT fnptr, GLenum  target, GLsizei  levels, GLenum  internalFormat, GLsizei  width, GLsizei  height, GLuint  memory, GLuint64  offset) {
//   (*fnptr)(target, levels, internalFormat, width, height, memory, offset);
// }
// static void  glowTexStorageMem2DMultisampleEXT(GPTEXSTORAGEMEM2DMULTISAMPLEEXT fnptr, GLenum  target, GLsizei  samples, GLenum  internalFormat, GLsizei  width, GLsizei  height, GLboolean  fixedSampleLocations, GLuint  memory, GLuint64  offset) {
//   (*fnptr)(target, samples, internalFormat, width, height, fixedSampleLocations, memory, offset);
// }
// static void  glowTexStorageMem3DEXT(GPTEXSTORAGEMEM3DEXT fnptr, GLenum  target, GLsizei  levels, GLenum  internalFormat, GLsizei  width, GLsizei  height, GLsizei  depth, GLuint  memory, GLuint64  offset) {
//   (*fnptr)(target, levels, internalFormat, width, height, depth, memory, offset);
// }
// static void  glowTexStorageMem3DMultisampleEXT(GPTEXSTORAGEMEM3DMULTISAMPLEEXT fnptr, GLenum  target, GLsizei  samples, GLenum  internalFormat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedSampleLocations, GLuint  memory, GLuint64  offset) {
//   (*fnptr)(target, samples, internalFormat, width, height, depth, fixedSampleLocations, memory, offset);
// }
// static void  glowTexSubImage2D(GPTEXSUBIMAGE2D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(target, level, xoffset, yoffset, width, height, format, type, pixels);
// }
// static void  glowTexSubImage3D(GPTEXSUBIMAGE3D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
// }
// static void  glowTexSubImage3DOES(GPTEXSUBIMAGE3DOES fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
// }
// static void  glowTextureFoveationParametersQCOM(GPTEXTUREFOVEATIONPARAMETERSQCOM fnptr, GLuint  texture, GLuint  layer, GLuint  focalPoint, GLfloat  focalX, GLfloat  focalY, GLfloat  gainX, GLfloat  gainY, GLfloat  foveaArea) {
//   (*fnptr)(texture, layer, focalPoint, focalX, focalY, gainX, gainY, foveaArea);
// }
// static void  glowTextureStorage1DEXT(GPTEXTURESTORAGE1DEXT fnptr, GLuint  texture, GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width) {
//   (*fnptr)(texture, target, levels, internalformat, width);
// }
// static void  glowTextureStorage2DEXT(GPTEXTURESTORAGE2DEXT fnptr, GLuint  texture, GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(texture, target, levels, internalformat, width, height);
// }
// static void  glowTextureStorage3DEXT(GPTEXTURESTORAGE3DEXT fnptr, GLuint  texture, GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth) {
//   (*fnptr)(texture, target, levels, internalformat, width, height, depth);
// }
// static void  glowTextureStorageMem1DEXT(GPTEXTURESTORAGEMEM1DEXT fnptr, GLuint  texture, GLsizei  levels, GLenum  internalFormat, GLsizei  width, GLuint  memory, GLuint64  offset) {
//   (*fnptr)(texture, levels, internalFormat, width, memory, offset);
// }
// static void  glowTextureStorageMem2DEXT(GPTEXTURESTORAGEMEM2DEXT fnptr, GLuint  texture, GLsizei  levels, GLenum  internalFormat, GLsizei  width, GLsizei  height, GLuint  memory, GLuint64  offset) {
//   (*fnptr)(texture, levels, internalFormat, width, height, memory, offset);
// }
// static void  glowTextureStorageMem2DMultisampleEXT(GPTEXTURESTORAGEMEM2DMULTISAMPLEEXT fnptr, GLuint  texture, GLsizei  samples, GLenum  internalFormat, GLsizei  width, GLsizei  height, GLboolean  fixedSampleLocations, GLuint  memory, GLuint64  offset) {
//   (*fnptr)(texture, samples, internalFormat, width, height, fixedSampleLocations, memory, offset);
// }
// static void  glowTextureStorageMem3DEXT(GPTEXTURESTORAGEMEM3DEXT fnptr, GLuint  texture, GLsizei  levels, GLenum  internalFormat, GLsizei  width, GLsizei  height, GLsizei  depth, GLuint  memory, GLuint64  offset) {
//   (*fnptr)(texture, levels, internalFormat, width, height, depth, memory, offset);
// }
// static void  glowTextureStorageMem3DMultisampleEXT(GPTEXTURESTORAGEMEM3DMULTISAMPLEEXT fnptr, GLuint  texture, GLsizei  samples, GLenum  internalFormat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedSampleLocations, GLuint  memory, GLuint64  offset) {
//   (*fnptr)(texture, samples, internalFormat, width, height, depth, fixedSampleLocations, memory, offset);
// }
// static void  glowTextureViewEXT(GPTEXTUREVIEWEXT fnptr, GLuint  texture, GLenum  target, GLuint  origtexture, GLenum  internalformat, GLuint  minlevel, GLuint  numlevels, GLuint  minlayer, GLuint  numlayers) {
//   (*fnptr)(texture, target, origtexture, internalformat, minlevel, numlevels, minlayer, numlayers);
// }
// static void  glowTextureViewOES(GPTEXTUREVIEWOES fnptr, GLuint  texture, GLenum  target, GLuint  origtexture, GLenum  internalformat, GLuint  minlevel, GLuint  numlevels, GLuint  minlayer, GLuint  numlayers) {
//   (*fnptr)(texture, target, origtexture, internalformat, minlevel, numlevels, minlayer, numlayers);
// }
// static void  glowTransformFeedbackVaryings(GPTRANSFORMFEEDBACKVARYINGS fnptr, GLuint  program, GLsizei  count, const GLchar *const* varyings, GLenum  bufferMode) {
//   (*fnptr)(program, count, varyings, bufferMode);
// }
// static void  glowTransformPathNV(GPTRANSFORMPATHNV fnptr, GLuint  resultPath, GLuint  srcPath, GLenum  transformType, const GLfloat * transformValues) {
//   (*fnptr)(resultPath, srcPath, transformType, transformValues);
// }
// static void  glowUniform1f(GPUNIFORM1F fnptr, GLint  location, GLfloat  v0) {
//   (*fnptr)(location, v0);
// }
// static void  glowUniform1fv(GPUNIFORM1FV fnptr, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform1i(GPUNIFORM1I fnptr, GLint  location, GLint  v0) {
//   (*fnptr)(location, v0);
// }
// static void  glowUniform1i64NV(GPUNIFORM1I64NV fnptr, GLint  location, GLint64EXT  x) {
//   (*fnptr)(location, x);
// }
// static void  glowUniform1i64vNV(GPUNIFORM1I64VNV fnptr, GLint  location, GLsizei  count, const GLint64EXT * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform1iv(GPUNIFORM1IV fnptr, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform1ui(GPUNIFORM1UI fnptr, GLint  location, GLuint  v0) {
//   (*fnptr)(location, v0);
// }
// static void  glowUniform1ui64NV(GPUNIFORM1UI64NV fnptr, GLint  location, GLuint64EXT  x) {
//   (*fnptr)(location, x);
// }
// static void  glowUniform1ui64vNV(GPUNIFORM1UI64VNV fnptr, GLint  location, GLsizei  count, const GLuint64EXT * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform1uiv(GPUNIFORM1UIV fnptr, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform2f(GPUNIFORM2F fnptr, GLint  location, GLfloat  v0, GLfloat  v1) {
//   (*fnptr)(location, v0, v1);
// }
// static void  glowUniform2fv(GPUNIFORM2FV fnptr, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform2i(GPUNIFORM2I fnptr, GLint  location, GLint  v0, GLint  v1) {
//   (*fnptr)(location, v0, v1);
// }
// static void  glowUniform2i64NV(GPUNIFORM2I64NV fnptr, GLint  location, GLint64EXT  x, GLint64EXT  y) {
//   (*fnptr)(location, x, y);
// }
// static void  glowUniform2i64vNV(GPUNIFORM2I64VNV fnptr, GLint  location, GLsizei  count, const GLint64EXT * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform2iv(GPUNIFORM2IV fnptr, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform2ui(GPUNIFORM2UI fnptr, GLint  location, GLuint  v0, GLuint  v1) {
//   (*fnptr)(location, v0, v1);
// }
// static void  glowUniform2ui64NV(GPUNIFORM2UI64NV fnptr, GLint  location, GLuint64EXT  x, GLuint64EXT  y) {
//   (*fnptr)(location, x, y);
// }
// static void  glowUniform2ui64vNV(GPUNIFORM2UI64VNV fnptr, GLint  location, GLsizei  count, const GLuint64EXT * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform2uiv(GPUNIFORM2UIV fnptr, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform3f(GPUNIFORM3F fnptr, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2) {
//   (*fnptr)(location, v0, v1, v2);
// }
// static void  glowUniform3fv(GPUNIFORM3FV fnptr, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform3i(GPUNIFORM3I fnptr, GLint  location, GLint  v0, GLint  v1, GLint  v2) {
//   (*fnptr)(location, v0, v1, v2);
// }
// static void  glowUniform3i64NV(GPUNIFORM3I64NV fnptr, GLint  location, GLint64EXT  x, GLint64EXT  y, GLint64EXT  z) {
//   (*fnptr)(location, x, y, z);
// }
// static void  glowUniform3i64vNV(GPUNIFORM3I64VNV fnptr, GLint  location, GLsizei  count, const GLint64EXT * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform3iv(GPUNIFORM3IV fnptr, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform3ui(GPUNIFORM3UI fnptr, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2) {
//   (*fnptr)(location, v0, v1, v2);
// }
// static void  glowUniform3ui64NV(GPUNIFORM3UI64NV fnptr, GLint  location, GLuint64EXT  x, GLuint64EXT  y, GLuint64EXT  z) {
//   (*fnptr)(location, x, y, z);
// }
// static void  glowUniform3ui64vNV(GPUNIFORM3UI64VNV fnptr, GLint  location, GLsizei  count, const GLuint64EXT * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform3uiv(GPUNIFORM3UIV fnptr, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform4f(GPUNIFORM4F fnptr, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2, GLfloat  v3) {
//   (*fnptr)(location, v0, v1, v2, v3);
// }
// static void  glowUniform4fv(GPUNIFORM4FV fnptr, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform4i(GPUNIFORM4I fnptr, GLint  location, GLint  v0, GLint  v1, GLint  v2, GLint  v3) {
//   (*fnptr)(location, v0, v1, v2, v3);
// }
// static void  glowUniform4i64NV(GPUNIFORM4I64NV fnptr, GLint  location, GLint64EXT  x, GLint64EXT  y, GLint64EXT  z, GLint64EXT  w) {
//   (*fnptr)(location, x, y, z, w);
// }
// static void  glowUniform4i64vNV(GPUNIFORM4I64VNV fnptr, GLint  location, GLsizei  count, const GLint64EXT * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform4iv(GPUNIFORM4IV fnptr, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform4ui(GPUNIFORM4UI fnptr, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2, GLuint  v3) {
//   (*fnptr)(location, v0, v1, v2, v3);
// }
// static void  glowUniform4ui64NV(GPUNIFORM4UI64NV fnptr, GLint  location, GLuint64EXT  x, GLuint64EXT  y, GLuint64EXT  z, GLuint64EXT  w) {
//   (*fnptr)(location, x, y, z, w);
// }
// static void  glowUniform4ui64vNV(GPUNIFORM4UI64VNV fnptr, GLint  location, GLsizei  count, const GLuint64EXT * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform4uiv(GPUNIFORM4UIV fnptr, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniformBlockBinding(GPUNIFORMBLOCKBINDING fnptr, GLuint  program, GLuint  uniformBlockIndex, GLuint  uniformBlockBinding) {
//   (*fnptr)(program, uniformBlockIndex, uniformBlockBinding);
// }
// static void  glowUniformHandleui64IMG(GPUNIFORMHANDLEUI64IMG fnptr, GLint  location, GLuint64  value) {
//   (*fnptr)(location, value);
// }
// static void  glowUniformHandleui64NV(GPUNIFORMHANDLEUI64NV fnptr, GLint  location, GLuint64  value) {
//   (*fnptr)(location, value);
// }
// static void  glowUniformHandleui64vIMG(GPUNIFORMHANDLEUI64VIMG fnptr, GLint  location, GLsizei  count, const GLuint64 * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniformHandleui64vNV(GPUNIFORMHANDLEUI64VNV fnptr, GLint  location, GLsizei  count, const GLuint64 * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniformMatrix2fv(GPUNIFORMMATRIX2FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix2x3fv(GPUNIFORMMATRIX2X3FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix2x3fvNV(GPUNIFORMMATRIX2X3FVNV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix2x4fv(GPUNIFORMMATRIX2X4FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix2x4fvNV(GPUNIFORMMATRIX2X4FVNV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix3fv(GPUNIFORMMATRIX3FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix3x2fv(GPUNIFORMMATRIX3X2FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix3x2fvNV(GPUNIFORMMATRIX3X2FVNV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix3x4fv(GPUNIFORMMATRIX3X4FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix3x4fvNV(GPUNIFORMMATRIX3X4FVNV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix4fv(GPUNIFORMMATRIX4FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix4x2fv(GPUNIFORMMATRIX4X2FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix4x2fvNV(GPUNIFORMMATRIX4X2FVNV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix4x3fv(GPUNIFORMMATRIX4X3FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix4x3fvNV(GPUNIFORMMATRIX4X3FVNV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static GLboolean  glowUnmapBuffer(GPUNMAPBUFFER fnptr, GLenum  target) {
//   return (*fnptr)(target);
// }
// static GLboolean  glowUnmapBufferOES(GPUNMAPBUFFEROES fnptr, GLenum  target) {
//   return (*fnptr)(target);
// }
// static void  glowUseProgram(GPUSEPROGRAM fnptr, GLuint  program) {
//   (*fnptr)(program);
// }
// static void  glowUseProgramStages(GPUSEPROGRAMSTAGES fnptr, GLuint  pipeline, GLbitfield  stages, GLuint  program) {
//   (*fnptr)(pipeline, stages, program);
// }
// static void  glowUseProgramStagesEXT(GPUSEPROGRAMSTAGESEXT fnptr, GLuint  pipeline, GLbitfield  stages, GLuint  program) {
//   (*fnptr)(pipeline, stages, program);
// }
// static void  glowUseShaderProgramEXT(GPUSESHADERPROGRAMEXT fnptr, GLenum  type, GLuint  program) {
//   (*fnptr)(type, program);
// }
// static void  glowValidateProgram(GPVALIDATEPROGRAM fnptr, GLuint  program) {
//   (*fnptr)(program);
// }
// static void  glowValidateProgramPipeline(GPVALIDATEPROGRAMPIPELINE fnptr, GLuint  pipeline) {
//   (*fnptr)(pipeline);
// }
// static void  glowValidateProgramPipelineEXT(GPVALIDATEPROGRAMPIPELINEEXT fnptr, GLuint  pipeline) {
//   (*fnptr)(pipeline);
// }
// static void  glowVertexAttrib1f(GPVERTEXATTRIB1F fnptr, GLuint  index, GLfloat  x) {
//   (*fnptr)(index, x);
// }
// static void  glowVertexAttrib1fv(GPVERTEXATTRIB1FV fnptr, GLuint  index, const GLfloat * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib2f(GPVERTEXATTRIB2F fnptr, GLuint  index, GLfloat  x, GLfloat  y) {
//   (*fnptr)(index, x, y);
// }
// static void  glowVertexAttrib2fv(GPVERTEXATTRIB2FV fnptr, GLuint  index, const GLfloat * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib3f(GPVERTEXATTRIB3F fnptr, GLuint  index, GLfloat  x, GLfloat  y, GLfloat  z) {
//   (*fnptr)(index, x, y, z);
// }
// static void  glowVertexAttrib3fv(GPVERTEXATTRIB3FV fnptr, GLuint  index, const GLfloat * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4f(GPVERTEXATTRIB4F fnptr, GLuint  index, GLfloat  x, GLfloat  y, GLfloat  z, GLfloat  w) {
//   (*fnptr)(index, x, y, z, w);
// }
// static void  glowVertexAttrib4fv(GPVERTEXATTRIB4FV fnptr, GLuint  index, const GLfloat * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribBinding(GPVERTEXATTRIBBINDING fnptr, GLuint  attribindex, GLuint  bindingindex) {
//   (*fnptr)(attribindex, bindingindex);
// }
// static void  glowVertexAttribDivisor(GPVERTEXATTRIBDIVISOR fnptr, GLuint  index, GLuint  divisor) {
//   (*fnptr)(index, divisor);
// }
// static void  glowVertexAttribDivisorANGLE(GPVERTEXATTRIBDIVISORANGLE fnptr, GLuint  index, GLuint  divisor) {
//   (*fnptr)(index, divisor);
// }
// static void  glowVertexAttribDivisorEXT(GPVERTEXATTRIBDIVISOREXT fnptr, GLuint  index, GLuint  divisor) {
//   (*fnptr)(index, divisor);
// }
// static void  glowVertexAttribDivisorNV(GPVERTEXATTRIBDIVISORNV fnptr, GLuint  index, GLuint  divisor) {
//   (*fnptr)(index, divisor);
// }
// static void  glowVertexAttribFormat(GPVERTEXATTRIBFORMAT fnptr, GLuint  attribindex, GLint  size, GLenum  type, GLboolean  normalized, GLuint  relativeoffset) {
//   (*fnptr)(attribindex, size, type, normalized, relativeoffset);
// }
// static void  glowVertexAttribI4i(GPVERTEXATTRIBI4I fnptr, GLuint  index, GLint  x, GLint  y, GLint  z, GLint  w) {
//   (*fnptr)(index, x, y, z, w);
// }
// static void  glowVertexAttribI4iv(GPVERTEXATTRIBI4IV fnptr, GLuint  index, const GLint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI4ui(GPVERTEXATTRIBI4UI fnptr, GLuint  index, GLuint  x, GLuint  y, GLuint  z, GLuint  w) {
//   (*fnptr)(index, x, y, z, w);
// }
// static void  glowVertexAttribI4uiv(GPVERTEXATTRIBI4UIV fnptr, GLuint  index, const GLuint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribIFormat(GPVERTEXATTRIBIFORMAT fnptr, GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset) {
//   (*fnptr)(attribindex, size, type, relativeoffset);
// }
// static void  glowVertexAttribIPointer(GPVERTEXATTRIBIPOINTER fnptr, GLuint  index, GLint  size, GLenum  type, GLsizei  stride, const void * pointer) {
//   (*fnptr)(index, size, type, stride, pointer);
// }
// static void  glowVertexAttribPointer(GPVERTEXATTRIBPOINTER fnptr, GLuint  index, GLint  size, GLenum  type, GLboolean  normalized, GLsizei  stride, const void * pointer) {
//   (*fnptr)(index, size, type, normalized, stride, pointer);
// }
// static void  glowVertexBindingDivisor(GPVERTEXBINDINGDIVISOR fnptr, GLuint  bindingindex, GLuint  divisor) {
//   (*fnptr)(bindingindex, divisor);
// }
// static void  glowViewport(GPVIEWPORT fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(x, y, width, height);
// }
// static void  glowViewportArrayvNV(GPVIEWPORTARRAYVNV fnptr, GLuint  first, GLsizei  count, const GLfloat * v) {
//   (*fnptr)(first, count, v);
// }
// static void  glowViewportArrayvOES(GPVIEWPORTARRAYVOES fnptr, GLuint  first, GLsizei  count, const GLfloat * v) {
//   (*fnptr)(first, count, v);
// }
// static void  glowViewportIndexedfNV(GPVIEWPORTINDEXEDFNV fnptr, GLuint  index, GLfloat  x, GLfloat  y, GLfloat  w, GLfloat  h) {
//   (*fnptr)(index, x, y, w, h);
// }
// static void  glowViewportIndexedfOES(GPVIEWPORTINDEXEDFOES fnptr, GLuint  index, GLfloat  x, GLfloat  y, GLfloat  w, GLfloat  h) {
//   (*fnptr)(index, x, y, w, h);
// }
// static void  glowViewportIndexedfvNV(GPVIEWPORTINDEXEDFVNV fnptr, GLuint  index, const GLfloat * v) {
//   (*fnptr)(index, v);
// }
// static void  glowViewportIndexedfvOES(GPVIEWPORTINDEXEDFVOES fnptr, GLuint  index, const GLfloat * v) {
//   (*fnptr)(index, v);
// }
// static void  glowViewportPositionWScaleNV(GPVIEWPORTPOSITIONWSCALENV fnptr, GLuint  index, GLfloat  xcoeff, GLfloat  ycoeff) {
//   (*fnptr)(index, xcoeff, ycoeff);
// }
// static void  glowViewportSwizzleNV(GPVIEWPORTSWIZZLENV fnptr, GLuint  index, GLenum  swizzlex, GLenum  swizzley, GLenum  swizzlez, GLenum  swizzlew) {
//   (*fnptr)(index, swizzlex, swizzley, swizzlez, swizzlew);
// }
// static void  glowWaitSemaphoreEXT(GPWAITSEMAPHOREEXT fnptr, GLuint  semaphore, GLuint  numBufferBarriers, const GLuint * buffers, GLuint  numTextureBarriers, const GLuint * textures, const GLenum * srcLayouts) {
//   (*fnptr)(semaphore, numBufferBarriers, buffers, numTextureBarriers, textures, srcLayouts);
// }
// static void  glowWaitSync(GPWAITSYNC fnptr, GLsync  sync, GLbitfield  flags, GLuint64  timeout) {
//   (*fnptr)(sync, flags, timeout);
// }
// static void  glowWaitSyncAPPLE(GPWAITSYNCAPPLE fnptr, GLsync  sync, GLbitfield  flags, GLuint64  timeout) {
//   (*fnptr)(sync, flags, timeout);
// }
// static void  glowWaitVkSemaphoreNV(GPWAITVKSEMAPHORENV fnptr, GLuint64  vkSemaphore) {
//   (*fnptr)(vkSemaphore);
// }
// static void  glowWeightPathsNV(GPWEIGHTPATHSNV fnptr, GLuint  resultPath, GLsizei  numPaths, const GLuint * paths, const GLfloat * weights) {
//   (*fnptr)(resultPath, numPaths, paths, weights);
// }
// static void  glowWindowRectanglesEXT(GPWINDOWRECTANGLESEXT fnptr, GLenum  mode, GLsizei  count, const GLint * box) {
//   (*fnptr)(mode, count, box);
// }
import "C"
import (
	"errors"
	"unsafe"
)

const (
	GL_3DC_XY_AMD                                                         = 0x87FA
	GL_3DC_X_AMD                                                          = 0x87F9
	ACCUM_ADJACENT_PAIRS_NV                                               = 0x90AD
	ACTIVE_ATOMIC_COUNTER_BUFFERS                                         = 0x92D9
	ACTIVE_ATTRIBUTES                                                     = 0x8B89
	ACTIVE_ATTRIBUTE_MAX_LENGTH                                           = 0x8B8A
	ACTIVE_PROGRAM                                                        = 0x8259
	ACTIVE_PROGRAM_EXT                                                    = 0x8259
	ACTIVE_RESOURCES                                                      = 0x92F5
	ACTIVE_TEXTURE                                                        = 0x84E0
	ACTIVE_UNIFORMS                                                       = 0x8B86
	ACTIVE_UNIFORM_BLOCKS                                                 = 0x8A36
	ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH                                  = 0x8A35
	ACTIVE_UNIFORM_MAX_LENGTH                                             = 0x8B87
	ACTIVE_VARIABLES                                                      = 0x9305
	ADJACENT_PAIRS_NV                                                     = 0x90AE
	AFFINE_2D_NV                                                          = 0x9092
	AFFINE_3D_NV                                                          = 0x9094
	ALIASED_LINE_WIDTH_RANGE                                              = 0x846E
	ALIASED_POINT_SIZE_RANGE                                              = 0x846D
	ALL_BARRIER_BITS                                                      = 0xFFFFFFFF
	ALL_COMPLETED_NV                                                      = 0x84F2
	ALL_SHADER_BITS                                                       = 0xFFFFFFFF
	ALL_SHADER_BITS_EXT                                                   = 0xFFFFFFFF
	ALPHA                                                                 = 0x1906
	ALPHA16F_EXT                                                          = 0x881C
	ALPHA32F_EXT                                                          = 0x8816
	ALPHA8_EXT                                                            = 0x803C
	ALPHA8_OES                                                            = 0x803C
	ALPHA_BITS                                                            = 0x0D55
	ALPHA_TEST_FUNC_QCOM                                                  = 0x0BC1
	ALPHA_TEST_QCOM                                                       = 0x0BC0
	ALPHA_TEST_REF_QCOM                                                   = 0x0BC2
	ALREADY_SIGNALED                                                      = 0x911A
	ALREADY_SIGNALED_APPLE                                                = 0x911A
	ALWAYS                                                                = 0x0207
	ANY_SAMPLES_PASSED                                                    = 0x8C2F
	ANY_SAMPLES_PASSED_CONSERVATIVE                                       = 0x8D6A
	ANY_SAMPLES_PASSED_CONSERVATIVE_EXT                                   = 0x8D6A
	ANY_SAMPLES_PASSED_EXT                                                = 0x8C2F
	ARC_TO_NV                                                             = 0xFE
	ARRAY_BUFFER                                                          = 0x8892
	ARRAY_BUFFER_BINDING                                                  = 0x8894
	ARRAY_SIZE                                                            = 0x92FB
	ARRAY_STRIDE                                                          = 0x92FE
	ATC_RGBA_EXPLICIT_ALPHA_AMD                                           = 0x8C93
	ATC_RGBA_INTERPOLATED_ALPHA_AMD                                       = 0x87EE
	ATC_RGB_AMD                                                           = 0x8C92
	ATOMIC_COUNTER_BARRIER_BIT                                            = 0x00001000
	ATOMIC_COUNTER_BUFFER                                                 = 0x92C0
	ATOMIC_COUNTER_BUFFER_BINDING                                         = 0x92C1
	ATOMIC_COUNTER_BUFFER_INDEX                                           = 0x9301
	ATOMIC_COUNTER_BUFFER_SIZE                                            = 0x92C3
	ATOMIC_COUNTER_BUFFER_START                                           = 0x92C2
	ATTACHED_SHADERS                                                      = 0x8B85
	BACK                                                                  = 0x0405
	BEVEL_NV                                                              = 0x90A6
	BGRA8_EXT                                                             = 0x93A1
	BGRA_EXT                                                              = 0x80E1
	BGRA_IMG                                                              = 0x80E1
	BINNING_CONTROL_HINT_QCOM                                             = 0x8FB0
	BLACKHOLE_RENDER_INTEL                                                = 0x83FC
	BLEND                                                                 = 0x0BE2
	BLEND_ADVANCED_COHERENT_KHR                                           = 0x9285
	BLEND_ADVANCED_COHERENT_NV                                            = 0x9285
	BLEND_COLOR                                                           = 0x8005
	BLEND_DST_ALPHA                                                       = 0x80CA
	BLEND_DST_RGB                                                         = 0x80C8
	BLEND_EQUATION                                                        = 0x8009
	BLEND_EQUATION_ALPHA                                                  = 0x883D
	BLEND_EQUATION_EXT                                                    = 0x8009
	BLEND_EQUATION_RGB                                                    = 0x8009
	BLEND_OVERLAP_NV                                                      = 0x9281
	BLEND_PREMULTIPLIED_SRC_NV                                            = 0x9280
	BLEND_SRC_ALPHA                                                       = 0x80CB
	BLEND_SRC_RGB                                                         = 0x80C9
	BLOCK_INDEX                                                           = 0x92FD
	BLUE                                                                  = 0x1905
	BLUE_BITS                                                             = 0x0D54
	BLUE_NV                                                               = 0x1905
	BOLD_BIT_NV                                                           = 0x01
	BOOL                                                                  = 0x8B56
	BOOL_VEC2                                                             = 0x8B57
	BOOL_VEC3                                                             = 0x8B58
	BOOL_VEC4                                                             = 0x8B59
	BOUNDING_BOX_NV                                                       = 0x908D
	BOUNDING_BOX_OF_BOUNDING_BOXES_NV                                     = 0x909C
	BUFFER                                                                = 0x82E0
	BUFFER_ACCESS_FLAGS                                                   = 0x911F
	BUFFER_ACCESS_OES                                                     = 0x88BB
	BUFFER_BINDING                                                        = 0x9302
	BUFFER_DATA_SIZE                                                      = 0x9303
	BUFFER_IMMUTABLE_STORAGE_EXT                                          = 0x821F
	BUFFER_KHR                                                            = 0x82E0
	BUFFER_MAPPED                                                         = 0x88BC
	BUFFER_MAPPED_OES                                                     = 0x88BC
	BUFFER_MAP_LENGTH                                                     = 0x9120
	BUFFER_MAP_OFFSET                                                     = 0x9121
	BUFFER_MAP_POINTER                                                    = 0x88BD
	BUFFER_MAP_POINTER_OES                                                = 0x88BD
	BUFFER_OBJECT_EXT                                                     = 0x9151
	BUFFER_SIZE                                                           = 0x8764
	BUFFER_STORAGE_FLAGS_EXT                                              = 0x8220
	BUFFER_UPDATE_BARRIER_BIT                                             = 0x00000200
	BUFFER_USAGE                                                          = 0x8765
	BUFFER_VARIABLE                                                       = 0x92E5
	BYTE                                                                  = 0x1400
	CCW                                                                   = 0x0901
	CIRCULAR_CCW_ARC_TO_NV                                                = 0xF8
	CIRCULAR_CW_ARC_TO_NV                                                 = 0xFA
	CIRCULAR_TANGENT_ARC_TO_NV                                            = 0xFC
	CLAMP_TO_BORDER_EXT                                                   = 0x812D
	CLAMP_TO_BORDER_NV                                                    = 0x812D
	CLAMP_TO_BORDER_OES                                                   = 0x812D
	CLAMP_TO_EDGE                                                         = 0x812F
	CLIENT_MAPPED_BUFFER_BARRIER_BIT_EXT                                  = 0x00004000
	CLIENT_STORAGE_BIT_EXT                                                = 0x0200
	CLIP_DEPTH_MODE_EXT                                                   = 0x935D
	CLIP_DISTANCE0_APPLE                                                  = 0x3000
	CLIP_DISTANCE0_EXT                                                    = 0x3000
	CLIP_DISTANCE1_APPLE                                                  = 0x3001
	CLIP_DISTANCE1_EXT                                                    = 0x3001
	CLIP_DISTANCE2_APPLE                                                  = 0x3002
	CLIP_DISTANCE2_EXT                                                    = 0x3002
	CLIP_DISTANCE3_APPLE                                                  = 0x3003
	CLIP_DISTANCE3_EXT                                                    = 0x3003
	CLIP_DISTANCE4_APPLE                                                  = 0x3004
	CLIP_DISTANCE4_EXT                                                    = 0x3004
	CLIP_DISTANCE5_APPLE                                                  = 0x3005
	CLIP_DISTANCE5_EXT                                                    = 0x3005
	CLIP_DISTANCE6_APPLE                                                  = 0x3006
	CLIP_DISTANCE6_EXT                                                    = 0x3006
	CLIP_DISTANCE7_APPLE                                                  = 0x3007
	CLIP_DISTANCE7_EXT                                                    = 0x3007
	CLIP_ORIGIN_EXT                                                       = 0x935C
	CLOSE_PATH_NV                                                         = 0x00
	COLOR                                                                 = 0x1800
	COLORBURN_KHR                                                         = 0x929A
	COLORBURN_NV                                                          = 0x929A
	COLORDODGE_KHR                                                        = 0x9299
	COLORDODGE_NV                                                         = 0x9299
	COLOR_ATTACHMENT0                                                     = 0x8CE0
	COLOR_ATTACHMENT0_EXT                                                 = 0x8CE0
	COLOR_ATTACHMENT0_NV                                                  = 0x8CE0
	COLOR_ATTACHMENT1                                                     = 0x8CE1
	COLOR_ATTACHMENT10                                                    = 0x8CEA
	COLOR_ATTACHMENT10_EXT                                                = 0x8CEA
	COLOR_ATTACHMENT10_NV                                                 = 0x8CEA
	COLOR_ATTACHMENT11                                                    = 0x8CEB
	COLOR_ATTACHMENT11_EXT                                                = 0x8CEB
	COLOR_ATTACHMENT11_NV                                                 = 0x8CEB
	COLOR_ATTACHMENT12                                                    = 0x8CEC
	COLOR_ATTACHMENT12_EXT                                                = 0x8CEC
	COLOR_ATTACHMENT12_NV                                                 = 0x8CEC
	COLOR_ATTACHMENT13                                                    = 0x8CED
	COLOR_ATTACHMENT13_EXT                                                = 0x8CED
	COLOR_ATTACHMENT13_NV                                                 = 0x8CED
	COLOR_ATTACHMENT14                                                    = 0x8CEE
	COLOR_ATTACHMENT14_EXT                                                = 0x8CEE
	COLOR_ATTACHMENT14_NV                                                 = 0x8CEE
	COLOR_ATTACHMENT15                                                    = 0x8CEF
	COLOR_ATTACHMENT15_EXT                                                = 0x8CEF
	COLOR_ATTACHMENT15_NV                                                 = 0x8CEF
	COLOR_ATTACHMENT16                                                    = 0x8CF0
	COLOR_ATTACHMENT17                                                    = 0x8CF1
	COLOR_ATTACHMENT18                                                    = 0x8CF2
	COLOR_ATTACHMENT19                                                    = 0x8CF3
	COLOR_ATTACHMENT1_EXT                                                 = 0x8CE1
	COLOR_ATTACHMENT1_NV                                                  = 0x8CE1
	COLOR_ATTACHMENT2                                                     = 0x8CE2
	COLOR_ATTACHMENT20                                                    = 0x8CF4
	COLOR_ATTACHMENT21                                                    = 0x8CF5
	COLOR_ATTACHMENT22                                                    = 0x8CF6
	COLOR_ATTACHMENT23                                                    = 0x8CF7
	COLOR_ATTACHMENT24                                                    = 0x8CF8
	COLOR_ATTACHMENT25                                                    = 0x8CF9
	COLOR_ATTACHMENT26                                                    = 0x8CFA
	COLOR_ATTACHMENT27                                                    = 0x8CFB
	COLOR_ATTACHMENT28                                                    = 0x8CFC
	COLOR_ATTACHMENT29                                                    = 0x8CFD
	COLOR_ATTACHMENT2_EXT                                                 = 0x8CE2
	COLOR_ATTACHMENT2_NV                                                  = 0x8CE2
	COLOR_ATTACHMENT3                                                     = 0x8CE3
	COLOR_ATTACHMENT30                                                    = 0x8CFE
	COLOR_ATTACHMENT31                                                    = 0x8CFF
	COLOR_ATTACHMENT3_EXT                                                 = 0x8CE3
	COLOR_ATTACHMENT3_NV                                                  = 0x8CE3
	COLOR_ATTACHMENT4                                                     = 0x8CE4
	COLOR_ATTACHMENT4_EXT                                                 = 0x8CE4
	COLOR_ATTACHMENT4_NV                                                  = 0x8CE4
	COLOR_ATTACHMENT5                                                     = 0x8CE5
	COLOR_ATTACHMENT5_EXT                                                 = 0x8CE5
	COLOR_ATTACHMENT5_NV                                                  = 0x8CE5
	COLOR_ATTACHMENT6                                                     = 0x8CE6
	COLOR_ATTACHMENT6_EXT                                                 = 0x8CE6
	COLOR_ATTACHMENT6_NV                                                  = 0x8CE6
	COLOR_ATTACHMENT7                                                     = 0x8CE7
	COLOR_ATTACHMENT7_EXT                                                 = 0x8CE7
	COLOR_ATTACHMENT7_NV                                                  = 0x8CE7
	COLOR_ATTACHMENT8                                                     = 0x8CE8
	COLOR_ATTACHMENT8_EXT                                                 = 0x8CE8
	COLOR_ATTACHMENT8_NV                                                  = 0x8CE8
	COLOR_ATTACHMENT9                                                     = 0x8CE9
	COLOR_ATTACHMENT9_EXT                                                 = 0x8CE9
	COLOR_ATTACHMENT9_NV                                                  = 0x8CE9
	COLOR_ATTACHMENT_EXT                                                  = 0x90F0
	COLOR_BUFFER_BIT                                                      = 0x00004000
	COLOR_BUFFER_BIT0_QCOM                                                = 0x00000001
	COLOR_BUFFER_BIT1_QCOM                                                = 0x00000002
	COLOR_BUFFER_BIT2_QCOM                                                = 0x00000004
	COLOR_BUFFER_BIT3_QCOM                                                = 0x00000008
	COLOR_BUFFER_BIT4_QCOM                                                = 0x00000010
	COLOR_BUFFER_BIT5_QCOM                                                = 0x00000020
	COLOR_BUFFER_BIT6_QCOM                                                = 0x00000040
	COLOR_BUFFER_BIT7_QCOM                                                = 0x00000080
	COLOR_CLEAR_VALUE                                                     = 0x0C22
	COLOR_EXT                                                             = 0x1800
	COLOR_SAMPLES_NV                                                      = 0x8E20
	COLOR_WRITEMASK                                                       = 0x0C23
	COMMAND_BARRIER_BIT                                                   = 0x00000040
	COMPARE_REF_TO_TEXTURE                                                = 0x884E
	COMPARE_REF_TO_TEXTURE_EXT                                            = 0x884E
	COMPILE_STATUS                                                        = 0x8B81
	COMPLETION_STATUS_KHR                                                 = 0x91B1
	COMPRESSED_R11_EAC                                                    = 0x9270
	COMPRESSED_RED_GREEN_RGTC2_EXT                                        = 0x8DBD
	COMPRESSED_RED_RGTC1_EXT                                              = 0x8DBB
	COMPRESSED_RG11_EAC                                                   = 0x9272
	COMPRESSED_RGB8_ETC2                                                  = 0x9274
	COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2                              = 0x9276
	COMPRESSED_RGBA8_ETC2_EAC                                             = 0x9278
	COMPRESSED_RGBA_ASTC_10x10_KHR                                        = 0x93BB
	COMPRESSED_RGBA_ASTC_10x5_KHR                                         = 0x93B8
	COMPRESSED_RGBA_ASTC_10x6_KHR                                         = 0x93B9
	COMPRESSED_RGBA_ASTC_10x8_KHR                                         = 0x93BA
	COMPRESSED_RGBA_ASTC_12x10_KHR                                        = 0x93BC
	COMPRESSED_RGBA_ASTC_12x12_KHR                                        = 0x93BD
	COMPRESSED_RGBA_ASTC_3x3x3_OES                                        = 0x93C0
	COMPRESSED_RGBA_ASTC_4x3x3_OES                                        = 0x93C1
	COMPRESSED_RGBA_ASTC_4x4_KHR                                          = 0x93B0
	COMPRESSED_RGBA_ASTC_4x4x3_OES                                        = 0x93C2
	COMPRESSED_RGBA_ASTC_4x4x4_OES                                        = 0x93C3
	COMPRESSED_RGBA_ASTC_5x4_KHR                                          = 0x93B1
	COMPRESSED_RGBA_ASTC_5x4x4_OES                                        = 0x93C4
	COMPRESSED_RGBA_ASTC_5x5_KHR                                          = 0x93B2
	COMPRESSED_RGBA_ASTC_5x5x4_OES                                        = 0x93C5
	COMPRESSED_RGBA_ASTC_5x5x5_OES                                        = 0x93C6
	COMPRESSED_RGBA_ASTC_6x5_KHR                                          = 0x93B3
	COMPRESSED_RGBA_ASTC_6x5x5_OES                                        = 0x93C7
	COMPRESSED_RGBA_ASTC_6x6_KHR                                          = 0x93B4
	COMPRESSED_RGBA_ASTC_6x6x5_OES                                        = 0x93C8
	COMPRESSED_RGBA_ASTC_6x6x6_OES                                        = 0x93C9
	COMPRESSED_RGBA_ASTC_8x5_KHR                                          = 0x93B5
	COMPRESSED_RGBA_ASTC_8x6_KHR                                          = 0x93B6
	COMPRESSED_RGBA_ASTC_8x8_KHR                                          = 0x93B7
	COMPRESSED_RGBA_BPTC_UNORM_EXT                                        = 0x8E8C
	COMPRESSED_RGBA_PVRTC_2BPPV1_IMG                                      = 0x8C03
	COMPRESSED_RGBA_PVRTC_2BPPV2_IMG                                      = 0x9137
	COMPRESSED_RGBA_PVRTC_4BPPV1_IMG                                      = 0x8C02
	COMPRESSED_RGBA_PVRTC_4BPPV2_IMG                                      = 0x9138
	COMPRESSED_RGBA_S3TC_DXT1_EXT                                         = 0x83F1
	COMPRESSED_RGBA_S3TC_DXT3_ANGLE                                       = 0x83F2
	COMPRESSED_RGBA_S3TC_DXT3_EXT                                         = 0x83F2
	COMPRESSED_RGBA_S3TC_DXT5_ANGLE                                       = 0x83F3
	COMPRESSED_RGBA_S3TC_DXT5_EXT                                         = 0x83F3
	COMPRESSED_RGB_BPTC_SIGNED_FLOAT_EXT                                  = 0x8E8E
	COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT_EXT                                = 0x8E8F
	COMPRESSED_RGB_PVRTC_2BPPV1_IMG                                       = 0x8C01
	COMPRESSED_RGB_PVRTC_4BPPV1_IMG                                       = 0x8C00
	COMPRESSED_RGB_S3TC_DXT1_EXT                                          = 0x83F0
	COMPRESSED_SIGNED_R11_EAC                                             = 0x9271
	COMPRESSED_SIGNED_RED_GREEN_RGTC2_EXT                                 = 0x8DBE
	COMPRESSED_SIGNED_RED_RGTC1_EXT                                       = 0x8DBC
	COMPRESSED_SIGNED_RG11_EAC                                            = 0x9273
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR                                = 0x93DB
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR                                 = 0x93D8
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR                                 = 0x93D9
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR                                 = 0x93DA
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR                                = 0x93DC
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR                                = 0x93DD
	COMPRESSED_SRGB8_ALPHA8_ASTC_3x3x3_OES                                = 0x93E0
	COMPRESSED_SRGB8_ALPHA8_ASTC_4x3x3_OES                                = 0x93E1
	COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR                                  = 0x93D0
	COMPRESSED_SRGB8_ALPHA8_ASTC_4x4x3_OES                                = 0x93E2
	COMPRESSED_SRGB8_ALPHA8_ASTC_4x4x4_OES                                = 0x93E3
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR                                  = 0x93D1
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x4x4_OES                                = 0x93E4
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR                                  = 0x93D2
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x5x4_OES                                = 0x93E5
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x5x5_OES                                = 0x93E6
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR                                  = 0x93D3
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x5x5_OES                                = 0x93E7
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR                                  = 0x93D4
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x6x5_OES                                = 0x93E8
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x6x6_OES                                = 0x93E9
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR                                  = 0x93D5
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR                                  = 0x93D6
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR                                  = 0x93D7
	COMPRESSED_SRGB8_ALPHA8_ETC2_EAC                                      = 0x9279
	COMPRESSED_SRGB8_ETC2                                                 = 0x9275
	COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2                             = 0x9277
	COMPRESSED_SRGB_ALPHA_BPTC_UNORM_EXT                                  = 0x8E8D
	COMPRESSED_SRGB_ALPHA_PVRTC_2BPPV1_EXT                                = 0x8A56
	COMPRESSED_SRGB_ALPHA_PVRTC_2BPPV2_IMG                                = 0x93F0
	COMPRESSED_SRGB_ALPHA_PVRTC_4BPPV1_EXT                                = 0x8A57
	COMPRESSED_SRGB_ALPHA_PVRTC_4BPPV2_IMG                                = 0x93F1
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_EXT                                   = 0x8C4D
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_NV                                    = 0x8C4D
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_EXT                                   = 0x8C4E
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_NV                                    = 0x8C4E
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_EXT                                   = 0x8C4F
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_NV                                    = 0x8C4F
	COMPRESSED_SRGB_PVRTC_2BPPV1_EXT                                      = 0x8A54
	COMPRESSED_SRGB_PVRTC_4BPPV1_EXT                                      = 0x8A55
	COMPRESSED_SRGB_S3TC_DXT1_EXT                                         = 0x8C4C
	COMPRESSED_SRGB_S3TC_DXT1_NV                                          = 0x8C4C
	COMPRESSED_TEXTURE_FORMATS                                            = 0x86A3
	COMPUTE_SHADER                                                        = 0x91B9
	COMPUTE_SHADER_BIT                                                    = 0x00000020
	COMPUTE_WORK_GROUP_SIZE                                               = 0x8267
	CONDITION_SATISFIED                                                   = 0x911C
	CONDITION_SATISFIED_APPLE                                             = 0x911C
	CONFORMANT_NV                                                         = 0x9374
	CONIC_CURVE_TO_NV                                                     = 0x1A
	CONJOINT_NV                                                           = 0x9284
	CONSERVATIVE_RASTERIZATION_INTEL                                      = 0x83FE
	CONSERVATIVE_RASTERIZATION_NV                                         = 0x9346
	CONSERVATIVE_RASTER_MODE_NV                                           = 0x954D
	CONSERVATIVE_RASTER_MODE_POST_SNAP_NV                                 = 0x954E
	CONSERVATIVE_RASTER_MODE_PRE_SNAP_NV                                  = 0x9550
	CONSERVATIVE_RASTER_MODE_PRE_SNAP_TRIANGLES_NV                        = 0x954F
	CONSTANT_ALPHA                                                        = 0x8003
	CONSTANT_COLOR                                                        = 0x8001
	CONTEXT_FLAG_DEBUG_BIT                                                = 0x00000002
	CONTEXT_FLAG_DEBUG_BIT_KHR                                            = 0x00000002
	CONTEXT_FLAG_NO_ERROR_BIT_KHR                                         = 0x00000008
	CONTEXT_FLAG_PROTECTED_CONTENT_BIT_EXT                                = 0x00000010
	CONTEXT_LOST                                                          = 0x0507
	CONTEXT_LOST_KHR                                                      = 0x0507
	CONTEXT_RELEASE_BEHAVIOR                                              = 0x82FB
	CONTEXT_RELEASE_BEHAVIOR_FLUSH                                        = 0x82FC
	CONTEXT_RELEASE_BEHAVIOR_FLUSH_KHR                                    = 0x82FC
	CONTEXT_RELEASE_BEHAVIOR_KHR                                          = 0x82FB
	CONTEXT_ROBUST_ACCESS                                                 = 0x90F3
	CONTEXT_ROBUST_ACCESS_EXT                                             = 0x90F3
	CONTEXT_ROBUST_ACCESS_KHR                                             = 0x90F3
	CONTRAST_NV                                                           = 0x92A1
	CONVEX_HULL_NV                                                        = 0x908B
	COPY_READ_BUFFER                                                      = 0x8F36
	COPY_READ_BUFFER_BINDING                                              = 0x8F36
	COPY_READ_BUFFER_NV                                                   = 0x8F36
	COPY_WRITE_BUFFER                                                     = 0x8F37
	COPY_WRITE_BUFFER_BINDING                                             = 0x8F37
	COPY_WRITE_BUFFER_NV                                                  = 0x8F37
	COUNTER_RANGE_AMD                                                     = 0x8BC1
	COUNTER_TYPE_AMD                                                      = 0x8BC0
	COUNT_DOWN_NV                                                         = 0x9089
	COUNT_UP_NV                                                           = 0x9088
	COVERAGE_ALL_FRAGMENTS_NV                                             = 0x8ED5
	COVERAGE_ATTACHMENT_NV                                                = 0x8ED2
	COVERAGE_AUTOMATIC_NV                                                 = 0x8ED7
	COVERAGE_BUFFERS_NV                                                   = 0x8ED3
	COVERAGE_BUFFER_BIT_NV                                                = 0x00008000
	COVERAGE_COMPONENT4_NV                                                = 0x8ED1
	COVERAGE_COMPONENT_NV                                                 = 0x8ED0
	COVERAGE_EDGE_FRAGMENTS_NV                                            = 0x8ED6
	COVERAGE_MODULATION_NV                                                = 0x9332
	COVERAGE_MODULATION_TABLE_NV                                          = 0x9331
	COVERAGE_MODULATION_TABLE_SIZE_NV                                     = 0x9333
	COVERAGE_SAMPLES_NV                                                   = 0x8ED4
	CPU_OPTIMIZED_QCOM                                                    = 0x8FB1
	CUBIC_CURVE_TO_NV                                                     = 0x0C
	CUBIC_IMG                                                             = 0x9139
	CUBIC_MIPMAP_LINEAR_IMG                                               = 0x913B
	CUBIC_MIPMAP_NEAREST_IMG                                              = 0x913A
	CULL_FACE                                                             = 0x0B44
	CULL_FACE_MODE                                                        = 0x0B45
	CURRENT_PROGRAM                                                       = 0x8B8D
	CURRENT_QUERY                                                         = 0x8865
	CURRENT_QUERY_EXT                                                     = 0x8865
	CURRENT_VERTEX_ATTRIB                                                 = 0x8626
	CW                                                                    = 0x0900
	D3D12_FENCE_VALUE_EXT                                                 = 0x9595
	DARKEN_KHR                                                            = 0x9297
	DARKEN_NV                                                             = 0x9297
	DEBUG_CALLBACK_FUNCTION                                               = 0x8244
	DEBUG_CALLBACK_FUNCTION_KHR                                           = 0x8244
	DEBUG_CALLBACK_USER_PARAM                                             = 0x8245
	DEBUG_CALLBACK_USER_PARAM_KHR                                         = 0x8245
	DEBUG_GROUP_STACK_DEPTH                                               = 0x826D
	DEBUG_GROUP_STACK_DEPTH_KHR                                           = 0x826D
	DEBUG_LOGGED_MESSAGES                                                 = 0x9145
	DEBUG_LOGGED_MESSAGES_KHR                                             = 0x9145
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH                                      = 0x8243
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_KHR                                  = 0x8243
	DEBUG_OUTPUT                                                          = 0x92E0
	DEBUG_OUTPUT_KHR                                                      = 0x92E0
	DEBUG_OUTPUT_SYNCHRONOUS                                              = 0x8242
	DEBUG_OUTPUT_SYNCHRONOUS_KHR                                          = 0x8242
	DEBUG_SEVERITY_HIGH                                                   = 0x9146
	DEBUG_SEVERITY_HIGH_KHR                                               = 0x9146
	DEBUG_SEVERITY_LOW                                                    = 0x9148
	DEBUG_SEVERITY_LOW_KHR                                                = 0x9148
	DEBUG_SEVERITY_MEDIUM                                                 = 0x9147
	DEBUG_SEVERITY_MEDIUM_KHR                                             = 0x9147
	DEBUG_SEVERITY_NOTIFICATION                                           = 0x826B
	DEBUG_SEVERITY_NOTIFICATION_KHR                                       = 0x826B
	DEBUG_SOURCE_API                                                      = 0x8246
	DEBUG_SOURCE_API_KHR                                                  = 0x8246
	DEBUG_SOURCE_APPLICATION                                              = 0x824A
	DEBUG_SOURCE_APPLICATION_KHR                                          = 0x824A
	DEBUG_SOURCE_OTHER                                                    = 0x824B
	DEBUG_SOURCE_OTHER_KHR                                                = 0x824B
	DEBUG_SOURCE_SHADER_COMPILER                                          = 0x8248
	DEBUG_SOURCE_SHADER_COMPILER_KHR                                      = 0x8248
	DEBUG_SOURCE_THIRD_PARTY                                              = 0x8249
	DEBUG_SOURCE_THIRD_PARTY_KHR                                          = 0x8249
	DEBUG_SOURCE_WINDOW_SYSTEM                                            = 0x8247
	DEBUG_SOURCE_WINDOW_SYSTEM_KHR                                        = 0x8247
	DEBUG_TYPE_DEPRECATED_BEHAVIOR                                        = 0x824D
	DEBUG_TYPE_DEPRECATED_BEHAVIOR_KHR                                    = 0x824D
	DEBUG_TYPE_ERROR                                                      = 0x824C
	DEBUG_TYPE_ERROR_KHR                                                  = 0x824C
	DEBUG_TYPE_MARKER                                                     = 0x8268
	DEBUG_TYPE_MARKER_KHR                                                 = 0x8268
	DEBUG_TYPE_OTHER                                                      = 0x8251
	DEBUG_TYPE_OTHER_KHR                                                  = 0x8251
	DEBUG_TYPE_PERFORMANCE                                                = 0x8250
	DEBUG_TYPE_PERFORMANCE_KHR                                            = 0x8250
	DEBUG_TYPE_POP_GROUP                                                  = 0x826A
	DEBUG_TYPE_POP_GROUP_KHR                                              = 0x826A
	DEBUG_TYPE_PORTABILITY                                                = 0x824F
	DEBUG_TYPE_PORTABILITY_KHR                                            = 0x824F
	DEBUG_TYPE_PUSH_GROUP                                                 = 0x8269
	DEBUG_TYPE_PUSH_GROUP_KHR                                             = 0x8269
	DEBUG_TYPE_UNDEFINED_BEHAVIOR                                         = 0x824E
	DEBUG_TYPE_UNDEFINED_BEHAVIOR_KHR                                     = 0x824E
	DECODE_EXT                                                            = 0x8A49
	DECR                                                                  = 0x1E03
	DECR_WRAP                                                             = 0x8508
	DEDICATED_MEMORY_OBJECT_EXT                                           = 0x9581
	DELETE_STATUS                                                         = 0x8B80
	DEPTH                                                                 = 0x1801
	DEPTH24_STENCIL8                                                      = 0x88F0
	DEPTH24_STENCIL8_OES                                                  = 0x88F0
	DEPTH32F_STENCIL8                                                     = 0x8CAD
	DEPTH_ATTACHMENT                                                      = 0x8D00
	DEPTH_BITS                                                            = 0x0D56
	DEPTH_BUFFER_BIT                                                      = 0x00000100
	DEPTH_BUFFER_BIT0_QCOM                                                = 0x00000100
	DEPTH_BUFFER_BIT1_QCOM                                                = 0x00000200
	DEPTH_BUFFER_BIT2_QCOM                                                = 0x00000400
	DEPTH_BUFFER_BIT3_QCOM                                                = 0x00000800
	DEPTH_BUFFER_BIT4_QCOM                                                = 0x00001000
	DEPTH_BUFFER_BIT5_QCOM                                                = 0x00002000
	DEPTH_BUFFER_BIT6_QCOM                                                = 0x00004000
	DEPTH_BUFFER_BIT7_QCOM                                                = 0x00008000
	DEPTH_CLEAR_VALUE                                                     = 0x0B73
	DEPTH_COMPONENT                                                       = 0x1902
	DEPTH_COMPONENT16                                                     = 0x81A5
	DEPTH_COMPONENT16_NONLINEAR_NV                                        = 0x8E2C
	DEPTH_COMPONENT16_OES                                                 = 0x81A5
	DEPTH_COMPONENT24                                                     = 0x81A6
	DEPTH_COMPONENT24_OES                                                 = 0x81A6
	DEPTH_COMPONENT32F                                                    = 0x8CAC
	DEPTH_COMPONENT32_OES                                                 = 0x81A7
	DEPTH_EXT                                                             = 0x1801
	DEPTH_FUNC                                                            = 0x0B74
	DEPTH_RANGE                                                           = 0x0B70
	DEPTH_SAMPLES_NV                                                      = 0x932D
	DEPTH_STENCIL                                                         = 0x84F9
	DEPTH_STENCIL_ATTACHMENT                                              = 0x821A
	DEPTH_STENCIL_OES                                                     = 0x84F9
	DEPTH_STENCIL_TEXTURE_MODE                                            = 0x90EA
	DEPTH_TEST                                                            = 0x0B71
	DEPTH_WRITEMASK                                                       = 0x0B72
	DEVICE_LUID_EXT                                                       = 0x9599
	DEVICE_NODE_MASK_EXT                                                  = 0x959A
	DEVICE_UUID_EXT                                                       = 0x9597
	DIFFERENCE_KHR                                                        = 0x929E
	DIFFERENCE_NV                                                         = 0x929E
	DISJOINT_NV                                                           = 0x9283
	DISPATCH_INDIRECT_BUFFER                                              = 0x90EE
	DISPATCH_INDIRECT_BUFFER_BINDING                                      = 0x90EF
	DITHER                                                                = 0x0BD0
	DMP_PROGRAM_BINARY_DMP                                                = 0x9253
	DONT_CARE                                                             = 0x1100
	DOWNSAMPLE_SCALES_IMG                                                 = 0x913E
	DRAW_BUFFER0                                                          = 0x8825
	DRAW_BUFFER0_EXT                                                      = 0x8825
	DRAW_BUFFER0_NV                                                       = 0x8825
	DRAW_BUFFER1                                                          = 0x8826
	DRAW_BUFFER10                                                         = 0x882F
	DRAW_BUFFER10_EXT                                                     = 0x882F
	DRAW_BUFFER10_NV                                                      = 0x882F
	DRAW_BUFFER11                                                         = 0x8830
	DRAW_BUFFER11_EXT                                                     = 0x8830
	DRAW_BUFFER11_NV                                                      = 0x8830
	DRAW_BUFFER12                                                         = 0x8831
	DRAW_BUFFER12_EXT                                                     = 0x8831
	DRAW_BUFFER12_NV                                                      = 0x8831
	DRAW_BUFFER13                                                         = 0x8832
	DRAW_BUFFER13_EXT                                                     = 0x8832
	DRAW_BUFFER13_NV                                                      = 0x8832
	DRAW_BUFFER14                                                         = 0x8833
	DRAW_BUFFER14_EXT                                                     = 0x8833
	DRAW_BUFFER14_NV                                                      = 0x8833
	DRAW_BUFFER15                                                         = 0x8834
	DRAW_BUFFER15_EXT                                                     = 0x8834
	DRAW_BUFFER15_NV                                                      = 0x8834
	DRAW_BUFFER1_EXT                                                      = 0x8826
	DRAW_BUFFER1_NV                                                       = 0x8826
	DRAW_BUFFER2                                                          = 0x8827
	DRAW_BUFFER2_EXT                                                      = 0x8827
	DRAW_BUFFER2_NV                                                       = 0x8827
	DRAW_BUFFER3                                                          = 0x8828
	DRAW_BUFFER3_EXT                                                      = 0x8828
	DRAW_BUFFER3_NV                                                       = 0x8828
	DRAW_BUFFER4                                                          = 0x8829
	DRAW_BUFFER4_EXT                                                      = 0x8829
	DRAW_BUFFER4_NV                                                       = 0x8829
	DRAW_BUFFER5                                                          = 0x882A
	DRAW_BUFFER5_EXT                                                      = 0x882A
	DRAW_BUFFER5_NV                                                       = 0x882A
	DRAW_BUFFER6                                                          = 0x882B
	DRAW_BUFFER6_EXT                                                      = 0x882B
	DRAW_BUFFER6_NV                                                       = 0x882B
	DRAW_BUFFER7                                                          = 0x882C
	DRAW_BUFFER7_EXT                                                      = 0x882C
	DRAW_BUFFER7_NV                                                       = 0x882C
	DRAW_BUFFER8                                                          = 0x882D
	DRAW_BUFFER8_EXT                                                      = 0x882D
	DRAW_BUFFER8_NV                                                       = 0x882D
	DRAW_BUFFER9                                                          = 0x882E
	DRAW_BUFFER9_EXT                                                      = 0x882E
	DRAW_BUFFER9_NV                                                       = 0x882E
	DRAW_BUFFER_EXT                                                       = 0x0C01
	DRAW_FRAMEBUFFER                                                      = 0x8CA9
	DRAW_FRAMEBUFFER_ANGLE                                                = 0x8CA9
	DRAW_FRAMEBUFFER_APPLE                                                = 0x8CA9
	DRAW_FRAMEBUFFER_BINDING                                              = 0x8CA6
	DRAW_FRAMEBUFFER_BINDING_ANGLE                                        = 0x8CA6
	DRAW_FRAMEBUFFER_BINDING_APPLE                                        = 0x8CA6
	DRAW_FRAMEBUFFER_BINDING_NV                                           = 0x8CA6
	DRAW_FRAMEBUFFER_NV                                                   = 0x8CA9
	DRAW_INDIRECT_BUFFER                                                  = 0x8F3F
	DRAW_INDIRECT_BUFFER_BINDING                                          = 0x8F43
	DRIVER_UUID_EXT                                                       = 0x9598
	DST_ALPHA                                                             = 0x0304
	DST_ATOP_NV                                                           = 0x928F
	DST_COLOR                                                             = 0x0306
	DST_IN_NV                                                             = 0x928B
	DST_NV                                                                = 0x9287
	DST_OUT_NV                                                            = 0x928D
	DST_OVER_NV                                                           = 0x9289
	DUP_FIRST_CUBIC_CURVE_TO_NV                                           = 0xF2
	DUP_LAST_CUBIC_CURVE_TO_NV                                            = 0xF4
	DYNAMIC_COPY                                                          = 0x88EA
	DYNAMIC_DRAW                                                          = 0x88E8
	DYNAMIC_READ                                                          = 0x88E9
	DYNAMIC_STORAGE_BIT_EXT                                               = 0x0100
	EFFECTIVE_RASTER_SAMPLES_EXT                                          = 0x932C
	ELEMENT_ARRAY_BARRIER_BIT                                             = 0x00000002
	ELEMENT_ARRAY_BUFFER                                                  = 0x8893
	ELEMENT_ARRAY_BUFFER_BINDING                                          = 0x8895
	EQUAL                                                                 = 0x0202
	ETC1_RGB8_OES                                                         = 0x8D64
	ETC1_SRGB8_NV                                                         = 0x88EE
	EXCLUSION_KHR                                                         = 0x92A0
	EXCLUSION_NV                                                          = 0x92A0
	EXCLUSIVE_EXT                                                         = 0x8F11
	EXTENSIONS                                                            = 0x1F03
	FACTOR_MAX_AMD                                                        = 0x901D
	FACTOR_MIN_AMD                                                        = 0x901C
	FALSE                                                                 = 0
	FASTEST                                                               = 0x1101
	FENCE_CONDITION_NV                                                    = 0x84F4
	FENCE_STATUS_NV                                                       = 0x84F3
	FETCH_PER_SAMPLE_ARM                                                  = 0x8F65
	FILE_NAME_NV                                                          = 0x9074
	FILL_NV                                                               = 0x1B02
	FILL_RECTANGLE_NV                                                     = 0x933C
	FIRST_TO_REST_NV                                                      = 0x90AF
	FIRST_VERTEX_CONVENTION_EXT                                           = 0x8E4D
	FIRST_VERTEX_CONVENTION_OES                                           = 0x8E4D
	FIXED                                                                 = 0x140C
	FLOAT                                                                 = 0x1406
	FLOAT16_NV                                                            = 0x8FF8
	FLOAT16_VEC2_NV                                                       = 0x8FF9
	FLOAT16_VEC3_NV                                                       = 0x8FFA
	FLOAT16_VEC4_NV                                                       = 0x8FFB
	FLOAT_32_UNSIGNED_INT_24_8_REV                                        = 0x8DAD
	FLOAT_MAT2                                                            = 0x8B5A
	FLOAT_MAT2x3                                                          = 0x8B65
	FLOAT_MAT2x3_NV                                                       = 0x8B65
	FLOAT_MAT2x4                                                          = 0x8B66
	FLOAT_MAT2x4_NV                                                       = 0x8B66
	FLOAT_MAT3                                                            = 0x8B5B
	FLOAT_MAT3x2                                                          = 0x8B67
	FLOAT_MAT3x2_NV                                                       = 0x8B67
	FLOAT_MAT3x4                                                          = 0x8B68
	FLOAT_MAT3x4_NV                                                       = 0x8B68
	FLOAT_MAT4                                                            = 0x8B5C
	FLOAT_MAT4x2                                                          = 0x8B69
	FLOAT_MAT4x2_NV                                                       = 0x8B69
	FLOAT_MAT4x3                                                          = 0x8B6A
	FLOAT_MAT4x3_NV                                                       = 0x8B6A
	FLOAT_VEC2                                                            = 0x8B50
	FLOAT_VEC3                                                            = 0x8B51
	FLOAT_VEC4                                                            = 0x8B52
	FONT_ASCENDER_BIT_NV                                                  = 0x00200000
	FONT_DESCENDER_BIT_NV                                                 = 0x00400000
	FONT_GLYPHS_AVAILABLE_NV                                              = 0x9368
	FONT_HAS_KERNING_BIT_NV                                               = 0x10000000
	FONT_HEIGHT_BIT_NV                                                    = 0x00800000
	FONT_MAX_ADVANCE_HEIGHT_BIT_NV                                        = 0x02000000
	FONT_MAX_ADVANCE_WIDTH_BIT_NV                                         = 0x01000000
	FONT_NUM_GLYPH_INDICES_BIT_NV                                         = 0x20000000
	FONT_TARGET_UNAVAILABLE_NV                                            = 0x9369
	FONT_UNAVAILABLE_NV                                                   = 0x936A
	FONT_UNDERLINE_POSITION_BIT_NV                                        = 0x04000000
	FONT_UNDERLINE_THICKNESS_BIT_NV                                       = 0x08000000
	FONT_UNINTELLIGIBLE_NV                                                = 0x936B
	FONT_UNITS_PER_EM_BIT_NV                                              = 0x00100000
	FONT_X_MAX_BOUNDS_BIT_NV                                              = 0x00040000
	FONT_X_MIN_BOUNDS_BIT_NV                                              = 0x00010000
	FONT_Y_MAX_BOUNDS_BIT_NV                                              = 0x00080000
	FONT_Y_MIN_BOUNDS_BIT_NV                                              = 0x00020000
	FOVEATION_ENABLE_BIT_QCOM                                             = 0x00000001
	FOVEATION_SCALED_BIN_METHOD_BIT_QCOM                                  = 0x00000002
	FRACTIONAL_EVEN_EXT                                                   = 0x8E7C
	FRACTIONAL_EVEN_OES                                                   = 0x8E7C
	FRACTIONAL_ODD_EXT                                                    = 0x8E7B
	FRACTIONAL_ODD_OES                                                    = 0x8E7B
	FRAGMENT_COVERAGE_COLOR_NV                                            = 0x92DE
	FRAGMENT_COVERAGE_TO_COLOR_NV                                         = 0x92DD
	FRAGMENT_INPUT_NV                                                     = 0x936D
	FRAGMENT_INTERPOLATION_OFFSET_BITS_OES                                = 0x8E5D
	FRAGMENT_SHADER                                                       = 0x8B30
	FRAGMENT_SHADER_BIT                                                   = 0x00000002
	FRAGMENT_SHADER_BIT_EXT                                               = 0x00000002
	FRAGMENT_SHADER_DERIVATIVE_HINT                                       = 0x8B8B
	FRAGMENT_SHADER_DERIVATIVE_HINT_OES                                   = 0x8B8B
	FRAGMENT_SHADER_DISCARDS_SAMPLES_EXT                                  = 0x8A52
	FRAGMENT_SHADER_FRAMEBUFFER_FETCH_MRT_ARM                             = 0x8F66
	FRAMEBUFFER                                                           = 0x8D40
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE                                     = 0x8215
	FRAMEBUFFER_ATTACHMENT_ANGLE                                          = 0x93A3
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE                                      = 0x8214
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING                                 = 0x8210
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING_EXT                             = 0x8210
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE                                 = 0x8211
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE_EXT                             = 0x8211
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE                                     = 0x8216
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE                                     = 0x8213
	FRAMEBUFFER_ATTACHMENT_LAYERED_EXT                                    = 0x8DA7
	FRAMEBUFFER_ATTACHMENT_LAYERED_OES                                    = 0x8DA7
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME                                    = 0x8CD1
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE                                    = 0x8CD0
	FRAMEBUFFER_ATTACHMENT_RED_SIZE                                       = 0x8212
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE                                   = 0x8217
	FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_OES                         = 0x8CD4
	FRAMEBUFFER_ATTACHMENT_TEXTURE_BASE_VIEW_INDEX_OVR                    = 0x9632
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE                          = 0x8CD3
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER                                  = 0x8CD4
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL                                  = 0x8CD2
	FRAMEBUFFER_ATTACHMENT_TEXTURE_NUM_VIEWS_OVR                          = 0x9630
	FRAMEBUFFER_ATTACHMENT_TEXTURE_SAMPLES_EXT                            = 0x8D6C
	FRAMEBUFFER_ATTACHMENT_TEXTURE_SCALE_IMG                              = 0x913F
	FRAMEBUFFER_BARRIER_BIT                                               = 0x00000400
	FRAMEBUFFER_BINDING                                                   = 0x8CA6
	FRAMEBUFFER_COMPLETE                                                  = 0x8CD5
	FRAMEBUFFER_DEFAULT                                                   = 0x8218
	FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS                            = 0x9314
	FRAMEBUFFER_DEFAULT_HEIGHT                                            = 0x9311
	FRAMEBUFFER_DEFAULT_LAYERS_EXT                                        = 0x9312
	FRAMEBUFFER_DEFAULT_LAYERS_OES                                        = 0x9312
	FRAMEBUFFER_DEFAULT_SAMPLES                                           = 0x9313
	FRAMEBUFFER_DEFAULT_WIDTH                                             = 0x9310
	FRAMEBUFFER_FETCH_NONCOHERENT_QCOM                                    = 0x96A2
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT                                     = 0x8CD6
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS                                     = 0x8CD9
	FRAMEBUFFER_INCOMPLETE_FOVEATION_QCOM                                 = 0x8BFF
	FRAMEBUFFER_INCOMPLETE_INSUFFICIENT_SHADER_COMBINED_LOCAL_STORAGE_EXT = 0x9652
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_EXT                              = 0x8DA8
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_OES                              = 0x8DA8
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT                             = 0x8CD7
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE                                    = 0x8D56
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_AND_DOWNSAMPLE_IMG                 = 0x913C
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_ANGLE                              = 0x8D56
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_APPLE                              = 0x8D56
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_EXT                                = 0x8D56
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_IMG                                = 0x9134
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_NV                                 = 0x8D56
	FRAMEBUFFER_INCOMPLETE_VIEW_TARGETS_OVR                               = 0x9633
	FRAMEBUFFER_PROGRAMMABLE_SAMPLE_LOCATIONS_NV                          = 0x9342
	FRAMEBUFFER_SAMPLE_LOCATION_PIXEL_GRID_NV                             = 0x9343
	FRAMEBUFFER_SRGB_EXT                                                  = 0x8DB9
	FRAMEBUFFER_UNDEFINED                                                 = 0x8219
	FRAMEBUFFER_UNDEFINED_OES                                             = 0x8219
	FRAMEBUFFER_UNSUPPORTED                                               = 0x8CDD
	FRONT                                                                 = 0x0404
	FRONT_AND_BACK                                                        = 0x0408
	FRONT_FACE                                                            = 0x0B46
	FUNC_ADD                                                              = 0x8006
	FUNC_ADD_EXT                                                          = 0x8006
	FUNC_REVERSE_SUBTRACT                                                 = 0x800B
	FUNC_SUBTRACT                                                         = 0x800A
	GCCSO_SHADER_BINARY_FJ                                                = 0x9260
	GENERATE_MIPMAP_HINT                                                  = 0x8192
	GEOMETRY_LINKED_INPUT_TYPE_EXT                                        = 0x8917
	GEOMETRY_LINKED_INPUT_TYPE_OES                                        = 0x8917
	GEOMETRY_LINKED_OUTPUT_TYPE_EXT                                       = 0x8918
	GEOMETRY_LINKED_OUTPUT_TYPE_OES                                       = 0x8918
	GEOMETRY_LINKED_VERTICES_OUT_EXT                                      = 0x8916
	GEOMETRY_LINKED_VERTICES_OUT_OES                                      = 0x8916
	GEOMETRY_SHADER_BIT_EXT                                               = 0x00000004
	GEOMETRY_SHADER_BIT_OES                                               = 0x00000004
	GEOMETRY_SHADER_EXT                                                   = 0x8DD9
	GEOMETRY_SHADER_INVOCATIONS_EXT                                       = 0x887F
	GEOMETRY_SHADER_INVOCATIONS_OES                                       = 0x887F
	GEOMETRY_SHADER_OES                                                   = 0x8DD9
	GEQUAL                                                                = 0x0206
	GLYPH_HAS_KERNING_BIT_NV                                              = 0x100
	GLYPH_HEIGHT_BIT_NV                                                   = 0x02
	GLYPH_HORIZONTAL_BEARING_ADVANCE_BIT_NV                               = 0x10
	GLYPH_HORIZONTAL_BEARING_X_BIT_NV                                     = 0x04
	GLYPH_HORIZONTAL_BEARING_Y_BIT_NV                                     = 0x08
	GLYPH_VERTICAL_BEARING_ADVANCE_BIT_NV                                 = 0x80
	GLYPH_VERTICAL_BEARING_X_BIT_NV                                       = 0x20
	GLYPH_VERTICAL_BEARING_Y_BIT_NV                                       = 0x40
	GLYPH_WIDTH_BIT_NV                                                    = 0x01
	GPU_DISJOINT_EXT                                                      = 0x8FBB
	GPU_OPTIMIZED_QCOM                                                    = 0x8FB2
	GREATER                                                               = 0x0204
	GREEN                                                                 = 0x1904
	GREEN_BITS                                                            = 0x0D53
	GREEN_NV                                                              = 0x1904
	GUILTY_CONTEXT_RESET                                                  = 0x8253
	GUILTY_CONTEXT_RESET_EXT                                              = 0x8253
	GUILTY_CONTEXT_RESET_KHR                                              = 0x8253
	HALF_FLOAT                                                            = 0x140B
	HALF_FLOAT_OES                                                        = 0x8D61
	HANDLE_TYPE_D3D11_IMAGE_EXT                                           = 0x958B
	HANDLE_TYPE_D3D11_IMAGE_KMT_EXT                                       = 0x958C
	HANDLE_TYPE_D3D12_FENCE_EXT                                           = 0x9594
	HANDLE_TYPE_D3D12_RESOURCE_EXT                                        = 0x958A
	HANDLE_TYPE_D3D12_TILEPOOL_EXT                                        = 0x9589
	HANDLE_TYPE_OPAQUE_FD_EXT                                             = 0x9586
	HANDLE_TYPE_OPAQUE_WIN32_EXT                                          = 0x9587
	HANDLE_TYPE_OPAQUE_WIN32_KMT_EXT                                      = 0x9588
	HARDLIGHT_KHR                                                         = 0x929B
	HARDLIGHT_NV                                                          = 0x929B
	HARDMIX_NV                                                            = 0x92A9
	HIGH_FLOAT                                                            = 0x8DF2
	HIGH_INT                                                              = 0x8DF5
	HORIZONTAL_LINE_TO_NV                                                 = 0x06
	HSL_COLOR_KHR                                                         = 0x92AF
	HSL_COLOR_NV                                                          = 0x92AF
	HSL_HUE_KHR                                                           = 0x92AD
	HSL_HUE_NV                                                            = 0x92AD
	HSL_LUMINOSITY_KHR                                                    = 0x92B0
	HSL_LUMINOSITY_NV                                                     = 0x92B0
	HSL_SATURATION_KHR                                                    = 0x92AE
	HSL_SATURATION_NV                                                     = 0x92AE
	IMAGE_2D                                                              = 0x904D
	IMAGE_2D_ARRAY                                                        = 0x9053
	IMAGE_3D                                                              = 0x904E
	IMAGE_BINDING_ACCESS                                                  = 0x8F3E
	IMAGE_BINDING_FORMAT                                                  = 0x906E
	IMAGE_BINDING_LAYER                                                   = 0x8F3D
	IMAGE_BINDING_LAYERED                                                 = 0x8F3C
	IMAGE_BINDING_LEVEL                                                   = 0x8F3B
	IMAGE_BINDING_NAME                                                    = 0x8F3A
	IMAGE_BUFFER_EXT                                                      = 0x9051
	IMAGE_BUFFER_OES                                                      = 0x9051
	IMAGE_CUBE                                                            = 0x9050
	IMAGE_CUBE_MAP_ARRAY_EXT                                              = 0x9054
	IMAGE_CUBE_MAP_ARRAY_OES                                              = 0x9054
	IMAGE_FORMAT_COMPATIBILITY_BY_CLASS                                   = 0x90C9
	IMAGE_FORMAT_COMPATIBILITY_BY_SIZE                                    = 0x90C8
	IMAGE_FORMAT_COMPATIBILITY_TYPE                                       = 0x90C7
	IMPLEMENTATION_COLOR_READ_FORMAT                                      = 0x8B9B
	IMPLEMENTATION_COLOR_READ_TYPE                                        = 0x8B9A
	INCLUSIVE_EXT                                                         = 0x8F10
	INCR                                                                  = 0x1E02
	INCR_WRAP                                                             = 0x8507
	INFO_LOG_LENGTH                                                       = 0x8B84
	INNOCENT_CONTEXT_RESET                                                = 0x8254
	INNOCENT_CONTEXT_RESET_EXT                                            = 0x8254
	INNOCENT_CONTEXT_RESET_KHR                                            = 0x8254
	INT                                                                   = 0x1404
	INT16_NV                                                              = 0x8FE4
	INT16_VEC2_NV                                                         = 0x8FE5
	INT16_VEC3_NV                                                         = 0x8FE6
	INT16_VEC4_NV                                                         = 0x8FE7
	INT64_NV                                                              = 0x140E
	INT64_VEC2_NV                                                         = 0x8FE9
	INT64_VEC3_NV                                                         = 0x8FEA
	INT64_VEC4_NV                                                         = 0x8FEB
	INT8_NV                                                               = 0x8FE0
	INT8_VEC2_NV                                                          = 0x8FE1
	INT8_VEC3_NV                                                          = 0x8FE2
	INT8_VEC4_NV                                                          = 0x8FE3
	INTERLEAVED_ATTRIBS                                                   = 0x8C8C
	INT_10_10_10_2_OES                                                    = 0x8DF7
	INT_2_10_10_10_REV                                                    = 0x8D9F
	INT_IMAGE_2D                                                          = 0x9058
	INT_IMAGE_2D_ARRAY                                                    = 0x905E
	INT_IMAGE_3D                                                          = 0x9059
	INT_IMAGE_BUFFER_EXT                                                  = 0x905C
	INT_IMAGE_BUFFER_OES                                                  = 0x905C
	INT_IMAGE_CUBE                                                        = 0x905B
	INT_IMAGE_CUBE_MAP_ARRAY_EXT                                          = 0x905F
	INT_IMAGE_CUBE_MAP_ARRAY_OES                                          = 0x905F
	INT_SAMPLER_2D                                                        = 0x8DCA
	INT_SAMPLER_2D_ARRAY                                                  = 0x8DCF
	INT_SAMPLER_2D_MULTISAMPLE                                            = 0x9109
	INT_SAMPLER_2D_MULTISAMPLE_ARRAY_OES                                  = 0x910C
	INT_SAMPLER_3D                                                        = 0x8DCB
	INT_SAMPLER_BUFFER_EXT                                                = 0x8DD0
	INT_SAMPLER_BUFFER_OES                                                = 0x8DD0
	INT_SAMPLER_CUBE                                                      = 0x8DCC
	INT_SAMPLER_CUBE_MAP_ARRAY_EXT                                        = 0x900E
	INT_SAMPLER_CUBE_MAP_ARRAY_OES                                        = 0x900E
	INT_VEC2                                                              = 0x8B53
	INT_VEC3                                                              = 0x8B54
	INT_VEC4                                                              = 0x8B55
	INVALID_ENUM                                                          = 0x0500
	INVALID_FRAMEBUFFER_OPERATION                                         = 0x0506
	INVALID_INDEX                                                         = 0xFFFFFFFF
	INVALID_OPERATION                                                     = 0x0502
	INVALID_VALUE                                                         = 0x0501
	INVERT                                                                = 0x150A
	INVERT_OVG_NV                                                         = 0x92B4
	INVERT_RGB_NV                                                         = 0x92A3
	ISOLINES_EXT                                                          = 0x8E7A
	ISOLINES_OES                                                          = 0x8E7A
	IS_PER_PATCH_EXT                                                      = 0x92E7
	IS_PER_PATCH_OES                                                      = 0x92E7
	IS_ROW_MAJOR                                                          = 0x9300
	ITALIC_BIT_NV                                                         = 0x02
	KEEP                                                                  = 0x1E00
	LARGE_CCW_ARC_TO_NV                                                   = 0x16
	LARGE_CW_ARC_TO_NV                                                    = 0x18
	LAST_VERTEX_CONVENTION_EXT                                            = 0x8E4E
	LAST_VERTEX_CONVENTION_OES                                            = 0x8E4E
	LAYER_PROVOKING_VERTEX_EXT                                            = 0x825E
	LAYER_PROVOKING_VERTEX_OES                                            = 0x825E
	LAYOUT_COLOR_ATTACHMENT_EXT                                           = 0x958E
	LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_EXT                         = 0x9531
	LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_EXT                         = 0x9530
	LAYOUT_DEPTH_STENCIL_ATTACHMENT_EXT                                   = 0x958F
	LAYOUT_DEPTH_STENCIL_READ_ONLY_EXT                                    = 0x9590
	LAYOUT_GENERAL_EXT                                                    = 0x958D
	LAYOUT_SHADER_READ_ONLY_EXT                                           = 0x9591
	LAYOUT_TRANSFER_DST_EXT                                               = 0x9593
	LAYOUT_TRANSFER_SRC_EXT                                               = 0x9592
	LEQUAL                                                                = 0x0203
	LESS                                                                  = 0x0201
	LIGHTEN_KHR                                                           = 0x9298
	LIGHTEN_NV                                                            = 0x9298
	LINEAR                                                                = 0x2601
	LINEARBURN_NV                                                         = 0x92A5
	LINEARDODGE_NV                                                        = 0x92A4
	LINEARLIGHT_NV                                                        = 0x92A7
	LINEAR_MIPMAP_LINEAR                                                  = 0x2703
	LINEAR_MIPMAP_NEAREST                                                 = 0x2701
	LINEAR_TILING_EXT                                                     = 0x9585
	LINES                                                                 = 0x0001
	LINES_ADJACENCY_EXT                                                   = 0x000A
	LINES_ADJACENCY_OES                                                   = 0x000A
	LINE_LOOP                                                             = 0x0002
	LINE_NV                                                               = 0x1B01
	LINE_STRIP                                                            = 0x0003
	LINE_STRIP_ADJACENCY_EXT                                              = 0x000B
	LINE_STRIP_ADJACENCY_OES                                              = 0x000B
	LINE_TO_NV                                                            = 0x04
	LINE_WIDTH                                                            = 0x0B21
	LINK_STATUS                                                           = 0x8B82
	LOCATION                                                              = 0x930E
	LOCATION_INDEX_EXT                                                    = 0x930F
	LOSE_CONTEXT_ON_RESET                                                 = 0x8252
	LOSE_CONTEXT_ON_RESET_EXT                                             = 0x8252
	LOSE_CONTEXT_ON_RESET_KHR                                             = 0x8252
	LOWER_LEFT_EXT                                                        = 0x8CA1
	LOW_FLOAT                                                             = 0x8DF0
	LOW_INT                                                               = 0x8DF3
	LUID_SIZE_EXT                                                         = 8
	LUMINANCE                                                             = 0x1909
	LUMINANCE16F_EXT                                                      = 0x881E
	LUMINANCE32F_EXT                                                      = 0x8818
	LUMINANCE4_ALPHA4_OES                                                 = 0x8043
	LUMINANCE8_ALPHA8_EXT                                                 = 0x8045
	LUMINANCE8_ALPHA8_OES                                                 = 0x8045
	LUMINANCE8_EXT                                                        = 0x8040
	LUMINANCE8_OES                                                        = 0x8040
	LUMINANCE_ALPHA                                                       = 0x190A
	LUMINANCE_ALPHA16F_EXT                                                = 0x881F
	LUMINANCE_ALPHA32F_EXT                                                = 0x8819
	MAJOR_VERSION                                                         = 0x821B
	MALI_PROGRAM_BINARY_ARM                                               = 0x8F61
	MALI_SHADER_BINARY_ARM                                                = 0x8F60
	MAP_COHERENT_BIT_EXT                                                  = 0x0080
	MAP_FLUSH_EXPLICIT_BIT                                                = 0x0010
	MAP_FLUSH_EXPLICIT_BIT_EXT                                            = 0x0010
	MAP_INVALIDATE_BUFFER_BIT                                             = 0x0008
	MAP_INVALIDATE_BUFFER_BIT_EXT                                         = 0x0008
	MAP_INVALIDATE_RANGE_BIT                                              = 0x0004
	MAP_INVALIDATE_RANGE_BIT_EXT                                          = 0x0004
	MAP_PERSISTENT_BIT_EXT                                                = 0x0040
	MAP_READ_BIT                                                          = 0x0001
	MAP_READ_BIT_EXT                                                      = 0x0001
	MAP_UNSYNCHRONIZED_BIT                                                = 0x0020
	MAP_UNSYNCHRONIZED_BIT_EXT                                            = 0x0020
	MAP_WRITE_BIT                                                         = 0x0002
	MAP_WRITE_BIT_EXT                                                     = 0x0002
	MATRIX_STRIDE                                                         = 0x92FF
	MAX                                                                   = 0x8008
	MAX_3D_TEXTURE_SIZE                                                   = 0x8073
	MAX_3D_TEXTURE_SIZE_OES                                               = 0x8073
	MAX_ARRAY_TEXTURE_LAYERS                                              = 0x88FF
	MAX_ATOMIC_COUNTER_BUFFER_BINDINGS                                    = 0x92DC
	MAX_ATOMIC_COUNTER_BUFFER_SIZE                                        = 0x92D8
	MAX_CLIP_DISTANCES_APPLE                                              = 0x0D32
	MAX_CLIP_DISTANCES_EXT                                                = 0x0D32
	MAX_COLOR_ATTACHMENTS                                                 = 0x8CDF
	MAX_COLOR_ATTACHMENTS_EXT                                             = 0x8CDF
	MAX_COLOR_ATTACHMENTS_NV                                              = 0x8CDF
	MAX_COLOR_TEXTURE_SAMPLES                                             = 0x910E
	MAX_COMBINED_ATOMIC_COUNTERS                                          = 0x92D7
	MAX_COMBINED_ATOMIC_COUNTER_BUFFERS                                   = 0x92D1
	MAX_COMBINED_CLIP_AND_CULL_DISTANCES_EXT                              = 0x82FA
	MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS                               = 0x8266
	MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS                              = 0x8A33
	MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS_EXT                          = 0x8A32
	MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS_OES                          = 0x8A32
	MAX_COMBINED_IMAGE_UNIFORMS                                           = 0x90CF
	MAX_COMBINED_SHADER_OUTPUT_RESOURCES                                  = 0x8F39
	MAX_COMBINED_SHADER_STORAGE_BLOCKS                                    = 0x90DC
	MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS_EXT                      = 0x8E1E
	MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS_OES                      = 0x8E1E
	MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS_EXT                   = 0x8E1F
	MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS_OES                   = 0x8E1F
	MAX_COMBINED_TEXTURE_IMAGE_UNITS                                      = 0x8B4D
	MAX_COMBINED_UNIFORM_BLOCKS                                           = 0x8A2E
	MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS                                = 0x8A31
	MAX_COMPUTE_ATOMIC_COUNTERS                                           = 0x8265
	MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS                                    = 0x8264
	MAX_COMPUTE_IMAGE_UNIFORMS                                            = 0x91BD
	MAX_COMPUTE_SHADER_STORAGE_BLOCKS                                     = 0x90DB
	MAX_COMPUTE_SHARED_MEMORY_SIZE                                        = 0x8262
	MAX_COMPUTE_TEXTURE_IMAGE_UNITS                                       = 0x91BC
	MAX_COMPUTE_UNIFORM_BLOCKS                                            = 0x91BB
	MAX_COMPUTE_UNIFORM_COMPONENTS                                        = 0x8263
	MAX_COMPUTE_WORK_GROUP_COUNT                                          = 0x91BE
	MAX_COMPUTE_WORK_GROUP_INVOCATIONS                                    = 0x90EB
	MAX_COMPUTE_WORK_GROUP_SIZE                                           = 0x91BF
	MAX_CUBE_MAP_TEXTURE_SIZE                                             = 0x851C
	MAX_CULL_DISTANCES_EXT                                                = 0x82F9
	MAX_DEBUG_GROUP_STACK_DEPTH                                           = 0x826C
	MAX_DEBUG_GROUP_STACK_DEPTH_KHR                                       = 0x826C
	MAX_DEBUG_LOGGED_MESSAGES                                             = 0x9144
	MAX_DEBUG_LOGGED_MESSAGES_KHR                                         = 0x9144
	MAX_DEBUG_MESSAGE_LENGTH                                              = 0x9143
	MAX_DEBUG_MESSAGE_LENGTH_KHR                                          = 0x9143
	MAX_DEPTH_TEXTURE_SAMPLES                                             = 0x910F
	MAX_DRAW_BUFFERS                                                      = 0x8824
	MAX_DRAW_BUFFERS_EXT                                                  = 0x8824
	MAX_DRAW_BUFFERS_NV                                                   = 0x8824
	MAX_DUAL_SOURCE_DRAW_BUFFERS_EXT                                      = 0x88FC
	MAX_ELEMENTS_INDICES                                                  = 0x80E9
	MAX_ELEMENTS_VERTICES                                                 = 0x80E8
	MAX_ELEMENT_INDEX                                                     = 0x8D6B
	MAX_EXT                                                               = 0x8008
	MAX_FRAGMENT_ATOMIC_COUNTERS                                          = 0x92D6
	MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS                                   = 0x92D0
	MAX_FRAGMENT_IMAGE_UNIFORMS                                           = 0x90CE
	MAX_FRAGMENT_INPUT_COMPONENTS                                         = 0x9125
	MAX_FRAGMENT_INTERPOLATION_OFFSET_OES                                 = 0x8E5C
	MAX_FRAGMENT_SHADER_STORAGE_BLOCKS                                    = 0x90DA
	MAX_FRAGMENT_UNIFORM_BLOCKS                                           = 0x8A2D
	MAX_FRAGMENT_UNIFORM_COMPONENTS                                       = 0x8B49
	MAX_FRAGMENT_UNIFORM_VECTORS                                          = 0x8DFD
	MAX_FRAMEBUFFER_HEIGHT                                                = 0x9316
	MAX_FRAMEBUFFER_LAYERS_EXT                                            = 0x9317
	MAX_FRAMEBUFFER_LAYERS_OES                                            = 0x9317
	MAX_FRAMEBUFFER_SAMPLES                                               = 0x9318
	MAX_FRAMEBUFFER_WIDTH                                                 = 0x9315
	MAX_GEOMETRY_ATOMIC_COUNTERS_EXT                                      = 0x92D5
	MAX_GEOMETRY_ATOMIC_COUNTERS_OES                                      = 0x92D5
	MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS_EXT                               = 0x92CF
	MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS_OES                               = 0x92CF
	MAX_GEOMETRY_IMAGE_UNIFORMS_EXT                                       = 0x90CD
	MAX_GEOMETRY_IMAGE_UNIFORMS_OES                                       = 0x90CD
	MAX_GEOMETRY_INPUT_COMPONENTS_EXT                                     = 0x9123
	MAX_GEOMETRY_INPUT_COMPONENTS_OES                                     = 0x9123
	MAX_GEOMETRY_OUTPUT_COMPONENTS_EXT                                    = 0x9124
	MAX_GEOMETRY_OUTPUT_COMPONENTS_OES                                    = 0x9124
	MAX_GEOMETRY_OUTPUT_VERTICES_EXT                                      = 0x8DE0
	MAX_GEOMETRY_OUTPUT_VERTICES_OES                                      = 0x8DE0
	MAX_GEOMETRY_SHADER_INVOCATIONS_EXT                                   = 0x8E5A
	MAX_GEOMETRY_SHADER_INVOCATIONS_OES                                   = 0x8E5A
	MAX_GEOMETRY_SHADER_STORAGE_BLOCKS_EXT                                = 0x90D7
	MAX_GEOMETRY_SHADER_STORAGE_BLOCKS_OES                                = 0x90D7
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_EXT                                  = 0x8C29
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_OES                                  = 0x8C29
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_EXT                              = 0x8DE1
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_OES                              = 0x8DE1
	MAX_GEOMETRY_UNIFORM_BLOCKS_EXT                                       = 0x8A2C
	MAX_GEOMETRY_UNIFORM_BLOCKS_OES                                       = 0x8A2C
	MAX_GEOMETRY_UNIFORM_COMPONENTS_EXT                                   = 0x8DDF
	MAX_GEOMETRY_UNIFORM_COMPONENTS_OES                                   = 0x8DDF
	MAX_IMAGE_UNITS                                                       = 0x8F38
	MAX_INTEGER_SAMPLES                                                   = 0x9110
	MAX_LABEL_LENGTH                                                      = 0x82E8
	MAX_LABEL_LENGTH_KHR                                                  = 0x82E8
	MAX_MULTIVIEW_BUFFERS_EXT                                             = 0x90F2
	MAX_NAME_LENGTH                                                       = 0x92F6
	MAX_NUM_ACTIVE_VARIABLES                                              = 0x92F7
	MAX_PATCH_VERTICES_EXT                                                = 0x8E7D
	MAX_PATCH_VERTICES_OES                                                = 0x8E7D
	MAX_PROGRAM_TEXEL_OFFSET                                              = 0x8905
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET                                     = 0x8E5F
	MAX_RASTER_SAMPLES_EXT                                                = 0x9329
	MAX_RENDERBUFFER_SIZE                                                 = 0x84E8
	MAX_SAMPLES                                                           = 0x8D57
	MAX_SAMPLES_ANGLE                                                     = 0x8D57
	MAX_SAMPLES_APPLE                                                     = 0x8D57
	MAX_SAMPLES_EXT                                                       = 0x8D57
	MAX_SAMPLES_IMG                                                       = 0x9135
	MAX_SAMPLES_NV                                                        = 0x8D57
	MAX_SAMPLE_MASK_WORDS                                                 = 0x8E59
	MAX_SERVER_WAIT_TIMEOUT                                               = 0x9111
	MAX_SERVER_WAIT_TIMEOUT_APPLE                                         = 0x9111
	MAX_SHADER_COMBINED_LOCAL_STORAGE_FAST_SIZE_EXT                       = 0x9650
	MAX_SHADER_COMBINED_LOCAL_STORAGE_SIZE_EXT                            = 0x9651
	MAX_SHADER_COMPILER_THREADS_KHR                                       = 0x91B0
	MAX_SHADER_PIXEL_LOCAL_STORAGE_FAST_SIZE_EXT                          = 0x8F63
	MAX_SHADER_PIXEL_LOCAL_STORAGE_SIZE_EXT                               = 0x8F67
	MAX_SHADER_STORAGE_BLOCK_SIZE                                         = 0x90DE
	MAX_SHADER_STORAGE_BUFFER_BINDINGS                                    = 0x90DD
	MAX_SPARSE_3D_TEXTURE_SIZE_EXT                                        = 0x9199
	MAX_SPARSE_ARRAY_TEXTURE_LAYERS_EXT                                   = 0x919A
	MAX_SPARSE_TEXTURE_SIZE_EXT                                           = 0x9198
	MAX_SUBPIXEL_PRECISION_BIAS_BITS_NV                                   = 0x9349
	MAX_TESS_CONTROL_ATOMIC_COUNTERS_EXT                                  = 0x92D3
	MAX_TESS_CONTROL_ATOMIC_COUNTERS_OES                                  = 0x92D3
	MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS_EXT                           = 0x92CD
	MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS_OES                           = 0x92CD
	MAX_TESS_CONTROL_IMAGE_UNIFORMS_EXT                                   = 0x90CB
	MAX_TESS_CONTROL_IMAGE_UNIFORMS_OES                                   = 0x90CB
	MAX_TESS_CONTROL_INPUT_COMPONENTS_EXT                                 = 0x886C
	MAX_TESS_CONTROL_INPUT_COMPONENTS_OES                                 = 0x886C
	MAX_TESS_CONTROL_OUTPUT_COMPONENTS_EXT                                = 0x8E83
	MAX_TESS_CONTROL_OUTPUT_COMPONENTS_OES                                = 0x8E83
	MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS_EXT                            = 0x90D8
	MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS_OES                            = 0x90D8
	MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS_EXT                              = 0x8E81
	MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS_OES                              = 0x8E81
	MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS_EXT                          = 0x8E85
	MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS_OES                          = 0x8E85
	MAX_TESS_CONTROL_UNIFORM_BLOCKS_EXT                                   = 0x8E89
	MAX_TESS_CONTROL_UNIFORM_BLOCKS_OES                                   = 0x8E89
	MAX_TESS_CONTROL_UNIFORM_COMPONENTS_EXT                               = 0x8E7F
	MAX_TESS_CONTROL_UNIFORM_COMPONENTS_OES                               = 0x8E7F
	MAX_TESS_EVALUATION_ATOMIC_COUNTERS_EXT                               = 0x92D4
	MAX_TESS_EVALUATION_ATOMIC_COUNTERS_OES                               = 0x92D4
	MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS_EXT                        = 0x92CE
	MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS_OES                        = 0x92CE
	MAX_TESS_EVALUATION_IMAGE_UNIFORMS_EXT                                = 0x90CC
	MAX_TESS_EVALUATION_IMAGE_UNIFORMS_OES                                = 0x90CC
	MAX_TESS_EVALUATION_INPUT_COMPONENTS_EXT                              = 0x886D
	MAX_TESS_EVALUATION_INPUT_COMPONENTS_OES                              = 0x886D
	MAX_TESS_EVALUATION_OUTPUT_COMPONENTS_EXT                             = 0x8E86
	MAX_TESS_EVALUATION_OUTPUT_COMPONENTS_OES                             = 0x8E86
	MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS_EXT                         = 0x90D9
	MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS_OES                         = 0x90D9
	MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS_EXT                           = 0x8E82
	MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS_OES                           = 0x8E82
	MAX_TESS_EVALUATION_UNIFORM_BLOCKS_EXT                                = 0x8E8A
	MAX_TESS_EVALUATION_UNIFORM_BLOCKS_OES                                = 0x8E8A
	MAX_TESS_EVALUATION_UNIFORM_COMPONENTS_EXT                            = 0x8E80
	MAX_TESS_EVALUATION_UNIFORM_COMPONENTS_OES                            = 0x8E80
	MAX_TESS_GEN_LEVEL_EXT                                                = 0x8E7E
	MAX_TESS_GEN_LEVEL_OES                                                = 0x8E7E
	MAX_TESS_PATCH_COMPONENTS_EXT                                         = 0x8E84
	MAX_TESS_PATCH_COMPONENTS_OES                                         = 0x8E84
	MAX_TEXTURE_BUFFER_SIZE_EXT                                           = 0x8C2B
	MAX_TEXTURE_BUFFER_SIZE_OES                                           = 0x8C2B
	MAX_TEXTURE_IMAGE_UNITS                                               = 0x8872
	MAX_TEXTURE_LOD_BIAS                                                  = 0x84FD
	MAX_TEXTURE_MAX_ANISOTROPY_EXT                                        = 0x84FF
	MAX_TEXTURE_SIZE                                                      = 0x0D33
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS                         = 0x8C8A
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS                               = 0x8C8B
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS                            = 0x8C80
	MAX_UNIFORM_BLOCK_SIZE                                                = 0x8A30
	MAX_UNIFORM_BUFFER_BINDINGS                                           = 0x8A2F
	MAX_UNIFORM_LOCATIONS                                                 = 0x826E
	MAX_VARYING_COMPONENTS                                                = 0x8B4B
	MAX_VARYING_VECTORS                                                   = 0x8DFC
	MAX_VERTEX_ATOMIC_COUNTERS                                            = 0x92D2
	MAX_VERTEX_ATOMIC_COUNTER_BUFFERS                                     = 0x92CC
	MAX_VERTEX_ATTRIBS                                                    = 0x8869
	MAX_VERTEX_ATTRIB_BINDINGS                                            = 0x82DA
	MAX_VERTEX_ATTRIB_RELATIVE_OFFSET                                     = 0x82D9
	MAX_VERTEX_ATTRIB_STRIDE                                              = 0x82E5
	MAX_VERTEX_IMAGE_UNIFORMS                                             = 0x90CA
	MAX_VERTEX_OUTPUT_COMPONENTS                                          = 0x9122
	MAX_VERTEX_SHADER_STORAGE_BLOCKS                                      = 0x90D6
	MAX_VERTEX_TEXTURE_IMAGE_UNITS                                        = 0x8B4C
	MAX_VERTEX_UNIFORM_BLOCKS                                             = 0x8A2B
	MAX_VERTEX_UNIFORM_COMPONENTS                                         = 0x8B4A
	MAX_VERTEX_UNIFORM_VECTORS                                            = 0x8DFB
	MAX_VIEWPORTS_NV                                                      = 0x825B
	MAX_VIEWPORTS_OES                                                     = 0x825B
	MAX_VIEWPORT_DIMS                                                     = 0x0D3A
	MAX_VIEWS_OVR                                                         = 0x9631
	MAX_WINDOW_RECTANGLES_EXT                                             = 0x8F14
	MEDIUM_FLOAT                                                          = 0x8DF1
	MEDIUM_INT                                                            = 0x8DF4
	MIN                                                                   = 0x8007
	MINOR_VERSION                                                         = 0x821C
	MINUS_CLAMPED_NV                                                      = 0x92B3
	MINUS_NV                                                              = 0x929F
	MIN_EXT                                                               = 0x8007
	MIN_FRAGMENT_INTERPOLATION_OFFSET_OES                                 = 0x8E5B
	MIN_PROGRAM_TEXEL_OFFSET                                              = 0x8904
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET                                     = 0x8E5E
	MIN_SAMPLE_SHADING_VALUE_OES                                          = 0x8C37
	MIRRORED_REPEAT                                                       = 0x8370
	MIRROR_CLAMP_TO_EDGE_EXT                                              = 0x8743
	MITER_REVERT_NV                                                       = 0x90A7
	MITER_TRUNCATE_NV                                                     = 0x90A8
	MIXED_DEPTH_SAMPLES_SUPPORTED_NV                                      = 0x932F
	MIXED_STENCIL_SAMPLES_SUPPORTED_NV                                    = 0x9330
	MOVE_TO_CONTINUES_NV                                                  = 0x90B6
	MOVE_TO_NV                                                            = 0x02
	MOVE_TO_RESETS_NV                                                     = 0x90B5
	MULTIPLY_KHR                                                          = 0x9294
	MULTIPLY_NV                                                           = 0x9294
	MULTISAMPLES_NV                                                       = 0x9371
	MULTISAMPLE_BUFFER_BIT0_QCOM                                          = 0x01000000
	MULTISAMPLE_BUFFER_BIT1_QCOM                                          = 0x02000000
	MULTISAMPLE_BUFFER_BIT2_QCOM                                          = 0x04000000
	MULTISAMPLE_BUFFER_BIT3_QCOM                                          = 0x08000000
	MULTISAMPLE_BUFFER_BIT4_QCOM                                          = 0x10000000
	MULTISAMPLE_BUFFER_BIT5_QCOM                                          = 0x20000000
	MULTISAMPLE_BUFFER_BIT6_QCOM                                          = 0x40000000
	MULTISAMPLE_BUFFER_BIT7_QCOM                                          = 0x80000000
	MULTISAMPLE_EXT                                                       = 0x809D
	MULTISAMPLE_RASTERIZATION_ALLOWED_EXT                                 = 0x932B
	MULTIVIEW_EXT                                                         = 0x90F1
	NAME_LENGTH                                                           = 0x92F9
	NEAREST                                                               = 0x2600
	NEAREST_MIPMAP_LINEAR                                                 = 0x2702
	NEAREST_MIPMAP_NEAREST                                                = 0x2700
	NEGATIVE_ONE_TO_ONE_EXT                                               = 0x935E
	NEVER                                                                 = 0x0200
	NICEST                                                                = 0x1102
	NONE                                                                  = 0
	NOTEQUAL                                                              = 0x0205
	NO_ERROR                                                              = 0
	NO_RESET_NOTIFICATION                                                 = 0x8261
	NO_RESET_NOTIFICATION_EXT                                             = 0x8261
	NO_RESET_NOTIFICATION_KHR                                             = 0x8261
	NUM_ACTIVE_VARIABLES                                                  = 0x9304
	NUM_COMPRESSED_TEXTURE_FORMATS                                        = 0x86A2
	NUM_DEVICE_UUIDS_EXT                                                  = 0x9596
	NUM_DOWNSAMPLE_SCALES_IMG                                             = 0x913D
	NUM_EXTENSIONS                                                        = 0x821D
	NUM_PROGRAM_BINARY_FORMATS                                            = 0x87FE
	NUM_PROGRAM_BINARY_FORMATS_OES                                        = 0x87FE
	NUM_SAMPLE_COUNTS                                                     = 0x9380
	NUM_SHADER_BINARY_FORMATS                                             = 0x8DF9
	NUM_SPARSE_LEVELS_EXT                                                 = 0x91AA
	NUM_TILING_TYPES_EXT                                                  = 0x9582
	NUM_VIRTUAL_PAGE_SIZES_EXT                                            = 0x91A8
	NUM_WINDOW_RECTANGLES_EXT                                             = 0x8F15
	OBJECT_TYPE                                                           = 0x9112
	OBJECT_TYPE_APPLE                                                     = 0x9112
	OFFSET                                                                = 0x92FC
	ONE                                                                   = 1
	ONE_MINUS_CONSTANT_ALPHA                                              = 0x8004
	ONE_MINUS_CONSTANT_COLOR                                              = 0x8002
	ONE_MINUS_DST_ALPHA                                                   = 0x0305
	ONE_MINUS_DST_COLOR                                                   = 0x0307
	ONE_MINUS_SRC1_ALPHA_EXT                                              = 0x88FB
	ONE_MINUS_SRC1_COLOR_EXT                                              = 0x88FA
	ONE_MINUS_SRC_ALPHA                                                   = 0x0303
	ONE_MINUS_SRC_COLOR                                                   = 0x0301
	OPTIMAL_TILING_EXT                                                    = 0x9584
	OUT_OF_MEMORY                                                         = 0x0505
	OVERLAY_KHR                                                           = 0x9296
	OVERLAY_NV                                                            = 0x9296
	PACK_ALIGNMENT                                                        = 0x0D05
	PACK_REVERSE_ROW_ORDER_ANGLE                                          = 0x93A4
	PACK_ROW_LENGTH                                                       = 0x0D02
	PACK_SKIP_PIXELS                                                      = 0x0D04
	PACK_SKIP_ROWS                                                        = 0x0D03
	PALETTE4_R5_G6_B5_OES                                                 = 0x8B92
	PALETTE4_RGB5_A1_OES                                                  = 0x8B94
	PALETTE4_RGB8_OES                                                     = 0x8B90
	PALETTE4_RGBA4_OES                                                    = 0x8B93
	PALETTE4_RGBA8_OES                                                    = 0x8B91
	PALETTE8_R5_G6_B5_OES                                                 = 0x8B97
	PALETTE8_RGB5_A1_OES                                                  = 0x8B99
	PALETTE8_RGB8_OES                                                     = 0x8B95
	PALETTE8_RGBA4_OES                                                    = 0x8B98
	PALETTE8_RGBA8_OES                                                    = 0x8B96
	PATCHES                                                               = 0x000E
	PATCHES_EXT                                                           = 0x000E
	PATCHES_OES                                                           = 0x000E
	PATCH_VERTICES_EXT                                                    = 0x8E72
	PATCH_VERTICES_OES                                                    = 0x8E72
	PATH_CLIENT_LENGTH_NV                                                 = 0x907F
	PATH_COMMAND_COUNT_NV                                                 = 0x909D
	PATH_COMPUTED_LENGTH_NV                                               = 0x90A0
	PATH_COORD_COUNT_NV                                                   = 0x909E
	PATH_COVER_DEPTH_FUNC_NV                                              = 0x90BF
	PATH_DASH_ARRAY_COUNT_NV                                              = 0x909F
	PATH_DASH_CAPS_NV                                                     = 0x907B
	PATH_DASH_OFFSET_NV                                                   = 0x907E
	PATH_DASH_OFFSET_RESET_NV                                             = 0x90B4
	PATH_END_CAPS_NV                                                      = 0x9076
	PATH_ERROR_POSITION_NV                                                = 0x90AB
	PATH_FILL_BOUNDING_BOX_NV                                             = 0x90A1
	PATH_FILL_COVER_MODE_NV                                               = 0x9082
	PATH_FILL_MASK_NV                                                     = 0x9081
	PATH_FILL_MODE_NV                                                     = 0x9080
	PATH_FORMAT_PS_NV                                                     = 0x9071
	PATH_FORMAT_SVG_NV                                                    = 0x9070
	PATH_GEN_COEFF_NV                                                     = 0x90B1
	PATH_GEN_COMPONENTS_NV                                                = 0x90B3
	PATH_GEN_MODE_NV                                                      = 0x90B0
	PATH_INITIAL_DASH_CAP_NV                                              = 0x907C
	PATH_INITIAL_END_CAP_NV                                               = 0x9077
	PATH_JOIN_STYLE_NV                                                    = 0x9079
	PATH_MAX_MODELVIEW_STACK_DEPTH_NV                                     = 0x0D36
	PATH_MAX_PROJECTION_STACK_DEPTH_NV                                    = 0x0D38
	PATH_MITER_LIMIT_NV                                                   = 0x907A
	PATH_MODELVIEW_MATRIX_NV                                              = 0x0BA6
	PATH_MODELVIEW_NV                                                     = 0x1700
	PATH_MODELVIEW_STACK_DEPTH_NV                                         = 0x0BA3
	PATH_OBJECT_BOUNDING_BOX_NV                                           = 0x908A
	PATH_PROJECTION_MATRIX_NV                                             = 0x0BA7
	PATH_PROJECTION_NV                                                    = 0x1701
	PATH_PROJECTION_STACK_DEPTH_NV                                        = 0x0BA4
	PATH_STENCIL_DEPTH_OFFSET_FACTOR_NV                                   = 0x90BD
	PATH_STENCIL_DEPTH_OFFSET_UNITS_NV                                    = 0x90BE
	PATH_STENCIL_FUNC_NV                                                  = 0x90B7
	PATH_STENCIL_REF_NV                                                   = 0x90B8
	PATH_STENCIL_VALUE_MASK_NV                                            = 0x90B9
	PATH_STROKE_BOUNDING_BOX_NV                                           = 0x90A2
	PATH_STROKE_COVER_MODE_NV                                             = 0x9083
	PATH_STROKE_MASK_NV                                                   = 0x9084
	PATH_STROKE_WIDTH_NV                                                  = 0x9075
	PATH_TERMINAL_DASH_CAP_NV                                             = 0x907D
	PATH_TERMINAL_END_CAP_NV                                              = 0x9078
	PATH_TRANSPOSE_MODELVIEW_MATRIX_NV                                    = 0x84E3
	PATH_TRANSPOSE_PROJECTION_MATRIX_NV                                   = 0x84E4
	PERCENTAGE_AMD                                                        = 0x8BC3
	PERFMON_GLOBAL_MODE_QCOM                                              = 0x8FA0
	PERFMON_RESULT_AMD                                                    = 0x8BC6
	PERFMON_RESULT_AVAILABLE_AMD                                          = 0x8BC4
	PERFMON_RESULT_SIZE_AMD                                               = 0x8BC5
	PERFQUERY_COUNTER_DATA_BOOL32_INTEL                                   = 0x94FC
	PERFQUERY_COUNTER_DATA_DOUBLE_INTEL                                   = 0x94FB
	PERFQUERY_COUNTER_DATA_FLOAT_INTEL                                    = 0x94FA
	PERFQUERY_COUNTER_DATA_UINT32_INTEL                                   = 0x94F8
	PERFQUERY_COUNTER_DATA_UINT64_INTEL                                   = 0x94F9
	PERFQUERY_COUNTER_DESC_LENGTH_MAX_INTEL                               = 0x94FF
	PERFQUERY_COUNTER_DURATION_NORM_INTEL                                 = 0x94F1
	PERFQUERY_COUNTER_DURATION_RAW_INTEL                                  = 0x94F2
	PERFQUERY_COUNTER_EVENT_INTEL                                         = 0x94F0
	PERFQUERY_COUNTER_NAME_LENGTH_MAX_INTEL                               = 0x94FE
	PERFQUERY_COUNTER_RAW_INTEL                                           = 0x94F4
	PERFQUERY_COUNTER_THROUGHPUT_INTEL                                    = 0x94F3
	PERFQUERY_COUNTER_TIMESTAMP_INTEL                                     = 0x94F5
	PERFQUERY_DONOT_FLUSH_INTEL                                           = 0x83F9
	PERFQUERY_FLUSH_INTEL                                                 = 0x83FA
	PERFQUERY_GLOBAL_CONTEXT_INTEL                                        = 0x00000001
	PERFQUERY_GPA_EXTENDED_COUNTERS_INTEL                                 = 0x9500
	PERFQUERY_QUERY_NAME_LENGTH_MAX_INTEL                                 = 0x94FD
	PERFQUERY_SINGLE_CONTEXT_INTEL                                        = 0x00000000
	PERFQUERY_WAIT_INTEL                                                  = 0x83FB
	PINLIGHT_NV                                                           = 0x92A8
	PIXEL_BUFFER_BARRIER_BIT                                              = 0x00000080
	PIXEL_PACK_BUFFER                                                     = 0x88EB
	PIXEL_PACK_BUFFER_BINDING                                             = 0x88ED
	PIXEL_PACK_BUFFER_BINDING_NV                                          = 0x88ED
	PIXEL_PACK_BUFFER_NV                                                  = 0x88EB
	PIXEL_UNPACK_BUFFER                                                   = 0x88EC
	PIXEL_UNPACK_BUFFER_BINDING                                           = 0x88EF
	PIXEL_UNPACK_BUFFER_BINDING_NV                                        = 0x88EF
	PIXEL_UNPACK_BUFFER_NV                                                = 0x88EC
	PLUS_CLAMPED_ALPHA_NV                                                 = 0x92B2
	PLUS_CLAMPED_NV                                                       = 0x92B1
	PLUS_DARKER_NV                                                        = 0x9292
	PLUS_NV                                                               = 0x9291
	POINTS                                                                = 0x0000
	POINT_NV                                                              = 0x1B00
	POLYGON_MODE_NV                                                       = 0x0B40
	POLYGON_OFFSET_CLAMP_EXT                                              = 0x8E1B
	POLYGON_OFFSET_FACTOR                                                 = 0x8038
	POLYGON_OFFSET_FILL                                                   = 0x8037
	POLYGON_OFFSET_LINE_NV                                                = 0x2A02
	POLYGON_OFFSET_POINT_NV                                               = 0x2A01
	POLYGON_OFFSET_UNITS                                                  = 0x2A00
	PRIMITIVES_GENERATED_EXT                                              = 0x8C87
	PRIMITIVES_GENERATED_OES                                              = 0x8C87
	PRIMITIVE_BOUNDING_BOX_EXT                                            = 0x92BE
	PRIMITIVE_BOUNDING_BOX_OES                                            = 0x92BE
	PRIMITIVE_RESTART_FIXED_INDEX                                         = 0x8D69
	PRIMITIVE_RESTART_FOR_PATCHES_SUPPORTED                               = 0x8221
	PRIMITIVE_RESTART_FOR_PATCHES_SUPPORTED_OES                           = 0x8221
	PROGRAM                                                               = 0x82E2
	PROGRAMMABLE_SAMPLE_LOCATION_NV                                       = 0x9341
	PROGRAMMABLE_SAMPLE_LOCATION_TABLE_SIZE_NV                            = 0x9340
	PROGRAM_BINARY_ANGLE                                                  = 0x93A6
	PROGRAM_BINARY_FORMATS                                                = 0x87FF
	PROGRAM_BINARY_FORMATS_OES                                            = 0x87FF
	PROGRAM_BINARY_FORMAT_MESA                                            = 0x875F
	PROGRAM_BINARY_LENGTH                                                 = 0x8741
	PROGRAM_BINARY_LENGTH_OES                                             = 0x8741
	PROGRAM_BINARY_RETRIEVABLE_HINT                                       = 0x8257
	PROGRAM_INPUT                                                         = 0x92E3
	PROGRAM_KHR                                                           = 0x82E2
	PROGRAM_OBJECT_EXT                                                    = 0x8B40
	PROGRAM_OUTPUT                                                        = 0x92E4
	PROGRAM_PIPELINE                                                      = 0x82E4
	PROGRAM_PIPELINE_BINDING                                              = 0x825A
	PROGRAM_PIPELINE_BINDING_EXT                                          = 0x825A
	PROGRAM_PIPELINE_KHR                                                  = 0x82E4
	PROGRAM_PIPELINE_OBJECT_EXT                                           = 0x8A4F
	PROGRAM_SEPARABLE                                                     = 0x8258
	PROGRAM_SEPARABLE_EXT                                                 = 0x8258
	PROTECTED_MEMORY_OBJECT_EXT                                           = 0x959B
	QUADRATIC_CURVE_TO_NV                                                 = 0x0A
	QUADS_EXT                                                             = 0x0007
	QUADS_OES                                                             = 0x0007
	QUERY                                                                 = 0x82E3
	QUERY_BY_REGION_NO_WAIT_NV                                            = 0x8E16
	QUERY_BY_REGION_WAIT_NV                                               = 0x8E15
	QUERY_COUNTER_BITS_EXT                                                = 0x8864
	QUERY_KHR                                                             = 0x82E3
	QUERY_NO_WAIT_NV                                                      = 0x8E14
	QUERY_OBJECT_EXT                                                      = 0x9153
	QUERY_RESULT                                                          = 0x8866
	QUERY_RESULT_AVAILABLE                                                = 0x8867
	QUERY_RESULT_AVAILABLE_EXT                                            = 0x8867
	QUERY_RESULT_EXT                                                      = 0x8866
	QUERY_WAIT_NV                                                         = 0x8E13
	R11F_G11F_B10F                                                        = 0x8C3A
	R11F_G11F_B10F_APPLE                                                  = 0x8C3A
	R16F                                                                  = 0x822D
	R16F_EXT                                                              = 0x822D
	R16I                                                                  = 0x8233
	R16UI                                                                 = 0x8234
	R16_EXT                                                               = 0x822A
	R16_SNORM_EXT                                                         = 0x8F98
	R32F                                                                  = 0x822E
	R32F_EXT                                                              = 0x822E
	R32I                                                                  = 0x8235
	R32UI                                                                 = 0x8236
	R8                                                                    = 0x8229
	R8I                                                                   = 0x8231
	R8UI                                                                  = 0x8232
	R8_EXT                                                                = 0x8229
	R8_SNORM                                                              = 0x8F94
	RASTERIZER_DISCARD                                                    = 0x8C89
	RASTER_FIXED_SAMPLE_LOCATIONS_EXT                                     = 0x932A
	RASTER_MULTISAMPLE_EXT                                                = 0x9327
	RASTER_SAMPLES_EXT                                                    = 0x9328
	READ_BUFFER                                                           = 0x0C02
	READ_BUFFER_EXT                                                       = 0x0C02
	READ_BUFFER_NV                                                        = 0x0C02
	READ_FRAMEBUFFER                                                      = 0x8CA8
	READ_FRAMEBUFFER_ANGLE                                                = 0x8CA8
	READ_FRAMEBUFFER_APPLE                                                = 0x8CA8
	READ_FRAMEBUFFER_BINDING                                              = 0x8CAA
	READ_FRAMEBUFFER_BINDING_ANGLE                                        = 0x8CAA
	READ_FRAMEBUFFER_BINDING_APPLE                                        = 0x8CAA
	READ_FRAMEBUFFER_BINDING_NV                                           = 0x8CAA
	READ_FRAMEBUFFER_NV                                                   = 0x8CA8
	READ_ONLY                                                             = 0x88B8
	READ_WRITE                                                            = 0x88BA
	RECT_NV                                                               = 0xF6
	RED                                                                   = 0x1903
	RED_BITS                                                              = 0x0D52
	RED_EXT                                                               = 0x1903
	RED_INTEGER                                                           = 0x8D94
	RED_NV                                                                = 0x1903
	REFERENCED_BY_COMPUTE_SHADER                                          = 0x930B
	REFERENCED_BY_FRAGMENT_SHADER                                         = 0x930A
	REFERENCED_BY_GEOMETRY_SHADER_EXT                                     = 0x9309
	REFERENCED_BY_GEOMETRY_SHADER_OES                                     = 0x9309
	REFERENCED_BY_TESS_CONTROL_SHADER_EXT                                 = 0x9307
	REFERENCED_BY_TESS_CONTROL_SHADER_OES                                 = 0x9307
	REFERENCED_BY_TESS_EVALUATION_SHADER_EXT                              = 0x9308
	REFERENCED_BY_TESS_EVALUATION_SHADER_OES                              = 0x9308
	REFERENCED_BY_VERTEX_SHADER                                           = 0x9306
	RELATIVE_ARC_TO_NV                                                    = 0xFF
	RELATIVE_CONIC_CURVE_TO_NV                                            = 0x1B
	RELATIVE_CUBIC_CURVE_TO_NV                                            = 0x0D
	RELATIVE_HORIZONTAL_LINE_TO_NV                                        = 0x07
	RELATIVE_LARGE_CCW_ARC_TO_NV                                          = 0x17
	RELATIVE_LARGE_CW_ARC_TO_NV                                           = 0x19
	RELATIVE_LINE_TO_NV                                                   = 0x05
	RELATIVE_MOVE_TO_NV                                                   = 0x03
	RELATIVE_QUADRATIC_CURVE_TO_NV                                        = 0x0B
	RELATIVE_RECT_NV                                                      = 0xF7
	RELATIVE_ROUNDED_RECT2_NV                                             = 0xEB
	RELATIVE_ROUNDED_RECT4_NV                                             = 0xED
	RELATIVE_ROUNDED_RECT8_NV                                             = 0xEF
	RELATIVE_ROUNDED_RECT_NV                                              = 0xE9
	RELATIVE_SMALL_CCW_ARC_TO_NV                                          = 0x13
	RELATIVE_SMALL_CW_ARC_TO_NV                                           = 0x15
	RELATIVE_SMOOTH_CUBIC_CURVE_TO_NV                                     = 0x11
	RELATIVE_SMOOTH_QUADRATIC_CURVE_TO_NV                                 = 0x0F
	RELATIVE_VERTICAL_LINE_TO_NV                                          = 0x09
	RENDERBUFFER                                                          = 0x8D41
	RENDERBUFFER_ALPHA_SIZE                                               = 0x8D53
	RENDERBUFFER_BINDING                                                  = 0x8CA7
	RENDERBUFFER_BLUE_SIZE                                                = 0x8D52
	RENDERBUFFER_DEPTH_SIZE                                               = 0x8D54
	RENDERBUFFER_GREEN_SIZE                                               = 0x8D51
	RENDERBUFFER_HEIGHT                                                   = 0x8D43
	RENDERBUFFER_INTERNAL_FORMAT                                          = 0x8D44
	RENDERBUFFER_RED_SIZE                                                 = 0x8D50
	RENDERBUFFER_SAMPLES                                                  = 0x8CAB
	RENDERBUFFER_SAMPLES_ANGLE                                            = 0x8CAB
	RENDERBUFFER_SAMPLES_APPLE                                            = 0x8CAB
	RENDERBUFFER_SAMPLES_EXT                                              = 0x8CAB
	RENDERBUFFER_SAMPLES_IMG                                              = 0x9133
	RENDERBUFFER_SAMPLES_NV                                               = 0x8CAB
	RENDERBUFFER_STENCIL_SIZE                                             = 0x8D55
	RENDERBUFFER_WIDTH                                                    = 0x8D42
	RENDERER                                                              = 0x1F01
	RENDER_DIRECT_TO_FRAMEBUFFER_QCOM                                     = 0x8FB3
	REPEAT                                                                = 0x2901
	REPLACE                                                               = 0x1E01
	REQUIRED_TEXTURE_IMAGE_UNITS_OES                                      = 0x8D68
	RESET_NOTIFICATION_STRATEGY                                           = 0x8256
	RESET_NOTIFICATION_STRATEGY_EXT                                       = 0x8256
	RESET_NOTIFICATION_STRATEGY_KHR                                       = 0x8256
	RESTART_PATH_NV                                                       = 0xF0
	RG                                                                    = 0x8227
	RG16F                                                                 = 0x822F
	RG16F_EXT                                                             = 0x822F
	RG16I                                                                 = 0x8239
	RG16UI                                                                = 0x823A
	RG16_EXT                                                              = 0x822C
	RG16_SNORM_EXT                                                        = 0x8F99
	RG32F                                                                 = 0x8230
	RG32F_EXT                                                             = 0x8230
	RG32I                                                                 = 0x823B
	RG32UI                                                                = 0x823C
	RG8                                                                   = 0x822B
	RG8I                                                                  = 0x8237
	RG8UI                                                                 = 0x8238
	RG8_EXT                                                               = 0x822B
	RG8_SNORM                                                             = 0x8F95
	RGB                                                                   = 0x1907
	RGB10_A2                                                              = 0x8059
	RGB10_A2UI                                                            = 0x906F
	RGB10_A2_EXT                                                          = 0x8059
	RGB10_EXT                                                             = 0x8052
	RGB16F                                                                = 0x881B
	RGB16F_EXT                                                            = 0x881B
	RGB16I                                                                = 0x8D89
	RGB16UI                                                               = 0x8D77
	RGB16_EXT                                                             = 0x8054
	RGB16_SNORM_EXT                                                       = 0x8F9A
	RGB32F                                                                = 0x8815
	RGB32F_EXT                                                            = 0x8815
	RGB32I                                                                = 0x8D83
	RGB32UI                                                               = 0x8D71
	RGB565                                                                = 0x8D62
	RGB565_OES                                                            = 0x8D62
	RGB5_A1                                                               = 0x8057
	RGB5_A1_OES                                                           = 0x8057
	RGB8                                                                  = 0x8051
	RGB8I                                                                 = 0x8D8F
	RGB8UI                                                                = 0x8D7D
	RGB8_OES                                                              = 0x8051
	RGB8_SNORM                                                            = 0x8F96
	RGB9_E5                                                               = 0x8C3D
	RGB9_E5_APPLE                                                         = 0x8C3D
	RGBA                                                                  = 0x1908
	RGBA16F                                                               = 0x881A
	RGBA16F_EXT                                                           = 0x881A
	RGBA16I                                                               = 0x8D88
	RGBA16UI                                                              = 0x8D76
	RGBA16_EXT                                                            = 0x805B
	RGBA16_SNORM_EXT                                                      = 0x8F9B
	RGBA32F                                                               = 0x8814
	RGBA32F_EXT                                                           = 0x8814
	RGBA32I                                                               = 0x8D82
	RGBA32UI                                                              = 0x8D70
	RGBA4                                                                 = 0x8056
	RGBA4_OES                                                             = 0x8056
	RGBA8                                                                 = 0x8058
	RGBA8I                                                                = 0x8D8E
	RGBA8UI                                                               = 0x8D7C
	RGBA8_OES                                                             = 0x8058
	RGBA8_SNORM                                                           = 0x8F97
	RGBA_INTEGER                                                          = 0x8D99
	RGB_422_APPLE                                                         = 0x8A1F
	RGB_INTEGER                                                           = 0x8D98
	RGB_RAW_422_APPLE                                                     = 0x8A51
	RG_EXT                                                                = 0x8227
	RG_INTEGER                                                            = 0x8228
	ROUNDED_RECT2_NV                                                      = 0xEA
	ROUNDED_RECT4_NV                                                      = 0xEC
	ROUNDED_RECT8_NV                                                      = 0xEE
	ROUNDED_RECT_NV                                                       = 0xE8
	ROUND_NV                                                              = 0x90A4
	SAMPLER                                                               = 0x82E6
	SAMPLER_2D                                                            = 0x8B5E
	SAMPLER_2D_ARRAY                                                      = 0x8DC1
	SAMPLER_2D_ARRAY_SHADOW                                               = 0x8DC4
	SAMPLER_2D_ARRAY_SHADOW_NV                                            = 0x8DC4
	SAMPLER_2D_MULTISAMPLE                                                = 0x9108
	SAMPLER_2D_MULTISAMPLE_ARRAY_OES                                      = 0x910B
	SAMPLER_2D_SHADOW                                                     = 0x8B62
	SAMPLER_2D_SHADOW_EXT                                                 = 0x8B62
	SAMPLER_3D                                                            = 0x8B5F
	SAMPLER_3D_OES                                                        = 0x8B5F
	SAMPLER_BINDING                                                       = 0x8919
	SAMPLER_BUFFER_EXT                                                    = 0x8DC2
	SAMPLER_BUFFER_OES                                                    = 0x8DC2
	SAMPLER_CUBE                                                          = 0x8B60
	SAMPLER_CUBE_MAP_ARRAY_EXT                                            = 0x900C
	SAMPLER_CUBE_MAP_ARRAY_OES                                            = 0x900C
	SAMPLER_CUBE_MAP_ARRAY_SHADOW_EXT                                     = 0x900D
	SAMPLER_CUBE_MAP_ARRAY_SHADOW_OES                                     = 0x900D
	SAMPLER_CUBE_SHADOW                                                   = 0x8DC5
	SAMPLER_CUBE_SHADOW_NV                                                = 0x8DC5
	SAMPLER_EXTERNAL_2D_Y2Y_EXT                                           = 0x8BE7
	SAMPLER_EXTERNAL_OES                                                  = 0x8D66
	SAMPLER_KHR                                                           = 0x82E6
	SAMPLES                                                               = 0x80A9
	SAMPLE_ALPHA_TO_COVERAGE                                              = 0x809E
	SAMPLE_ALPHA_TO_ONE_EXT                                               = 0x809F
	SAMPLE_BUFFERS                                                        = 0x80A8
	SAMPLE_COVERAGE                                                       = 0x80A0
	SAMPLE_COVERAGE_INVERT                                                = 0x80AB
	SAMPLE_COVERAGE_VALUE                                                 = 0x80AA
	SAMPLE_LOCATION_NV                                                    = 0x8E50
	SAMPLE_LOCATION_PIXEL_GRID_HEIGHT_NV                                  = 0x933F
	SAMPLE_LOCATION_PIXEL_GRID_WIDTH_NV                                   = 0x933E
	SAMPLE_LOCATION_SUBPIXEL_BITS_NV                                      = 0x933D
	SAMPLE_MASK                                                           = 0x8E51
	SAMPLE_MASK_VALUE                                                     = 0x8E52
	SAMPLE_POSITION                                                       = 0x8E50
	SAMPLE_SHADING_OES                                                    = 0x8C36
	SCISSOR_BOX                                                           = 0x0C10
	SCISSOR_TEST                                                          = 0x0C11
	SCREEN_KHR                                                            = 0x9295
	SCREEN_NV                                                             = 0x9295
	SEPARATE_ATTRIBS                                                      = 0x8C8D
	SGX_BINARY_IMG                                                        = 0x8C0A
	SGX_PROGRAM_BINARY_IMG                                                = 0x9130
	SHADER                                                                = 0x82E1
	SHADER_BINARY_DMP                                                     = 0x9250
	SHADER_BINARY_FORMATS                                                 = 0x8DF8
	SHADER_BINARY_VIV                                                     = 0x8FC4
	SHADER_COMPILER                                                       = 0x8DFA
	SHADER_IMAGE_ACCESS_BARRIER_BIT                                       = 0x00000020
	SHADER_KHR                                                            = 0x82E1
	SHADER_OBJECT_EXT                                                     = 0x8B48
	SHADER_PIXEL_LOCAL_STORAGE_EXT                                        = 0x8F64
	SHADER_SOURCE_LENGTH                                                  = 0x8B88
	SHADER_STORAGE_BARRIER_BIT                                            = 0x00002000
	SHADER_STORAGE_BLOCK                                                  = 0x92E6
	SHADER_STORAGE_BUFFER                                                 = 0x90D2
	SHADER_STORAGE_BUFFER_BINDING                                         = 0x90D3
	SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT                                = 0x90DF
	SHADER_STORAGE_BUFFER_SIZE                                            = 0x90D5
	SHADER_STORAGE_BUFFER_START                                           = 0x90D4
	SHADER_TYPE                                                           = 0x8B4F
	SHADING_LANGUAGE_VERSION                                              = 0x8B8C
	SHARED_EDGE_NV                                                        = 0xC0
	SHORT                                                                 = 0x1402
	SIGNALED                                                              = 0x9119
	SIGNALED_APPLE                                                        = 0x9119
	SIGNED_NORMALIZED                                                     = 0x8F9C
	SKIP_DECODE_EXT                                                       = 0x8A4A
	SKIP_MISSING_GLYPH_NV                                                 = 0x90A9
	SLUMINANCE8_ALPHA8_NV                                                 = 0x8C45
	SLUMINANCE8_NV                                                        = 0x8C47
	SLUMINANCE_ALPHA_NV                                                   = 0x8C44
	SLUMINANCE_NV                                                         = 0x8C46
	SMALL_CCW_ARC_TO_NV                                                   = 0x12
	SMALL_CW_ARC_TO_NV                                                    = 0x14
	SMAPHS30_PROGRAM_BINARY_DMP                                           = 0x9251
	SMAPHS_PROGRAM_BINARY_DMP                                             = 0x9252
	SMOOTH_CUBIC_CURVE_TO_NV                                              = 0x10
	SMOOTH_QUADRATIC_CURVE_TO_NV                                          = 0x0E
	SOFTLIGHT_KHR                                                         = 0x929C
	SOFTLIGHT_NV                                                          = 0x929C
	SPARSE_TEXTURE_FULL_ARRAY_CUBE_MIPMAPS_EXT                            = 0x91A9
	SQUARE_NV                                                             = 0x90A3
	SR8_EXT                                                               = 0x8FBD
	SRC1_ALPHA_EXT                                                        = 0x8589
	SRC1_COLOR_EXT                                                        = 0x88F9
	SRC_ALPHA                                                             = 0x0302
	SRC_ALPHA_SATURATE                                                    = 0x0308
	SRC_ALPHA_SATURATE_EXT                                                = 0x0308
	SRC_ATOP_NV                                                           = 0x928E
	SRC_COLOR                                                             = 0x0300
	SRC_IN_NV                                                             = 0x928A
	SRC_NV                                                                = 0x9286
	SRC_OUT_NV                                                            = 0x928C
	SRC_OVER_NV                                                           = 0x9288
	SRG8_EXT                                                              = 0x8FBE
	SRGB                                                                  = 0x8C40
	SRGB8                                                                 = 0x8C41
	SRGB8_ALPHA8                                                          = 0x8C43
	SRGB8_ALPHA8_EXT                                                      = 0x8C43
	SRGB8_NV                                                              = 0x8C41
	SRGB_ALPHA_EXT                                                        = 0x8C42
	SRGB_EXT                                                              = 0x8C40
	STACK_OVERFLOW                                                        = 0x0503
	STACK_OVERFLOW_KHR                                                    = 0x0503
	STACK_UNDERFLOW                                                       = 0x0504
	STACK_UNDERFLOW_KHR                                                   = 0x0504
	STANDARD_FONT_FORMAT_NV                                               = 0x936C
	STANDARD_FONT_NAME_NV                                                 = 0x9072
	STATE_RESTORE                                                         = 0x8BDC
	STATIC_COPY                                                           = 0x88E6
	STATIC_DRAW                                                           = 0x88E4
	STATIC_READ                                                           = 0x88E5
	STENCIL                                                               = 0x1802
	STENCIL_ATTACHMENT                                                    = 0x8D20
	STENCIL_BACK_FAIL                                                     = 0x8801
	STENCIL_BACK_FUNC                                                     = 0x8800
	STENCIL_BACK_PASS_DEPTH_FAIL                                          = 0x8802
	STENCIL_BACK_PASS_DEPTH_PASS                                          = 0x8803
	STENCIL_BACK_REF                                                      = 0x8CA3
	STENCIL_BACK_VALUE_MASK                                               = 0x8CA4
	STENCIL_BACK_WRITEMASK                                                = 0x8CA5
	STENCIL_BITS                                                          = 0x0D57
	STENCIL_BUFFER_BIT                                                    = 0x00000400
	STENCIL_BUFFER_BIT0_QCOM                                              = 0x00010000
	STENCIL_BUFFER_BIT1_QCOM                                              = 0x00020000
	STENCIL_BUFFER_BIT2_QCOM                                              = 0x00040000
	STENCIL_BUFFER_BIT3_QCOM                                              = 0x00080000
	STENCIL_BUFFER_BIT4_QCOM                                              = 0x00100000
	STENCIL_BUFFER_BIT5_QCOM                                              = 0x00200000
	STENCIL_BUFFER_BIT6_QCOM                                              = 0x00400000
	STENCIL_BUFFER_BIT7_QCOM                                              = 0x00800000
	STENCIL_CLEAR_VALUE                                                   = 0x0B91
	STENCIL_EXT                                                           = 0x1802
	STENCIL_FAIL                                                          = 0x0B94
	STENCIL_FUNC                                                          = 0x0B92
	STENCIL_INDEX                                                         = 0x1901
	STENCIL_INDEX1_OES                                                    = 0x8D46
	STENCIL_INDEX4_OES                                                    = 0x8D47
	STENCIL_INDEX8                                                        = 0x8D48
	STENCIL_INDEX8_OES                                                    = 0x8D48
	STENCIL_INDEX_OES                                                     = 0x1901
	STENCIL_PASS_DEPTH_FAIL                                               = 0x0B95
	STENCIL_PASS_DEPTH_PASS                                               = 0x0B96
	STENCIL_REF                                                           = 0x0B97
	STENCIL_SAMPLES_NV                                                    = 0x932E
	STENCIL_TEST                                                          = 0x0B90
	STENCIL_VALUE_MASK                                                    = 0x0B93
	STENCIL_WRITEMASK                                                     = 0x0B98
	STREAM_COPY                                                           = 0x88E2
	STREAM_DRAW                                                           = 0x88E0
	STREAM_READ                                                           = 0x88E1
	SUBPIXEL_BITS                                                         = 0x0D50
	SUBPIXEL_PRECISION_BIAS_X_BITS_NV                                     = 0x9347
	SUBPIXEL_PRECISION_BIAS_Y_BITS_NV                                     = 0x9348
	SUPERSAMPLE_SCALE_X_NV                                                = 0x9372
	SUPERSAMPLE_SCALE_Y_NV                                                = 0x9373
	SYNC_CONDITION                                                        = 0x9113
	SYNC_CONDITION_APPLE                                                  = 0x9113
	SYNC_FENCE                                                            = 0x9116
	SYNC_FENCE_APPLE                                                      = 0x9116
	SYNC_FLAGS                                                            = 0x9115
	SYNC_FLAGS_APPLE                                                      = 0x9115
	SYNC_FLUSH_COMMANDS_BIT                                               = 0x00000001
	SYNC_FLUSH_COMMANDS_BIT_APPLE                                         = 0x00000001
	SYNC_GPU_COMMANDS_COMPLETE                                            = 0x9117
	SYNC_GPU_COMMANDS_COMPLETE_APPLE                                      = 0x9117
	SYNC_OBJECT_APPLE                                                     = 0x8A53
	SYNC_STATUS                                                           = 0x9114
	SYNC_STATUS_APPLE                                                     = 0x9114
	SYSTEM_FONT_NAME_NV                                                   = 0x9073
	TESS_CONTROL_OUTPUT_VERTICES_EXT                                      = 0x8E75
	TESS_CONTROL_OUTPUT_VERTICES_OES                                      = 0x8E75
	TESS_CONTROL_SHADER_BIT_EXT                                           = 0x00000008
	TESS_CONTROL_SHADER_BIT_OES                                           = 0x00000008
	TESS_CONTROL_SHADER_EXT                                               = 0x8E88
	TESS_CONTROL_SHADER_OES                                               = 0x8E88
	TESS_EVALUATION_SHADER_BIT_EXT                                        = 0x00000010
	TESS_EVALUATION_SHADER_BIT_OES                                        = 0x00000010
	TESS_EVALUATION_SHADER_EXT                                            = 0x8E87
	TESS_EVALUATION_SHADER_OES                                            = 0x8E87
	TESS_GEN_MODE_EXT                                                     = 0x8E76
	TESS_GEN_MODE_OES                                                     = 0x8E76
	TESS_GEN_POINT_MODE_EXT                                               = 0x8E79
	TESS_GEN_POINT_MODE_OES                                               = 0x8E79
	TESS_GEN_SPACING_EXT                                                  = 0x8E77
	TESS_GEN_SPACING_OES                                                  = 0x8E77
	TESS_GEN_VERTEX_ORDER_EXT                                             = 0x8E78
	TESS_GEN_VERTEX_ORDER_OES                                             = 0x8E78
	TEXTURE                                                               = 0x1702
	TEXTURE0                                                              = 0x84C0
	TEXTURE1                                                              = 0x84C1
	TEXTURE10                                                             = 0x84CA
	TEXTURE11                                                             = 0x84CB
	TEXTURE12                                                             = 0x84CC
	TEXTURE13                                                             = 0x84CD
	TEXTURE14                                                             = 0x84CE
	TEXTURE15                                                             = 0x84CF
	TEXTURE16                                                             = 0x84D0
	TEXTURE17                                                             = 0x84D1
	TEXTURE18                                                             = 0x84D2
	TEXTURE19                                                             = 0x84D3
	TEXTURE2                                                              = 0x84C2
	TEXTURE20                                                             = 0x84D4
	TEXTURE21                                                             = 0x84D5
	TEXTURE22                                                             = 0x84D6
	TEXTURE23                                                             = 0x84D7
	TEXTURE24                                                             = 0x84D8
	TEXTURE25                                                             = 0x84D9
	TEXTURE26                                                             = 0x84DA
	TEXTURE27                                                             = 0x84DB
	TEXTURE28                                                             = 0x84DC
	TEXTURE29                                                             = 0x84DD
	TEXTURE3                                                              = 0x84C3
	TEXTURE30                                                             = 0x84DE
	TEXTURE31                                                             = 0x84DF
	TEXTURE4                                                              = 0x84C4
	TEXTURE5                                                              = 0x84C5
	TEXTURE6                                                              = 0x84C6
	TEXTURE7                                                              = 0x84C7
	TEXTURE8                                                              = 0x84C8
	TEXTURE9                                                              = 0x84C9
	TEXTURE_2D                                                            = 0x0DE1
	TEXTURE_2D_ARRAY                                                      = 0x8C1A
	TEXTURE_2D_MULTISAMPLE                                                = 0x9100
	TEXTURE_2D_MULTISAMPLE_ARRAY                                          = 0x9102
	TEXTURE_2D_MULTISAMPLE_ARRAY_OES                                      = 0x9102
	TEXTURE_3D                                                            = 0x806F
	TEXTURE_3D_OES                                                        = 0x806F
	TEXTURE_ALPHA_SIZE                                                    = 0x805F
	TEXTURE_ALPHA_TYPE                                                    = 0x8C13
	TEXTURE_ASTC_DECODE_PRECISION_EXT                                     = 0x8F69
	TEXTURE_BASE_LEVEL                                                    = 0x813C
	TEXTURE_BINDING_2D                                                    = 0x8069
	TEXTURE_BINDING_2D_ARRAY                                              = 0x8C1D
	TEXTURE_BINDING_2D_MULTISAMPLE                                        = 0x9104
	TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY_OES                              = 0x9105
	TEXTURE_BINDING_3D                                                    = 0x806A
	TEXTURE_BINDING_3D_OES                                                = 0x806A
	TEXTURE_BINDING_BUFFER_EXT                                            = 0x8C2C
	TEXTURE_BINDING_BUFFER_OES                                            = 0x8C2C
	TEXTURE_BINDING_CUBE_MAP                                              = 0x8514
	TEXTURE_BINDING_CUBE_MAP_ARRAY_EXT                                    = 0x900A
	TEXTURE_BINDING_CUBE_MAP_ARRAY_OES                                    = 0x900A
	TEXTURE_BINDING_EXTERNAL_OES                                          = 0x8D67
	TEXTURE_BLUE_SIZE                                                     = 0x805E
	TEXTURE_BLUE_TYPE                                                     = 0x8C12
	TEXTURE_BORDER_COLOR_EXT                                              = 0x1004
	TEXTURE_BORDER_COLOR_NV                                               = 0x1004
	TEXTURE_BORDER_COLOR_OES                                              = 0x1004
	TEXTURE_BUFFER_BINDING_EXT                                            = 0x8C2A
	TEXTURE_BUFFER_BINDING_OES                                            = 0x8C2A
	TEXTURE_BUFFER_DATA_STORE_BINDING_EXT                                 = 0x8C2D
	TEXTURE_BUFFER_DATA_STORE_BINDING_OES                                 = 0x8C2D
	TEXTURE_BUFFER_EXT                                                    = 0x8C2A
	TEXTURE_BUFFER_OES                                                    = 0x8C2A
	TEXTURE_BUFFER_OFFSET_ALIGNMENT_EXT                                   = 0x919F
	TEXTURE_BUFFER_OFFSET_ALIGNMENT_OES                                   = 0x919F
	TEXTURE_BUFFER_OFFSET_EXT                                             = 0x919D
	TEXTURE_BUFFER_OFFSET_OES                                             = 0x919D
	TEXTURE_BUFFER_SIZE_EXT                                               = 0x919E
	TEXTURE_BUFFER_SIZE_OES                                               = 0x919E
	TEXTURE_COMPARE_FUNC                                                  = 0x884D
	TEXTURE_COMPARE_FUNC_EXT                                              = 0x884D
	TEXTURE_COMPARE_MODE                                                  = 0x884C
	TEXTURE_COMPARE_MODE_EXT                                              = 0x884C
	TEXTURE_COMPRESSED                                                    = 0x86A1
	TEXTURE_CUBE_MAP                                                      = 0x8513
	TEXTURE_CUBE_MAP_ARRAY_EXT                                            = 0x9009
	TEXTURE_CUBE_MAP_ARRAY_OES                                            = 0x9009
	TEXTURE_CUBE_MAP_NEGATIVE_X                                           = 0x8516
	TEXTURE_CUBE_MAP_NEGATIVE_Y                                           = 0x8518
	TEXTURE_CUBE_MAP_NEGATIVE_Z                                           = 0x851A
	TEXTURE_CUBE_MAP_POSITIVE_X                                           = 0x8515
	TEXTURE_CUBE_MAP_POSITIVE_Y                                           = 0x8517
	TEXTURE_CUBE_MAP_POSITIVE_Z                                           = 0x8519
	TEXTURE_DEPTH                                                         = 0x8071
	TEXTURE_DEPTH_QCOM                                                    = 0x8BD4
	TEXTURE_DEPTH_SIZE                                                    = 0x884A
	TEXTURE_DEPTH_TYPE                                                    = 0x8C16
	TEXTURE_EXTERNAL_OES                                                  = 0x8D65
	TEXTURE_FETCH_BARRIER_BIT                                             = 0x00000008
	TEXTURE_FIXED_SAMPLE_LOCATIONS                                        = 0x9107
	TEXTURE_FORMAT_QCOM                                                   = 0x8BD6
	TEXTURE_FORMAT_SRGB_OVERRIDE_EXT                                      = 0x8FBF
	TEXTURE_FOVEATED_FEATURE_BITS_QCOM                                    = 0x8BFB
	TEXTURE_FOVEATED_FEATURE_QUERY_QCOM                                   = 0x8BFD
	TEXTURE_FOVEATED_MIN_PIXEL_DENSITY_QCOM                               = 0x8BFC
	TEXTURE_FOVEATED_NUM_FOCAL_POINTS_QUERY_QCOM                          = 0x8BFE
	TEXTURE_GREEN_SIZE                                                    = 0x805D
	TEXTURE_GREEN_TYPE                                                    = 0x8C11
	TEXTURE_HEIGHT                                                        = 0x1001
	TEXTURE_HEIGHT_QCOM                                                   = 0x8BD3
	TEXTURE_IMAGE_VALID_QCOM                                              = 0x8BD8
	TEXTURE_IMMUTABLE_FORMAT                                              = 0x912F
	TEXTURE_IMMUTABLE_FORMAT_EXT                                          = 0x912F
	TEXTURE_IMMUTABLE_LEVELS                                              = 0x82DF
	TEXTURE_INTERNAL_FORMAT                                               = 0x1003
	TEXTURE_INTERNAL_FORMAT_QCOM                                          = 0x8BD5
	TEXTURE_MAG_FILTER                                                    = 0x2800
	TEXTURE_MAX_ANISOTROPY_EXT                                            = 0x84FE
	TEXTURE_MAX_LEVEL                                                     = 0x813D
	TEXTURE_MAX_LEVEL_APPLE                                               = 0x813D
	TEXTURE_MAX_LOD                                                       = 0x813B
	TEXTURE_MIN_FILTER                                                    = 0x2801
	TEXTURE_MIN_LOD                                                       = 0x813A
	TEXTURE_NUM_LEVELS_QCOM                                               = 0x8BD9
	TEXTURE_OBJECT_VALID_QCOM                                             = 0x8BDB
	TEXTURE_PROTECTED_EXT                                                 = 0x8BFA
	TEXTURE_REDUCTION_MODE_EXT                                            = 0x9366
	TEXTURE_RED_SIZE                                                      = 0x805C
	TEXTURE_RED_TYPE                                                      = 0x8C10
	TEXTURE_SAMPLES                                                       = 0x9106
	TEXTURE_SAMPLES_IMG                                                   = 0x9136
	TEXTURE_SHARED_SIZE                                                   = 0x8C3F
	TEXTURE_SPARSE_EXT                                                    = 0x91A6
	TEXTURE_SRGB_DECODE_EXT                                               = 0x8A48
	TEXTURE_STENCIL_SIZE                                                  = 0x88F1
	TEXTURE_SWIZZLE_A                                                     = 0x8E45
	TEXTURE_SWIZZLE_B                                                     = 0x8E44
	TEXTURE_SWIZZLE_G                                                     = 0x8E43
	TEXTURE_SWIZZLE_R                                                     = 0x8E42
	TEXTURE_TARGET_QCOM                                                   = 0x8BDA
	TEXTURE_TILING_EXT                                                    = 0x9580
	TEXTURE_TYPE_QCOM                                                     = 0x8BD7
	TEXTURE_UPDATE_BARRIER_BIT                                            = 0x00000100
	TEXTURE_USAGE_ANGLE                                                   = 0x93A2
	TEXTURE_VIEW_MIN_LAYER_EXT                                            = 0x82DD
	TEXTURE_VIEW_MIN_LAYER_OES                                            = 0x82DD
	TEXTURE_VIEW_MIN_LEVEL_EXT                                            = 0x82DB
	TEXTURE_VIEW_MIN_LEVEL_OES                                            = 0x82DB
	TEXTURE_VIEW_NUM_LAYERS_EXT                                           = 0x82DE
	TEXTURE_VIEW_NUM_LAYERS_OES                                           = 0x82DE
	TEXTURE_VIEW_NUM_LEVELS_EXT                                           = 0x82DC
	TEXTURE_VIEW_NUM_LEVELS_OES                                           = 0x82DC
	TEXTURE_WIDTH                                                         = 0x1000
	TEXTURE_WIDTH_QCOM                                                    = 0x8BD2
	TEXTURE_WRAP_R                                                        = 0x8072
	TEXTURE_WRAP_R_OES                                                    = 0x8072
	TEXTURE_WRAP_S                                                        = 0x2802
	TEXTURE_WRAP_T                                                        = 0x2803
	TILING_TYPES_EXT                                                      = 0x9583
	TIMEOUT_EXPIRED                                                       = 0x911B
	TIMEOUT_EXPIRED_APPLE                                                 = 0x911B
	TIMEOUT_IGNORED                                                       = 0xFFFFFFFFFFFFFFFF
	TIMEOUT_IGNORED_APPLE                                                 = 0xFFFFFFFFFFFFFFFF
	TIMESTAMP_EXT                                                         = 0x8E28
	TIME_ELAPSED_EXT                                                      = 0x88BF
	TOP_LEVEL_ARRAY_SIZE                                                  = 0x930C
	TOP_LEVEL_ARRAY_STRIDE                                                = 0x930D
	TRANSFORM_FEEDBACK                                                    = 0x8E22
	TRANSFORM_FEEDBACK_ACTIVE                                             = 0x8E24
	TRANSFORM_FEEDBACK_BARRIER_BIT                                        = 0x00000800
	TRANSFORM_FEEDBACK_BINDING                                            = 0x8E25
	TRANSFORM_FEEDBACK_BUFFER                                             = 0x8C8E
	TRANSFORM_FEEDBACK_BUFFER_BINDING                                     = 0x8C8F
	TRANSFORM_FEEDBACK_BUFFER_MODE                                        = 0x8C7F
	TRANSFORM_FEEDBACK_BUFFER_SIZE                                        = 0x8C85
	TRANSFORM_FEEDBACK_BUFFER_START                                       = 0x8C84
	TRANSFORM_FEEDBACK_PAUSED                                             = 0x8E23
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN                                 = 0x8C88
	TRANSFORM_FEEDBACK_VARYING                                            = 0x92F4
	TRANSFORM_FEEDBACK_VARYINGS                                           = 0x8C83
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH                                 = 0x8C76
	TRANSLATED_SHADER_SOURCE_LENGTH_ANGLE                                 = 0x93A0
	TRANSLATE_2D_NV                                                       = 0x9090
	TRANSLATE_3D_NV                                                       = 0x9091
	TRANSLATE_X_NV                                                        = 0x908E
	TRANSLATE_Y_NV                                                        = 0x908F
	TRANSPOSE_AFFINE_2D_NV                                                = 0x9096
	TRANSPOSE_AFFINE_3D_NV                                                = 0x9098
	TRIANGLES                                                             = 0x0004
	TRIANGLES_ADJACENCY_EXT                                               = 0x000C
	TRIANGLES_ADJACENCY_OES                                               = 0x000C
	TRIANGLE_FAN                                                          = 0x0006
	TRIANGLE_STRIP                                                        = 0x0005
	TRIANGLE_STRIP_ADJACENCY_EXT                                          = 0x000D
	TRIANGLE_STRIP_ADJACENCY_OES                                          = 0x000D
	TRIANGULAR_NV                                                         = 0x90A5
	TRUE                                                                  = 1
	TYPE                                                                  = 0x92FA
	UNCORRELATED_NV                                                       = 0x9282
	UNDEFINED_VERTEX_EXT                                                  = 0x8260
	UNDEFINED_VERTEX_OES                                                  = 0x8260
	UNIFORM                                                               = 0x92E1
	UNIFORM_ARRAY_STRIDE                                                  = 0x8A3C
	UNIFORM_BARRIER_BIT                                                   = 0x00000004
	UNIFORM_BLOCK                                                         = 0x92E2
	UNIFORM_BLOCK_ACTIVE_UNIFORMS                                         = 0x8A42
	UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES                                  = 0x8A43
	UNIFORM_BLOCK_BINDING                                                 = 0x8A3F
	UNIFORM_BLOCK_DATA_SIZE                                               = 0x8A40
	UNIFORM_BLOCK_INDEX                                                   = 0x8A3A
	UNIFORM_BLOCK_NAME_LENGTH                                             = 0x8A41
	UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER                           = 0x8A46
	UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER                             = 0x8A44
	UNIFORM_BUFFER                                                        = 0x8A11
	UNIFORM_BUFFER_BINDING                                                = 0x8A28
	UNIFORM_BUFFER_OFFSET_ALIGNMENT                                       = 0x8A34
	UNIFORM_BUFFER_SIZE                                                   = 0x8A2A
	UNIFORM_BUFFER_START                                                  = 0x8A29
	UNIFORM_IS_ROW_MAJOR                                                  = 0x8A3E
	UNIFORM_MATRIX_STRIDE                                                 = 0x8A3D
	UNIFORM_NAME_LENGTH                                                   = 0x8A39
	UNIFORM_OFFSET                                                        = 0x8A3B
	UNIFORM_SIZE                                                          = 0x8A38
	UNIFORM_TYPE                                                          = 0x8A37
	UNKNOWN_CONTEXT_RESET                                                 = 0x8255
	UNKNOWN_CONTEXT_RESET_EXT                                             = 0x8255
	UNKNOWN_CONTEXT_RESET_KHR                                             = 0x8255
	UNPACK_ALIGNMENT                                                      = 0x0CF5
	UNPACK_IMAGE_HEIGHT                                                   = 0x806E
	UNPACK_ROW_LENGTH                                                     = 0x0CF2
	UNPACK_ROW_LENGTH_EXT                                                 = 0x0CF2
	UNPACK_SKIP_IMAGES                                                    = 0x806D
	UNPACK_SKIP_PIXELS                                                    = 0x0CF4
	UNPACK_SKIP_PIXELS_EXT                                                = 0x0CF4
	UNPACK_SKIP_ROWS                                                      = 0x0CF3
	UNPACK_SKIP_ROWS_EXT                                                  = 0x0CF3
	UNSIGNALED                                                            = 0x9118
	UNSIGNALED_APPLE                                                      = 0x9118
	UNSIGNED_BYTE                                                         = 0x1401
	UNSIGNED_INT                                                          = 0x1405
	UNSIGNED_INT16_NV                                                     = 0x8FF0
	UNSIGNED_INT16_VEC2_NV                                                = 0x8FF1
	UNSIGNED_INT16_VEC3_NV                                                = 0x8FF2
	UNSIGNED_INT16_VEC4_NV                                                = 0x8FF3
	UNSIGNED_INT64_AMD                                                    = 0x8BC2
	UNSIGNED_INT64_NV                                                     = 0x140F
	UNSIGNED_INT64_VEC2_NV                                                = 0x8FF5
	UNSIGNED_INT64_VEC3_NV                                                = 0x8FF6
	UNSIGNED_INT64_VEC4_NV                                                = 0x8FF7
	UNSIGNED_INT8_NV                                                      = 0x8FEC
	UNSIGNED_INT8_VEC2_NV                                                 = 0x8FED
	UNSIGNED_INT8_VEC3_NV                                                 = 0x8FEE
	UNSIGNED_INT8_VEC4_NV                                                 = 0x8FEF
	UNSIGNED_INT_10F_11F_11F_REV                                          = 0x8C3B
	UNSIGNED_INT_10F_11F_11F_REV_APPLE                                    = 0x8C3B
	UNSIGNED_INT_10_10_10_2_OES                                           = 0x8DF6
	UNSIGNED_INT_24_8                                                     = 0x84FA
	UNSIGNED_INT_24_8_OES                                                 = 0x84FA
	UNSIGNED_INT_2_10_10_10_REV                                           = 0x8368
	UNSIGNED_INT_2_10_10_10_REV_EXT                                       = 0x8368
	UNSIGNED_INT_5_9_9_9_REV                                              = 0x8C3E
	UNSIGNED_INT_5_9_9_9_REV_APPLE                                        = 0x8C3E
	UNSIGNED_INT_ATOMIC_COUNTER                                           = 0x92DB
	UNSIGNED_INT_IMAGE_2D                                                 = 0x9063
	UNSIGNED_INT_IMAGE_2D_ARRAY                                           = 0x9069
	UNSIGNED_INT_IMAGE_3D                                                 = 0x9064
	UNSIGNED_INT_IMAGE_BUFFER_EXT                                         = 0x9067
	UNSIGNED_INT_IMAGE_BUFFER_OES                                         = 0x9067
	UNSIGNED_INT_IMAGE_CUBE                                               = 0x9066
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY_EXT                                 = 0x906A
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY_OES                                 = 0x906A
	UNSIGNED_INT_SAMPLER_2D                                               = 0x8DD2
	UNSIGNED_INT_SAMPLER_2D_ARRAY                                         = 0x8DD7
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE                                   = 0x910A
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY_OES                         = 0x910D
	UNSIGNED_INT_SAMPLER_3D                                               = 0x8DD3
	UNSIGNED_INT_SAMPLER_BUFFER_EXT                                       = 0x8DD8
	UNSIGNED_INT_SAMPLER_BUFFER_OES                                       = 0x8DD8
	UNSIGNED_INT_SAMPLER_CUBE                                             = 0x8DD4
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY_EXT                               = 0x900F
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY_OES                               = 0x900F
	UNSIGNED_INT_VEC2                                                     = 0x8DC6
	UNSIGNED_INT_VEC3                                                     = 0x8DC7
	UNSIGNED_INT_VEC4                                                     = 0x8DC8
	UNSIGNED_NORMALIZED                                                   = 0x8C17
	UNSIGNED_NORMALIZED_EXT                                               = 0x8C17
	UNSIGNED_SHORT                                                        = 0x1403
	UNSIGNED_SHORT_1_5_5_5_REV_EXT                                        = 0x8366
	UNSIGNED_SHORT_4_4_4_4                                                = 0x8033
	UNSIGNED_SHORT_4_4_4_4_REV_EXT                                        = 0x8365
	UNSIGNED_SHORT_4_4_4_4_REV_IMG                                        = 0x8365
	UNSIGNED_SHORT_5_5_5_1                                                = 0x8034
	UNSIGNED_SHORT_5_6_5                                                  = 0x8363
	UNSIGNED_SHORT_8_8_APPLE                                              = 0x85BA
	UNSIGNED_SHORT_8_8_REV_APPLE                                          = 0x85BB
	UPPER_LEFT_EXT                                                        = 0x8CA2
	USE_MISSING_GLYPH_NV                                                  = 0x90AA
	UTF16_NV                                                              = 0x909B
	UTF8_NV                                                               = 0x909A
	UUID_SIZE_EXT                                                         = 16
	VALIDATE_STATUS                                                       = 0x8B83
	VENDOR                                                                = 0x1F00
	VERSION                                                               = 0x1F02
	VERTEX_ARRAY                                                          = 0x8074
	VERTEX_ARRAY_BINDING                                                  = 0x85B5
	VERTEX_ARRAY_BINDING_OES                                              = 0x85B5
	VERTEX_ARRAY_KHR                                                      = 0x8074
	VERTEX_ARRAY_OBJECT_EXT                                               = 0x9154
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT                                       = 0x00000001
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING                                    = 0x889F
	VERTEX_ATTRIB_ARRAY_DIVISOR                                           = 0x88FE
	VERTEX_ATTRIB_ARRAY_DIVISOR_ANGLE                                     = 0x88FE
	VERTEX_ATTRIB_ARRAY_DIVISOR_EXT                                       = 0x88FE
	VERTEX_ATTRIB_ARRAY_DIVISOR_NV                                        = 0x88FE
	VERTEX_ATTRIB_ARRAY_ENABLED                                           = 0x8622
	VERTEX_ATTRIB_ARRAY_INTEGER                                           = 0x88FD
	VERTEX_ATTRIB_ARRAY_NORMALIZED                                        = 0x886A
	VERTEX_ATTRIB_ARRAY_POINTER                                           = 0x8645
	VERTEX_ATTRIB_ARRAY_SIZE                                              = 0x8623
	VERTEX_ATTRIB_ARRAY_STRIDE                                            = 0x8624
	VERTEX_ATTRIB_ARRAY_TYPE                                              = 0x8625
	VERTEX_ATTRIB_BINDING                                                 = 0x82D4
	VERTEX_ATTRIB_RELATIVE_OFFSET                                         = 0x82D5
	VERTEX_BINDING_BUFFER                                                 = 0x8F4F
	VERTEX_BINDING_DIVISOR                                                = 0x82D6
	VERTEX_BINDING_OFFSET                                                 = 0x82D7
	VERTEX_BINDING_STRIDE                                                 = 0x82D8
	VERTEX_SHADER                                                         = 0x8B31
	VERTEX_SHADER_BIT                                                     = 0x00000001
	VERTEX_SHADER_BIT_EXT                                                 = 0x00000001
	VERTICAL_LINE_TO_NV                                                   = 0x08
	VIEWPORT                                                              = 0x0BA2
	VIEWPORT_BOUNDS_RANGE_NV                                              = 0x825D
	VIEWPORT_BOUNDS_RANGE_OES                                             = 0x825D
	VIEWPORT_INDEX_PROVOKING_VERTEX_NV                                    = 0x825F
	VIEWPORT_INDEX_PROVOKING_VERTEX_OES                                   = 0x825F
	VIEWPORT_POSITION_W_SCALE_NV                                          = 0x937C
	VIEWPORT_POSITION_W_SCALE_X_COEFF_NV                                  = 0x937D
	VIEWPORT_POSITION_W_SCALE_Y_COEFF_NV                                  = 0x937E
	VIEWPORT_SUBPIXEL_BITS_NV                                             = 0x825C
	VIEWPORT_SUBPIXEL_BITS_OES                                            = 0x825C
	VIEWPORT_SWIZZLE_NEGATIVE_W_NV                                        = 0x9357
	VIEWPORT_SWIZZLE_NEGATIVE_X_NV                                        = 0x9351
	VIEWPORT_SWIZZLE_NEGATIVE_Y_NV                                        = 0x9353
	VIEWPORT_SWIZZLE_NEGATIVE_Z_NV                                        = 0x9355
	VIEWPORT_SWIZZLE_POSITIVE_W_NV                                        = 0x9356
	VIEWPORT_SWIZZLE_POSITIVE_X_NV                                        = 0x9350
	VIEWPORT_SWIZZLE_POSITIVE_Y_NV                                        = 0x9352
	VIEWPORT_SWIZZLE_POSITIVE_Z_NV                                        = 0x9354
	VIEWPORT_SWIZZLE_W_NV                                                 = 0x935B
	VIEWPORT_SWIZZLE_X_NV                                                 = 0x9358
	VIEWPORT_SWIZZLE_Y_NV                                                 = 0x9359
	VIEWPORT_SWIZZLE_Z_NV                                                 = 0x935A
	VIRTUAL_PAGE_SIZE_INDEX_EXT                                           = 0x91A7
	VIRTUAL_PAGE_SIZE_X_EXT                                               = 0x9195
	VIRTUAL_PAGE_SIZE_Y_EXT                                               = 0x9196
	VIRTUAL_PAGE_SIZE_Z_EXT                                               = 0x9197
	VIVIDLIGHT_NV                                                         = 0x92A6
	WAIT_FAILED                                                           = 0x911D
	WAIT_FAILED_APPLE                                                     = 0x911D
	WEIGHTED_AVERAGE_EXT                                                  = 0x9367
	WINDOW_RECTANGLE_EXT                                                  = 0x8F12
	WINDOW_RECTANGLE_MODE_EXT                                             = 0x8F13
	WRITEONLY_RENDERING_QCOM                                              = 0x8823
	WRITE_ONLY                                                            = 0x88B9
	WRITE_ONLY_OES                                                        = 0x88B9
	XOR_NV                                                                = 0x1506
	Z400_BINARY_AMD                                                       = 0x8740
	ZERO                                                                  = 0
	ZERO_TO_ONE_EXT                                                       = 0x935F
)

var (
	gpAcquireKeyedMutexWin32EXT                      C.GPACQUIREKEYEDMUTEXWIN32EXT
	gpActiveProgramEXT                               C.GPACTIVEPROGRAMEXT
	gpActiveShaderProgram                            C.GPACTIVESHADERPROGRAM
	gpActiveShaderProgramEXT                         C.GPACTIVESHADERPROGRAMEXT
	gpActiveTexture                                  C.GPACTIVETEXTURE
	gpAlphaFuncQCOM                                  C.GPALPHAFUNCQCOM
	gpApplyFramebufferAttachmentCMAAINTEL            C.GPAPPLYFRAMEBUFFERATTACHMENTCMAAINTEL
	gpAttachShader                                   C.GPATTACHSHADER
	gpBeginConditionalRenderNV                       C.GPBEGINCONDITIONALRENDERNV
	gpBeginPerfMonitorAMD                            C.GPBEGINPERFMONITORAMD
	gpBeginPerfQueryINTEL                            C.GPBEGINPERFQUERYINTEL
	gpBeginQuery                                     C.GPBEGINQUERY
	gpBeginQueryEXT                                  C.GPBEGINQUERYEXT
	gpBeginTransformFeedback                         C.GPBEGINTRANSFORMFEEDBACK
	gpBindAttribLocation                             C.GPBINDATTRIBLOCATION
	gpBindBuffer                                     C.GPBINDBUFFER
	gpBindBufferBase                                 C.GPBINDBUFFERBASE
	gpBindBufferRange                                C.GPBINDBUFFERRANGE
	gpBindFragDataLocationEXT                        C.GPBINDFRAGDATALOCATIONEXT
	gpBindFragDataLocationIndexedEXT                 C.GPBINDFRAGDATALOCATIONINDEXEDEXT
	gpBindFramebuffer                                C.GPBINDFRAMEBUFFER
	gpBindImageTexture                               C.GPBINDIMAGETEXTURE
	gpBindProgramPipeline                            C.GPBINDPROGRAMPIPELINE
	gpBindProgramPipelineEXT                         C.GPBINDPROGRAMPIPELINEEXT
	gpBindRenderbuffer                               C.GPBINDRENDERBUFFER
	gpBindSampler                                    C.GPBINDSAMPLER
	gpBindTexture                                    C.GPBINDTEXTURE
	gpBindTransformFeedback                          C.GPBINDTRANSFORMFEEDBACK
	gpBindVertexArray                                C.GPBINDVERTEXARRAY
	gpBindVertexArrayOES                             C.GPBINDVERTEXARRAYOES
	gpBindVertexBuffer                               C.GPBINDVERTEXBUFFER
	gpBlendBarrierKHR                                C.GPBLENDBARRIERKHR
	gpBlendBarrierNV                                 C.GPBLENDBARRIERNV
	gpBlendColor                                     C.GPBLENDCOLOR
	gpBlendEquation                                  C.GPBLENDEQUATION
	gpBlendEquationEXT                               C.GPBLENDEQUATIONEXT
	gpBlendEquationSeparate                          C.GPBLENDEQUATIONSEPARATE
	gpBlendEquationSeparateiEXT                      C.GPBLENDEQUATIONSEPARATEIEXT
	gpBlendEquationSeparateiOES                      C.GPBLENDEQUATIONSEPARATEIOES
	gpBlendEquationiEXT                              C.GPBLENDEQUATIONIEXT
	gpBlendEquationiOES                              C.GPBLENDEQUATIONIOES
	gpBlendFunc                                      C.GPBLENDFUNC
	gpBlendFuncSeparate                              C.GPBLENDFUNCSEPARATE
	gpBlendFuncSeparateiEXT                          C.GPBLENDFUNCSEPARATEIEXT
	gpBlendFuncSeparateiOES                          C.GPBLENDFUNCSEPARATEIOES
	gpBlendFunciEXT                                  C.GPBLENDFUNCIEXT
	gpBlendFunciOES                                  C.GPBLENDFUNCIOES
	gpBlendParameteriNV                              C.GPBLENDPARAMETERINV
	gpBlitFramebuffer                                C.GPBLITFRAMEBUFFER
	gpBlitFramebufferANGLE                           C.GPBLITFRAMEBUFFERANGLE
	gpBlitFramebufferNV                              C.GPBLITFRAMEBUFFERNV
	gpBufferData                                     C.GPBUFFERDATA
	gpBufferStorageEXT                               C.GPBUFFERSTORAGEEXT
	gpBufferStorageExternalEXT                       C.GPBUFFERSTORAGEEXTERNALEXT
	gpBufferStorageMemEXT                            C.GPBUFFERSTORAGEMEMEXT
	gpBufferSubData                                  C.GPBUFFERSUBDATA
	gpCheckFramebufferStatus                         C.GPCHECKFRAMEBUFFERSTATUS
	gpClear                                          C.GPCLEAR
	gpClearBufferfi                                  C.GPCLEARBUFFERFI
	gpClearBufferfv                                  C.GPCLEARBUFFERFV
	gpClearBufferiv                                  C.GPCLEARBUFFERIV
	gpClearBufferuiv                                 C.GPCLEARBUFFERUIV
	gpClearColor                                     C.GPCLEARCOLOR
	gpClearDepthf                                    C.GPCLEARDEPTHF
	gpClearPixelLocalStorageuiEXT                    C.GPCLEARPIXELLOCALSTORAGEUIEXT
	gpClearStencil                                   C.GPCLEARSTENCIL
	gpClearTexImageEXT                               C.GPCLEARTEXIMAGEEXT
	gpClearTexSubImageEXT                            C.GPCLEARTEXSUBIMAGEEXT
	gpClientWaitSync                                 C.GPCLIENTWAITSYNC
	gpClientWaitSyncAPPLE                            C.GPCLIENTWAITSYNCAPPLE
	gpClipControlEXT                                 C.GPCLIPCONTROLEXT
	gpColorMask                                      C.GPCOLORMASK
	gpColorMaskiEXT                                  C.GPCOLORMASKIEXT
	gpColorMaskiOES                                  C.GPCOLORMASKIOES
	gpCompileShader                                  C.GPCOMPILESHADER
	gpCompressedTexImage2D                           C.GPCOMPRESSEDTEXIMAGE2D
	gpCompressedTexImage3D                           C.GPCOMPRESSEDTEXIMAGE3D
	gpCompressedTexImage3DOES                        C.GPCOMPRESSEDTEXIMAGE3DOES
	gpCompressedTexSubImage2D                        C.GPCOMPRESSEDTEXSUBIMAGE2D
	gpCompressedTexSubImage3D                        C.GPCOMPRESSEDTEXSUBIMAGE3D
	gpCompressedTexSubImage3DOES                     C.GPCOMPRESSEDTEXSUBIMAGE3DOES
	gpConservativeRasterParameteriNV                 C.GPCONSERVATIVERASTERPARAMETERINV
	gpCopyBufferSubData                              C.GPCOPYBUFFERSUBDATA
	gpCopyBufferSubDataNV                            C.GPCOPYBUFFERSUBDATANV
	gpCopyImageSubDataEXT                            C.GPCOPYIMAGESUBDATAEXT
	gpCopyImageSubDataOES                            C.GPCOPYIMAGESUBDATAOES
	gpCopyPathNV                                     C.GPCOPYPATHNV
	gpCopyTexImage2D                                 C.GPCOPYTEXIMAGE2D
	gpCopyTexSubImage2D                              C.GPCOPYTEXSUBIMAGE2D
	gpCopyTexSubImage3D                              C.GPCOPYTEXSUBIMAGE3D
	gpCopyTexSubImage3DOES                           C.GPCOPYTEXSUBIMAGE3DOES
	gpCopyTextureLevelsAPPLE                         C.GPCOPYTEXTURELEVELSAPPLE
	gpCoverFillPathInstancedNV                       C.GPCOVERFILLPATHINSTANCEDNV
	gpCoverFillPathNV                                C.GPCOVERFILLPATHNV
	gpCoverStrokePathInstancedNV                     C.GPCOVERSTROKEPATHINSTANCEDNV
	gpCoverStrokePathNV                              C.GPCOVERSTROKEPATHNV
	gpCoverageMaskNV                                 C.GPCOVERAGEMASKNV
	gpCoverageModulationNV                           C.GPCOVERAGEMODULATIONNV
	gpCoverageModulationTableNV                      C.GPCOVERAGEMODULATIONTABLENV
	gpCoverageOperationNV                            C.GPCOVERAGEOPERATIONNV
	gpCreateMemoryObjectsEXT                         C.GPCREATEMEMORYOBJECTSEXT
	gpCreatePerfQueryINTEL                           C.GPCREATEPERFQUERYINTEL
	gpCreateProgram                                  C.GPCREATEPROGRAM
	gpCreateShader                                   C.GPCREATESHADER
	gpCreateShaderProgramEXT                         C.GPCREATESHADERPROGRAMEXT
	gpCreateShaderProgramv                           C.GPCREATESHADERPROGRAMV
	gpCreateShaderProgramvEXT                        C.GPCREATESHADERPROGRAMVEXT
	gpCullFace                                       C.GPCULLFACE
	gpDebugMessageCallback                           C.GPDEBUGMESSAGECALLBACK
	gpDebugMessageCallbackKHR                        C.GPDEBUGMESSAGECALLBACKKHR
	gpDebugMessageControl                            C.GPDEBUGMESSAGECONTROL
	gpDebugMessageControlKHR                         C.GPDEBUGMESSAGECONTROLKHR
	gpDebugMessageInsert                             C.GPDEBUGMESSAGEINSERT
	gpDebugMessageInsertKHR                          C.GPDEBUGMESSAGEINSERTKHR
	gpDeleteBuffers                                  C.GPDELETEBUFFERS
	gpDeleteFencesNV                                 C.GPDELETEFENCESNV
	gpDeleteFramebuffers                             C.GPDELETEFRAMEBUFFERS
	gpDeleteMemoryObjectsEXT                         C.GPDELETEMEMORYOBJECTSEXT
	gpDeletePathsNV                                  C.GPDELETEPATHSNV
	gpDeletePerfMonitorsAMD                          C.GPDELETEPERFMONITORSAMD
	gpDeletePerfQueryINTEL                           C.GPDELETEPERFQUERYINTEL
	gpDeleteProgram                                  C.GPDELETEPROGRAM
	gpDeleteProgramPipelines                         C.GPDELETEPROGRAMPIPELINES
	gpDeleteProgramPipelinesEXT                      C.GPDELETEPROGRAMPIPELINESEXT
	gpDeleteQueries                                  C.GPDELETEQUERIES
	gpDeleteQueriesEXT                               C.GPDELETEQUERIESEXT
	gpDeleteRenderbuffers                            C.GPDELETERENDERBUFFERS
	gpDeleteSamplers                                 C.GPDELETESAMPLERS
	gpDeleteSemaphoresEXT                            C.GPDELETESEMAPHORESEXT
	gpDeleteShader                                   C.GPDELETESHADER
	gpDeleteSync                                     C.GPDELETESYNC
	gpDeleteSyncAPPLE                                C.GPDELETESYNCAPPLE
	gpDeleteTextures                                 C.GPDELETETEXTURES
	gpDeleteTransformFeedbacks                       C.GPDELETETRANSFORMFEEDBACKS
	gpDeleteVertexArrays                             C.GPDELETEVERTEXARRAYS
	gpDeleteVertexArraysOES                          C.GPDELETEVERTEXARRAYSOES
	gpDepthFunc                                      C.GPDEPTHFUNC
	gpDepthMask                                      C.GPDEPTHMASK
	gpDepthRangeArrayfvNV                            C.GPDEPTHRANGEARRAYFVNV
	gpDepthRangeArrayfvOES                           C.GPDEPTHRANGEARRAYFVOES
	gpDepthRangeIndexedfNV                           C.GPDEPTHRANGEINDEXEDFNV
	gpDepthRangeIndexedfOES                          C.GPDEPTHRANGEINDEXEDFOES
	gpDepthRangef                                    C.GPDEPTHRANGEF
	gpDetachShader                                   C.GPDETACHSHADER
	gpDisable                                        C.GPDISABLE
	gpDisableDriverControlQCOM                       C.GPDISABLEDRIVERCONTROLQCOM
	gpDisableVertexAttribArray                       C.GPDISABLEVERTEXATTRIBARRAY
	gpDisableiEXT                                    C.GPDISABLEIEXT
	gpDisableiNV                                     C.GPDISABLEINV
	gpDisableiOES                                    C.GPDISABLEIOES
	gpDiscardFramebufferEXT                          C.GPDISCARDFRAMEBUFFEREXT
	gpDispatchCompute                                C.GPDISPATCHCOMPUTE
	gpDispatchComputeIndirect                        C.GPDISPATCHCOMPUTEINDIRECT
	gpDrawArrays                                     C.GPDRAWARRAYS
	gpDrawArraysIndirect                             C.GPDRAWARRAYSINDIRECT
	gpDrawArraysInstanced                            C.GPDRAWARRAYSINSTANCED
	gpDrawArraysInstancedANGLE                       C.GPDRAWARRAYSINSTANCEDANGLE
	gpDrawArraysInstancedBaseInstanceEXT             C.GPDRAWARRAYSINSTANCEDBASEINSTANCEEXT
	gpDrawArraysInstancedEXT                         C.GPDRAWARRAYSINSTANCEDEXT
	gpDrawArraysInstancedNV                          C.GPDRAWARRAYSINSTANCEDNV
	gpDrawBuffers                                    C.GPDRAWBUFFERS
	gpDrawBuffersEXT                                 C.GPDRAWBUFFERSEXT
	gpDrawBuffersIndexedEXT                          C.GPDRAWBUFFERSINDEXEDEXT
	gpDrawBuffersNV                                  C.GPDRAWBUFFERSNV
	gpDrawElements                                   C.GPDRAWELEMENTS
	gpDrawElementsBaseVertexEXT                      C.GPDRAWELEMENTSBASEVERTEXEXT
	gpDrawElementsBaseVertexOES                      C.GPDRAWELEMENTSBASEVERTEXOES
	gpDrawElementsIndirect                           C.GPDRAWELEMENTSINDIRECT
	gpDrawElementsInstanced                          C.GPDRAWELEMENTSINSTANCED
	gpDrawElementsInstancedANGLE                     C.GPDRAWELEMENTSINSTANCEDANGLE
	gpDrawElementsInstancedBaseInstanceEXT           C.GPDRAWELEMENTSINSTANCEDBASEINSTANCEEXT
	gpDrawElementsInstancedBaseVertexBaseInstanceEXT C.GPDRAWELEMENTSINSTANCEDBASEVERTEXBASEINSTANCEEXT
	gpDrawElementsInstancedBaseVertexEXT             C.GPDRAWELEMENTSINSTANCEDBASEVERTEXEXT
	gpDrawElementsInstancedBaseVertexOES             C.GPDRAWELEMENTSINSTANCEDBASEVERTEXOES
	gpDrawElementsInstancedEXT                       C.GPDRAWELEMENTSINSTANCEDEXT
	gpDrawElementsInstancedNV                        C.GPDRAWELEMENTSINSTANCEDNV
	gpDrawRangeElements                              C.GPDRAWRANGEELEMENTS
	gpDrawRangeElementsBaseVertexEXT                 C.GPDRAWRANGEELEMENTSBASEVERTEXEXT
	gpDrawRangeElementsBaseVertexOES                 C.GPDRAWRANGEELEMENTSBASEVERTEXOES
	gpDrawTransformFeedbackEXT                       C.GPDRAWTRANSFORMFEEDBACKEXT
	gpDrawTransformFeedbackInstancedEXT              C.GPDRAWTRANSFORMFEEDBACKINSTANCEDEXT
	gpDrawVkImageNV                                  C.GPDRAWVKIMAGENV
	gpEGLImageTargetRenderbufferStorageOES           C.GPEGLIMAGETARGETRENDERBUFFERSTORAGEOES
	gpEGLImageTargetTexStorageEXT                    C.GPEGLIMAGETARGETTEXSTORAGEEXT
	gpEGLImageTargetTexture2DOES                     C.GPEGLIMAGETARGETTEXTURE2DOES
	gpEGLImageTargetTextureStorageEXT                C.GPEGLIMAGETARGETTEXTURESTORAGEEXT
	gpEnable                                         C.GPENABLE
	gpEnableDriverControlQCOM                        C.GPENABLEDRIVERCONTROLQCOM
	gpEnableVertexAttribArray                        C.GPENABLEVERTEXATTRIBARRAY
	gpEnableiEXT                                     C.GPENABLEIEXT
	gpEnableiNV                                      C.GPENABLEINV
	gpEnableiOES                                     C.GPENABLEIOES
	gpEndConditionalRenderNV                         C.GPENDCONDITIONALRENDERNV
	gpEndPerfMonitorAMD                              C.GPENDPERFMONITORAMD
	gpEndPerfQueryINTEL                              C.GPENDPERFQUERYINTEL
	gpEndQuery                                       C.GPENDQUERY
	gpEndQueryEXT                                    C.GPENDQUERYEXT
	gpEndTilingQCOM                                  C.GPENDTILINGQCOM
	gpEndTransformFeedback                           C.GPENDTRANSFORMFEEDBACK
	gpExtGetBufferPointervQCOM                       C.GPEXTGETBUFFERPOINTERVQCOM
	gpExtGetBuffersQCOM                              C.GPEXTGETBUFFERSQCOM
	gpExtGetFramebuffersQCOM                         C.GPEXTGETFRAMEBUFFERSQCOM
	gpExtGetProgramBinarySourceQCOM                  C.GPEXTGETPROGRAMBINARYSOURCEQCOM
	gpExtGetProgramsQCOM                             C.GPEXTGETPROGRAMSQCOM
	gpExtGetRenderbuffersQCOM                        C.GPEXTGETRENDERBUFFERSQCOM
	gpExtGetShadersQCOM                              C.GPEXTGETSHADERSQCOM
	gpExtGetTexLevelParameterivQCOM                  C.GPEXTGETTEXLEVELPARAMETERIVQCOM
	gpExtGetTexSubImageQCOM                          C.GPEXTGETTEXSUBIMAGEQCOM
	gpExtGetTexturesQCOM                             C.GPEXTGETTEXTURESQCOM
	gpExtIsProgramBinaryQCOM                         C.GPEXTISPROGRAMBINARYQCOM
	gpExtTexObjectStateOverrideiQCOM                 C.GPEXTTEXOBJECTSTATEOVERRIDEIQCOM
	gpFenceSync                                      C.GPFENCESYNC
	gpFenceSyncAPPLE                                 C.GPFENCESYNCAPPLE
	gpFinish                                         C.GPFINISH
	gpFinishFenceNV                                  C.GPFINISHFENCENV
	gpFlush                                          C.GPFLUSH
	gpFlushMappedBufferRange                         C.GPFLUSHMAPPEDBUFFERRANGE
	gpFlushMappedBufferRangeEXT                      C.GPFLUSHMAPPEDBUFFERRANGEEXT
	gpFragmentCoverageColorNV                        C.GPFRAGMENTCOVERAGECOLORNV
	gpFramebufferFetchBarrierEXT                     C.GPFRAMEBUFFERFETCHBARRIEREXT
	gpFramebufferFetchBarrierQCOM                    C.GPFRAMEBUFFERFETCHBARRIERQCOM
	gpFramebufferFoveationConfigQCOM                 C.GPFRAMEBUFFERFOVEATIONCONFIGQCOM
	gpFramebufferFoveationParametersQCOM             C.GPFRAMEBUFFERFOVEATIONPARAMETERSQCOM
	gpFramebufferParameteri                          C.GPFRAMEBUFFERPARAMETERI
	gpFramebufferPixelLocalStorageSizeEXT            C.GPFRAMEBUFFERPIXELLOCALSTORAGESIZEEXT
	gpFramebufferRenderbuffer                        C.GPFRAMEBUFFERRENDERBUFFER
	gpFramebufferSampleLocationsfvNV                 C.GPFRAMEBUFFERSAMPLELOCATIONSFVNV
	gpFramebufferTexture2D                           C.GPFRAMEBUFFERTEXTURE2D
	gpFramebufferTexture2DDownsampleIMG              C.GPFRAMEBUFFERTEXTURE2DDOWNSAMPLEIMG
	gpFramebufferTexture2DMultisampleEXT             C.GPFRAMEBUFFERTEXTURE2DMULTISAMPLEEXT
	gpFramebufferTexture2DMultisampleIMG             C.GPFRAMEBUFFERTEXTURE2DMULTISAMPLEIMG
	gpFramebufferTexture3DOES                        C.GPFRAMEBUFFERTEXTURE3DOES
	gpFramebufferTextureEXT                          C.GPFRAMEBUFFERTEXTUREEXT
	gpFramebufferTextureLayer                        C.GPFRAMEBUFFERTEXTURELAYER
	gpFramebufferTextureLayerDownsampleIMG           C.GPFRAMEBUFFERTEXTURELAYERDOWNSAMPLEIMG
	gpFramebufferTextureMultisampleMultiviewOVR      C.GPFRAMEBUFFERTEXTUREMULTISAMPLEMULTIVIEWOVR
	gpFramebufferTextureMultiviewOVR                 C.GPFRAMEBUFFERTEXTUREMULTIVIEWOVR
	gpFramebufferTextureOES                          C.GPFRAMEBUFFERTEXTUREOES
	gpFrontFace                                      C.GPFRONTFACE
	gpGenBuffers                                     C.GPGENBUFFERS
	gpGenFencesNV                                    C.GPGENFENCESNV
	gpGenFramebuffers                                C.GPGENFRAMEBUFFERS
	gpGenPathsNV                                     C.GPGENPATHSNV
	gpGenPerfMonitorsAMD                             C.GPGENPERFMONITORSAMD
	gpGenProgramPipelines                            C.GPGENPROGRAMPIPELINES
	gpGenProgramPipelinesEXT                         C.GPGENPROGRAMPIPELINESEXT
	gpGenQueries                                     C.GPGENQUERIES
	gpGenQueriesEXT                                  C.GPGENQUERIESEXT
	gpGenRenderbuffers                               C.GPGENRENDERBUFFERS
	gpGenSamplers                                    C.GPGENSAMPLERS
	gpGenSemaphoresEXT                               C.GPGENSEMAPHORESEXT
	gpGenTextures                                    C.GPGENTEXTURES
	gpGenTransformFeedbacks                          C.GPGENTRANSFORMFEEDBACKS
	gpGenVertexArrays                                C.GPGENVERTEXARRAYS
	gpGenVertexArraysOES                             C.GPGENVERTEXARRAYSOES
	gpGenerateMipmap                                 C.GPGENERATEMIPMAP
	gpGetActiveAttrib                                C.GPGETACTIVEATTRIB
	gpGetActiveUniform                               C.GPGETACTIVEUNIFORM
	gpGetActiveUniformBlockName                      C.GPGETACTIVEUNIFORMBLOCKNAME
	gpGetActiveUniformBlockiv                        C.GPGETACTIVEUNIFORMBLOCKIV
	gpGetActiveUniformsiv                            C.GPGETACTIVEUNIFORMSIV
	gpGetAttachedShaders                             C.GPGETATTACHEDSHADERS
	gpGetAttribLocation                              C.GPGETATTRIBLOCATION
	gpGetBooleani_v                                  C.GPGETBOOLEANI_V
	gpGetBooleanv                                    C.GPGETBOOLEANV
	gpGetBufferParameteri64v                         C.GPGETBUFFERPARAMETERI64V
	gpGetBufferParameteriv                           C.GPGETBUFFERPARAMETERIV
	gpGetBufferPointerv                              C.GPGETBUFFERPOINTERV
	gpGetBufferPointervOES                           C.GPGETBUFFERPOINTERVOES
	gpGetCoverageModulationTableNV                   C.GPGETCOVERAGEMODULATIONTABLENV
	gpGetDebugMessageLog                             C.GPGETDEBUGMESSAGELOG
	gpGetDebugMessageLogKHR                          C.GPGETDEBUGMESSAGELOGKHR
	gpGetDriverControlStringQCOM                     C.GPGETDRIVERCONTROLSTRINGQCOM
	gpGetDriverControlsQCOM                          C.GPGETDRIVERCONTROLSQCOM
	gpGetError                                       C.GPGETERROR
	gpGetFenceivNV                                   C.GPGETFENCEIVNV
	gpGetFirstPerfQueryIdINTEL                       C.GPGETFIRSTPERFQUERYIDINTEL
	gpGetFloati_vNV                                  C.GPGETFLOATI_VNV
	gpGetFloati_vOES                                 C.GPGETFLOATI_VOES
	gpGetFloatv                                      C.GPGETFLOATV
	gpGetFragDataIndexEXT                            C.GPGETFRAGDATAINDEXEXT
	gpGetFragDataLocation                            C.GPGETFRAGDATALOCATION
	gpGetFramebufferAttachmentParameteriv            C.GPGETFRAMEBUFFERATTACHMENTPARAMETERIV
	gpGetFramebufferParameteriv                      C.GPGETFRAMEBUFFERPARAMETERIV
	gpGetFramebufferPixelLocalStorageSizeEXT         C.GPGETFRAMEBUFFERPIXELLOCALSTORAGESIZEEXT
	gpGetGraphicsResetStatus                         C.GPGETGRAPHICSRESETSTATUS
	gpGetGraphicsResetStatusEXT                      C.GPGETGRAPHICSRESETSTATUSEXT
	gpGetGraphicsResetStatusKHR                      C.GPGETGRAPHICSRESETSTATUSKHR
	gpGetImageHandleNV                               C.GPGETIMAGEHANDLENV
	gpGetInteger64i_v                                C.GPGETINTEGER64I_V
	gpGetInteger64v                                  C.GPGETINTEGER64V
	gpGetInteger64vAPPLE                             C.GPGETINTEGER64VAPPLE
	gpGetIntegeri_v                                  C.GPGETINTEGERI_V
	gpGetIntegeri_vEXT                               C.GPGETINTEGERI_VEXT
	gpGetIntegerv                                    C.GPGETINTEGERV
	gpGetInternalformatSampleivNV                    C.GPGETINTERNALFORMATSAMPLEIVNV
	gpGetInternalformativ                            C.GPGETINTERNALFORMATIV
	gpGetMemoryObjectParameterivEXT                  C.GPGETMEMORYOBJECTPARAMETERIVEXT
	gpGetMultisamplefv                               C.GPGETMULTISAMPLEFV
	gpGetNextPerfQueryIdINTEL                        C.GPGETNEXTPERFQUERYIDINTEL
	gpGetObjectLabel                                 C.GPGETOBJECTLABEL
	gpGetObjectLabelEXT                              C.GPGETOBJECTLABELEXT
	gpGetObjectLabelKHR                              C.GPGETOBJECTLABELKHR
	gpGetObjectPtrLabel                              C.GPGETOBJECTPTRLABEL
	gpGetObjectPtrLabelKHR                           C.GPGETOBJECTPTRLABELKHR
	gpGetPathCommandsNV                              C.GPGETPATHCOMMANDSNV
	gpGetPathCoordsNV                                C.GPGETPATHCOORDSNV
	gpGetPathDashArrayNV                             C.GPGETPATHDASHARRAYNV
	gpGetPathLengthNV                                C.GPGETPATHLENGTHNV
	gpGetPathMetricRangeNV                           C.GPGETPATHMETRICRANGENV
	gpGetPathMetricsNV                               C.GPGETPATHMETRICSNV
	gpGetPathParameterfvNV                           C.GPGETPATHPARAMETERFVNV
	gpGetPathParameterivNV                           C.GPGETPATHPARAMETERIVNV
	gpGetPathSpacingNV                               C.GPGETPATHSPACINGNV
	gpGetPerfCounterInfoINTEL                        C.GPGETPERFCOUNTERINFOINTEL
	gpGetPerfMonitorCounterDataAMD                   C.GPGETPERFMONITORCOUNTERDATAAMD
	gpGetPerfMonitorCounterInfoAMD                   C.GPGETPERFMONITORCOUNTERINFOAMD
	gpGetPerfMonitorCounterStringAMD                 C.GPGETPERFMONITORCOUNTERSTRINGAMD
	gpGetPerfMonitorCountersAMD                      C.GPGETPERFMONITORCOUNTERSAMD
	gpGetPerfMonitorGroupStringAMD                   C.GPGETPERFMONITORGROUPSTRINGAMD
	gpGetPerfMonitorGroupsAMD                        C.GPGETPERFMONITORGROUPSAMD
	gpGetPerfQueryDataINTEL                          C.GPGETPERFQUERYDATAINTEL
	gpGetPerfQueryIdByNameINTEL                      C.GPGETPERFQUERYIDBYNAMEINTEL
	gpGetPerfQueryInfoINTEL                          C.GPGETPERFQUERYINFOINTEL
	gpGetPointerv                                    C.GPGETPOINTERV
	gpGetPointervKHR                                 C.GPGETPOINTERVKHR
	gpGetProgramBinary                               C.GPGETPROGRAMBINARY
	gpGetProgramBinaryOES                            C.GPGETPROGRAMBINARYOES
	gpGetProgramInfoLog                              C.GPGETPROGRAMINFOLOG
	gpGetProgramInterfaceiv                          C.GPGETPROGRAMINTERFACEIV
	gpGetProgramPipelineInfoLog                      C.GPGETPROGRAMPIPELINEINFOLOG
	gpGetProgramPipelineInfoLogEXT                   C.GPGETPROGRAMPIPELINEINFOLOGEXT
	gpGetProgramPipelineiv                           C.GPGETPROGRAMPIPELINEIV
	gpGetProgramPipelineivEXT                        C.GPGETPROGRAMPIPELINEIVEXT
	gpGetProgramResourceIndex                        C.GPGETPROGRAMRESOURCEINDEX
	gpGetProgramResourceLocation                     C.GPGETPROGRAMRESOURCELOCATION
	gpGetProgramResourceLocationIndexEXT             C.GPGETPROGRAMRESOURCELOCATIONINDEXEXT
	gpGetProgramResourceName                         C.GPGETPROGRAMRESOURCENAME
	gpGetProgramResourcefvNV                         C.GPGETPROGRAMRESOURCEFVNV
	gpGetProgramResourceiv                           C.GPGETPROGRAMRESOURCEIV
	gpGetProgramiv                                   C.GPGETPROGRAMIV
	gpGetQueryObjecti64vEXT                          C.GPGETQUERYOBJECTI64VEXT
	gpGetQueryObjectivEXT                            C.GPGETQUERYOBJECTIVEXT
	gpGetQueryObjectui64vEXT                         C.GPGETQUERYOBJECTUI64VEXT
	gpGetQueryObjectuiv                              C.GPGETQUERYOBJECTUIV
	gpGetQueryObjectuivEXT                           C.GPGETQUERYOBJECTUIVEXT
	gpGetQueryiv                                     C.GPGETQUERYIV
	gpGetQueryivEXT                                  C.GPGETQUERYIVEXT
	gpGetRenderbufferParameteriv                     C.GPGETRENDERBUFFERPARAMETERIV
	gpGetSamplerParameterIivEXT                      C.GPGETSAMPLERPARAMETERIIVEXT
	gpGetSamplerParameterIivOES                      C.GPGETSAMPLERPARAMETERIIVOES
	gpGetSamplerParameterIuivEXT                     C.GPGETSAMPLERPARAMETERIUIVEXT
	gpGetSamplerParameterIuivOES                     C.GPGETSAMPLERPARAMETERIUIVOES
	gpGetSamplerParameterfv                          C.GPGETSAMPLERPARAMETERFV
	gpGetSamplerParameteriv                          C.GPGETSAMPLERPARAMETERIV
	gpGetSemaphoreParameterui64vEXT                  C.GPGETSEMAPHOREPARAMETERUI64VEXT
	gpGetShaderInfoLog                               C.GPGETSHADERINFOLOG
	gpGetShaderPrecisionFormat                       C.GPGETSHADERPRECISIONFORMAT
	gpGetShaderSource                                C.GPGETSHADERSOURCE
	gpGetShaderiv                                    C.GPGETSHADERIV
	gpGetString                                      C.GPGETSTRING
	gpGetStringi                                     C.GPGETSTRINGI
	gpGetSynciv                                      C.GPGETSYNCIV
	gpGetSyncivAPPLE                                 C.GPGETSYNCIVAPPLE
	gpGetTexLevelParameterfv                         C.GPGETTEXLEVELPARAMETERFV
	gpGetTexLevelParameteriv                         C.GPGETTEXLEVELPARAMETERIV
	gpGetTexParameterIivEXT                          C.GPGETTEXPARAMETERIIVEXT
	gpGetTexParameterIivOES                          C.GPGETTEXPARAMETERIIVOES
	gpGetTexParameterIuivEXT                         C.GPGETTEXPARAMETERIUIVEXT
	gpGetTexParameterIuivOES                         C.GPGETTEXPARAMETERIUIVOES
	gpGetTexParameterfv                              C.GPGETTEXPARAMETERFV
	gpGetTexParameteriv                              C.GPGETTEXPARAMETERIV
	gpGetTextureHandleIMG                            C.GPGETTEXTUREHANDLEIMG
	gpGetTextureHandleNV                             C.GPGETTEXTUREHANDLENV
	gpGetTextureSamplerHandleIMG                     C.GPGETTEXTURESAMPLERHANDLEIMG
	gpGetTextureSamplerHandleNV                      C.GPGETTEXTURESAMPLERHANDLENV
	gpGetTransformFeedbackVarying                    C.GPGETTRANSFORMFEEDBACKVARYING
	gpGetTranslatedShaderSourceANGLE                 C.GPGETTRANSLATEDSHADERSOURCEANGLE
	gpGetUniformBlockIndex                           C.GPGETUNIFORMBLOCKINDEX
	gpGetUniformIndices                              C.GPGETUNIFORMINDICES
	gpGetUniformLocation                             C.GPGETUNIFORMLOCATION
	gpGetUniformfv                                   C.GPGETUNIFORMFV
	gpGetUniformi64vNV                               C.GPGETUNIFORMI64VNV
	gpGetUniformiv                                   C.GPGETUNIFORMIV
	gpGetUniformuiv                                  C.GPGETUNIFORMUIV
	gpGetUnsignedBytei_vEXT                          C.GPGETUNSIGNEDBYTEI_VEXT
	gpGetUnsignedBytevEXT                            C.GPGETUNSIGNEDBYTEVEXT
	gpGetVertexAttribIiv                             C.GPGETVERTEXATTRIBIIV
	gpGetVertexAttribIuiv                            C.GPGETVERTEXATTRIBIUIV
	gpGetVertexAttribPointerv                        C.GPGETVERTEXATTRIBPOINTERV
	gpGetVertexAttribfv                              C.GPGETVERTEXATTRIBFV
	gpGetVertexAttribiv                              C.GPGETVERTEXATTRIBIV
	gpGetVkProcAddrNV                                C.GPGETVKPROCADDRNV
	gpGetnUniformfv                                  C.GPGETNUNIFORMFV
	gpGetnUniformfvEXT                               C.GPGETNUNIFORMFVEXT
	gpGetnUniformfvKHR                               C.GPGETNUNIFORMFVKHR
	gpGetnUniformiv                                  C.GPGETNUNIFORMIV
	gpGetnUniformivEXT                               C.GPGETNUNIFORMIVEXT
	gpGetnUniformivKHR                               C.GPGETNUNIFORMIVKHR
	gpGetnUniformuiv                                 C.GPGETNUNIFORMUIV
	gpGetnUniformuivKHR                              C.GPGETNUNIFORMUIVKHR
	gpHint                                           C.GPHINT
	gpImportMemoryFdEXT                              C.GPIMPORTMEMORYFDEXT
	gpImportMemoryWin32HandleEXT                     C.GPIMPORTMEMORYWIN32HANDLEEXT
	gpImportMemoryWin32NameEXT                       C.GPIMPORTMEMORYWIN32NAMEEXT
	gpImportSemaphoreFdEXT                           C.GPIMPORTSEMAPHOREFDEXT
	gpImportSemaphoreWin32HandleEXT                  C.GPIMPORTSEMAPHOREWIN32HANDLEEXT
	gpImportSemaphoreWin32NameEXT                    C.GPIMPORTSEMAPHOREWIN32NAMEEXT
	gpInsertEventMarkerEXT                           C.GPINSERTEVENTMARKEREXT
	gpInterpolatePathsNV                             C.GPINTERPOLATEPATHSNV
	gpInvalidateFramebuffer                          C.GPINVALIDATEFRAMEBUFFER
	gpInvalidateSubFramebuffer                       C.GPINVALIDATESUBFRAMEBUFFER
	gpIsBuffer                                       C.GPISBUFFER
	gpIsEnabled                                      C.GPISENABLED
	gpIsEnablediEXT                                  C.GPISENABLEDIEXT
	gpIsEnablediNV                                   C.GPISENABLEDINV
	gpIsEnablediOES                                  C.GPISENABLEDIOES
	gpIsFenceNV                                      C.GPISFENCENV
	gpIsFramebuffer                                  C.GPISFRAMEBUFFER
	gpIsImageHandleResidentNV                        C.GPISIMAGEHANDLERESIDENTNV
	gpIsMemoryObjectEXT                              C.GPISMEMORYOBJECTEXT
	gpIsPathNV                                       C.GPISPATHNV
	gpIsPointInFillPathNV                            C.GPISPOINTINFILLPATHNV
	gpIsPointInStrokePathNV                          C.GPISPOINTINSTROKEPATHNV
	gpIsProgram                                      C.GPISPROGRAM
	gpIsProgramPipeline                              C.GPISPROGRAMPIPELINE
	gpIsProgramPipelineEXT                           C.GPISPROGRAMPIPELINEEXT
	gpIsQuery                                        C.GPISQUERY
	gpIsQueryEXT                                     C.GPISQUERYEXT
	gpIsRenderbuffer                                 C.GPISRENDERBUFFER
	gpIsSampler                                      C.GPISSAMPLER
	gpIsSemaphoreEXT                                 C.GPISSEMAPHOREEXT
	gpIsShader                                       C.GPISSHADER
	gpIsSync                                         C.GPISSYNC
	gpIsSyncAPPLE                                    C.GPISSYNCAPPLE
	gpIsTexture                                      C.GPISTEXTURE
	gpIsTextureHandleResidentNV                      C.GPISTEXTUREHANDLERESIDENTNV
	gpIsTransformFeedback                            C.GPISTRANSFORMFEEDBACK
	gpIsVertexArray                                  C.GPISVERTEXARRAY
	gpIsVertexArrayOES                               C.GPISVERTEXARRAYOES
	gpLabelObjectEXT                                 C.GPLABELOBJECTEXT
	gpLineWidth                                      C.GPLINEWIDTH
	gpLinkProgram                                    C.GPLINKPROGRAM
	gpMakeImageHandleNonResidentNV                   C.GPMAKEIMAGEHANDLENONRESIDENTNV
	gpMakeImageHandleResidentNV                      C.GPMAKEIMAGEHANDLERESIDENTNV
	gpMakeTextureHandleNonResidentNV                 C.GPMAKETEXTUREHANDLENONRESIDENTNV
	gpMakeTextureHandleResidentNV                    C.GPMAKETEXTUREHANDLERESIDENTNV
	gpMapBufferOES                                   C.GPMAPBUFFEROES
	gpMapBufferRange                                 C.GPMAPBUFFERRANGE
	gpMapBufferRangeEXT                              C.GPMAPBUFFERRANGEEXT
	gpMatrixFrustumEXT                               C.GPMATRIXFRUSTUMEXT
	gpMatrixLoad3x2fNV                               C.GPMATRIXLOAD3X2FNV
	gpMatrixLoad3x3fNV                               C.GPMATRIXLOAD3X3FNV
	gpMatrixLoadIdentityEXT                          C.GPMATRIXLOADIDENTITYEXT
	gpMatrixLoadTranspose3x3fNV                      C.GPMATRIXLOADTRANSPOSE3X3FNV
	gpMatrixLoadTransposedEXT                        C.GPMATRIXLOADTRANSPOSEDEXT
	gpMatrixLoadTransposefEXT                        C.GPMATRIXLOADTRANSPOSEFEXT
	gpMatrixLoaddEXT                                 C.GPMATRIXLOADDEXT
	gpMatrixLoadfEXT                                 C.GPMATRIXLOADFEXT
	gpMatrixMult3x2fNV                               C.GPMATRIXMULT3X2FNV
	gpMatrixMult3x3fNV                               C.GPMATRIXMULT3X3FNV
	gpMatrixMultTranspose3x3fNV                      C.GPMATRIXMULTTRANSPOSE3X3FNV
	gpMatrixMultTransposedEXT                        C.GPMATRIXMULTTRANSPOSEDEXT
	gpMatrixMultTransposefEXT                        C.GPMATRIXMULTTRANSPOSEFEXT
	gpMatrixMultdEXT                                 C.GPMATRIXMULTDEXT
	gpMatrixMultfEXT                                 C.GPMATRIXMULTFEXT
	gpMatrixOrthoEXT                                 C.GPMATRIXORTHOEXT
	gpMatrixPopEXT                                   C.GPMATRIXPOPEXT
	gpMatrixPushEXT                                  C.GPMATRIXPUSHEXT
	gpMatrixRotatedEXT                               C.GPMATRIXROTATEDEXT
	gpMatrixRotatefEXT                               C.GPMATRIXROTATEFEXT
	gpMatrixScaledEXT                                C.GPMATRIXSCALEDEXT
	gpMatrixScalefEXT                                C.GPMATRIXSCALEFEXT
	gpMatrixTranslatedEXT                            C.GPMATRIXTRANSLATEDEXT
	gpMatrixTranslatefEXT                            C.GPMATRIXTRANSLATEFEXT
	gpMaxShaderCompilerThreadsKHR                    C.GPMAXSHADERCOMPILERTHREADSKHR
	gpMemoryBarrier                                  C.GPMEMORYBARRIER
	gpMemoryBarrierByRegion                          C.GPMEMORYBARRIERBYREGION
	gpMemoryObjectParameterivEXT                     C.GPMEMORYOBJECTPARAMETERIVEXT
	gpMinSampleShadingOES                            C.GPMINSAMPLESHADINGOES
	gpMultiDrawArraysEXT                             C.GPMULTIDRAWARRAYSEXT
	gpMultiDrawArraysIndirectEXT                     C.GPMULTIDRAWARRAYSINDIRECTEXT
	gpMultiDrawElementsBaseVertexEXT                 C.GPMULTIDRAWELEMENTSBASEVERTEXEXT
	gpMultiDrawElementsEXT                           C.GPMULTIDRAWELEMENTSEXT
	gpMultiDrawElementsIndirectEXT                   C.GPMULTIDRAWELEMENTSINDIRECTEXT
	gpNamedBufferStorageExternalEXT                  C.GPNAMEDBUFFERSTORAGEEXTERNALEXT
	gpNamedBufferStorageMemEXT                       C.GPNAMEDBUFFERSTORAGEMEMEXT
	gpNamedFramebufferSampleLocationsfvNV            C.GPNAMEDFRAMEBUFFERSAMPLELOCATIONSFVNV
	gpObjectLabel                                    C.GPOBJECTLABEL
	gpObjectLabelKHR                                 C.GPOBJECTLABELKHR
	gpObjectPtrLabel                                 C.GPOBJECTPTRLABEL
	gpObjectPtrLabelKHR                              C.GPOBJECTPTRLABELKHR
	gpPatchParameteriEXT                             C.GPPATCHPARAMETERIEXT
	gpPatchParameteriOES                             C.GPPATCHPARAMETERIOES
	gpPathCommandsNV                                 C.GPPATHCOMMANDSNV
	gpPathCoordsNV                                   C.GPPATHCOORDSNV
	gpPathCoverDepthFuncNV                           C.GPPATHCOVERDEPTHFUNCNV
	gpPathDashArrayNV                                C.GPPATHDASHARRAYNV
	gpPathGlyphIndexArrayNV                          C.GPPATHGLYPHINDEXARRAYNV
	gpPathGlyphIndexRangeNV                          C.GPPATHGLYPHINDEXRANGENV
	gpPathGlyphRangeNV                               C.GPPATHGLYPHRANGENV
	gpPathGlyphsNV                                   C.GPPATHGLYPHSNV
	gpPathMemoryGlyphIndexArrayNV                    C.GPPATHMEMORYGLYPHINDEXARRAYNV
	gpPathParameterfNV                               C.GPPATHPARAMETERFNV
	gpPathParameterfvNV                              C.GPPATHPARAMETERFVNV
	gpPathParameteriNV                               C.GPPATHPARAMETERINV
	gpPathParameterivNV                              C.GPPATHPARAMETERIVNV
	gpPathStencilDepthOffsetNV                       C.GPPATHSTENCILDEPTHOFFSETNV
	gpPathStencilFuncNV                              C.GPPATHSTENCILFUNCNV
	gpPathStringNV                                   C.GPPATHSTRINGNV
	gpPathSubCommandsNV                              C.GPPATHSUBCOMMANDSNV
	gpPathSubCoordsNV                                C.GPPATHSUBCOORDSNV
	gpPauseTransformFeedback                         C.GPPAUSETRANSFORMFEEDBACK
	gpPixelStorei                                    C.GPPIXELSTOREI
	gpPointAlongPathNV                               C.GPPOINTALONGPATHNV
	gpPolygonModeNV                                  C.GPPOLYGONMODENV
	gpPolygonOffset                                  C.GPPOLYGONOFFSET
	gpPolygonOffsetClampEXT                          C.GPPOLYGONOFFSETCLAMPEXT
	gpPopDebugGroup                                  C.GPPOPDEBUGGROUP
	gpPopDebugGroupKHR                               C.GPPOPDEBUGGROUPKHR
	gpPopGroupMarkerEXT                              C.GPPOPGROUPMARKEREXT
	gpPrimitiveBoundingBoxEXT                        C.GPPRIMITIVEBOUNDINGBOXEXT
	gpPrimitiveBoundingBoxOES                        C.GPPRIMITIVEBOUNDINGBOXOES
	gpProgramBinary                                  C.GPPROGRAMBINARY
	gpProgramBinaryOES                               C.GPPROGRAMBINARYOES
	gpProgramParameteri                              C.GPPROGRAMPARAMETERI
	gpProgramParameteriEXT                           C.GPPROGRAMPARAMETERIEXT
	gpProgramPathFragmentInputGenNV                  C.GPPROGRAMPATHFRAGMENTINPUTGENNV
	gpProgramUniform1f                               C.GPPROGRAMUNIFORM1F
	gpProgramUniform1fEXT                            C.GPPROGRAMUNIFORM1FEXT
	gpProgramUniform1fv                              C.GPPROGRAMUNIFORM1FV
	gpProgramUniform1fvEXT                           C.GPPROGRAMUNIFORM1FVEXT
	gpProgramUniform1i                               C.GPPROGRAMUNIFORM1I
	gpProgramUniform1i64NV                           C.GPPROGRAMUNIFORM1I64NV
	gpProgramUniform1i64vNV                          C.GPPROGRAMUNIFORM1I64VNV
	gpProgramUniform1iEXT                            C.GPPROGRAMUNIFORM1IEXT
	gpProgramUniform1iv                              C.GPPROGRAMUNIFORM1IV
	gpProgramUniform1ivEXT                           C.GPPROGRAMUNIFORM1IVEXT
	gpProgramUniform1ui                              C.GPPROGRAMUNIFORM1UI
	gpProgramUniform1ui64NV                          C.GPPROGRAMUNIFORM1UI64NV
	gpProgramUniform1ui64vNV                         C.GPPROGRAMUNIFORM1UI64VNV
	gpProgramUniform1uiEXT                           C.GPPROGRAMUNIFORM1UIEXT
	gpProgramUniform1uiv                             C.GPPROGRAMUNIFORM1UIV
	gpProgramUniform1uivEXT                          C.GPPROGRAMUNIFORM1UIVEXT
	gpProgramUniform2f                               C.GPPROGRAMUNIFORM2F
	gpProgramUniform2fEXT                            C.GPPROGRAMUNIFORM2FEXT
	gpProgramUniform2fv                              C.GPPROGRAMUNIFORM2FV
	gpProgramUniform2fvEXT                           C.GPPROGRAMUNIFORM2FVEXT
	gpProgramUniform2i                               C.GPPROGRAMUNIFORM2I
	gpProgramUniform2i64NV                           C.GPPROGRAMUNIFORM2I64NV
	gpProgramUniform2i64vNV                          C.GPPROGRAMUNIFORM2I64VNV
	gpProgramUniform2iEXT                            C.GPPROGRAMUNIFORM2IEXT
	gpProgramUniform2iv                              C.GPPROGRAMUNIFORM2IV
	gpProgramUniform2ivEXT                           C.GPPROGRAMUNIFORM2IVEXT
	gpProgramUniform2ui                              C.GPPROGRAMUNIFORM2UI
	gpProgramUniform2ui64NV                          C.GPPROGRAMUNIFORM2UI64NV
	gpProgramUniform2ui64vNV                         C.GPPROGRAMUNIFORM2UI64VNV
	gpProgramUniform2uiEXT                           C.GPPROGRAMUNIFORM2UIEXT
	gpProgramUniform2uiv                             C.GPPROGRAMUNIFORM2UIV
	gpProgramUniform2uivEXT                          C.GPPROGRAMUNIFORM2UIVEXT
	gpProgramUniform3f                               C.GPPROGRAMUNIFORM3F
	gpProgramUniform3fEXT                            C.GPPROGRAMUNIFORM3FEXT
	gpProgramUniform3fv                              C.GPPROGRAMUNIFORM3FV
	gpProgramUniform3fvEXT                           C.GPPROGRAMUNIFORM3FVEXT
	gpProgramUniform3i                               C.GPPROGRAMUNIFORM3I
	gpProgramUniform3i64NV                           C.GPPROGRAMUNIFORM3I64NV
	gpProgramUniform3i64vNV                          C.GPPROGRAMUNIFORM3I64VNV
	gpProgramUniform3iEXT                            C.GPPROGRAMUNIFORM3IEXT
	gpProgramUniform3iv                              C.GPPROGRAMUNIFORM3IV
	gpProgramUniform3ivEXT                           C.GPPROGRAMUNIFORM3IVEXT
	gpProgramUniform3ui                              C.GPPROGRAMUNIFORM3UI
	gpProgramUniform3ui64NV                          C.GPPROGRAMUNIFORM3UI64NV
	gpProgramUniform3ui64vNV                         C.GPPROGRAMUNIFORM3UI64VNV
	gpProgramUniform3uiEXT                           C.GPPROGRAMUNIFORM3UIEXT
	gpProgramUniform3uiv                             C.GPPROGRAMUNIFORM3UIV
	gpProgramUniform3uivEXT                          C.GPPROGRAMUNIFORM3UIVEXT
	gpProgramUniform4f                               C.GPPROGRAMUNIFORM4F
	gpProgramUniform4fEXT                            C.GPPROGRAMUNIFORM4FEXT
	gpProgramUniform4fv                              C.GPPROGRAMUNIFORM4FV
	gpProgramUniform4fvEXT                           C.GPPROGRAMUNIFORM4FVEXT
	gpProgramUniform4i                               C.GPPROGRAMUNIFORM4I
	gpProgramUniform4i64NV                           C.GPPROGRAMUNIFORM4I64NV
	gpProgramUniform4i64vNV                          C.GPPROGRAMUNIFORM4I64VNV
	gpProgramUniform4iEXT                            C.GPPROGRAMUNIFORM4IEXT
	gpProgramUniform4iv                              C.GPPROGRAMUNIFORM4IV
	gpProgramUniform4ivEXT                           C.GPPROGRAMUNIFORM4IVEXT
	gpProgramUniform4ui                              C.GPPROGRAMUNIFORM4UI
	gpProgramUniform4ui64NV                          C.GPPROGRAMUNIFORM4UI64NV
	gpProgramUniform4ui64vNV                         C.GPPROGRAMUNIFORM4UI64VNV
	gpProgramUniform4uiEXT                           C.GPPROGRAMUNIFORM4UIEXT
	gpProgramUniform4uiv                             C.GPPROGRAMUNIFORM4UIV
	gpProgramUniform4uivEXT                          C.GPPROGRAMUNIFORM4UIVEXT
	gpProgramUniformHandleui64IMG                    C.GPPROGRAMUNIFORMHANDLEUI64IMG
	gpProgramUniformHandleui64NV                     C.GPPROGRAMUNIFORMHANDLEUI64NV
	gpProgramUniformHandleui64vIMG                   C.GPPROGRAMUNIFORMHANDLEUI64VIMG
	gpProgramUniformHandleui64vNV                    C.GPPROGRAMUNIFORMHANDLEUI64VNV
	gpProgramUniformMatrix2fv                        C.GPPROGRAMUNIFORMMATRIX2FV
	gpProgramUniformMatrix2fvEXT                     C.GPPROGRAMUNIFORMMATRIX2FVEXT
	gpProgramUniformMatrix2x3fv                      C.GPPROGRAMUNIFORMMATRIX2X3FV
	gpProgramUniformMatrix2x3fvEXT                   C.GPPROGRAMUNIFORMMATRIX2X3FVEXT
	gpProgramUniformMatrix2x4fv                      C.GPPROGRAMUNIFORMMATRIX2X4FV
	gpProgramUniformMatrix2x4fvEXT                   C.GPPROGRAMUNIFORMMATRIX2X4FVEXT
	gpProgramUniformMatrix3fv                        C.GPPROGRAMUNIFORMMATRIX3FV
	gpProgramUniformMatrix3fvEXT                     C.GPPROGRAMUNIFORMMATRIX3FVEXT
	gpProgramUniformMatrix3x2fv                      C.GPPROGRAMUNIFORMMATRIX3X2FV
	gpProgramUniformMatrix3x2fvEXT                   C.GPPROGRAMUNIFORMMATRIX3X2FVEXT
	gpProgramUniformMatrix3x4fv                      C.GPPROGRAMUNIFORMMATRIX3X4FV
	gpProgramUniformMatrix3x4fvEXT                   C.GPPROGRAMUNIFORMMATRIX3X4FVEXT
	gpProgramUniformMatrix4fv                        C.GPPROGRAMUNIFORMMATRIX4FV
	gpProgramUniformMatrix4fvEXT                     C.GPPROGRAMUNIFORMMATRIX4FVEXT
	gpProgramUniformMatrix4x2fv                      C.GPPROGRAMUNIFORMMATRIX4X2FV
	gpProgramUniformMatrix4x2fvEXT                   C.GPPROGRAMUNIFORMMATRIX4X2FVEXT
	gpProgramUniformMatrix4x3fv                      C.GPPROGRAMUNIFORMMATRIX4X3FV
	gpProgramUniformMatrix4x3fvEXT                   C.GPPROGRAMUNIFORMMATRIX4X3FVEXT
	gpPushDebugGroup                                 C.GPPUSHDEBUGGROUP
	gpPushDebugGroupKHR                              C.GPPUSHDEBUGGROUPKHR
	gpPushGroupMarkerEXT                             C.GPPUSHGROUPMARKEREXT
	gpQueryCounterEXT                                C.GPQUERYCOUNTEREXT
	gpRasterSamplesEXT                               C.GPRASTERSAMPLESEXT
	gpReadBuffer                                     C.GPREADBUFFER
	gpReadBufferIndexedEXT                           C.GPREADBUFFERINDEXEDEXT
	gpReadBufferNV                                   C.GPREADBUFFERNV
	gpReadPixels                                     C.GPREADPIXELS
	gpReadnPixels                                    C.GPREADNPIXELS
	gpReadnPixelsEXT                                 C.GPREADNPIXELSEXT
	gpReadnPixelsKHR                                 C.GPREADNPIXELSKHR
	gpReleaseKeyedMutexWin32EXT                      C.GPRELEASEKEYEDMUTEXWIN32EXT
	gpReleaseShaderCompiler                          C.GPRELEASESHADERCOMPILER
	gpRenderbufferStorage                            C.GPRENDERBUFFERSTORAGE
	gpRenderbufferStorageMultisample                 C.GPRENDERBUFFERSTORAGEMULTISAMPLE
	gpRenderbufferStorageMultisampleANGLE            C.GPRENDERBUFFERSTORAGEMULTISAMPLEANGLE
	gpRenderbufferStorageMultisampleAPPLE            C.GPRENDERBUFFERSTORAGEMULTISAMPLEAPPLE
	gpRenderbufferStorageMultisampleEXT              C.GPRENDERBUFFERSTORAGEMULTISAMPLEEXT
	gpRenderbufferStorageMultisampleIMG              C.GPRENDERBUFFERSTORAGEMULTISAMPLEIMG
	gpRenderbufferStorageMultisampleNV               C.GPRENDERBUFFERSTORAGEMULTISAMPLENV
	gpResolveDepthValuesNV                           C.GPRESOLVEDEPTHVALUESNV
	gpResolveMultisampleFramebufferAPPLE             C.GPRESOLVEMULTISAMPLEFRAMEBUFFERAPPLE
	gpResumeTransformFeedback                        C.GPRESUMETRANSFORMFEEDBACK
	gpSampleCoverage                                 C.GPSAMPLECOVERAGE
	gpSampleMaski                                    C.GPSAMPLEMASKI
	gpSamplerParameterIivEXT                         C.GPSAMPLERPARAMETERIIVEXT
	gpSamplerParameterIivOES                         C.GPSAMPLERPARAMETERIIVOES
	gpSamplerParameterIuivEXT                        C.GPSAMPLERPARAMETERIUIVEXT
	gpSamplerParameterIuivOES                        C.GPSAMPLERPARAMETERIUIVOES
	gpSamplerParameterf                              C.GPSAMPLERPARAMETERF
	gpSamplerParameterfv                             C.GPSAMPLERPARAMETERFV
	gpSamplerParameteri                              C.GPSAMPLERPARAMETERI
	gpSamplerParameteriv                             C.GPSAMPLERPARAMETERIV
	gpScissor                                        C.GPSCISSOR
	gpScissorArrayvNV                                C.GPSCISSORARRAYVNV
	gpScissorArrayvOES                               C.GPSCISSORARRAYVOES
	gpScissorIndexedNV                               C.GPSCISSORINDEXEDNV
	gpScissorIndexedOES                              C.GPSCISSORINDEXEDOES
	gpScissorIndexedvNV                              C.GPSCISSORINDEXEDVNV
	gpScissorIndexedvOES                             C.GPSCISSORINDEXEDVOES
	gpSelectPerfMonitorCountersAMD                   C.GPSELECTPERFMONITORCOUNTERSAMD
	gpSemaphoreParameterui64vEXT                     C.GPSEMAPHOREPARAMETERUI64VEXT
	gpSetFenceNV                                     C.GPSETFENCENV
	gpShaderBinary                                   C.GPSHADERBINARY
	gpShaderSource                                   C.GPSHADERSOURCE
	gpSignalSemaphoreEXT                             C.GPSIGNALSEMAPHOREEXT
	gpSignalVkFenceNV                                C.GPSIGNALVKFENCENV
	gpSignalVkSemaphoreNV                            C.GPSIGNALVKSEMAPHORENV
	gpStartTilingQCOM                                C.GPSTARTTILINGQCOM
	gpStencilFillPathInstancedNV                     C.GPSTENCILFILLPATHINSTANCEDNV
	gpStencilFillPathNV                              C.GPSTENCILFILLPATHNV
	gpStencilFunc                                    C.GPSTENCILFUNC
	gpStencilFuncSeparate                            C.GPSTENCILFUNCSEPARATE
	gpStencilMask                                    C.GPSTENCILMASK
	gpStencilMaskSeparate                            C.GPSTENCILMASKSEPARATE
	gpStencilOp                                      C.GPSTENCILOP
	gpStencilOpSeparate                              C.GPSTENCILOPSEPARATE
	gpStencilStrokePathInstancedNV                   C.GPSTENCILSTROKEPATHINSTANCEDNV
	gpStencilStrokePathNV                            C.GPSTENCILSTROKEPATHNV
	gpStencilThenCoverFillPathInstancedNV            C.GPSTENCILTHENCOVERFILLPATHINSTANCEDNV
	gpStencilThenCoverFillPathNV                     C.GPSTENCILTHENCOVERFILLPATHNV
	gpStencilThenCoverStrokePathInstancedNV          C.GPSTENCILTHENCOVERSTROKEPATHINSTANCEDNV
	gpStencilThenCoverStrokePathNV                   C.GPSTENCILTHENCOVERSTROKEPATHNV
	gpSubpixelPrecisionBiasNV                        C.GPSUBPIXELPRECISIONBIASNV
	gpTestFenceNV                                    C.GPTESTFENCENV
	gpTexBufferEXT                                   C.GPTEXBUFFEREXT
	gpTexBufferOES                                   C.GPTEXBUFFEROES
	gpTexBufferRangeEXT                              C.GPTEXBUFFERRANGEEXT
	gpTexBufferRangeOES                              C.GPTEXBUFFERRANGEOES
	gpTexImage2D                                     C.GPTEXIMAGE2D
	gpTexImage3D                                     C.GPTEXIMAGE3D
	gpTexImage3DOES                                  C.GPTEXIMAGE3DOES
	gpTexPageCommitmentEXT                           C.GPTEXPAGECOMMITMENTEXT
	gpTexParameterIivEXT                             C.GPTEXPARAMETERIIVEXT
	gpTexParameterIivOES                             C.GPTEXPARAMETERIIVOES
	gpTexParameterIuivEXT                            C.GPTEXPARAMETERIUIVEXT
	gpTexParameterIuivOES                            C.GPTEXPARAMETERIUIVOES
	gpTexParameterf                                  C.GPTEXPARAMETERF
	gpTexParameterfv                                 C.GPTEXPARAMETERFV
	gpTexParameteri                                  C.GPTEXPARAMETERI
	gpTexParameteriv                                 C.GPTEXPARAMETERIV
	gpTexStorage1DEXT                                C.GPTEXSTORAGE1DEXT
	gpTexStorage2D                                   C.GPTEXSTORAGE2D
	gpTexStorage2DEXT                                C.GPTEXSTORAGE2DEXT
	gpTexStorage2DMultisample                        C.GPTEXSTORAGE2DMULTISAMPLE
	gpTexStorage3D                                   C.GPTEXSTORAGE3D
	gpTexStorage3DEXT                                C.GPTEXSTORAGE3DEXT
	gpTexStorage3DMultisampleOES                     C.GPTEXSTORAGE3DMULTISAMPLEOES
	gpTexStorageMem1DEXT                             C.GPTEXSTORAGEMEM1DEXT
	gpTexStorageMem2DEXT                             C.GPTEXSTORAGEMEM2DEXT
	gpTexStorageMem2DMultisampleEXT                  C.GPTEXSTORAGEMEM2DMULTISAMPLEEXT
	gpTexStorageMem3DEXT                             C.GPTEXSTORAGEMEM3DEXT
	gpTexStorageMem3DMultisampleEXT                  C.GPTEXSTORAGEMEM3DMULTISAMPLEEXT
	gpTexSubImage2D                                  C.GPTEXSUBIMAGE2D
	gpTexSubImage3D                                  C.GPTEXSUBIMAGE3D
	gpTexSubImage3DOES                               C.GPTEXSUBIMAGE3DOES
	gpTextureFoveationParametersQCOM                 C.GPTEXTUREFOVEATIONPARAMETERSQCOM
	gpTextureStorage1DEXT                            C.GPTEXTURESTORAGE1DEXT
	gpTextureStorage2DEXT                            C.GPTEXTURESTORAGE2DEXT
	gpTextureStorage3DEXT                            C.GPTEXTURESTORAGE3DEXT
	gpTextureStorageMem1DEXT                         C.GPTEXTURESTORAGEMEM1DEXT
	gpTextureStorageMem2DEXT                         C.GPTEXTURESTORAGEMEM2DEXT
	gpTextureStorageMem2DMultisampleEXT              C.GPTEXTURESTORAGEMEM2DMULTISAMPLEEXT
	gpTextureStorageMem3DEXT                         C.GPTEXTURESTORAGEMEM3DEXT
	gpTextureStorageMem3DMultisampleEXT              C.GPTEXTURESTORAGEMEM3DMULTISAMPLEEXT
	gpTextureViewEXT                                 C.GPTEXTUREVIEWEXT
	gpTextureViewOES                                 C.GPTEXTUREVIEWOES
	gpTransformFeedbackVaryings                      C.GPTRANSFORMFEEDBACKVARYINGS
	gpTransformPathNV                                C.GPTRANSFORMPATHNV
	gpUniform1f                                      C.GPUNIFORM1F
	gpUniform1fv                                     C.GPUNIFORM1FV
	gpUniform1i                                      C.GPUNIFORM1I
	gpUniform1i64NV                                  C.GPUNIFORM1I64NV
	gpUniform1i64vNV                                 C.GPUNIFORM1I64VNV
	gpUniform1iv                                     C.GPUNIFORM1IV
	gpUniform1ui                                     C.GPUNIFORM1UI
	gpUniform1ui64NV                                 C.GPUNIFORM1UI64NV
	gpUniform1ui64vNV                                C.GPUNIFORM1UI64VNV
	gpUniform1uiv                                    C.GPUNIFORM1UIV
	gpUniform2f                                      C.GPUNIFORM2F
	gpUniform2fv                                     C.GPUNIFORM2FV
	gpUniform2i                                      C.GPUNIFORM2I
	gpUniform2i64NV                                  C.GPUNIFORM2I64NV
	gpUniform2i64vNV                                 C.GPUNIFORM2I64VNV
	gpUniform2iv                                     C.GPUNIFORM2IV
	gpUniform2ui                                     C.GPUNIFORM2UI
	gpUniform2ui64NV                                 C.GPUNIFORM2UI64NV
	gpUniform2ui64vNV                                C.GPUNIFORM2UI64VNV
	gpUniform2uiv                                    C.GPUNIFORM2UIV
	gpUniform3f                                      C.GPUNIFORM3F
	gpUniform3fv                                     C.GPUNIFORM3FV
	gpUniform3i                                      C.GPUNIFORM3I
	gpUniform3i64NV                                  C.GPUNIFORM3I64NV
	gpUniform3i64vNV                                 C.GPUNIFORM3I64VNV
	gpUniform3iv                                     C.GPUNIFORM3IV
	gpUniform3ui                                     C.GPUNIFORM3UI
	gpUniform3ui64NV                                 C.GPUNIFORM3UI64NV
	gpUniform3ui64vNV                                C.GPUNIFORM3UI64VNV
	gpUniform3uiv                                    C.GPUNIFORM3UIV
	gpUniform4f                                      C.GPUNIFORM4F
	gpUniform4fv                                     C.GPUNIFORM4FV
	gpUniform4i                                      C.GPUNIFORM4I
	gpUniform4i64NV                                  C.GPUNIFORM4I64NV
	gpUniform4i64vNV                                 C.GPUNIFORM4I64VNV
	gpUniform4iv                                     C.GPUNIFORM4IV
	gpUniform4ui                                     C.GPUNIFORM4UI
	gpUniform4ui64NV                                 C.GPUNIFORM4UI64NV
	gpUniform4ui64vNV                                C.GPUNIFORM4UI64VNV
	gpUniform4uiv                                    C.GPUNIFORM4UIV
	gpUniformBlockBinding                            C.GPUNIFORMBLOCKBINDING
	gpUniformHandleui64IMG                           C.GPUNIFORMHANDLEUI64IMG
	gpUniformHandleui64NV                            C.GPUNIFORMHANDLEUI64NV
	gpUniformHandleui64vIMG                          C.GPUNIFORMHANDLEUI64VIMG
	gpUniformHandleui64vNV                           C.GPUNIFORMHANDLEUI64VNV
	gpUniformMatrix2fv                               C.GPUNIFORMMATRIX2FV
	gpUniformMatrix2x3fv                             C.GPUNIFORMMATRIX2X3FV
	gpUniformMatrix2x3fvNV                           C.GPUNIFORMMATRIX2X3FVNV
	gpUniformMatrix2x4fv                             C.GPUNIFORMMATRIX2X4FV
	gpUniformMatrix2x4fvNV                           C.GPUNIFORMMATRIX2X4FVNV
	gpUniformMatrix3fv                               C.GPUNIFORMMATRIX3FV
	gpUniformMatrix3x2fv                             C.GPUNIFORMMATRIX3X2FV
	gpUniformMatrix3x2fvNV                           C.GPUNIFORMMATRIX3X2FVNV
	gpUniformMatrix3x4fv                             C.GPUNIFORMMATRIX3X4FV
	gpUniformMatrix3x4fvNV                           C.GPUNIFORMMATRIX3X4FVNV
	gpUniformMatrix4fv                               C.GPUNIFORMMATRIX4FV
	gpUniformMatrix4x2fv                             C.GPUNIFORMMATRIX4X2FV
	gpUniformMatrix4x2fvNV                           C.GPUNIFORMMATRIX4X2FVNV
	gpUniformMatrix4x3fv                             C.GPUNIFORMMATRIX4X3FV
	gpUniformMatrix4x3fvNV                           C.GPUNIFORMMATRIX4X3FVNV
	gpUnmapBuffer                                    C.GPUNMAPBUFFER
	gpUnmapBufferOES                                 C.GPUNMAPBUFFEROES
	gpUseProgram                                     C.GPUSEPROGRAM
	gpUseProgramStages                               C.GPUSEPROGRAMSTAGES
	gpUseProgramStagesEXT                            C.GPUSEPROGRAMSTAGESEXT
	gpUseShaderProgramEXT                            C.GPUSESHADERPROGRAMEXT
	gpValidateProgram                                C.GPVALIDATEPROGRAM
	gpValidateProgramPipeline                        C.GPVALIDATEPROGRAMPIPELINE
	gpValidateProgramPipelineEXT                     C.GPVALIDATEPROGRAMPIPELINEEXT
	gpVertexAttrib1f                                 C.GPVERTEXATTRIB1F
	gpVertexAttrib1fv                                C.GPVERTEXATTRIB1FV
	gpVertexAttrib2f                                 C.GPVERTEXATTRIB2F
	gpVertexAttrib2fv                                C.GPVERTEXATTRIB2FV
	gpVertexAttrib3f                                 C.GPVERTEXATTRIB3F
	gpVertexAttrib3fv                                C.GPVERTEXATTRIB3FV
	gpVertexAttrib4f                                 C.GPVERTEXATTRIB4F
	gpVertexAttrib4fv                                C.GPVERTEXATTRIB4FV
	gpVertexAttribBinding                            C.GPVERTEXATTRIBBINDING
	gpVertexAttribDivisor                            C.GPVERTEXATTRIBDIVISOR
	gpVertexAttribDivisorANGLE                       C.GPVERTEXATTRIBDIVISORANGLE
	gpVertexAttribDivisorEXT                         C.GPVERTEXATTRIBDIVISOREXT
	gpVertexAttribDivisorNV                          C.GPVERTEXATTRIBDIVISORNV
	gpVertexAttribFormat                             C.GPVERTEXATTRIBFORMAT
	gpVertexAttribI4i                                C.GPVERTEXATTRIBI4I
	gpVertexAttribI4iv                               C.GPVERTEXATTRIBI4IV
	gpVertexAttribI4ui                               C.GPVERTEXATTRIBI4UI
	gpVertexAttribI4uiv                              C.GPVERTEXATTRIBI4UIV
	gpVertexAttribIFormat                            C.GPVERTEXATTRIBIFORMAT
	gpVertexAttribIPointer                           C.GPVERTEXATTRIBIPOINTER
	gpVertexAttribPointer                            C.GPVERTEXATTRIBPOINTER
	gpVertexBindingDivisor                           C.GPVERTEXBINDINGDIVISOR
	gpViewport                                       C.GPVIEWPORT
	gpViewportArrayvNV                               C.GPVIEWPORTARRAYVNV
	gpViewportArrayvOES                              C.GPVIEWPORTARRAYVOES
	gpViewportIndexedfNV                             C.GPVIEWPORTINDEXEDFNV
	gpViewportIndexedfOES                            C.GPVIEWPORTINDEXEDFOES
	gpViewportIndexedfvNV                            C.GPVIEWPORTINDEXEDFVNV
	gpViewportIndexedfvOES                           C.GPVIEWPORTINDEXEDFVOES
	gpViewportPositionWScaleNV                       C.GPVIEWPORTPOSITIONWSCALENV
	gpViewportSwizzleNV                              C.GPVIEWPORTSWIZZLENV
	gpWaitSemaphoreEXT                               C.GPWAITSEMAPHOREEXT
	gpWaitSync                                       C.GPWAITSYNC
	gpWaitSyncAPPLE                                  C.GPWAITSYNCAPPLE
	gpWaitVkSemaphoreNV                              C.GPWAITVKSEMAPHORENV
	gpWeightPathsNV                                  C.GPWEIGHTPATHSNV
	gpWindowRectanglesEXT                            C.GPWINDOWRECTANGLESEXT
)

// Helper functions
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
func AcquireKeyedMutexWin32EXT(memory uint32, key uint64, timeout uint32) bool {
	ret := C.glowAcquireKeyedMutexWin32EXT(gpAcquireKeyedMutexWin32EXT, (C.GLuint)(memory), (C.GLuint64)(key), (C.GLuint)(timeout))
	return ret == TRUE
}
func ActiveProgramEXT(program uint32) {
	C.glowActiveProgramEXT(gpActiveProgramEXT, (C.GLuint)(program))
}

// set the active program object for a program pipeline object
func ActiveShaderProgram(pipeline uint32, program uint32) {
	C.glowActiveShaderProgram(gpActiveShaderProgram, (C.GLuint)(pipeline), (C.GLuint)(program))
}
func ActiveShaderProgramEXT(pipeline uint32, program uint32) {
	C.glowActiveShaderProgramEXT(gpActiveShaderProgramEXT, (C.GLuint)(pipeline), (C.GLuint)(program))
}

// select active texture unit
func ActiveTexture(texture uint32) {
	C.glowActiveTexture(gpActiveTexture, (C.GLenum)(texture))
}
func AlphaFuncQCOM(xfunc uint32, ref float32) {
	C.glowAlphaFuncQCOM(gpAlphaFuncQCOM, (C.GLenum)(xfunc), (C.GLclampf)(ref))
}
func ApplyFramebufferAttachmentCMAAINTEL() {
	C.glowApplyFramebufferAttachmentCMAAINTEL(gpApplyFramebufferAttachmentCMAAINTEL)
}

// Attaches a shader object to a program object
func AttachShader(program uint32, shader uint32) {
	C.glowAttachShader(gpAttachShader, (C.GLuint)(program), (C.GLuint)(shader))
}
func BeginConditionalRenderNV(id uint32, mode uint32) {
	C.glowBeginConditionalRenderNV(gpBeginConditionalRenderNV, (C.GLuint)(id), (C.GLenum)(mode))
}
func BeginPerfMonitorAMD(monitor uint32) {
	C.glowBeginPerfMonitorAMD(gpBeginPerfMonitorAMD, (C.GLuint)(monitor))
}
func BeginPerfQueryINTEL(queryHandle uint32) {
	C.glowBeginPerfQueryINTEL(gpBeginPerfQueryINTEL, (C.GLuint)(queryHandle))
}

// delimit the boundaries of a query object
func BeginQuery(target uint32, id uint32) {
	C.glowBeginQuery(gpBeginQuery, (C.GLenum)(target), (C.GLuint)(id))
}
func BeginQueryEXT(target uint32, id uint32) {
	C.glowBeginQueryEXT(gpBeginQueryEXT, (C.GLenum)(target), (C.GLuint)(id))
}

// start transform feedback operation
func BeginTransformFeedback(primitiveMode uint32) {
	C.glowBeginTransformFeedback(gpBeginTransformFeedback, (C.GLenum)(primitiveMode))
}

// Associates a generic vertex attribute index with a named attribute variable
func BindAttribLocation(program uint32, index uint32, name *uint8) {
	C.glowBindAttribLocation(gpBindAttribLocation, (C.GLuint)(program), (C.GLuint)(index), (*C.GLchar)(unsafe.Pointer(name)))
}

// bind a named buffer object
func BindBuffer(target uint32, buffer uint32) {
	C.glowBindBuffer(gpBindBuffer, (C.GLenum)(target), (C.GLuint)(buffer))
}

// bind a buffer object to an indexed buffer target
func BindBufferBase(target uint32, index uint32, buffer uint32) {
	C.glowBindBufferBase(gpBindBufferBase, (C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(buffer))
}

// bind a range within a buffer object to an indexed buffer target
func BindBufferRange(target uint32, index uint32, buffer uint32, offset int, size int) {
	C.glowBindBufferRange(gpBindBufferRange, (C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(size))
}
func BindFragDataLocationEXT(program uint32, color uint32, name *uint8) {
	C.glowBindFragDataLocationEXT(gpBindFragDataLocationEXT, (C.GLuint)(program), (C.GLuint)(color), (*C.GLchar)(unsafe.Pointer(name)))
}
func BindFragDataLocationIndexedEXT(program uint32, colorNumber uint32, index uint32, name *uint8) {
	C.glowBindFragDataLocationIndexedEXT(gpBindFragDataLocationIndexedEXT, (C.GLuint)(program), (C.GLuint)(colorNumber), (C.GLuint)(index), (*C.GLchar)(unsafe.Pointer(name)))
}

// bind a framebuffer to a framebuffer target
func BindFramebuffer(target uint32, framebuffer uint32) {
	C.glowBindFramebuffer(gpBindFramebuffer, (C.GLenum)(target), (C.GLuint)(framebuffer))
}

// bind a level of a texture to an image unit
func BindImageTexture(unit uint32, texture uint32, level int32, layered bool, layer int32, access uint32, format uint32) {
	C.glowBindImageTexture(gpBindImageTexture, (C.GLuint)(unit), (C.GLuint)(texture), (C.GLint)(level), (C.GLboolean)(boolToInt(layered)), (C.GLint)(layer), (C.GLenum)(access), (C.GLenum)(format))
}

// bind a program pipeline to the current context
func BindProgramPipeline(pipeline uint32) {
	C.glowBindProgramPipeline(gpBindProgramPipeline, (C.GLuint)(pipeline))
}
func BindProgramPipelineEXT(pipeline uint32) {
	C.glowBindProgramPipelineEXT(gpBindProgramPipelineEXT, (C.GLuint)(pipeline))
}

// bind a renderbuffer to a renderbuffer target
func BindRenderbuffer(target uint32, renderbuffer uint32) {
	C.glowBindRenderbuffer(gpBindRenderbuffer, (C.GLenum)(target), (C.GLuint)(renderbuffer))
}

// bind a named sampler to a texturing target
func BindSampler(unit uint32, sampler uint32) {
	C.glowBindSampler(gpBindSampler, (C.GLuint)(unit), (C.GLuint)(sampler))
}

// bind a named texture to a texturing target
func BindTexture(target uint32, texture uint32) {
	C.glowBindTexture(gpBindTexture, (C.GLenum)(target), (C.GLuint)(texture))
}

// bind a transform feedback object
func BindTransformFeedback(target uint32, id uint32) {
	C.glowBindTransformFeedback(gpBindTransformFeedback, (C.GLenum)(target), (C.GLuint)(id))
}

// bind a vertex array object
func BindVertexArray(array uint32) {
	C.glowBindVertexArray(gpBindVertexArray, (C.GLuint)(array))
}
func BindVertexArrayOES(array uint32) {
	C.glowBindVertexArrayOES(gpBindVertexArrayOES, (C.GLuint)(array))
}

// bind a buffer to a vertex buffer bind point
func BindVertexBuffer(bindingindex uint32, buffer uint32, offset int, stride int32) {
	C.glowBindVertexBuffer(gpBindVertexBuffer, (C.GLuint)(bindingindex), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(stride))
}
func BlendBarrierKHR() {
	C.glowBlendBarrierKHR(gpBlendBarrierKHR)
}
func BlendBarrierNV() {
	C.glowBlendBarrierNV(gpBlendBarrierNV)
}

// set the blend color
func BlendColor(red float32, green float32, blue float32, alpha float32) {
	C.glowBlendColor(gpBlendColor, (C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}

// specify the equation used for both the RGB blend equation and the Alpha blend equation
func BlendEquation(mode uint32) {
	C.glowBlendEquation(gpBlendEquation, (C.GLenum)(mode))
}
func BlendEquationEXT(mode uint32) {
	C.glowBlendEquationEXT(gpBlendEquationEXT, (C.GLenum)(mode))
}

// set the RGB blend equation and the alpha blend equation separately
func BlendEquationSeparate(modeRGB uint32, modeAlpha uint32) {
	C.glowBlendEquationSeparate(gpBlendEquationSeparate, (C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}
func BlendEquationSeparateiEXT(buf uint32, modeRGB uint32, modeAlpha uint32) {
	C.glowBlendEquationSeparateiEXT(gpBlendEquationSeparateiEXT, (C.GLuint)(buf), (C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}
func BlendEquationSeparateiOES(buf uint32, modeRGB uint32, modeAlpha uint32) {
	C.glowBlendEquationSeparateiOES(gpBlendEquationSeparateiOES, (C.GLuint)(buf), (C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}
func BlendEquationiEXT(buf uint32, mode uint32) {
	C.glowBlendEquationiEXT(gpBlendEquationiEXT, (C.GLuint)(buf), (C.GLenum)(mode))
}
func BlendEquationiOES(buf uint32, mode uint32) {
	C.glowBlendEquationiOES(gpBlendEquationiOES, (C.GLuint)(buf), (C.GLenum)(mode))
}

// specify pixel arithmetic
func BlendFunc(sfactor uint32, dfactor uint32) {
	C.glowBlendFunc(gpBlendFunc, (C.GLenum)(sfactor), (C.GLenum)(dfactor))
}

// specify pixel arithmetic for RGB and alpha components separately
func BlendFuncSeparate(sfactorRGB uint32, dfactorRGB uint32, sfactorAlpha uint32, dfactorAlpha uint32) {
	C.glowBlendFuncSeparate(gpBlendFuncSeparate, (C.GLenum)(sfactorRGB), (C.GLenum)(dfactorRGB), (C.GLenum)(sfactorAlpha), (C.GLenum)(dfactorAlpha))
}
func BlendFuncSeparateiEXT(buf uint32, srcRGB uint32, dstRGB uint32, srcAlpha uint32, dstAlpha uint32) {
	C.glowBlendFuncSeparateiEXT(gpBlendFuncSeparateiEXT, (C.GLuint)(buf), (C.GLenum)(srcRGB), (C.GLenum)(dstRGB), (C.GLenum)(srcAlpha), (C.GLenum)(dstAlpha))
}
func BlendFuncSeparateiOES(buf uint32, srcRGB uint32, dstRGB uint32, srcAlpha uint32, dstAlpha uint32) {
	C.glowBlendFuncSeparateiOES(gpBlendFuncSeparateiOES, (C.GLuint)(buf), (C.GLenum)(srcRGB), (C.GLenum)(dstRGB), (C.GLenum)(srcAlpha), (C.GLenum)(dstAlpha))
}
func BlendFunciEXT(buf uint32, src uint32, dst uint32) {
	C.glowBlendFunciEXT(gpBlendFunciEXT, (C.GLuint)(buf), (C.GLenum)(src), (C.GLenum)(dst))
}
func BlendFunciOES(buf uint32, src uint32, dst uint32) {
	C.glowBlendFunciOES(gpBlendFunciOES, (C.GLuint)(buf), (C.GLenum)(src), (C.GLenum)(dst))
}
func BlendParameteriNV(pname uint32, value int32) {
	C.glowBlendParameteriNV(gpBlendParameteriNV, (C.GLenum)(pname), (C.GLint)(value))
}

// copy a block of pixels from one framebuffer object to another
func BlitFramebuffer(srcX0 int32, srcY0 int32, srcX1 int32, srcY1 int32, dstX0 int32, dstY0 int32, dstX1 int32, dstY1 int32, mask uint32, filter uint32) {
	C.glowBlitFramebuffer(gpBlitFramebuffer, (C.GLint)(srcX0), (C.GLint)(srcY0), (C.GLint)(srcX1), (C.GLint)(srcY1), (C.GLint)(dstX0), (C.GLint)(dstY0), (C.GLint)(dstX1), (C.GLint)(dstY1), (C.GLbitfield)(mask), (C.GLenum)(filter))
}
func BlitFramebufferANGLE(srcX0 int32, srcY0 int32, srcX1 int32, srcY1 int32, dstX0 int32, dstY0 int32, dstX1 int32, dstY1 int32, mask uint32, filter uint32) {
	C.glowBlitFramebufferANGLE(gpBlitFramebufferANGLE, (C.GLint)(srcX0), (C.GLint)(srcY0), (C.GLint)(srcX1), (C.GLint)(srcY1), (C.GLint)(dstX0), (C.GLint)(dstY0), (C.GLint)(dstX1), (C.GLint)(dstY1), (C.GLbitfield)(mask), (C.GLenum)(filter))
}
func BlitFramebufferNV(srcX0 int32, srcY0 int32, srcX1 int32, srcY1 int32, dstX0 int32, dstY0 int32, dstX1 int32, dstY1 int32, mask uint32, filter uint32) {
	C.glowBlitFramebufferNV(gpBlitFramebufferNV, (C.GLint)(srcX0), (C.GLint)(srcY0), (C.GLint)(srcX1), (C.GLint)(srcY1), (C.GLint)(dstX0), (C.GLint)(dstY0), (C.GLint)(dstX1), (C.GLint)(dstY1), (C.GLbitfield)(mask), (C.GLenum)(filter))
}

// creates and initializes a buffer object's data     store
func BufferData(target uint32, size int, data unsafe.Pointer, usage uint32) {
	C.glowBufferData(gpBufferData, (C.GLenum)(target), (C.GLsizeiptr)(size), data, (C.GLenum)(usage))
}
func BufferStorageEXT(target uint32, size int, data unsafe.Pointer, flags uint32) {
	C.glowBufferStorageEXT(gpBufferStorageEXT, (C.GLenum)(target), (C.GLsizeiptr)(size), data, (C.GLbitfield)(flags))
}
func BufferStorageExternalEXT(target uint32, offset int, size int, clientBuffer C.GLeglClientBufferEXT, flags uint32) {
	C.glowBufferStorageExternalEXT(gpBufferStorageExternalEXT, (C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), (C.GLeglClientBufferEXT)(clientBuffer), (C.GLbitfield)(flags))
}
func BufferStorageMemEXT(target uint32, size int, memory uint32, offset uint64) {
	C.glowBufferStorageMemEXT(gpBufferStorageMemEXT, (C.GLenum)(target), (C.GLsizeiptr)(size), (C.GLuint)(memory), (C.GLuint64)(offset))
}

// updates a subset of a buffer object's data store
func BufferSubData(target uint32, offset int, size int, data unsafe.Pointer) {
	C.glowBufferSubData(gpBufferSubData, (C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), data)
}

// check the completeness status of a framebuffer
func CheckFramebufferStatus(target uint32) uint32 {
	ret := C.glowCheckFramebufferStatus(gpCheckFramebufferStatus, (C.GLenum)(target))
	return (uint32)(ret)
}

// clear buffers to preset values
func Clear(mask uint32) {
	C.glowClear(gpClear, (C.GLbitfield)(mask))
}
func ClearBufferfi(buffer uint32, drawbuffer int32, depth float32, stencil int32) {
	C.glowClearBufferfi(gpClearBufferfi, (C.GLenum)(buffer), (C.GLint)(drawbuffer), (C.GLfloat)(depth), (C.GLint)(stencil))
}
func ClearBufferfv(buffer uint32, drawbuffer int32, value *float32) {
	C.glowClearBufferfv(gpClearBufferfv, (C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ClearBufferiv(buffer uint32, drawbuffer int32, value *int32) {
	C.glowClearBufferiv(gpClearBufferiv, (C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLint)(unsafe.Pointer(value)))
}
func ClearBufferuiv(buffer uint32, drawbuffer int32, value *uint32) {
	C.glowClearBufferuiv(gpClearBufferuiv, (C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLuint)(unsafe.Pointer(value)))
}

// specify clear values for the color buffers
func ClearColor(red float32, green float32, blue float32, alpha float32) {
	C.glowClearColor(gpClearColor, (C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}

// specify the clear value for the depth buffer
func ClearDepthf(d float32) {
	C.glowClearDepthf(gpClearDepthf, (C.GLfloat)(d))
}
func ClearPixelLocalStorageuiEXT(offset int32, n int32, values *uint32) {
	C.glowClearPixelLocalStorageuiEXT(gpClearPixelLocalStorageuiEXT, (C.GLsizei)(offset), (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(values)))
}

// specify the clear value for the stencil buffer
func ClearStencil(s int32) {
	C.glowClearStencil(gpClearStencil, (C.GLint)(s))
}
func ClearTexImageEXT(texture uint32, level int32, format uint32, xtype uint32, data unsafe.Pointer) {
	C.glowClearTexImageEXT(gpClearTexImageEXT, (C.GLuint)(texture), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(xtype), data)
}
func ClearTexSubImageEXT(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, data unsafe.Pointer) {
	C.glowClearTexSubImageEXT(gpClearTexSubImageEXT, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(xtype), data)
}

// block and wait for a sync object to become signaled
func ClientWaitSync(sync uintptr, flags uint32, timeout uint64) uint32 {
	ret := C.glowClientWaitSync(gpClientWaitSync, (C.GLsync)(sync), (C.GLbitfield)(flags), (C.GLuint64)(timeout))
	return (uint32)(ret)
}
func ClientWaitSyncAPPLE(sync uintptr, flags uint32, timeout uint64) uint32 {
	ret := C.glowClientWaitSyncAPPLE(gpClientWaitSyncAPPLE, (C.GLsync)(sync), (C.GLbitfield)(flags), (C.GLuint64)(timeout))
	return (uint32)(ret)
}
func ClipControlEXT(origin uint32, depth uint32) {
	C.glowClipControlEXT(gpClipControlEXT, (C.GLenum)(origin), (C.GLenum)(depth))
}
func ColorMask(red bool, green bool, blue bool, alpha bool) {
	C.glowColorMask(gpColorMask, (C.GLboolean)(boolToInt(red)), (C.GLboolean)(boolToInt(green)), (C.GLboolean)(boolToInt(blue)), (C.GLboolean)(boolToInt(alpha)))
}
func ColorMaskiEXT(index uint32, r bool, g bool, b bool, a bool) {
	C.glowColorMaskiEXT(gpColorMaskiEXT, (C.GLuint)(index), (C.GLboolean)(boolToInt(r)), (C.GLboolean)(boolToInt(g)), (C.GLboolean)(boolToInt(b)), (C.GLboolean)(boolToInt(a)))
}
func ColorMaskiOES(index uint32, r bool, g bool, b bool, a bool) {
	C.glowColorMaskiOES(gpColorMaskiOES, (C.GLuint)(index), (C.GLboolean)(boolToInt(r)), (C.GLboolean)(boolToInt(g)), (C.GLboolean)(boolToInt(b)), (C.GLboolean)(boolToInt(a)))
}

// Compiles a shader object
func CompileShader(shader uint32) {
	C.glowCompileShader(gpCompileShader, (C.GLuint)(shader))
}

// specify a two-dimensional texture image in a compressed format
func CompressedTexImage2D(target uint32, level int32, internalformat uint32, width int32, height int32, border int32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTexImage2D(gpCompressedTexImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLsizei)(imageSize), data)
}

// specify a three-dimensional texture image in a compressed format
func CompressedTexImage3D(target uint32, level int32, internalformat uint32, width int32, height int32, depth int32, border int32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTexImage3D(gpCompressedTexImage3D, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLsizei)(imageSize), data)
}
func CompressedTexImage3DOES(target uint32, level int32, internalformat uint32, width int32, height int32, depth int32, border int32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTexImage3DOES(gpCompressedTexImage3DOES, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLsizei)(imageSize), data)
}

// specify a two-dimensional texture subimage in a compressed format
func CompressedTexSubImage2D(target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTexSubImage2D(gpCompressedTexSubImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLsizei)(imageSize), data)
}

// specify a three-dimensional texture subimage in a compressed format
func CompressedTexSubImage3D(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTexSubImage3D(gpCompressedTexSubImage3D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLsizei)(imageSize), data)
}
func CompressedTexSubImage3DOES(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTexSubImage3DOES(gpCompressedTexSubImage3DOES, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLsizei)(imageSize), data)
}
func ConservativeRasterParameteriNV(pname uint32, param int32) {
	C.glowConservativeRasterParameteriNV(gpConservativeRasterParameteriNV, (C.GLenum)(pname), (C.GLint)(param))
}

// copy all or part of the data store of a buffer object to the data store of another buffer object
func CopyBufferSubData(readTarget uint32, writeTarget uint32, readOffset int, writeOffset int, size int) {
	C.glowCopyBufferSubData(gpCopyBufferSubData, (C.GLenum)(readTarget), (C.GLenum)(writeTarget), (C.GLintptr)(readOffset), (C.GLintptr)(writeOffset), (C.GLsizeiptr)(size))
}
func CopyBufferSubDataNV(readTarget uint32, writeTarget uint32, readOffset int, writeOffset int, size int) {
	C.glowCopyBufferSubDataNV(gpCopyBufferSubDataNV, (C.GLenum)(readTarget), (C.GLenum)(writeTarget), (C.GLintptr)(readOffset), (C.GLintptr)(writeOffset), (C.GLsizeiptr)(size))
}
func CopyImageSubDataEXT(srcName uint32, srcTarget uint32, srcLevel int32, srcX int32, srcY int32, srcZ int32, dstName uint32, dstTarget uint32, dstLevel int32, dstX int32, dstY int32, dstZ int32, srcWidth int32, srcHeight int32, srcDepth int32) {
	C.glowCopyImageSubDataEXT(gpCopyImageSubDataEXT, (C.GLuint)(srcName), (C.GLenum)(srcTarget), (C.GLint)(srcLevel), (C.GLint)(srcX), (C.GLint)(srcY), (C.GLint)(srcZ), (C.GLuint)(dstName), (C.GLenum)(dstTarget), (C.GLint)(dstLevel), (C.GLint)(dstX), (C.GLint)(dstY), (C.GLint)(dstZ), (C.GLsizei)(srcWidth), (C.GLsizei)(srcHeight), (C.GLsizei)(srcDepth))
}
func CopyImageSubDataOES(srcName uint32, srcTarget uint32, srcLevel int32, srcX int32, srcY int32, srcZ int32, dstName uint32, dstTarget uint32, dstLevel int32, dstX int32, dstY int32, dstZ int32, srcWidth int32, srcHeight int32, srcDepth int32) {
	C.glowCopyImageSubDataOES(gpCopyImageSubDataOES, (C.GLuint)(srcName), (C.GLenum)(srcTarget), (C.GLint)(srcLevel), (C.GLint)(srcX), (C.GLint)(srcY), (C.GLint)(srcZ), (C.GLuint)(dstName), (C.GLenum)(dstTarget), (C.GLint)(dstLevel), (C.GLint)(dstX), (C.GLint)(dstY), (C.GLint)(dstZ), (C.GLsizei)(srcWidth), (C.GLsizei)(srcHeight), (C.GLsizei)(srcDepth))
}
func CopyPathNV(resultPath uint32, srcPath uint32) {
	C.glowCopyPathNV(gpCopyPathNV, (C.GLuint)(resultPath), (C.GLuint)(srcPath))
}

// copy pixels into a 2D texture image
func CopyTexImage2D(target uint32, level int32, internalformat uint32, x int32, y int32, width int32, height int32, border int32) {
	C.glowCopyTexImage2D(gpCopyTexImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border))
}

// copy a two-dimensional texture subimage
func CopyTexSubImage2D(target uint32, level int32, xoffset int32, yoffset int32, x int32, y int32, width int32, height int32) {
	C.glowCopyTexSubImage2D(gpCopyTexSubImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

// copy a three-dimensional texture subimage
func CopyTexSubImage3D(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, x int32, y int32, width int32, height int32) {
	C.glowCopyTexSubImage3D(gpCopyTexSubImage3D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
func CopyTexSubImage3DOES(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, x int32, y int32, width int32, height int32) {
	C.glowCopyTexSubImage3DOES(gpCopyTexSubImage3DOES, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
func CopyTextureLevelsAPPLE(destinationTexture uint32, sourceTexture uint32, sourceBaseLevel int32, sourceLevelCount int32) {
	C.glowCopyTextureLevelsAPPLE(gpCopyTextureLevelsAPPLE, (C.GLuint)(destinationTexture), (C.GLuint)(sourceTexture), (C.GLint)(sourceBaseLevel), (C.GLsizei)(sourceLevelCount))
}
func CoverFillPathInstancedNV(numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, coverMode uint32, transformType uint32, transformValues *float32) {
	C.glowCoverFillPathInstancedNV(gpCoverFillPathInstancedNV, (C.GLsizei)(numPaths), (C.GLenum)(pathNameType), paths, (C.GLuint)(pathBase), (C.GLenum)(coverMode), (C.GLenum)(transformType), (*C.GLfloat)(unsafe.Pointer(transformValues)))
}
func CoverFillPathNV(path uint32, coverMode uint32) {
	C.glowCoverFillPathNV(gpCoverFillPathNV, (C.GLuint)(path), (C.GLenum)(coverMode))
}
func CoverStrokePathInstancedNV(numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, coverMode uint32, transformType uint32, transformValues *float32) {
	C.glowCoverStrokePathInstancedNV(gpCoverStrokePathInstancedNV, (C.GLsizei)(numPaths), (C.GLenum)(pathNameType), paths, (C.GLuint)(pathBase), (C.GLenum)(coverMode), (C.GLenum)(transformType), (*C.GLfloat)(unsafe.Pointer(transformValues)))
}
func CoverStrokePathNV(path uint32, coverMode uint32) {
	C.glowCoverStrokePathNV(gpCoverStrokePathNV, (C.GLuint)(path), (C.GLenum)(coverMode))
}
func CoverageMaskNV(mask bool) {
	C.glowCoverageMaskNV(gpCoverageMaskNV, (C.GLboolean)(boolToInt(mask)))
}
func CoverageModulationNV(components uint32) {
	C.glowCoverageModulationNV(gpCoverageModulationNV, (C.GLenum)(components))
}
func CoverageModulationTableNV(n int32, v *float32) {
	C.glowCoverageModulationTableNV(gpCoverageModulationTableNV, (C.GLsizei)(n), (*C.GLfloat)(unsafe.Pointer(v)))
}
func CoverageOperationNV(operation uint32) {
	C.glowCoverageOperationNV(gpCoverageOperationNV, (C.GLenum)(operation))
}
func CreateMemoryObjectsEXT(n int32, memoryObjects *uint32) {
	C.glowCreateMemoryObjectsEXT(gpCreateMemoryObjectsEXT, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(memoryObjects)))
}
func CreatePerfQueryINTEL(queryId uint32, queryHandle *uint32) {
	C.glowCreatePerfQueryINTEL(gpCreatePerfQueryINTEL, (C.GLuint)(queryId), (*C.GLuint)(unsafe.Pointer(queryHandle)))
}

// Creates a program object
func CreateProgram() uint32 {
	ret := C.glowCreateProgram(gpCreateProgram)
	return (uint32)(ret)
}

// Creates a shader object
func CreateShader(xtype uint32) uint32 {
	ret := C.glowCreateShader(gpCreateShader, (C.GLenum)(xtype))
	return (uint32)(ret)
}
func CreateShaderProgramEXT(xtype uint32, xstring *uint8) uint32 {
	ret := C.glowCreateShaderProgramEXT(gpCreateShaderProgramEXT, (C.GLenum)(xtype), (*C.GLchar)(unsafe.Pointer(xstring)))
	return (uint32)(ret)
}

// create a stand-alone program from an array of null-terminated source code strings
func CreateShaderProgramv(xtype uint32, count int32, strings **uint8) uint32 {
	ret := C.glowCreateShaderProgramv(gpCreateShaderProgramv, (C.GLenum)(xtype), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(strings)))
	return (uint32)(ret)
}
func CreateShaderProgramvEXT(xtype uint32, count int32, strings **uint8) uint32 {
	ret := C.glowCreateShaderProgramvEXT(gpCreateShaderProgramvEXT, (C.GLenum)(xtype), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(strings)))
	return (uint32)(ret)
}

// specify whether front- or back-facing facets can be culled
func CullFace(mode uint32) {
	C.glowCullFace(gpCullFace, (C.GLenum)(mode))
}

// specify a callback to receive debugging messages from the GL
func DebugMessageCallback(callback DebugProc, userParam unsafe.Pointer) {
	userDebugCallback = callback
	C.glowDebugMessageCallback(gpDebugMessageCallback, (C.GLDEBUGPROC)(unsafe.Pointer(&callback)), userParam)
}
func DebugMessageCallbackKHR(callback DebugProc, userParam unsafe.Pointer) {
	userDebugCallback = callback
	C.glowDebugMessageCallbackKHR(gpDebugMessageCallbackKHR, (C.GLDEBUGPROCKHR)(unsafe.Pointer(&callback)), userParam)
}

// control the reporting of debug messages in a debug context
func DebugMessageControl(source uint32, xtype uint32, severity uint32, count int32, ids *uint32, enabled bool) {
	C.glowDebugMessageControl(gpDebugMessageControl, (C.GLenum)(source), (C.GLenum)(xtype), (C.GLenum)(severity), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(ids)), (C.GLboolean)(boolToInt(enabled)))
}
func DebugMessageControlKHR(source uint32, xtype uint32, severity uint32, count int32, ids *uint32, enabled bool) {
	C.glowDebugMessageControlKHR(gpDebugMessageControlKHR, (C.GLenum)(source), (C.GLenum)(xtype), (C.GLenum)(severity), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(ids)), (C.GLboolean)(boolToInt(enabled)))
}

// inject an application-supplied message into the debug message queue
func DebugMessageInsert(source uint32, xtype uint32, id uint32, severity uint32, length int32, buf *uint8) {
	C.glowDebugMessageInsert(gpDebugMessageInsert, (C.GLenum)(source), (C.GLenum)(xtype), (C.GLuint)(id), (C.GLenum)(severity), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(buf)))
}
func DebugMessageInsertKHR(source uint32, xtype uint32, id uint32, severity uint32, length int32, buf *uint8) {
	C.glowDebugMessageInsertKHR(gpDebugMessageInsertKHR, (C.GLenum)(source), (C.GLenum)(xtype), (C.GLuint)(id), (C.GLenum)(severity), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(buf)))
}

// delete named buffer objects
func DeleteBuffers(n int32, buffers *uint32) {
	C.glowDeleteBuffers(gpDeleteBuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(buffers)))
}
func DeleteFencesNV(n int32, fences *uint32) {
	C.glowDeleteFencesNV(gpDeleteFencesNV, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(fences)))
}

// delete framebuffer objects
func DeleteFramebuffers(n int32, framebuffers *uint32) {
	C.glowDeleteFramebuffers(gpDeleteFramebuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(framebuffers)))
}
func DeleteMemoryObjectsEXT(n int32, memoryObjects *uint32) {
	C.glowDeleteMemoryObjectsEXT(gpDeleteMemoryObjectsEXT, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(memoryObjects)))
}
func DeletePathsNV(path uint32, xrange int32) {
	C.glowDeletePathsNV(gpDeletePathsNV, (C.GLuint)(path), (C.GLsizei)(xrange))
}
func DeletePerfMonitorsAMD(n int32, monitors *uint32) {
	C.glowDeletePerfMonitorsAMD(gpDeletePerfMonitorsAMD, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(monitors)))
}
func DeletePerfQueryINTEL(queryHandle uint32) {
	C.glowDeletePerfQueryINTEL(gpDeletePerfQueryINTEL, (C.GLuint)(queryHandle))
}

// Deletes a program object
func DeleteProgram(program uint32) {
	C.glowDeleteProgram(gpDeleteProgram, (C.GLuint)(program))
}

// delete program pipeline objects
func DeleteProgramPipelines(n int32, pipelines *uint32) {
	C.glowDeleteProgramPipelines(gpDeleteProgramPipelines, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(pipelines)))
}
func DeleteProgramPipelinesEXT(n int32, pipelines *uint32) {
	C.glowDeleteProgramPipelinesEXT(gpDeleteProgramPipelinesEXT, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(pipelines)))
}

// delete named query objects
func DeleteQueries(n int32, ids *uint32) {
	C.glowDeleteQueries(gpDeleteQueries, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}
func DeleteQueriesEXT(n int32, ids *uint32) {
	C.glowDeleteQueriesEXT(gpDeleteQueriesEXT, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}

// delete renderbuffer objects
func DeleteRenderbuffers(n int32, renderbuffers *uint32) {
	C.glowDeleteRenderbuffers(gpDeleteRenderbuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(renderbuffers)))
}

// delete named sampler objects
func DeleteSamplers(count int32, samplers *uint32) {
	C.glowDeleteSamplers(gpDeleteSamplers, (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(samplers)))
}
func DeleteSemaphoresEXT(n int32, semaphores *uint32) {
	C.glowDeleteSemaphoresEXT(gpDeleteSemaphoresEXT, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(semaphores)))
}

// Deletes a shader object
func DeleteShader(shader uint32) {
	C.glowDeleteShader(gpDeleteShader, (C.GLuint)(shader))
}

// delete a sync object
func DeleteSync(sync uintptr) {
	C.glowDeleteSync(gpDeleteSync, (C.GLsync)(sync))
}
func DeleteSyncAPPLE(sync uintptr) {
	C.glowDeleteSyncAPPLE(gpDeleteSyncAPPLE, (C.GLsync)(sync))
}

// delete named textures
func DeleteTextures(n int32, textures *uint32) {
	C.glowDeleteTextures(gpDeleteTextures, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(textures)))
}

// delete transform feedback objects
func DeleteTransformFeedbacks(n int32, ids *uint32) {
	C.glowDeleteTransformFeedbacks(gpDeleteTransformFeedbacks, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}

// delete vertex array objects
func DeleteVertexArrays(n int32, arrays *uint32) {
	C.glowDeleteVertexArrays(gpDeleteVertexArrays, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(arrays)))
}
func DeleteVertexArraysOES(n int32, arrays *uint32) {
	C.glowDeleteVertexArraysOES(gpDeleteVertexArraysOES, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(arrays)))
}

// specify the value used for depth buffer comparisons
func DepthFunc(xfunc uint32) {
	C.glowDepthFunc(gpDepthFunc, (C.GLenum)(xfunc))
}

// enable or disable writing into the depth buffer
func DepthMask(flag bool) {
	C.glowDepthMask(gpDepthMask, (C.GLboolean)(boolToInt(flag)))
}
func DepthRangeArrayfvNV(first uint32, count int32, v *float32) {
	C.glowDepthRangeArrayfvNV(gpDepthRangeArrayfvNV, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(v)))
}
func DepthRangeArrayfvOES(first uint32, count int32, v *float32) {
	C.glowDepthRangeArrayfvOES(gpDepthRangeArrayfvOES, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(v)))
}
func DepthRangeIndexedfNV(index uint32, n float32, f float32) {
	C.glowDepthRangeIndexedfNV(gpDepthRangeIndexedfNV, (C.GLuint)(index), (C.GLfloat)(n), (C.GLfloat)(f))
}
func DepthRangeIndexedfOES(index uint32, n float32, f float32) {
	C.glowDepthRangeIndexedfOES(gpDepthRangeIndexedfOES, (C.GLuint)(index), (C.GLfloat)(n), (C.GLfloat)(f))
}

// specify mapping of depth values from normalized device coordinates to window coordinates
func DepthRangef(n float32, f float32) {
	C.glowDepthRangef(gpDepthRangef, (C.GLfloat)(n), (C.GLfloat)(f))
}

// Detaches a shader object from a program object to which it is attached
func DetachShader(program uint32, shader uint32) {
	C.glowDetachShader(gpDetachShader, (C.GLuint)(program), (C.GLuint)(shader))
}
func Disable(cap uint32) {
	C.glowDisable(gpDisable, (C.GLenum)(cap))
}
func DisableDriverControlQCOM(driverControl uint32) {
	C.glowDisableDriverControlQCOM(gpDisableDriverControlQCOM, (C.GLuint)(driverControl))
}

// Enable or disable a generic vertex attribute     array
func DisableVertexAttribArray(index uint32) {
	C.glowDisableVertexAttribArray(gpDisableVertexAttribArray, (C.GLuint)(index))
}
func DisableiEXT(target uint32, index uint32) {
	C.glowDisableiEXT(gpDisableiEXT, (C.GLenum)(target), (C.GLuint)(index))
}
func DisableiNV(target uint32, index uint32) {
	C.glowDisableiNV(gpDisableiNV, (C.GLenum)(target), (C.GLuint)(index))
}
func DisableiOES(target uint32, index uint32) {
	C.glowDisableiOES(gpDisableiOES, (C.GLenum)(target), (C.GLuint)(index))
}
func DiscardFramebufferEXT(target uint32, numAttachments int32, attachments *uint32) {
	C.glowDiscardFramebufferEXT(gpDiscardFramebufferEXT, (C.GLenum)(target), (C.GLsizei)(numAttachments), (*C.GLenum)(unsafe.Pointer(attachments)))
}

// launch one or more compute work groups
func DispatchCompute(num_groups_x uint32, num_groups_y uint32, num_groups_z uint32) {
	C.glowDispatchCompute(gpDispatchCompute, (C.GLuint)(num_groups_x), (C.GLuint)(num_groups_y), (C.GLuint)(num_groups_z))
}

// launch one or more compute work groups using parameters stored in a buffer
func DispatchComputeIndirect(indirect int) {
	C.glowDispatchComputeIndirect(gpDispatchComputeIndirect, (C.GLintptr)(indirect))
}

// render primitives from array data
func DrawArrays(mode uint32, first int32, count int32) {
	C.glowDrawArrays(gpDrawArrays, (C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count))
}

// render primitives from array data, taking parameters from memory
func DrawArraysIndirect(mode uint32, indirect unsafe.Pointer) {
	C.glowDrawArraysIndirect(gpDrawArraysIndirect, (C.GLenum)(mode), indirect)
}

// draw multiple instances of a range of elements
func DrawArraysInstanced(mode uint32, first int32, count int32, instancecount int32) {
	C.glowDrawArraysInstanced(gpDrawArraysInstanced, (C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count), (C.GLsizei)(instancecount))
}
func DrawArraysInstancedANGLE(mode uint32, first int32, count int32, primcount int32) {
	C.glowDrawArraysInstancedANGLE(gpDrawArraysInstancedANGLE, (C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count), (C.GLsizei)(primcount))
}
func DrawArraysInstancedBaseInstanceEXT(mode uint32, first int32, count int32, instancecount int32, baseinstance uint32) {
	C.glowDrawArraysInstancedBaseInstanceEXT(gpDrawArraysInstancedBaseInstanceEXT, (C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count), (C.GLsizei)(instancecount), (C.GLuint)(baseinstance))
}
func DrawArraysInstancedEXT(mode uint32, start int32, count int32, primcount int32) {
	C.glowDrawArraysInstancedEXT(gpDrawArraysInstancedEXT, (C.GLenum)(mode), (C.GLint)(start), (C.GLsizei)(count), (C.GLsizei)(primcount))
}
func DrawArraysInstancedNV(mode uint32, first int32, count int32, primcount int32) {
	C.glowDrawArraysInstancedNV(gpDrawArraysInstancedNV, (C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count), (C.GLsizei)(primcount))
}

// Specifies a list of color buffers to be drawn     into
func DrawBuffers(n int32, bufs *uint32) {
	C.glowDrawBuffers(gpDrawBuffers, (C.GLsizei)(n), (*C.GLenum)(unsafe.Pointer(bufs)))
}
func DrawBuffersEXT(n int32, bufs *uint32) {
	C.glowDrawBuffersEXT(gpDrawBuffersEXT, (C.GLsizei)(n), (*C.GLenum)(unsafe.Pointer(bufs)))
}
func DrawBuffersIndexedEXT(n int32, location *uint32, indices *int32) {
	C.glowDrawBuffersIndexedEXT(gpDrawBuffersIndexedEXT, (C.GLint)(n), (*C.GLenum)(unsafe.Pointer(location)), (*C.GLint)(unsafe.Pointer(indices)))
}
func DrawBuffersNV(n int32, bufs *uint32) {
	C.glowDrawBuffersNV(gpDrawBuffersNV, (C.GLsizei)(n), (*C.GLenum)(unsafe.Pointer(bufs)))
}

// render primitives from array data
func DrawElements(mode uint32, count int32, xtype uint32, indices unsafe.Pointer) {
	C.glowDrawElements(gpDrawElements, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices)
}
func DrawElementsBaseVertexEXT(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, basevertex int32) {
	C.glowDrawElementsBaseVertexEXT(gpDrawElementsBaseVertexEXT, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLint)(basevertex))
}
func DrawElementsBaseVertexOES(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, basevertex int32) {
	C.glowDrawElementsBaseVertexOES(gpDrawElementsBaseVertexOES, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLint)(basevertex))
}

// render indexed primitives from array data, taking parameters from memory
func DrawElementsIndirect(mode uint32, xtype uint32, indirect unsafe.Pointer) {
	C.glowDrawElementsIndirect(gpDrawElementsIndirect, (C.GLenum)(mode), (C.GLenum)(xtype), indirect)
}

// draw multiple instances of a set of elements
func DrawElementsInstanced(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, instancecount int32) {
	C.glowDrawElementsInstanced(gpDrawElementsInstanced, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLsizei)(instancecount))
}
func DrawElementsInstancedANGLE(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, primcount int32) {
	C.glowDrawElementsInstancedANGLE(gpDrawElementsInstancedANGLE, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLsizei)(primcount))
}
func DrawElementsInstancedBaseInstanceEXT(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, instancecount int32, baseinstance uint32) {
	C.glowDrawElementsInstancedBaseInstanceEXT(gpDrawElementsInstancedBaseInstanceEXT, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLsizei)(instancecount), (C.GLuint)(baseinstance))
}
func DrawElementsInstancedBaseVertexBaseInstanceEXT(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, instancecount int32, basevertex int32, baseinstance uint32) {
	C.glowDrawElementsInstancedBaseVertexBaseInstanceEXT(gpDrawElementsInstancedBaseVertexBaseInstanceEXT, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLsizei)(instancecount), (C.GLint)(basevertex), (C.GLuint)(baseinstance))
}
func DrawElementsInstancedBaseVertexEXT(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, instancecount int32, basevertex int32) {
	C.glowDrawElementsInstancedBaseVertexEXT(gpDrawElementsInstancedBaseVertexEXT, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLsizei)(instancecount), (C.GLint)(basevertex))
}
func DrawElementsInstancedBaseVertexOES(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, instancecount int32, basevertex int32) {
	C.glowDrawElementsInstancedBaseVertexOES(gpDrawElementsInstancedBaseVertexOES, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLsizei)(instancecount), (C.GLint)(basevertex))
}
func DrawElementsInstancedEXT(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, primcount int32) {
	C.glowDrawElementsInstancedEXT(gpDrawElementsInstancedEXT, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLsizei)(primcount))
}
func DrawElementsInstancedNV(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, primcount int32) {
	C.glowDrawElementsInstancedNV(gpDrawElementsInstancedNV, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLsizei)(primcount))
}

// render primitives from array data
func DrawRangeElements(mode uint32, start uint32, end uint32, count int32, xtype uint32, indices unsafe.Pointer) {
	C.glowDrawRangeElements(gpDrawRangeElements, (C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(xtype), indices)
}
func DrawRangeElementsBaseVertexEXT(mode uint32, start uint32, end uint32, count int32, xtype uint32, indices unsafe.Pointer, basevertex int32) {
	C.glowDrawRangeElementsBaseVertexEXT(gpDrawRangeElementsBaseVertexEXT, (C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLint)(basevertex))
}
func DrawRangeElementsBaseVertexOES(mode uint32, start uint32, end uint32, count int32, xtype uint32, indices unsafe.Pointer, basevertex int32) {
	C.glowDrawRangeElementsBaseVertexOES(gpDrawRangeElementsBaseVertexOES, (C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLint)(basevertex))
}
func DrawTransformFeedbackEXT(mode uint32, id uint32) {
	C.glowDrawTransformFeedbackEXT(gpDrawTransformFeedbackEXT, (C.GLenum)(mode), (C.GLuint)(id))
}
func DrawTransformFeedbackInstancedEXT(mode uint32, id uint32, instancecount int32) {
	C.glowDrawTransformFeedbackInstancedEXT(gpDrawTransformFeedbackInstancedEXT, (C.GLenum)(mode), (C.GLuint)(id), (C.GLsizei)(instancecount))
}
func DrawVkImageNV(vkImage uint64, sampler uint32, x0 float32, y0 float32, x1 float32, y1 float32, z float32, s0 float32, t0 float32, s1 float32, t1 float32) {
	C.glowDrawVkImageNV(gpDrawVkImageNV, (C.GLuint64)(vkImage), (C.GLuint)(sampler), (C.GLfloat)(x0), (C.GLfloat)(y0), (C.GLfloat)(x1), (C.GLfloat)(y1), (C.GLfloat)(z), (C.GLfloat)(s0), (C.GLfloat)(t0), (C.GLfloat)(s1), (C.GLfloat)(t1))
}
func EGLImageTargetRenderbufferStorageOES(target uint32, image C.GLeglImageOES) {
	C.glowEGLImageTargetRenderbufferStorageOES(gpEGLImageTargetRenderbufferStorageOES, (C.GLenum)(target), (C.GLeglImageOES)(image))
}
func EGLImageTargetTexStorageEXT(target uint32, image C.GLeglImageOES, attrib_list *int32) {
	C.glowEGLImageTargetTexStorageEXT(gpEGLImageTargetTexStorageEXT, (C.GLenum)(target), (C.GLeglImageOES)(image), (*C.GLint)(unsafe.Pointer(attrib_list)))
}
func EGLImageTargetTexture2DOES(target uint32, image C.GLeglImageOES) {
	C.glowEGLImageTargetTexture2DOES(gpEGLImageTargetTexture2DOES, (C.GLenum)(target), (C.GLeglImageOES)(image))
}
func EGLImageTargetTextureStorageEXT(texture uint32, image C.GLeglImageOES, attrib_list *int32) {
	C.glowEGLImageTargetTextureStorageEXT(gpEGLImageTargetTextureStorageEXT, (C.GLuint)(texture), (C.GLeglImageOES)(image), (*C.GLint)(unsafe.Pointer(attrib_list)))
}

// enable or disable server-side GL capabilities
func Enable(cap uint32) {
	C.glowEnable(gpEnable, (C.GLenum)(cap))
}
func EnableDriverControlQCOM(driverControl uint32) {
	C.glowEnableDriverControlQCOM(gpEnableDriverControlQCOM, (C.GLuint)(driverControl))
}

// Enable or disable a generic vertex attribute     array
func EnableVertexAttribArray(index uint32) {
	C.glowEnableVertexAttribArray(gpEnableVertexAttribArray, (C.GLuint)(index))
}
func EnableiEXT(target uint32, index uint32) {
	C.glowEnableiEXT(gpEnableiEXT, (C.GLenum)(target), (C.GLuint)(index))
}
func EnableiNV(target uint32, index uint32) {
	C.glowEnableiNV(gpEnableiNV, (C.GLenum)(target), (C.GLuint)(index))
}
func EnableiOES(target uint32, index uint32) {
	C.glowEnableiOES(gpEnableiOES, (C.GLenum)(target), (C.GLuint)(index))
}
func EndConditionalRenderNV() {
	C.glowEndConditionalRenderNV(gpEndConditionalRenderNV)
}
func EndPerfMonitorAMD(monitor uint32) {
	C.glowEndPerfMonitorAMD(gpEndPerfMonitorAMD, (C.GLuint)(monitor))
}
func EndPerfQueryINTEL(queryHandle uint32) {
	C.glowEndPerfQueryINTEL(gpEndPerfQueryINTEL, (C.GLuint)(queryHandle))
}
func EndQuery(target uint32) {
	C.glowEndQuery(gpEndQuery, (C.GLenum)(target))
}
func EndQueryEXT(target uint32) {
	C.glowEndQueryEXT(gpEndQueryEXT, (C.GLenum)(target))
}
func EndTilingQCOM(preserveMask uint32) {
	C.glowEndTilingQCOM(gpEndTilingQCOM, (C.GLbitfield)(preserveMask))
}
func EndTransformFeedback() {
	C.glowEndTransformFeedback(gpEndTransformFeedback)
}
func ExtGetBufferPointervQCOM(target uint32, params *unsafe.Pointer) {
	C.glowExtGetBufferPointervQCOM(gpExtGetBufferPointervQCOM, (C.GLenum)(target), params)
}
func ExtGetBuffersQCOM(buffers *uint32, maxBuffers int32, numBuffers *int32) {
	C.glowExtGetBuffersQCOM(gpExtGetBuffersQCOM, (*C.GLuint)(unsafe.Pointer(buffers)), (C.GLint)(maxBuffers), (*C.GLint)(unsafe.Pointer(numBuffers)))
}
func ExtGetFramebuffersQCOM(framebuffers *uint32, maxFramebuffers int32, numFramebuffers *int32) {
	C.glowExtGetFramebuffersQCOM(gpExtGetFramebuffersQCOM, (*C.GLuint)(unsafe.Pointer(framebuffers)), (C.GLint)(maxFramebuffers), (*C.GLint)(unsafe.Pointer(numFramebuffers)))
}
func ExtGetProgramBinarySourceQCOM(program uint32, shadertype uint32, source *uint8, length *int32) {
	C.glowExtGetProgramBinarySourceQCOM(gpExtGetProgramBinarySourceQCOM, (C.GLuint)(program), (C.GLenum)(shadertype), (*C.GLchar)(unsafe.Pointer(source)), (*C.GLint)(unsafe.Pointer(length)))
}
func ExtGetProgramsQCOM(programs *uint32, maxPrograms int32, numPrograms *int32) {
	C.glowExtGetProgramsQCOM(gpExtGetProgramsQCOM, (*C.GLuint)(unsafe.Pointer(programs)), (C.GLint)(maxPrograms), (*C.GLint)(unsafe.Pointer(numPrograms)))
}
func ExtGetRenderbuffersQCOM(renderbuffers *uint32, maxRenderbuffers int32, numRenderbuffers *int32) {
	C.glowExtGetRenderbuffersQCOM(gpExtGetRenderbuffersQCOM, (*C.GLuint)(unsafe.Pointer(renderbuffers)), (C.GLint)(maxRenderbuffers), (*C.GLint)(unsafe.Pointer(numRenderbuffers)))
}
func ExtGetShadersQCOM(shaders *uint32, maxShaders int32, numShaders *int32) {
	C.glowExtGetShadersQCOM(gpExtGetShadersQCOM, (*C.GLuint)(unsafe.Pointer(shaders)), (C.GLint)(maxShaders), (*C.GLint)(unsafe.Pointer(numShaders)))
}
func ExtGetTexLevelParameterivQCOM(texture uint32, face uint32, level int32, pname uint32, params *int32) {
	C.glowExtGetTexLevelParameterivQCOM(gpExtGetTexLevelParameterivQCOM, (C.GLuint)(texture), (C.GLenum)(face), (C.GLint)(level), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func ExtGetTexSubImageQCOM(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, texels unsafe.Pointer) {
	C.glowExtGetTexSubImageQCOM(gpExtGetTexSubImageQCOM, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(xtype), texels)
}
func ExtGetTexturesQCOM(textures *uint32, maxTextures int32, numTextures *int32) {
	C.glowExtGetTexturesQCOM(gpExtGetTexturesQCOM, (*C.GLuint)(unsafe.Pointer(textures)), (C.GLint)(maxTextures), (*C.GLint)(unsafe.Pointer(numTextures)))
}
func ExtIsProgramBinaryQCOM(program uint32) bool {
	ret := C.glowExtIsProgramBinaryQCOM(gpExtIsProgramBinaryQCOM, (C.GLuint)(program))
	return ret == TRUE
}
func ExtTexObjectStateOverrideiQCOM(target uint32, pname uint32, param int32) {
	C.glowExtTexObjectStateOverrideiQCOM(gpExtTexObjectStateOverrideiQCOM, (C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}

// create a new sync object and insert it into the GL command stream
func FenceSync(condition uint32, flags uint32) uintptr {
	ret := C.glowFenceSync(gpFenceSync, (C.GLenum)(condition), (C.GLbitfield)(flags))
	return (uintptr)(ret)
}
func FenceSyncAPPLE(condition uint32, flags uint32) uintptr {
	ret := C.glowFenceSyncAPPLE(gpFenceSyncAPPLE, (C.GLenum)(condition), (C.GLbitfield)(flags))
	return (uintptr)(ret)
}

// block until all GL execution is complete
func Finish() {
	C.glowFinish(gpFinish)
}
func FinishFenceNV(fence uint32) {
	C.glowFinishFenceNV(gpFinishFenceNV, (C.GLuint)(fence))
}

// force execution of GL commands in finite time
func Flush() {
	C.glowFlush(gpFlush)
}

// indicate modifications to a range of a mapped buffer
func FlushMappedBufferRange(target uint32, offset int, length int) {
	C.glowFlushMappedBufferRange(gpFlushMappedBufferRange, (C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(length))
}
func FlushMappedBufferRangeEXT(target uint32, offset int, length int) {
	C.glowFlushMappedBufferRangeEXT(gpFlushMappedBufferRangeEXT, (C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(length))
}
func FragmentCoverageColorNV(color uint32) {
	C.glowFragmentCoverageColorNV(gpFragmentCoverageColorNV, (C.GLuint)(color))
}
func FramebufferFetchBarrierEXT() {
	C.glowFramebufferFetchBarrierEXT(gpFramebufferFetchBarrierEXT)
}
func FramebufferFetchBarrierQCOM() {
	C.glowFramebufferFetchBarrierQCOM(gpFramebufferFetchBarrierQCOM)
}
func FramebufferFoveationConfigQCOM(framebuffer uint32, numLayers uint32, focalPointsPerLayer uint32, requestedFeatures uint32, providedFeatures *uint32) {
	C.glowFramebufferFoveationConfigQCOM(gpFramebufferFoveationConfigQCOM, (C.GLuint)(framebuffer), (C.GLuint)(numLayers), (C.GLuint)(focalPointsPerLayer), (C.GLuint)(requestedFeatures), (*C.GLuint)(unsafe.Pointer(providedFeatures)))
}
func FramebufferFoveationParametersQCOM(framebuffer uint32, layer uint32, focalPoint uint32, focalX float32, focalY float32, gainX float32, gainY float32, foveaArea float32) {
	C.glowFramebufferFoveationParametersQCOM(gpFramebufferFoveationParametersQCOM, (C.GLuint)(framebuffer), (C.GLuint)(layer), (C.GLuint)(focalPoint), (C.GLfloat)(focalX), (C.GLfloat)(focalY), (C.GLfloat)(gainX), (C.GLfloat)(gainY), (C.GLfloat)(foveaArea))
}

// set a named parameter of a framebuffer object
func FramebufferParameteri(target uint32, pname uint32, param int32) {
	C.glowFramebufferParameteri(gpFramebufferParameteri, (C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}
func FramebufferPixelLocalStorageSizeEXT(target uint32, size int32) {
	C.glowFramebufferPixelLocalStorageSizeEXT(gpFramebufferPixelLocalStorageSizeEXT, (C.GLuint)(target), (C.GLsizei)(size))
}

// attach a renderbuffer as a logical buffer of a framebuffer object
func FramebufferRenderbuffer(target uint32, attachment uint32, renderbuffertarget uint32, renderbuffer uint32) {
	C.glowFramebufferRenderbuffer(gpFramebufferRenderbuffer, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(renderbuffertarget), (C.GLuint)(renderbuffer))
}
func FramebufferSampleLocationsfvNV(target uint32, start uint32, count int32, v *float32) {
	C.glowFramebufferSampleLocationsfvNV(gpFramebufferSampleLocationsfvNV, (C.GLenum)(target), (C.GLuint)(start), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(v)))
}

// attach a level of a texture object as a logical buffer to the currently bound framebuffer object
func FramebufferTexture2D(target uint32, attachment uint32, textarget uint32, texture uint32, level int32) {
	C.glowFramebufferTexture2D(gpFramebufferTexture2D, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level))
}
func FramebufferTexture2DDownsampleIMG(target uint32, attachment uint32, textarget uint32, texture uint32, level int32, xscale int32, yscale int32) {
	C.glowFramebufferTexture2DDownsampleIMG(gpFramebufferTexture2DDownsampleIMG, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xscale), (C.GLint)(yscale))
}
func FramebufferTexture2DMultisampleEXT(target uint32, attachment uint32, textarget uint32, texture uint32, level int32, samples int32) {
	C.glowFramebufferTexture2DMultisampleEXT(gpFramebufferTexture2DMultisampleEXT, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level), (C.GLsizei)(samples))
}
func FramebufferTexture2DMultisampleIMG(target uint32, attachment uint32, textarget uint32, texture uint32, level int32, samples int32) {
	C.glowFramebufferTexture2DMultisampleIMG(gpFramebufferTexture2DMultisampleIMG, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level), (C.GLsizei)(samples))
}
func FramebufferTexture3DOES(target uint32, attachment uint32, textarget uint32, texture uint32, level int32, zoffset int32) {
	C.glowFramebufferTexture3DOES(gpFramebufferTexture3DOES, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(zoffset))
}
func FramebufferTextureEXT(target uint32, attachment uint32, texture uint32, level int32) {
	C.glowFramebufferTextureEXT(gpFramebufferTextureEXT, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level))
}

// attach a single layer of a texture object as a logical buffer of a framebuffer object
func FramebufferTextureLayer(target uint32, attachment uint32, texture uint32, level int32, layer int32) {
	C.glowFramebufferTextureLayer(gpFramebufferTextureLayer, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(layer))
}
func FramebufferTextureLayerDownsampleIMG(target uint32, attachment uint32, texture uint32, level int32, layer int32, xscale int32, yscale int32) {
	C.glowFramebufferTextureLayerDownsampleIMG(gpFramebufferTextureLayerDownsampleIMG, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(layer), (C.GLint)(xscale), (C.GLint)(yscale))
}
func FramebufferTextureMultisampleMultiviewOVR(target uint32, attachment uint32, texture uint32, level int32, samples int32, baseViewIndex int32, numViews int32) {
	C.glowFramebufferTextureMultisampleMultiviewOVR(gpFramebufferTextureMultisampleMultiviewOVR, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level), (C.GLsizei)(samples), (C.GLint)(baseViewIndex), (C.GLsizei)(numViews))
}
func FramebufferTextureMultiviewOVR(target uint32, attachment uint32, texture uint32, level int32, baseViewIndex int32, numViews int32) {
	C.glowFramebufferTextureMultiviewOVR(gpFramebufferTextureMultiviewOVR, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(baseViewIndex), (C.GLsizei)(numViews))
}
func FramebufferTextureOES(target uint32, attachment uint32, texture uint32, level int32) {
	C.glowFramebufferTextureOES(gpFramebufferTextureOES, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level))
}

// define front- and back-facing polygons
func FrontFace(mode uint32) {
	C.glowFrontFace(gpFrontFace, (C.GLenum)(mode))
}

// generate buffer object names
func GenBuffers(n int32, buffers *uint32) {
	C.glowGenBuffers(gpGenBuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(buffers)))
}
func GenFencesNV(n int32, fences *uint32) {
	C.glowGenFencesNV(gpGenFencesNV, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(fences)))
}

// generate framebuffer object names
func GenFramebuffers(n int32, framebuffers *uint32) {
	C.glowGenFramebuffers(gpGenFramebuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(framebuffers)))
}
func GenPathsNV(xrange int32) uint32 {
	ret := C.glowGenPathsNV(gpGenPathsNV, (C.GLsizei)(xrange))
	return (uint32)(ret)
}
func GenPerfMonitorsAMD(n int32, monitors *uint32) {
	C.glowGenPerfMonitorsAMD(gpGenPerfMonitorsAMD, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(monitors)))
}

// reserve program pipeline object names
func GenProgramPipelines(n int32, pipelines *uint32) {
	C.glowGenProgramPipelines(gpGenProgramPipelines, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(pipelines)))
}
func GenProgramPipelinesEXT(n int32, pipelines *uint32) {
	C.glowGenProgramPipelinesEXT(gpGenProgramPipelinesEXT, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(pipelines)))
}

// generate query object names
func GenQueries(n int32, ids *uint32) {
	C.glowGenQueries(gpGenQueries, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}
func GenQueriesEXT(n int32, ids *uint32) {
	C.glowGenQueriesEXT(gpGenQueriesEXT, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}

// generate renderbuffer object names
func GenRenderbuffers(n int32, renderbuffers *uint32) {
	C.glowGenRenderbuffers(gpGenRenderbuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(renderbuffers)))
}

// generate sampler object names
func GenSamplers(count int32, samplers *uint32) {
	C.glowGenSamplers(gpGenSamplers, (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(samplers)))
}
func GenSemaphoresEXT(n int32, semaphores *uint32) {
	C.glowGenSemaphoresEXT(gpGenSemaphoresEXT, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(semaphores)))
}

// generate texture names
func GenTextures(n int32, textures *uint32) {
	C.glowGenTextures(gpGenTextures, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(textures)))
}

// reserve transform feedback object names
func GenTransformFeedbacks(n int32, ids *uint32) {
	C.glowGenTransformFeedbacks(gpGenTransformFeedbacks, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}

// generate vertex array object names
func GenVertexArrays(n int32, arrays *uint32) {
	C.glowGenVertexArrays(gpGenVertexArrays, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(arrays)))
}
func GenVertexArraysOES(n int32, arrays *uint32) {
	C.glowGenVertexArraysOES(gpGenVertexArraysOES, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(arrays)))
}

// generate mipmaps for a specified texture object
func GenerateMipmap(target uint32) {
	C.glowGenerateMipmap(gpGenerateMipmap, (C.GLenum)(target))
}

// Returns information about an active attribute variable for the specified program object
func GetActiveAttrib(program uint32, index uint32, bufSize int32, length *int32, size *int32, xtype *uint32, name *uint8) {
	C.glowGetActiveAttrib(gpGetActiveAttrib, (C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(size)), (*C.GLenum)(unsafe.Pointer(xtype)), (*C.GLchar)(unsafe.Pointer(name)))
}

// Returns information about an active uniform variable for the specified program object
func GetActiveUniform(program uint32, index uint32, bufSize int32, length *int32, size *int32, xtype *uint32, name *uint8) {
	C.glowGetActiveUniform(gpGetActiveUniform, (C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(size)), (*C.GLenum)(unsafe.Pointer(xtype)), (*C.GLchar)(unsafe.Pointer(name)))
}

// retrieve the name of an active uniform block
func GetActiveUniformBlockName(program uint32, uniformBlockIndex uint32, bufSize int32, length *int32, uniformBlockName *uint8) {
	C.glowGetActiveUniformBlockName(gpGetActiveUniformBlockName, (C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(uniformBlockName)))
}

// query information about an active uniform block
func GetActiveUniformBlockiv(program uint32, uniformBlockIndex uint32, pname uint32, params *int32) {
	C.glowGetActiveUniformBlockiv(gpGetActiveUniformBlockiv, (C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// Returns information about several active uniform variables for the specified program object
func GetActiveUniformsiv(program uint32, uniformCount int32, uniformIndices *uint32, pname uint32, params *int32) {
	C.glowGetActiveUniformsiv(gpGetActiveUniformsiv, (C.GLuint)(program), (C.GLsizei)(uniformCount), (*C.GLuint)(unsafe.Pointer(uniformIndices)), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// Returns the handles of the shader objects attached to a program object
func GetAttachedShaders(program uint32, maxCount int32, count *int32, shaders *uint32) {
	C.glowGetAttachedShaders(gpGetAttachedShaders, (C.GLuint)(program), (C.GLsizei)(maxCount), (*C.GLsizei)(unsafe.Pointer(count)), (*C.GLuint)(unsafe.Pointer(shaders)))
}

// Returns the location of an attribute variable
func GetAttribLocation(program uint32, name *uint8) int32 {
	ret := C.glowGetAttribLocation(gpGetAttribLocation, (C.GLuint)(program), (*C.GLchar)(unsafe.Pointer(name)))
	return (int32)(ret)
}
func GetBooleani_v(target uint32, index uint32, data *bool) {
	C.glowGetBooleani_v(gpGetBooleani_v, (C.GLenum)(target), (C.GLuint)(index), (*C.GLboolean)(unsafe.Pointer(data)))
}
func GetBooleanv(pname uint32, data *bool) {
	C.glowGetBooleanv(gpGetBooleanv, (C.GLenum)(pname), (*C.GLboolean)(unsafe.Pointer(data)))
}

// return parameters of a buffer object
func GetBufferParameteri64v(target uint32, pname uint32, params *int64) {
	C.glowGetBufferParameteri64v(gpGetBufferParameteri64v, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint64)(unsafe.Pointer(params)))
}

// return parameters of a buffer object
func GetBufferParameteriv(target uint32, pname uint32, params *int32) {
	C.glowGetBufferParameteriv(gpGetBufferParameteriv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// return the pointer to a mapped buffer object's data store
func GetBufferPointerv(target uint32, pname uint32, params *unsafe.Pointer) {
	C.glowGetBufferPointerv(gpGetBufferPointerv, (C.GLenum)(target), (C.GLenum)(pname), params)
}
func GetBufferPointervOES(target uint32, pname uint32, params *unsafe.Pointer) {
	C.glowGetBufferPointervOES(gpGetBufferPointervOES, (C.GLenum)(target), (C.GLenum)(pname), params)
}
func GetCoverageModulationTableNV(bufsize int32, v *float32) {
	C.glowGetCoverageModulationTableNV(gpGetCoverageModulationTableNV, (C.GLsizei)(bufsize), (*C.GLfloat)(unsafe.Pointer(v)))
}

// retrieve messages from the debug message log
func GetDebugMessageLog(count uint32, bufSize int32, sources *uint32, types *uint32, ids *uint32, severities *uint32, lengths *int32, messageLog *uint8) uint32 {
	ret := C.glowGetDebugMessageLog(gpGetDebugMessageLog, (C.GLuint)(count), (C.GLsizei)(bufSize), (*C.GLenum)(unsafe.Pointer(sources)), (*C.GLenum)(unsafe.Pointer(types)), (*C.GLuint)(unsafe.Pointer(ids)), (*C.GLenum)(unsafe.Pointer(severities)), (*C.GLsizei)(unsafe.Pointer(lengths)), (*C.GLchar)(unsafe.Pointer(messageLog)))
	return (uint32)(ret)
}
func GetDebugMessageLogKHR(count uint32, bufSize int32, sources *uint32, types *uint32, ids *uint32, severities *uint32, lengths *int32, messageLog *uint8) uint32 {
	ret := C.glowGetDebugMessageLogKHR(gpGetDebugMessageLogKHR, (C.GLuint)(count), (C.GLsizei)(bufSize), (*C.GLenum)(unsafe.Pointer(sources)), (*C.GLenum)(unsafe.Pointer(types)), (*C.GLuint)(unsafe.Pointer(ids)), (*C.GLenum)(unsafe.Pointer(severities)), (*C.GLsizei)(unsafe.Pointer(lengths)), (*C.GLchar)(unsafe.Pointer(messageLog)))
	return (uint32)(ret)
}
func GetDriverControlStringQCOM(driverControl uint32, bufSize int32, length *int32, driverControlString *uint8) {
	C.glowGetDriverControlStringQCOM(gpGetDriverControlStringQCOM, (C.GLuint)(driverControl), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(driverControlString)))
}
func GetDriverControlsQCOM(num *int32, size int32, driverControls *uint32) {
	C.glowGetDriverControlsQCOM(gpGetDriverControlsQCOM, (*C.GLint)(unsafe.Pointer(num)), (C.GLsizei)(size), (*C.GLuint)(unsafe.Pointer(driverControls)))
}

// return error information
func GetError() uint32 {
	ret := C.glowGetError(gpGetError)
	return (uint32)(ret)
}
func GetFenceivNV(fence uint32, pname uint32, params *int32) {
	C.glowGetFenceivNV(gpGetFenceivNV, (C.GLuint)(fence), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetFirstPerfQueryIdINTEL(queryId *uint32) {
	C.glowGetFirstPerfQueryIdINTEL(gpGetFirstPerfQueryIdINTEL, (*C.GLuint)(unsafe.Pointer(queryId)))
}
func GetFloati_vNV(target uint32, index uint32, data *float32) {
	C.glowGetFloati_vNV(gpGetFloati_vNV, (C.GLenum)(target), (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(data)))
}
func GetFloati_vOES(target uint32, index uint32, data *float32) {
	C.glowGetFloati_vOES(gpGetFloati_vOES, (C.GLenum)(target), (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(data)))
}
func GetFloatv(pname uint32, data *float32) {
	C.glowGetFloatv(gpGetFloatv, (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(data)))
}
func GetFragDataIndexEXT(program uint32, name *uint8) int32 {
	ret := C.glowGetFragDataIndexEXT(gpGetFragDataIndexEXT, (C.GLuint)(program), (*C.GLchar)(unsafe.Pointer(name)))
	return (int32)(ret)
}

// query the bindings of color numbers to user-defined varying out variables
func GetFragDataLocation(program uint32, name *uint8) int32 {
	ret := C.glowGetFragDataLocation(gpGetFragDataLocation, (C.GLuint)(program), (*C.GLchar)(unsafe.Pointer(name)))
	return (int32)(ret)
}

// retrieve information about attachments of a bound framebuffer object
func GetFramebufferAttachmentParameteriv(target uint32, attachment uint32, pname uint32, params *int32) {
	C.glowGetFramebufferAttachmentParameteriv(gpGetFramebufferAttachmentParameteriv, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// retrieve a named parameter from a framebuffer
func GetFramebufferParameteriv(target uint32, pname uint32, params *int32) {
	C.glowGetFramebufferParameteriv(gpGetFramebufferParameteriv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetFramebufferPixelLocalStorageSizeEXT(target uint32) int32 {
	ret := C.glowGetFramebufferPixelLocalStorageSizeEXT(gpGetFramebufferPixelLocalStorageSizeEXT, (C.GLuint)(target))
	return (int32)(ret)
}

// check if the rendering context has not been lost due to software or hardware issues
func GetGraphicsResetStatus() uint32 {
	ret := C.glowGetGraphicsResetStatus(gpGetGraphicsResetStatus)
	return (uint32)(ret)
}
func GetGraphicsResetStatusEXT() uint32 {
	ret := C.glowGetGraphicsResetStatusEXT(gpGetGraphicsResetStatusEXT)
	return (uint32)(ret)
}
func GetGraphicsResetStatusKHR() uint32 {
	ret := C.glowGetGraphicsResetStatusKHR(gpGetGraphicsResetStatusKHR)
	return (uint32)(ret)
}
func GetImageHandleNV(texture uint32, level int32, layered bool, layer int32, format uint32) uint64 {
	ret := C.glowGetImageHandleNV(gpGetImageHandleNV, (C.GLuint)(texture), (C.GLint)(level), (C.GLboolean)(boolToInt(layered)), (C.GLint)(layer), (C.GLenum)(format))
	return (uint64)(ret)
}
func GetInteger64i_v(target uint32, index uint32, data *int64) {
	C.glowGetInteger64i_v(gpGetInteger64i_v, (C.GLenum)(target), (C.GLuint)(index), (*C.GLint64)(unsafe.Pointer(data)))
}
func GetInteger64v(pname uint32, data *int64) {
	C.glowGetInteger64v(gpGetInteger64v, (C.GLenum)(pname), (*C.GLint64)(unsafe.Pointer(data)))
}
func GetInteger64vAPPLE(pname uint32, params *int64) {
	C.glowGetInteger64vAPPLE(gpGetInteger64vAPPLE, (C.GLenum)(pname), (*C.GLint64)(unsafe.Pointer(params)))
}
func GetIntegeri_v(target uint32, index uint32, data *int32) {
	C.glowGetIntegeri_v(gpGetIntegeri_v, (C.GLenum)(target), (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(data)))
}
func GetIntegeri_vEXT(target uint32, index uint32, data *int32) {
	C.glowGetIntegeri_vEXT(gpGetIntegeri_vEXT, (C.GLenum)(target), (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(data)))
}
func GetIntegerv(pname uint32, data *int32) {
	C.glowGetIntegerv(gpGetIntegerv, (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(data)))
}
func GetInternalformatSampleivNV(target uint32, internalformat uint32, samples int32, pname uint32, bufSize int32, params *int32) {
	C.glowGetInternalformatSampleivNV(gpGetInternalformatSampleivNV, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLsizei)(samples), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLint)(unsafe.Pointer(params)))
}

// retrieve information about implementation-dependent support for internal formats
func GetInternalformativ(target uint32, internalformat uint32, pname uint32, bufSize int32, params *int32) {
	C.glowGetInternalformativ(gpGetInternalformativ, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLint)(unsafe.Pointer(params)))
}
func GetMemoryObjectParameterivEXT(memoryObject uint32, pname uint32, params *int32) {
	C.glowGetMemoryObjectParameterivEXT(gpGetMemoryObjectParameterivEXT, (C.GLuint)(memoryObject), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// retrieve the location of a sample
func GetMultisamplefv(pname uint32, index uint32, val *float32) {
	C.glowGetMultisamplefv(gpGetMultisamplefv, (C.GLenum)(pname), (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(val)))
}
func GetNextPerfQueryIdINTEL(queryId uint32, nextQueryId *uint32) {
	C.glowGetNextPerfQueryIdINTEL(gpGetNextPerfQueryIdINTEL, (C.GLuint)(queryId), (*C.GLuint)(unsafe.Pointer(nextQueryId)))
}

// retrieve the label of a named object identified within a namespace
func GetObjectLabel(identifier uint32, name uint32, bufSize int32, length *int32, label *uint8) {
	C.glowGetObjectLabel(gpGetObjectLabel, (C.GLenum)(identifier), (C.GLuint)(name), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(label)))
}
func GetObjectLabelEXT(xtype uint32, object uint32, bufSize int32, length *int32, label *uint8) {
	C.glowGetObjectLabelEXT(gpGetObjectLabelEXT, (C.GLenum)(xtype), (C.GLuint)(object), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(label)))
}
func GetObjectLabelKHR(identifier uint32, name uint32, bufSize int32, length *int32, label *uint8) {
	C.glowGetObjectLabelKHR(gpGetObjectLabelKHR, (C.GLenum)(identifier), (C.GLuint)(name), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(label)))
}

// retrieve the label of a sync object identified by a pointer
func GetObjectPtrLabel(ptr unsafe.Pointer, bufSize int32, length *int32, label *uint8) {
	C.glowGetObjectPtrLabel(gpGetObjectPtrLabel, ptr, (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(label)))
}
func GetObjectPtrLabelKHR(ptr unsafe.Pointer, bufSize int32, length *int32, label *uint8) {
	C.glowGetObjectPtrLabelKHR(gpGetObjectPtrLabelKHR, ptr, (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(label)))
}
func GetPathCommandsNV(path uint32, commands *uint8) {
	C.glowGetPathCommandsNV(gpGetPathCommandsNV, (C.GLuint)(path), (*C.GLubyte)(unsafe.Pointer(commands)))
}
func GetPathCoordsNV(path uint32, coords *float32) {
	C.glowGetPathCoordsNV(gpGetPathCoordsNV, (C.GLuint)(path), (*C.GLfloat)(unsafe.Pointer(coords)))
}
func GetPathDashArrayNV(path uint32, dashArray *float32) {
	C.glowGetPathDashArrayNV(gpGetPathDashArrayNV, (C.GLuint)(path), (*C.GLfloat)(unsafe.Pointer(dashArray)))
}
func GetPathLengthNV(path uint32, startSegment int32, numSegments int32) float32 {
	ret := C.glowGetPathLengthNV(gpGetPathLengthNV, (C.GLuint)(path), (C.GLsizei)(startSegment), (C.GLsizei)(numSegments))
	return (float32)(ret)
}
func GetPathMetricRangeNV(metricQueryMask uint32, firstPathName uint32, numPaths int32, stride int32, metrics *float32) {
	C.glowGetPathMetricRangeNV(gpGetPathMetricRangeNV, (C.GLbitfield)(metricQueryMask), (C.GLuint)(firstPathName), (C.GLsizei)(numPaths), (C.GLsizei)(stride), (*C.GLfloat)(unsafe.Pointer(metrics)))
}
func GetPathMetricsNV(metricQueryMask uint32, numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, stride int32, metrics *float32) {
	C.glowGetPathMetricsNV(gpGetPathMetricsNV, (C.GLbitfield)(metricQueryMask), (C.GLsizei)(numPaths), (C.GLenum)(pathNameType), paths, (C.GLuint)(pathBase), (C.GLsizei)(stride), (*C.GLfloat)(unsafe.Pointer(metrics)))
}
func GetPathParameterfvNV(path uint32, pname uint32, value *float32) {
	C.glowGetPathParameterfvNV(gpGetPathParameterfvNV, (C.GLuint)(path), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(value)))
}
func GetPathParameterivNV(path uint32, pname uint32, value *int32) {
	C.glowGetPathParameterivNV(gpGetPathParameterivNV, (C.GLuint)(path), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(value)))
}
func GetPathSpacingNV(pathListMode uint32, numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, advanceScale float32, kerningScale float32, transformType uint32, returnedSpacing *float32) {
	C.glowGetPathSpacingNV(gpGetPathSpacingNV, (C.GLenum)(pathListMode), (C.GLsizei)(numPaths), (C.GLenum)(pathNameType), paths, (C.GLuint)(pathBase), (C.GLfloat)(advanceScale), (C.GLfloat)(kerningScale), (C.GLenum)(transformType), (*C.GLfloat)(unsafe.Pointer(returnedSpacing)))
}
func GetPerfCounterInfoINTEL(queryId uint32, counterId uint32, counterNameLength uint32, counterName *uint8, counterDescLength uint32, counterDesc *uint8, counterOffset *uint32, counterDataSize *uint32, counterTypeEnum *uint32, counterDataTypeEnum *uint32, rawCounterMaxValue *uint64) {
	C.glowGetPerfCounterInfoINTEL(gpGetPerfCounterInfoINTEL, (C.GLuint)(queryId), (C.GLuint)(counterId), (C.GLuint)(counterNameLength), (*C.GLchar)(unsafe.Pointer(counterName)), (C.GLuint)(counterDescLength), (*C.GLchar)(unsafe.Pointer(counterDesc)), (*C.GLuint)(unsafe.Pointer(counterOffset)), (*C.GLuint)(unsafe.Pointer(counterDataSize)), (*C.GLuint)(unsafe.Pointer(counterTypeEnum)), (*C.GLuint)(unsafe.Pointer(counterDataTypeEnum)), (*C.GLuint64)(unsafe.Pointer(rawCounterMaxValue)))
}
func GetPerfMonitorCounterDataAMD(monitor uint32, pname uint32, dataSize int32, data *uint32, bytesWritten *int32) {
	C.glowGetPerfMonitorCounterDataAMD(gpGetPerfMonitorCounterDataAMD, (C.GLuint)(monitor), (C.GLenum)(pname), (C.GLsizei)(dataSize), (*C.GLuint)(unsafe.Pointer(data)), (*C.GLint)(unsafe.Pointer(bytesWritten)))
}
func GetPerfMonitorCounterInfoAMD(group uint32, counter uint32, pname uint32, data unsafe.Pointer) {
	C.glowGetPerfMonitorCounterInfoAMD(gpGetPerfMonitorCounterInfoAMD, (C.GLuint)(group), (C.GLuint)(counter), (C.GLenum)(pname), data)
}
func GetPerfMonitorCounterStringAMD(group uint32, counter uint32, bufSize int32, length *int32, counterString *uint8) {
	C.glowGetPerfMonitorCounterStringAMD(gpGetPerfMonitorCounterStringAMD, (C.GLuint)(group), (C.GLuint)(counter), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(counterString)))
}
func GetPerfMonitorCountersAMD(group uint32, numCounters *int32, maxActiveCounters *int32, counterSize int32, counters *uint32) {
	C.glowGetPerfMonitorCountersAMD(gpGetPerfMonitorCountersAMD, (C.GLuint)(group), (*C.GLint)(unsafe.Pointer(numCounters)), (*C.GLint)(unsafe.Pointer(maxActiveCounters)), (C.GLsizei)(counterSize), (*C.GLuint)(unsafe.Pointer(counters)))
}
func GetPerfMonitorGroupStringAMD(group uint32, bufSize int32, length *int32, groupString *uint8) {
	C.glowGetPerfMonitorGroupStringAMD(gpGetPerfMonitorGroupStringAMD, (C.GLuint)(group), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(groupString)))
}
func GetPerfMonitorGroupsAMD(numGroups *int32, groupsSize int32, groups *uint32) {
	C.glowGetPerfMonitorGroupsAMD(gpGetPerfMonitorGroupsAMD, (*C.GLint)(unsafe.Pointer(numGroups)), (C.GLsizei)(groupsSize), (*C.GLuint)(unsafe.Pointer(groups)))
}
func GetPerfQueryDataINTEL(queryHandle uint32, flags uint32, dataSize int32, data unsafe.Pointer, bytesWritten *uint32) {
	C.glowGetPerfQueryDataINTEL(gpGetPerfQueryDataINTEL, (C.GLuint)(queryHandle), (C.GLuint)(flags), (C.GLsizei)(dataSize), data, (*C.GLuint)(unsafe.Pointer(bytesWritten)))
}
func GetPerfQueryIdByNameINTEL(queryName *uint8, queryId *uint32) {
	C.glowGetPerfQueryIdByNameINTEL(gpGetPerfQueryIdByNameINTEL, (*C.GLchar)(unsafe.Pointer(queryName)), (*C.GLuint)(unsafe.Pointer(queryId)))
}
func GetPerfQueryInfoINTEL(queryId uint32, queryNameLength uint32, queryName *uint8, dataSize *uint32, noCounters *uint32, noInstances *uint32, capsMask *uint32) {
	C.glowGetPerfQueryInfoINTEL(gpGetPerfQueryInfoINTEL, (C.GLuint)(queryId), (C.GLuint)(queryNameLength), (*C.GLchar)(unsafe.Pointer(queryName)), (*C.GLuint)(unsafe.Pointer(dataSize)), (*C.GLuint)(unsafe.Pointer(noCounters)), (*C.GLuint)(unsafe.Pointer(noInstances)), (*C.GLuint)(unsafe.Pointer(capsMask)))
}

// return the address of the specified pointer
func GetPointerv(pname uint32, params *unsafe.Pointer) {
	C.glowGetPointerv(gpGetPointerv, (C.GLenum)(pname), params)
}
func GetPointervKHR(pname uint32, params *unsafe.Pointer) {
	C.glowGetPointervKHR(gpGetPointervKHR, (C.GLenum)(pname), params)
}

// return a binary representation of a program object's compiled and linked executable source
func GetProgramBinary(program uint32, bufSize int32, length *int32, binaryFormat *uint32, binary unsafe.Pointer) {
	C.glowGetProgramBinary(gpGetProgramBinary, (C.GLuint)(program), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLenum)(unsafe.Pointer(binaryFormat)), binary)
}
func GetProgramBinaryOES(program uint32, bufSize int32, length *int32, binaryFormat *uint32, binary unsafe.Pointer) {
	C.glowGetProgramBinaryOES(gpGetProgramBinaryOES, (C.GLuint)(program), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLenum)(unsafe.Pointer(binaryFormat)), binary)
}

// Returns the information log for a program object
func GetProgramInfoLog(program uint32, bufSize int32, length *int32, infoLog *uint8) {
	C.glowGetProgramInfoLog(gpGetProgramInfoLog, (C.GLuint)(program), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
}
func GetProgramInterfaceiv(program uint32, programInterface uint32, pname uint32, params *int32) {
	C.glowGetProgramInterfaceiv(gpGetProgramInterfaceiv, (C.GLuint)(program), (C.GLenum)(programInterface), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// retrieve the info log string from a program pipeline object
func GetProgramPipelineInfoLog(pipeline uint32, bufSize int32, length *int32, infoLog *uint8) {
	C.glowGetProgramPipelineInfoLog(gpGetProgramPipelineInfoLog, (C.GLuint)(pipeline), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
}
func GetProgramPipelineInfoLogEXT(pipeline uint32, bufSize int32, length *int32, infoLog *uint8) {
	C.glowGetProgramPipelineInfoLogEXT(gpGetProgramPipelineInfoLogEXT, (C.GLuint)(pipeline), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
}
func GetProgramPipelineiv(pipeline uint32, pname uint32, params *int32) {
	C.glowGetProgramPipelineiv(gpGetProgramPipelineiv, (C.GLuint)(pipeline), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetProgramPipelineivEXT(pipeline uint32, pname uint32, params *int32) {
	C.glowGetProgramPipelineivEXT(gpGetProgramPipelineivEXT, (C.GLuint)(pipeline), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// query the index of a named resource within a program
func GetProgramResourceIndex(program uint32, programInterface uint32, name *uint8) uint32 {
	ret := C.glowGetProgramResourceIndex(gpGetProgramResourceIndex, (C.GLuint)(program), (C.GLenum)(programInterface), (*C.GLchar)(unsafe.Pointer(name)))
	return (uint32)(ret)
}

// query the location of a named resource within a program
func GetProgramResourceLocation(program uint32, programInterface uint32, name *uint8) int32 {
	ret := C.glowGetProgramResourceLocation(gpGetProgramResourceLocation, (C.GLuint)(program), (C.GLenum)(programInterface), (*C.GLchar)(unsafe.Pointer(name)))
	return (int32)(ret)
}
func GetProgramResourceLocationIndexEXT(program uint32, programInterface uint32, name *uint8) int32 {
	ret := C.glowGetProgramResourceLocationIndexEXT(gpGetProgramResourceLocationIndexEXT, (C.GLuint)(program), (C.GLenum)(programInterface), (*C.GLchar)(unsafe.Pointer(name)))
	return (int32)(ret)
}

// query the name of an indexed resource within a program
func GetProgramResourceName(program uint32, programInterface uint32, index uint32, bufSize int32, length *int32, name *uint8) {
	C.glowGetProgramResourceName(gpGetProgramResourceName, (C.GLuint)(program), (C.GLenum)(programInterface), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(name)))
}
func GetProgramResourcefvNV(program uint32, programInterface uint32, index uint32, propCount int32, props *uint32, bufSize int32, length *int32, params *float32) {
	C.glowGetProgramResourcefvNV(gpGetProgramResourcefvNV, (C.GLuint)(program), (C.GLenum)(programInterface), (C.GLuint)(index), (C.GLsizei)(propCount), (*C.GLenum)(unsafe.Pointer(props)), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetProgramResourceiv(program uint32, programInterface uint32, index uint32, propCount int32, props *uint32, bufSize int32, length *int32, params *int32) {
	C.glowGetProgramResourceiv(gpGetProgramResourceiv, (C.GLuint)(program), (C.GLenum)(programInterface), (C.GLuint)(index), (C.GLsizei)(propCount), (*C.GLenum)(unsafe.Pointer(props)), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(params)))
}

// Returns a parameter from a program object
func GetProgramiv(program uint32, pname uint32, params *int32) {
	C.glowGetProgramiv(gpGetProgramiv, (C.GLuint)(program), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetQueryObjecti64vEXT(id uint32, pname uint32, params *int64) {
	C.glowGetQueryObjecti64vEXT(gpGetQueryObjecti64vEXT, (C.GLuint)(id), (C.GLenum)(pname), (*C.GLint64)(unsafe.Pointer(params)))
}
func GetQueryObjectivEXT(id uint32, pname uint32, params *int32) {
	C.glowGetQueryObjectivEXT(gpGetQueryObjectivEXT, (C.GLuint)(id), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetQueryObjectui64vEXT(id uint32, pname uint32, params *uint64) {
	C.glowGetQueryObjectui64vEXT(gpGetQueryObjectui64vEXT, (C.GLuint)(id), (C.GLenum)(pname), (*C.GLuint64)(unsafe.Pointer(params)))
}

// return parameters of a query object
func GetQueryObjectuiv(id uint32, pname uint32, params *uint32) {
	C.glowGetQueryObjectuiv(gpGetQueryObjectuiv, (C.GLuint)(id), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetQueryObjectuivEXT(id uint32, pname uint32, params *uint32) {
	C.glowGetQueryObjectuivEXT(gpGetQueryObjectuivEXT, (C.GLuint)(id), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}

// return parameters of a query object target
func GetQueryiv(target uint32, pname uint32, params *int32) {
	C.glowGetQueryiv(gpGetQueryiv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetQueryivEXT(target uint32, pname uint32, params *int32) {
	C.glowGetQueryivEXT(gpGetQueryivEXT, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// retrieve information about a bound renderbuffer object
func GetRenderbufferParameteriv(target uint32, pname uint32, params *int32) {
	C.glowGetRenderbufferParameteriv(gpGetRenderbufferParameteriv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetSamplerParameterIivEXT(sampler uint32, pname uint32, params *int32) {
	C.glowGetSamplerParameterIivEXT(gpGetSamplerParameterIivEXT, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetSamplerParameterIivOES(sampler uint32, pname uint32, params *int32) {
	C.glowGetSamplerParameterIivOES(gpGetSamplerParameterIivOES, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetSamplerParameterIuivEXT(sampler uint32, pname uint32, params *uint32) {
	C.glowGetSamplerParameterIuivEXT(gpGetSamplerParameterIuivEXT, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetSamplerParameterIuivOES(sampler uint32, pname uint32, params *uint32) {
	C.glowGetSamplerParameterIuivOES(gpGetSamplerParameterIuivOES, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetSamplerParameterfv(sampler uint32, pname uint32, params *float32) {
	C.glowGetSamplerParameterfv(gpGetSamplerParameterfv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetSamplerParameteriv(sampler uint32, pname uint32, params *int32) {
	C.glowGetSamplerParameteriv(gpGetSamplerParameteriv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetSemaphoreParameterui64vEXT(semaphore uint32, pname uint32, params *uint64) {
	C.glowGetSemaphoreParameterui64vEXT(gpGetSemaphoreParameterui64vEXT, (C.GLuint)(semaphore), (C.GLenum)(pname), (*C.GLuint64)(unsafe.Pointer(params)))
}

// Returns the information log for a shader object
func GetShaderInfoLog(shader uint32, bufSize int32, length *int32, infoLog *uint8) {
	C.glowGetShaderInfoLog(gpGetShaderInfoLog, (C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
}

// retrieve the range and precision for numeric formats supported by the shader compiler
func GetShaderPrecisionFormat(shadertype uint32, precisiontype uint32, xrange *int32, precision *int32) {
	C.glowGetShaderPrecisionFormat(gpGetShaderPrecisionFormat, (C.GLenum)(shadertype), (C.GLenum)(precisiontype), (*C.GLint)(unsafe.Pointer(xrange)), (*C.GLint)(unsafe.Pointer(precision)))
}

// Returns the source code string from a shader object
func GetShaderSource(shader uint32, bufSize int32, length *int32, source *uint8) {
	C.glowGetShaderSource(gpGetShaderSource, (C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(source)))
}

// Returns a parameter from a shader object
func GetShaderiv(shader uint32, pname uint32, params *int32) {
	C.glowGetShaderiv(gpGetShaderiv, (C.GLuint)(shader), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// return a string describing the current GL connection
func GetString(name uint32) *uint8 {
	ret := C.glowGetString(gpGetString, (C.GLenum)(name))
	return (*uint8)(ret)
}
func GetStringi(name uint32, index uint32) *uint8 {
	ret := C.glowGetStringi(gpGetStringi, (C.GLenum)(name), (C.GLuint)(index))
	return (*uint8)(ret)
}

// query the properties of a sync object
func GetSynciv(sync uintptr, pname uint32, bufSize int32, length *int32, values *int32) {
	C.glowGetSynciv(gpGetSynciv, (C.GLsync)(sync), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(values)))
}
func GetSyncivAPPLE(sync uintptr, pname uint32, bufSize int32, length *int32, values *int32) {
	C.glowGetSyncivAPPLE(gpGetSyncivAPPLE, (C.GLsync)(sync), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(values)))
}
func GetTexLevelParameterfv(target uint32, level int32, pname uint32, params *float32) {
	C.glowGetTexLevelParameterfv(gpGetTexLevelParameterfv, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetTexLevelParameteriv(target uint32, level int32, pname uint32, params *int32) {
	C.glowGetTexLevelParameteriv(gpGetTexLevelParameteriv, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTexParameterIivEXT(target uint32, pname uint32, params *int32) {
	C.glowGetTexParameterIivEXT(gpGetTexParameterIivEXT, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTexParameterIivOES(target uint32, pname uint32, params *int32) {
	C.glowGetTexParameterIivOES(gpGetTexParameterIivOES, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTexParameterIuivEXT(target uint32, pname uint32, params *uint32) {
	C.glowGetTexParameterIuivEXT(gpGetTexParameterIuivEXT, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetTexParameterIuivOES(target uint32, pname uint32, params *uint32) {
	C.glowGetTexParameterIuivOES(gpGetTexParameterIuivOES, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetTexParameterfv(target uint32, pname uint32, params *float32) {
	C.glowGetTexParameterfv(gpGetTexParameterfv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetTexParameteriv(target uint32, pname uint32, params *int32) {
	C.glowGetTexParameteriv(gpGetTexParameteriv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTextureHandleIMG(texture uint32) uint64 {
	ret := C.glowGetTextureHandleIMG(gpGetTextureHandleIMG, (C.GLuint)(texture))
	return (uint64)(ret)
}
func GetTextureHandleNV(texture uint32) uint64 {
	ret := C.glowGetTextureHandleNV(gpGetTextureHandleNV, (C.GLuint)(texture))
	return (uint64)(ret)
}
func GetTextureSamplerHandleIMG(texture uint32, sampler uint32) uint64 {
	ret := C.glowGetTextureSamplerHandleIMG(gpGetTextureSamplerHandleIMG, (C.GLuint)(texture), (C.GLuint)(sampler))
	return (uint64)(ret)
}
func GetTextureSamplerHandleNV(texture uint32, sampler uint32) uint64 {
	ret := C.glowGetTextureSamplerHandleNV(gpGetTextureSamplerHandleNV, (C.GLuint)(texture), (C.GLuint)(sampler))
	return (uint64)(ret)
}

// retrieve information about varying variables selected for transform feedback
func GetTransformFeedbackVarying(program uint32, index uint32, bufSize int32, length *int32, size *int32, xtype *uint32, name *uint8) {
	C.glowGetTransformFeedbackVarying(gpGetTransformFeedbackVarying, (C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLsizei)(unsafe.Pointer(size)), (*C.GLenum)(unsafe.Pointer(xtype)), (*C.GLchar)(unsafe.Pointer(name)))
}
func GetTranslatedShaderSourceANGLE(shader uint32, bufsize int32, length *int32, source *uint8) {
	C.glowGetTranslatedShaderSourceANGLE(gpGetTranslatedShaderSourceANGLE, (C.GLuint)(shader), (C.GLsizei)(bufsize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(source)))
}

// retrieve the index of a named uniform block
func GetUniformBlockIndex(program uint32, uniformBlockName *uint8) uint32 {
	ret := C.glowGetUniformBlockIndex(gpGetUniformBlockIndex, (C.GLuint)(program), (*C.GLchar)(unsafe.Pointer(uniformBlockName)))
	return (uint32)(ret)
}

// retrieve the index of a named uniform block
func GetUniformIndices(program uint32, uniformCount int32, uniformNames **uint8, uniformIndices *uint32) {
	C.glowGetUniformIndices(gpGetUniformIndices, (C.GLuint)(program), (C.GLsizei)(uniformCount), (**C.GLchar)(unsafe.Pointer(uniformNames)), (*C.GLuint)(unsafe.Pointer(uniformIndices)))
}

// Returns the location of a uniform variable
func GetUniformLocation(program uint32, name *uint8) int32 {
	ret := C.glowGetUniformLocation(gpGetUniformLocation, (C.GLuint)(program), (*C.GLchar)(unsafe.Pointer(name)))
	return (int32)(ret)
}

// Returns the value of a uniform variable
func GetUniformfv(program uint32, location int32, params *float32) {
	C.glowGetUniformfv(gpGetUniformfv, (C.GLuint)(program), (C.GLint)(location), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetUniformi64vNV(program uint32, location int32, params *int64) {
	C.glowGetUniformi64vNV(gpGetUniformi64vNV, (C.GLuint)(program), (C.GLint)(location), (*C.GLint64EXT)(unsafe.Pointer(params)))
}

// Returns the value of a uniform variable
func GetUniformiv(program uint32, location int32, params *int32) {
	C.glowGetUniformiv(gpGetUniformiv, (C.GLuint)(program), (C.GLint)(location), (*C.GLint)(unsafe.Pointer(params)))
}
func GetUniformuiv(program uint32, location int32, params *uint32) {
	C.glowGetUniformuiv(gpGetUniformuiv, (C.GLuint)(program), (C.GLint)(location), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetUnsignedBytei_vEXT(target uint32, index uint32, data *uint8) {
	C.glowGetUnsignedBytei_vEXT(gpGetUnsignedBytei_vEXT, (C.GLenum)(target), (C.GLuint)(index), (*C.GLubyte)(unsafe.Pointer(data)))
}
func GetUnsignedBytevEXT(pname uint32, data *uint8) {
	C.glowGetUnsignedBytevEXT(gpGetUnsignedBytevEXT, (C.GLenum)(pname), (*C.GLubyte)(unsafe.Pointer(data)))
}

// Return a generic vertex attribute parameter
func GetVertexAttribIiv(index uint32, pname uint32, params *int32) {
	C.glowGetVertexAttribIiv(gpGetVertexAttribIiv, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// Return a generic vertex attribute parameter
func GetVertexAttribIuiv(index uint32, pname uint32, params *uint32) {
	C.glowGetVertexAttribIuiv(gpGetVertexAttribIuiv, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}

// return the address of the specified generic vertex attribute pointer
func GetVertexAttribPointerv(index uint32, pname uint32, pointer *unsafe.Pointer) {
	C.glowGetVertexAttribPointerv(gpGetVertexAttribPointerv, (C.GLuint)(index), (C.GLenum)(pname), pointer)
}

// Return a generic vertex attribute parameter
func GetVertexAttribfv(index uint32, pname uint32, params *float32) {
	C.glowGetVertexAttribfv(gpGetVertexAttribfv, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}

// Return a generic vertex attribute parameter
func GetVertexAttribiv(index uint32, pname uint32, params *int32) {
	C.glowGetVertexAttribiv(gpGetVertexAttribiv, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetVkProcAddrNV(name *uint8) C.GLVULKANPROCNV {
	ret := C.glowGetVkProcAddrNV(gpGetVkProcAddrNV, (*C.GLchar)(unsafe.Pointer(name)))
	return (C.GLVULKANPROCNV)(ret)
}
func GetnUniformfv(program uint32, location int32, bufSize int32, params *float32) {
	C.glowGetnUniformfv(gpGetnUniformfv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetnUniformfvEXT(program uint32, location int32, bufSize int32, params *float32) {
	C.glowGetnUniformfvEXT(gpGetnUniformfvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetnUniformfvKHR(program uint32, location int32, bufSize int32, params *float32) {
	C.glowGetnUniformfvKHR(gpGetnUniformfvKHR, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetnUniformiv(program uint32, location int32, bufSize int32, params *int32) {
	C.glowGetnUniformiv(gpGetnUniformiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLint)(unsafe.Pointer(params)))
}
func GetnUniformivEXT(program uint32, location int32, bufSize int32, params *int32) {
	C.glowGetnUniformivEXT(gpGetnUniformivEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLint)(unsafe.Pointer(params)))
}
func GetnUniformivKHR(program uint32, location int32, bufSize int32, params *int32) {
	C.glowGetnUniformivKHR(gpGetnUniformivKHR, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLint)(unsafe.Pointer(params)))
}
func GetnUniformuiv(program uint32, location int32, bufSize int32, params *uint32) {
	C.glowGetnUniformuiv(gpGetnUniformuiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetnUniformuivKHR(program uint32, location int32, bufSize int32, params *uint32) {
	C.glowGetnUniformuivKHR(gpGetnUniformuivKHR, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLuint)(unsafe.Pointer(params)))
}

// specify implementation-specific hints
func Hint(target uint32, mode uint32) {
	C.glowHint(gpHint, (C.GLenum)(target), (C.GLenum)(mode))
}
func ImportMemoryFdEXT(memory uint32, size uint64, handleType uint32, fd int32) {
	C.glowImportMemoryFdEXT(gpImportMemoryFdEXT, (C.GLuint)(memory), (C.GLuint64)(size), (C.GLenum)(handleType), (C.GLint)(fd))
}
func ImportMemoryWin32HandleEXT(memory uint32, size uint64, handleType uint32, handle unsafe.Pointer) {
	C.glowImportMemoryWin32HandleEXT(gpImportMemoryWin32HandleEXT, (C.GLuint)(memory), (C.GLuint64)(size), (C.GLenum)(handleType), handle)
}
func ImportMemoryWin32NameEXT(memory uint32, size uint64, handleType uint32, name unsafe.Pointer) {
	C.glowImportMemoryWin32NameEXT(gpImportMemoryWin32NameEXT, (C.GLuint)(memory), (C.GLuint64)(size), (C.GLenum)(handleType), name)
}
func ImportSemaphoreFdEXT(semaphore uint32, handleType uint32, fd int32) {
	C.glowImportSemaphoreFdEXT(gpImportSemaphoreFdEXT, (C.GLuint)(semaphore), (C.GLenum)(handleType), (C.GLint)(fd))
}
func ImportSemaphoreWin32HandleEXT(semaphore uint32, handleType uint32, handle unsafe.Pointer) {
	C.glowImportSemaphoreWin32HandleEXT(gpImportSemaphoreWin32HandleEXT, (C.GLuint)(semaphore), (C.GLenum)(handleType), handle)
}
func ImportSemaphoreWin32NameEXT(semaphore uint32, handleType uint32, name unsafe.Pointer) {
	C.glowImportSemaphoreWin32NameEXT(gpImportSemaphoreWin32NameEXT, (C.GLuint)(semaphore), (C.GLenum)(handleType), name)
}
func InsertEventMarkerEXT(length int32, marker *uint8) {
	C.glowInsertEventMarkerEXT(gpInsertEventMarkerEXT, (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(marker)))
}
func InterpolatePathsNV(resultPath uint32, pathA uint32, pathB uint32, weight float32) {
	C.glowInterpolatePathsNV(gpInterpolatePathsNV, (C.GLuint)(resultPath), (C.GLuint)(pathA), (C.GLuint)(pathB), (C.GLfloat)(weight))
}

// invalidate the content of some or all of a framebuffer's attachments
func InvalidateFramebuffer(target uint32, numAttachments int32, attachments *uint32) {
	C.glowInvalidateFramebuffer(gpInvalidateFramebuffer, (C.GLenum)(target), (C.GLsizei)(numAttachments), (*C.GLenum)(unsafe.Pointer(attachments)))
}

// invalidate the content of a region of some or all of a framebuffer's attachments
func InvalidateSubFramebuffer(target uint32, numAttachments int32, attachments *uint32, x int32, y int32, width int32, height int32) {
	C.glowInvalidateSubFramebuffer(gpInvalidateSubFramebuffer, (C.GLenum)(target), (C.GLsizei)(numAttachments), (*C.GLenum)(unsafe.Pointer(attachments)), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

// determine if a name corresponds to a buffer object
func IsBuffer(buffer uint32) bool {
	ret := C.glowIsBuffer(gpIsBuffer, (C.GLuint)(buffer))
	return ret == TRUE
}
func IsEnabled(cap uint32) bool {
	ret := C.glowIsEnabled(gpIsEnabled, (C.GLenum)(cap))
	return ret == TRUE
}
func IsEnablediEXT(target uint32, index uint32) bool {
	ret := C.glowIsEnablediEXT(gpIsEnablediEXT, (C.GLenum)(target), (C.GLuint)(index))
	return ret == TRUE
}
func IsEnablediNV(target uint32, index uint32) bool {
	ret := C.glowIsEnablediNV(gpIsEnablediNV, (C.GLenum)(target), (C.GLuint)(index))
	return ret == TRUE
}
func IsEnablediOES(target uint32, index uint32) bool {
	ret := C.glowIsEnablediOES(gpIsEnablediOES, (C.GLenum)(target), (C.GLuint)(index))
	return ret == TRUE
}
func IsFenceNV(fence uint32) bool {
	ret := C.glowIsFenceNV(gpIsFenceNV, (C.GLuint)(fence))
	return ret == TRUE
}

// determine if a name corresponds to a framebuffer object
func IsFramebuffer(framebuffer uint32) bool {
	ret := C.glowIsFramebuffer(gpIsFramebuffer, (C.GLuint)(framebuffer))
	return ret == TRUE
}
func IsImageHandleResidentNV(handle uint64) bool {
	ret := C.glowIsImageHandleResidentNV(gpIsImageHandleResidentNV, (C.GLuint64)(handle))
	return ret == TRUE
}
func IsMemoryObjectEXT(memoryObject uint32) bool {
	ret := C.glowIsMemoryObjectEXT(gpIsMemoryObjectEXT, (C.GLuint)(memoryObject))
	return ret == TRUE
}
func IsPathNV(path uint32) bool {
	ret := C.glowIsPathNV(gpIsPathNV, (C.GLuint)(path))
	return ret == TRUE
}
func IsPointInFillPathNV(path uint32, mask uint32, x float32, y float32) bool {
	ret := C.glowIsPointInFillPathNV(gpIsPointInFillPathNV, (C.GLuint)(path), (C.GLuint)(mask), (C.GLfloat)(x), (C.GLfloat)(y))
	return ret == TRUE
}
func IsPointInStrokePathNV(path uint32, x float32, y float32) bool {
	ret := C.glowIsPointInStrokePathNV(gpIsPointInStrokePathNV, (C.GLuint)(path), (C.GLfloat)(x), (C.GLfloat)(y))
	return ret == TRUE
}

// Determines if a name corresponds to a program object
func IsProgram(program uint32) bool {
	ret := C.glowIsProgram(gpIsProgram, (C.GLuint)(program))
	return ret == TRUE
}

// determine if a name corresponds to a program pipeline object
func IsProgramPipeline(pipeline uint32) bool {
	ret := C.glowIsProgramPipeline(gpIsProgramPipeline, (C.GLuint)(pipeline))
	return ret == TRUE
}
func IsProgramPipelineEXT(pipeline uint32) bool {
	ret := C.glowIsProgramPipelineEXT(gpIsProgramPipelineEXT, (C.GLuint)(pipeline))
	return ret == TRUE
}

// determine if a name corresponds to a query object
func IsQuery(id uint32) bool {
	ret := C.glowIsQuery(gpIsQuery, (C.GLuint)(id))
	return ret == TRUE
}
func IsQueryEXT(id uint32) bool {
	ret := C.glowIsQueryEXT(gpIsQueryEXT, (C.GLuint)(id))
	return ret == TRUE
}

// determine if a name corresponds to a renderbuffer object
func IsRenderbuffer(renderbuffer uint32) bool {
	ret := C.glowIsRenderbuffer(gpIsRenderbuffer, (C.GLuint)(renderbuffer))
	return ret == TRUE
}

// determine if a name corresponds to a sampler object
func IsSampler(sampler uint32) bool {
	ret := C.glowIsSampler(gpIsSampler, (C.GLuint)(sampler))
	return ret == TRUE
}
func IsSemaphoreEXT(semaphore uint32) bool {
	ret := C.glowIsSemaphoreEXT(gpIsSemaphoreEXT, (C.GLuint)(semaphore))
	return ret == TRUE
}

// Determines if a name corresponds to a shader object
func IsShader(shader uint32) bool {
	ret := C.glowIsShader(gpIsShader, (C.GLuint)(shader))
	return ret == TRUE
}

// determine if a name corresponds to a sync object
func IsSync(sync uintptr) bool {
	ret := C.glowIsSync(gpIsSync, (C.GLsync)(sync))
	return ret == TRUE
}
func IsSyncAPPLE(sync uintptr) bool {
	ret := C.glowIsSyncAPPLE(gpIsSyncAPPLE, (C.GLsync)(sync))
	return ret == TRUE
}

// determine if a name corresponds to a texture
func IsTexture(texture uint32) bool {
	ret := C.glowIsTexture(gpIsTexture, (C.GLuint)(texture))
	return ret == TRUE
}
func IsTextureHandleResidentNV(handle uint64) bool {
	ret := C.glowIsTextureHandleResidentNV(gpIsTextureHandleResidentNV, (C.GLuint64)(handle))
	return ret == TRUE
}

// determine if a name corresponds to a transform feedback object
func IsTransformFeedback(id uint32) bool {
	ret := C.glowIsTransformFeedback(gpIsTransformFeedback, (C.GLuint)(id))
	return ret == TRUE
}

// determine if a name corresponds to a vertex array object
func IsVertexArray(array uint32) bool {
	ret := C.glowIsVertexArray(gpIsVertexArray, (C.GLuint)(array))
	return ret == TRUE
}
func IsVertexArrayOES(array uint32) bool {
	ret := C.glowIsVertexArrayOES(gpIsVertexArrayOES, (C.GLuint)(array))
	return ret == TRUE
}
func LabelObjectEXT(xtype uint32, object uint32, length int32, label *uint8) {
	C.glowLabelObjectEXT(gpLabelObjectEXT, (C.GLenum)(xtype), (C.GLuint)(object), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(label)))
}

// specify the width of rasterized lines
func LineWidth(width float32) {
	C.glowLineWidth(gpLineWidth, (C.GLfloat)(width))
}

// Links a program object
func LinkProgram(program uint32) {
	C.glowLinkProgram(gpLinkProgram, (C.GLuint)(program))
}
func MakeImageHandleNonResidentNV(handle uint64) {
	C.glowMakeImageHandleNonResidentNV(gpMakeImageHandleNonResidentNV, (C.GLuint64)(handle))
}
func MakeImageHandleResidentNV(handle uint64, access uint32) {
	C.glowMakeImageHandleResidentNV(gpMakeImageHandleResidentNV, (C.GLuint64)(handle), (C.GLenum)(access))
}
func MakeTextureHandleNonResidentNV(handle uint64) {
	C.glowMakeTextureHandleNonResidentNV(gpMakeTextureHandleNonResidentNV, (C.GLuint64)(handle))
}
func MakeTextureHandleResidentNV(handle uint64) {
	C.glowMakeTextureHandleResidentNV(gpMakeTextureHandleResidentNV, (C.GLuint64)(handle))
}
func MapBufferOES(target uint32, access uint32) unsafe.Pointer {
	ret := C.glowMapBufferOES(gpMapBufferOES, (C.GLenum)(target), (C.GLenum)(access))
	return (unsafe.Pointer)(ret)
}

// map all or part of a buffer object's data store into the client's address space
func MapBufferRange(target uint32, offset int, length int, access uint32) unsafe.Pointer {
	ret := C.glowMapBufferRange(gpMapBufferRange, (C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(length), (C.GLbitfield)(access))
	return (unsafe.Pointer)(ret)
}
func MapBufferRangeEXT(target uint32, offset int, length int, access uint32) unsafe.Pointer {
	ret := C.glowMapBufferRangeEXT(gpMapBufferRangeEXT, (C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(length), (C.GLbitfield)(access))
	return (unsafe.Pointer)(ret)
}
func MatrixFrustumEXT(mode uint32, left float64, right float64, bottom float64, top float64, zNear float64, zFar float64) {
	C.glowMatrixFrustumEXT(gpMatrixFrustumEXT, (C.GLenum)(mode), (C.GLdouble)(left), (C.GLdouble)(right), (C.GLdouble)(bottom), (C.GLdouble)(top), (C.GLdouble)(zNear), (C.GLdouble)(zFar))
}
func MatrixLoad3x2fNV(matrixMode uint32, m *float32) {
	C.glowMatrixLoad3x2fNV(gpMatrixLoad3x2fNV, (C.GLenum)(matrixMode), (*C.GLfloat)(unsafe.Pointer(m)))
}
func MatrixLoad3x3fNV(matrixMode uint32, m *float32) {
	C.glowMatrixLoad3x3fNV(gpMatrixLoad3x3fNV, (C.GLenum)(matrixMode), (*C.GLfloat)(unsafe.Pointer(m)))
}
func MatrixLoadIdentityEXT(mode uint32) {
	C.glowMatrixLoadIdentityEXT(gpMatrixLoadIdentityEXT, (C.GLenum)(mode))
}
func MatrixLoadTranspose3x3fNV(matrixMode uint32, m *float32) {
	C.glowMatrixLoadTranspose3x3fNV(gpMatrixLoadTranspose3x3fNV, (C.GLenum)(matrixMode), (*C.GLfloat)(unsafe.Pointer(m)))
}
func MatrixLoadTransposedEXT(mode uint32, m *float64) {
	C.glowMatrixLoadTransposedEXT(gpMatrixLoadTransposedEXT, (C.GLenum)(mode), (*C.GLdouble)(unsafe.Pointer(m)))
}
func MatrixLoadTransposefEXT(mode uint32, m *float32) {
	C.glowMatrixLoadTransposefEXT(gpMatrixLoadTransposefEXT, (C.GLenum)(mode), (*C.GLfloat)(unsafe.Pointer(m)))
}
func MatrixLoaddEXT(mode uint32, m *float64) {
	C.glowMatrixLoaddEXT(gpMatrixLoaddEXT, (C.GLenum)(mode), (*C.GLdouble)(unsafe.Pointer(m)))
}
func MatrixLoadfEXT(mode uint32, m *float32) {
	C.glowMatrixLoadfEXT(gpMatrixLoadfEXT, (C.GLenum)(mode), (*C.GLfloat)(unsafe.Pointer(m)))
}
func MatrixMult3x2fNV(matrixMode uint32, m *float32) {
	C.glowMatrixMult3x2fNV(gpMatrixMult3x2fNV, (C.GLenum)(matrixMode), (*C.GLfloat)(unsafe.Pointer(m)))
}
func MatrixMult3x3fNV(matrixMode uint32, m *float32) {
	C.glowMatrixMult3x3fNV(gpMatrixMult3x3fNV, (C.GLenum)(matrixMode), (*C.GLfloat)(unsafe.Pointer(m)))
}
func MatrixMultTranspose3x3fNV(matrixMode uint32, m *float32) {
	C.glowMatrixMultTranspose3x3fNV(gpMatrixMultTranspose3x3fNV, (C.GLenum)(matrixMode), (*C.GLfloat)(unsafe.Pointer(m)))
}
func MatrixMultTransposedEXT(mode uint32, m *float64) {
	C.glowMatrixMultTransposedEXT(gpMatrixMultTransposedEXT, (C.GLenum)(mode), (*C.GLdouble)(unsafe.Pointer(m)))
}
func MatrixMultTransposefEXT(mode uint32, m *float32) {
	C.glowMatrixMultTransposefEXT(gpMatrixMultTransposefEXT, (C.GLenum)(mode), (*C.GLfloat)(unsafe.Pointer(m)))
}
func MatrixMultdEXT(mode uint32, m *float64) {
	C.glowMatrixMultdEXT(gpMatrixMultdEXT, (C.GLenum)(mode), (*C.GLdouble)(unsafe.Pointer(m)))
}
func MatrixMultfEXT(mode uint32, m *float32) {
	C.glowMatrixMultfEXT(gpMatrixMultfEXT, (C.GLenum)(mode), (*C.GLfloat)(unsafe.Pointer(m)))
}
func MatrixOrthoEXT(mode uint32, left float64, right float64, bottom float64, top float64, zNear float64, zFar float64) {
	C.glowMatrixOrthoEXT(gpMatrixOrthoEXT, (C.GLenum)(mode), (C.GLdouble)(left), (C.GLdouble)(right), (C.GLdouble)(bottom), (C.GLdouble)(top), (C.GLdouble)(zNear), (C.GLdouble)(zFar))
}
func MatrixPopEXT(mode uint32) {
	C.glowMatrixPopEXT(gpMatrixPopEXT, (C.GLenum)(mode))
}
func MatrixPushEXT(mode uint32) {
	C.glowMatrixPushEXT(gpMatrixPushEXT, (C.GLenum)(mode))
}
func MatrixRotatedEXT(mode uint32, angle float64, x float64, y float64, z float64) {
	C.glowMatrixRotatedEXT(gpMatrixRotatedEXT, (C.GLenum)(mode), (C.GLdouble)(angle), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func MatrixRotatefEXT(mode uint32, angle float32, x float32, y float32, z float32) {
	C.glowMatrixRotatefEXT(gpMatrixRotatefEXT, (C.GLenum)(mode), (C.GLfloat)(angle), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func MatrixScaledEXT(mode uint32, x float64, y float64, z float64) {
	C.glowMatrixScaledEXT(gpMatrixScaledEXT, (C.GLenum)(mode), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func MatrixScalefEXT(mode uint32, x float32, y float32, z float32) {
	C.glowMatrixScalefEXT(gpMatrixScalefEXT, (C.GLenum)(mode), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func MatrixTranslatedEXT(mode uint32, x float64, y float64, z float64) {
	C.glowMatrixTranslatedEXT(gpMatrixTranslatedEXT, (C.GLenum)(mode), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func MatrixTranslatefEXT(mode uint32, x float32, y float32, z float32) {
	C.glowMatrixTranslatefEXT(gpMatrixTranslatefEXT, (C.GLenum)(mode), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func MaxShaderCompilerThreadsKHR(count uint32) {
	C.glowMaxShaderCompilerThreadsKHR(gpMaxShaderCompilerThreadsKHR, (C.GLuint)(count))
}

// defines a barrier ordering memory transactions
func MemoryBarrier(barriers uint32) {
	C.glowMemoryBarrier(gpMemoryBarrier, (C.GLbitfield)(barriers))
}
func MemoryBarrierByRegion(barriers uint32) {
	C.glowMemoryBarrierByRegion(gpMemoryBarrierByRegion, (C.GLbitfield)(barriers))
}
func MemoryObjectParameterivEXT(memoryObject uint32, pname uint32, params *int32) {
	C.glowMemoryObjectParameterivEXT(gpMemoryObjectParameterivEXT, (C.GLuint)(memoryObject), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func MinSampleShadingOES(value float32) {
	C.glowMinSampleShadingOES(gpMinSampleShadingOES, (C.GLfloat)(value))
}
func MultiDrawArraysEXT(mode uint32, first *int32, count *int32, primcount int32) {
	C.glowMultiDrawArraysEXT(gpMultiDrawArraysEXT, (C.GLenum)(mode), (*C.GLint)(unsafe.Pointer(first)), (*C.GLsizei)(unsafe.Pointer(count)), (C.GLsizei)(primcount))
}
func MultiDrawArraysIndirectEXT(mode uint32, indirect unsafe.Pointer, drawcount int32, stride int32) {
	C.glowMultiDrawArraysIndirectEXT(gpMultiDrawArraysIndirectEXT, (C.GLenum)(mode), indirect, (C.GLsizei)(drawcount), (C.GLsizei)(stride))
}
func MultiDrawElementsBaseVertexEXT(mode uint32, count *int32, xtype uint32, indices *unsafe.Pointer, primcount int32, basevertex *int32) {
	C.glowMultiDrawElementsBaseVertexEXT(gpMultiDrawElementsBaseVertexEXT, (C.GLenum)(mode), (*C.GLsizei)(unsafe.Pointer(count)), (C.GLenum)(xtype), indices, (C.GLsizei)(primcount), (*C.GLint)(unsafe.Pointer(basevertex)))
}
func MultiDrawElementsEXT(mode uint32, count *int32, xtype uint32, indices *unsafe.Pointer, primcount int32) {
	C.glowMultiDrawElementsEXT(gpMultiDrawElementsEXT, (C.GLenum)(mode), (*C.GLsizei)(unsafe.Pointer(count)), (C.GLenum)(xtype), indices, (C.GLsizei)(primcount))
}
func MultiDrawElementsIndirectEXT(mode uint32, xtype uint32, indirect unsafe.Pointer, drawcount int32, stride int32) {
	C.glowMultiDrawElementsIndirectEXT(gpMultiDrawElementsIndirectEXT, (C.GLenum)(mode), (C.GLenum)(xtype), indirect, (C.GLsizei)(drawcount), (C.GLsizei)(stride))
}
func NamedBufferStorageExternalEXT(buffer uint32, offset int, size int, clientBuffer C.GLeglClientBufferEXT, flags uint32) {
	C.glowNamedBufferStorageExternalEXT(gpNamedBufferStorageExternalEXT, (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(size), (C.GLeglClientBufferEXT)(clientBuffer), (C.GLbitfield)(flags))
}
func NamedBufferStorageMemEXT(buffer uint32, size int, memory uint32, offset uint64) {
	C.glowNamedBufferStorageMemEXT(gpNamedBufferStorageMemEXT, (C.GLuint)(buffer), (C.GLsizeiptr)(size), (C.GLuint)(memory), (C.GLuint64)(offset))
}
func NamedFramebufferSampleLocationsfvNV(framebuffer uint32, start uint32, count int32, v *float32) {
	C.glowNamedFramebufferSampleLocationsfvNV(gpNamedFramebufferSampleLocationsfvNV, (C.GLuint)(framebuffer), (C.GLuint)(start), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(v)))
}

// label a named object identified within a namespace
func ObjectLabel(identifier uint32, name uint32, length int32, label *uint8) {
	C.glowObjectLabel(gpObjectLabel, (C.GLenum)(identifier), (C.GLuint)(name), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(label)))
}
func ObjectLabelKHR(identifier uint32, name uint32, length int32, label *uint8) {
	C.glowObjectLabelKHR(gpObjectLabelKHR, (C.GLenum)(identifier), (C.GLuint)(name), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(label)))
}

// label a a sync object identified by a pointer
func ObjectPtrLabel(ptr unsafe.Pointer, length int32, label *uint8) {
	C.glowObjectPtrLabel(gpObjectPtrLabel, ptr, (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(label)))
}
func ObjectPtrLabelKHR(ptr unsafe.Pointer, length int32, label *uint8) {
	C.glowObjectPtrLabelKHR(gpObjectPtrLabelKHR, ptr, (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(label)))
}
func PatchParameteriEXT(pname uint32, value int32) {
	C.glowPatchParameteriEXT(gpPatchParameteriEXT, (C.GLenum)(pname), (C.GLint)(value))
}
func PatchParameteriOES(pname uint32, value int32) {
	C.glowPatchParameteriOES(gpPatchParameteriOES, (C.GLenum)(pname), (C.GLint)(value))
}
func PathCommandsNV(path uint32, numCommands int32, commands *uint8, numCoords int32, coordType uint32, coords unsafe.Pointer) {
	C.glowPathCommandsNV(gpPathCommandsNV, (C.GLuint)(path), (C.GLsizei)(numCommands), (*C.GLubyte)(unsafe.Pointer(commands)), (C.GLsizei)(numCoords), (C.GLenum)(coordType), coords)
}
func PathCoordsNV(path uint32, numCoords int32, coordType uint32, coords unsafe.Pointer) {
	C.glowPathCoordsNV(gpPathCoordsNV, (C.GLuint)(path), (C.GLsizei)(numCoords), (C.GLenum)(coordType), coords)
}
func PathCoverDepthFuncNV(xfunc uint32) {
	C.glowPathCoverDepthFuncNV(gpPathCoverDepthFuncNV, (C.GLenum)(xfunc))
}
func PathDashArrayNV(path uint32, dashCount int32, dashArray *float32) {
	C.glowPathDashArrayNV(gpPathDashArrayNV, (C.GLuint)(path), (C.GLsizei)(dashCount), (*C.GLfloat)(unsafe.Pointer(dashArray)))
}
func PathGlyphIndexArrayNV(firstPathName uint32, fontTarget uint32, fontName unsafe.Pointer, fontStyle uint32, firstGlyphIndex uint32, numGlyphs int32, pathParameterTemplate uint32, emScale float32) uint32 {
	ret := C.glowPathGlyphIndexArrayNV(gpPathGlyphIndexArrayNV, (C.GLuint)(firstPathName), (C.GLenum)(fontTarget), fontName, (C.GLbitfield)(fontStyle), (C.GLuint)(firstGlyphIndex), (C.GLsizei)(numGlyphs), (C.GLuint)(pathParameterTemplate), (C.GLfloat)(emScale))
	return (uint32)(ret)
}
func PathGlyphIndexRangeNV(fontTarget uint32, fontName unsafe.Pointer, fontStyle uint32, pathParameterTemplate uint32, emScale float32, baseAndCount *uint32) uint32 {
	ret := C.glowPathGlyphIndexRangeNV(gpPathGlyphIndexRangeNV, (C.GLenum)(fontTarget), fontName, (C.GLbitfield)(fontStyle), (C.GLuint)(pathParameterTemplate), (C.GLfloat)(emScale), (*C.GLuint)(unsafe.Pointer(baseAndCount)))
	return (uint32)(ret)
}
func PathGlyphRangeNV(firstPathName uint32, fontTarget uint32, fontName unsafe.Pointer, fontStyle uint32, firstGlyph uint32, numGlyphs int32, handleMissingGlyphs uint32, pathParameterTemplate uint32, emScale float32) {
	C.glowPathGlyphRangeNV(gpPathGlyphRangeNV, (C.GLuint)(firstPathName), (C.GLenum)(fontTarget), fontName, (C.GLbitfield)(fontStyle), (C.GLuint)(firstGlyph), (C.GLsizei)(numGlyphs), (C.GLenum)(handleMissingGlyphs), (C.GLuint)(pathParameterTemplate), (C.GLfloat)(emScale))
}
func PathGlyphsNV(firstPathName uint32, fontTarget uint32, fontName unsafe.Pointer, fontStyle uint32, numGlyphs int32, xtype uint32, charcodes unsafe.Pointer, handleMissingGlyphs uint32, pathParameterTemplate uint32, emScale float32) {
	C.glowPathGlyphsNV(gpPathGlyphsNV, (C.GLuint)(firstPathName), (C.GLenum)(fontTarget), fontName, (C.GLbitfield)(fontStyle), (C.GLsizei)(numGlyphs), (C.GLenum)(xtype), charcodes, (C.GLenum)(handleMissingGlyphs), (C.GLuint)(pathParameterTemplate), (C.GLfloat)(emScale))
}
func PathMemoryGlyphIndexArrayNV(firstPathName uint32, fontTarget uint32, fontSize int, fontData unsafe.Pointer, faceIndex int32, firstGlyphIndex uint32, numGlyphs int32, pathParameterTemplate uint32, emScale float32) uint32 {
	ret := C.glowPathMemoryGlyphIndexArrayNV(gpPathMemoryGlyphIndexArrayNV, (C.GLuint)(firstPathName), (C.GLenum)(fontTarget), (C.GLsizeiptr)(fontSize), fontData, (C.GLsizei)(faceIndex), (C.GLuint)(firstGlyphIndex), (C.GLsizei)(numGlyphs), (C.GLuint)(pathParameterTemplate), (C.GLfloat)(emScale))
	return (uint32)(ret)
}
func PathParameterfNV(path uint32, pname uint32, value float32) {
	C.glowPathParameterfNV(gpPathParameterfNV, (C.GLuint)(path), (C.GLenum)(pname), (C.GLfloat)(value))
}
func PathParameterfvNV(path uint32, pname uint32, value *float32) {
	C.glowPathParameterfvNV(gpPathParameterfvNV, (C.GLuint)(path), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(value)))
}
func PathParameteriNV(path uint32, pname uint32, value int32) {
	C.glowPathParameteriNV(gpPathParameteriNV, (C.GLuint)(path), (C.GLenum)(pname), (C.GLint)(value))
}
func PathParameterivNV(path uint32, pname uint32, value *int32) {
	C.glowPathParameterivNV(gpPathParameterivNV, (C.GLuint)(path), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(value)))
}
func PathStencilDepthOffsetNV(factor float32, units float32) {
	C.glowPathStencilDepthOffsetNV(gpPathStencilDepthOffsetNV, (C.GLfloat)(factor), (C.GLfloat)(units))
}
func PathStencilFuncNV(xfunc uint32, ref int32, mask uint32) {
	C.glowPathStencilFuncNV(gpPathStencilFuncNV, (C.GLenum)(xfunc), (C.GLint)(ref), (C.GLuint)(mask))
}
func PathStringNV(path uint32, format uint32, length int32, pathString unsafe.Pointer) {
	C.glowPathStringNV(gpPathStringNV, (C.GLuint)(path), (C.GLenum)(format), (C.GLsizei)(length), pathString)
}
func PathSubCommandsNV(path uint32, commandStart int32, commandsToDelete int32, numCommands int32, commands *uint8, numCoords int32, coordType uint32, coords unsafe.Pointer) {
	C.glowPathSubCommandsNV(gpPathSubCommandsNV, (C.GLuint)(path), (C.GLsizei)(commandStart), (C.GLsizei)(commandsToDelete), (C.GLsizei)(numCommands), (*C.GLubyte)(unsafe.Pointer(commands)), (C.GLsizei)(numCoords), (C.GLenum)(coordType), coords)
}
func PathSubCoordsNV(path uint32, coordStart int32, numCoords int32, coordType uint32, coords unsafe.Pointer) {
	C.glowPathSubCoordsNV(gpPathSubCoordsNV, (C.GLuint)(path), (C.GLsizei)(coordStart), (C.GLsizei)(numCoords), (C.GLenum)(coordType), coords)
}

// pause transform feedback operations
func PauseTransformFeedback() {
	C.glowPauseTransformFeedback(gpPauseTransformFeedback)
}

// set pixel storage modes
func PixelStorei(pname uint32, param int32) {
	C.glowPixelStorei(gpPixelStorei, (C.GLenum)(pname), (C.GLint)(param))
}
func PointAlongPathNV(path uint32, startSegment int32, numSegments int32, distance float32, x *float32, y *float32, tangentX *float32, tangentY *float32) bool {
	ret := C.glowPointAlongPathNV(gpPointAlongPathNV, (C.GLuint)(path), (C.GLsizei)(startSegment), (C.GLsizei)(numSegments), (C.GLfloat)(distance), (*C.GLfloat)(unsafe.Pointer(x)), (*C.GLfloat)(unsafe.Pointer(y)), (*C.GLfloat)(unsafe.Pointer(tangentX)), (*C.GLfloat)(unsafe.Pointer(tangentY)))
	return ret == TRUE
}
func PolygonModeNV(face uint32, mode uint32) {
	C.glowPolygonModeNV(gpPolygonModeNV, (C.GLenum)(face), (C.GLenum)(mode))
}

// set the scale and units used to calculate depth values
func PolygonOffset(factor float32, units float32) {
	C.glowPolygonOffset(gpPolygonOffset, (C.GLfloat)(factor), (C.GLfloat)(units))
}
func PolygonOffsetClampEXT(factor float32, units float32, clamp float32) {
	C.glowPolygonOffsetClampEXT(gpPolygonOffsetClampEXT, (C.GLfloat)(factor), (C.GLfloat)(units), (C.GLfloat)(clamp))
}

// pop the active debug group
func PopDebugGroup() {
	C.glowPopDebugGroup(gpPopDebugGroup)
}
func PopDebugGroupKHR() {
	C.glowPopDebugGroupKHR(gpPopDebugGroupKHR)
}
func PopGroupMarkerEXT() {
	C.glowPopGroupMarkerEXT(gpPopGroupMarkerEXT)
}
func PrimitiveBoundingBoxEXT(minX float32, minY float32, minZ float32, minW float32, maxX float32, maxY float32, maxZ float32, maxW float32) {
	C.glowPrimitiveBoundingBoxEXT(gpPrimitiveBoundingBoxEXT, (C.GLfloat)(minX), (C.GLfloat)(minY), (C.GLfloat)(minZ), (C.GLfloat)(minW), (C.GLfloat)(maxX), (C.GLfloat)(maxY), (C.GLfloat)(maxZ), (C.GLfloat)(maxW))
}
func PrimitiveBoundingBoxOES(minX float32, minY float32, minZ float32, minW float32, maxX float32, maxY float32, maxZ float32, maxW float32) {
	C.glowPrimitiveBoundingBoxOES(gpPrimitiveBoundingBoxOES, (C.GLfloat)(minX), (C.GLfloat)(minY), (C.GLfloat)(minZ), (C.GLfloat)(minW), (C.GLfloat)(maxX), (C.GLfloat)(maxY), (C.GLfloat)(maxZ), (C.GLfloat)(maxW))
}

// load a program object with a program binary
func ProgramBinary(program uint32, binaryFormat uint32, binary unsafe.Pointer, length int32) {
	C.glowProgramBinary(gpProgramBinary, (C.GLuint)(program), (C.GLenum)(binaryFormat), binary, (C.GLsizei)(length))
}
func ProgramBinaryOES(program uint32, binaryFormat uint32, binary unsafe.Pointer, length int32) {
	C.glowProgramBinaryOES(gpProgramBinaryOES, (C.GLuint)(program), (C.GLenum)(binaryFormat), binary, (C.GLint)(length))
}

// specify a parameter for a program object
func ProgramParameteri(program uint32, pname uint32, value int32) {
	C.glowProgramParameteri(gpProgramParameteri, (C.GLuint)(program), (C.GLenum)(pname), (C.GLint)(value))
}
func ProgramParameteriEXT(program uint32, pname uint32, value int32) {
	C.glowProgramParameteriEXT(gpProgramParameteriEXT, (C.GLuint)(program), (C.GLenum)(pname), (C.GLint)(value))
}
func ProgramPathFragmentInputGenNV(program uint32, location int32, genMode uint32, components int32, coeffs *float32) {
	C.glowProgramPathFragmentInputGenNV(gpProgramPathFragmentInputGenNV, (C.GLuint)(program), (C.GLint)(location), (C.GLenum)(genMode), (C.GLint)(components), (*C.GLfloat)(unsafe.Pointer(coeffs)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1f(program uint32, location int32, v0 float32) {
	C.glowProgramUniform1f(gpProgramUniform1f, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0))
}
func ProgramUniform1fEXT(program uint32, location int32, v0 float32) {
	C.glowProgramUniform1fEXT(gpProgramUniform1fEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1fv(program uint32, location int32, count int32, value *float32) {
	C.glowProgramUniform1fv(gpProgramUniform1fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniform1fvEXT(program uint32, location int32, count int32, value *float32) {
	C.glowProgramUniform1fvEXT(gpProgramUniform1fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1i(program uint32, location int32, v0 int32) {
	C.glowProgramUniform1i(gpProgramUniform1i, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0))
}
func ProgramUniform1i64NV(program uint32, location int32, x int64) {
	C.glowProgramUniform1i64NV(gpProgramUniform1i64NV, (C.GLuint)(program), (C.GLint)(location), (C.GLint64EXT)(x))
}
func ProgramUniform1i64vNV(program uint32, location int32, count int32, value *int64) {
	C.glowProgramUniform1i64vNV(gpProgramUniform1i64vNV, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint64EXT)(unsafe.Pointer(value)))
}
func ProgramUniform1iEXT(program uint32, location int32, v0 int32) {
	C.glowProgramUniform1iEXT(gpProgramUniform1iEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1iv(program uint32, location int32, count int32, value *int32) {
	C.glowProgramUniform1iv(gpProgramUniform1iv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}
func ProgramUniform1ivEXT(program uint32, location int32, count int32, value *int32) {
	C.glowProgramUniform1ivEXT(gpProgramUniform1ivEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1ui(program uint32, location int32, v0 uint32) {
	C.glowProgramUniform1ui(gpProgramUniform1ui, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0))
}
func ProgramUniform1ui64NV(program uint32, location int32, x uint64) {
	C.glowProgramUniform1ui64NV(gpProgramUniform1ui64NV, (C.GLuint)(program), (C.GLint)(location), (C.GLuint64EXT)(x))
}
func ProgramUniform1ui64vNV(program uint32, location int32, count int32, value *uint64) {
	C.glowProgramUniform1ui64vNV(gpProgramUniform1ui64vNV, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64EXT)(unsafe.Pointer(value)))
}
func ProgramUniform1uiEXT(program uint32, location int32, v0 uint32) {
	C.glowProgramUniform1uiEXT(gpProgramUniform1uiEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1uiv(program uint32, location int32, count int32, value *uint32) {
	C.glowProgramUniform1uiv(gpProgramUniform1uiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func ProgramUniform1uivEXT(program uint32, location int32, count int32, value *uint32) {
	C.glowProgramUniform1uivEXT(gpProgramUniform1uivEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2f(program uint32, location int32, v0 float32, v1 float32) {
	C.glowProgramUniform2f(gpProgramUniform2f, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1))
}
func ProgramUniform2fEXT(program uint32, location int32, v0 float32, v1 float32) {
	C.glowProgramUniform2fEXT(gpProgramUniform2fEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2fv(program uint32, location int32, count int32, value *float32) {
	C.glowProgramUniform2fv(gpProgramUniform2fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniform2fvEXT(program uint32, location int32, count int32, value *float32) {
	C.glowProgramUniform2fvEXT(gpProgramUniform2fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2i(program uint32, location int32, v0 int32, v1 int32) {
	C.glowProgramUniform2i(gpProgramUniform2i, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1))
}
func ProgramUniform2i64NV(program uint32, location int32, x int64, y int64) {
	C.glowProgramUniform2i64NV(gpProgramUniform2i64NV, (C.GLuint)(program), (C.GLint)(location), (C.GLint64EXT)(x), (C.GLint64EXT)(y))
}
func ProgramUniform2i64vNV(program uint32, location int32, count int32, value *int64) {
	C.glowProgramUniform2i64vNV(gpProgramUniform2i64vNV, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint64EXT)(unsafe.Pointer(value)))
}
func ProgramUniform2iEXT(program uint32, location int32, v0 int32, v1 int32) {
	C.glowProgramUniform2iEXT(gpProgramUniform2iEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2iv(program uint32, location int32, count int32, value *int32) {
	C.glowProgramUniform2iv(gpProgramUniform2iv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}
func ProgramUniform2ivEXT(program uint32, location int32, count int32, value *int32) {
	C.glowProgramUniform2ivEXT(gpProgramUniform2ivEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2ui(program uint32, location int32, v0 uint32, v1 uint32) {
	C.glowProgramUniform2ui(gpProgramUniform2ui, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1))
}
func ProgramUniform2ui64NV(program uint32, location int32, x uint64, y uint64) {
	C.glowProgramUniform2ui64NV(gpProgramUniform2ui64NV, (C.GLuint)(program), (C.GLint)(location), (C.GLuint64EXT)(x), (C.GLuint64EXT)(y))
}
func ProgramUniform2ui64vNV(program uint32, location int32, count int32, value *uint64) {
	C.glowProgramUniform2ui64vNV(gpProgramUniform2ui64vNV, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64EXT)(unsafe.Pointer(value)))
}
func ProgramUniform2uiEXT(program uint32, location int32, v0 uint32, v1 uint32) {
	C.glowProgramUniform2uiEXT(gpProgramUniform2uiEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2uiv(program uint32, location int32, count int32, value *uint32) {
	C.glowProgramUniform2uiv(gpProgramUniform2uiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func ProgramUniform2uivEXT(program uint32, location int32, count int32, value *uint32) {
	C.glowProgramUniform2uivEXT(gpProgramUniform2uivEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3f(program uint32, location int32, v0 float32, v1 float32, v2 float32) {
	C.glowProgramUniform3f(gpProgramUniform3f, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2))
}
func ProgramUniform3fEXT(program uint32, location int32, v0 float32, v1 float32, v2 float32) {
	C.glowProgramUniform3fEXT(gpProgramUniform3fEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3fv(program uint32, location int32, count int32, value *float32) {
	C.glowProgramUniform3fv(gpProgramUniform3fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniform3fvEXT(program uint32, location int32, count int32, value *float32) {
	C.glowProgramUniform3fvEXT(gpProgramUniform3fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3i(program uint32, location int32, v0 int32, v1 int32, v2 int32) {
	C.glowProgramUniform3i(gpProgramUniform3i, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2))
}
func ProgramUniform3i64NV(program uint32, location int32, x int64, y int64, z int64) {
	C.glowProgramUniform3i64NV(gpProgramUniform3i64NV, (C.GLuint)(program), (C.GLint)(location), (C.GLint64EXT)(x), (C.GLint64EXT)(y), (C.GLint64EXT)(z))
}
func ProgramUniform3i64vNV(program uint32, location int32, count int32, value *int64) {
	C.glowProgramUniform3i64vNV(gpProgramUniform3i64vNV, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint64EXT)(unsafe.Pointer(value)))
}
func ProgramUniform3iEXT(program uint32, location int32, v0 int32, v1 int32, v2 int32) {
	C.glowProgramUniform3iEXT(gpProgramUniform3iEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3iv(program uint32, location int32, count int32, value *int32) {
	C.glowProgramUniform3iv(gpProgramUniform3iv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}
func ProgramUniform3ivEXT(program uint32, location int32, count int32, value *int32) {
	C.glowProgramUniform3ivEXT(gpProgramUniform3ivEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3ui(program uint32, location int32, v0 uint32, v1 uint32, v2 uint32) {
	C.glowProgramUniform3ui(gpProgramUniform3ui, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2))
}
func ProgramUniform3ui64NV(program uint32, location int32, x uint64, y uint64, z uint64) {
	C.glowProgramUniform3ui64NV(gpProgramUniform3ui64NV, (C.GLuint)(program), (C.GLint)(location), (C.GLuint64EXT)(x), (C.GLuint64EXT)(y), (C.GLuint64EXT)(z))
}
func ProgramUniform3ui64vNV(program uint32, location int32, count int32, value *uint64) {
	C.glowProgramUniform3ui64vNV(gpProgramUniform3ui64vNV, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64EXT)(unsafe.Pointer(value)))
}
func ProgramUniform3uiEXT(program uint32, location int32, v0 uint32, v1 uint32, v2 uint32) {
	C.glowProgramUniform3uiEXT(gpProgramUniform3uiEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3uiv(program uint32, location int32, count int32, value *uint32) {
	C.glowProgramUniform3uiv(gpProgramUniform3uiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func ProgramUniform3uivEXT(program uint32, location int32, count int32, value *uint32) {
	C.glowProgramUniform3uivEXT(gpProgramUniform3uivEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4f(program uint32, location int32, v0 float32, v1 float32, v2 float32, v3 float32) {
	C.glowProgramUniform4f(gpProgramUniform4f, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLfloat)(v3))
}
func ProgramUniform4fEXT(program uint32, location int32, v0 float32, v1 float32, v2 float32, v3 float32) {
	C.glowProgramUniform4fEXT(gpProgramUniform4fEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLfloat)(v3))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4fv(program uint32, location int32, count int32, value *float32) {
	C.glowProgramUniform4fv(gpProgramUniform4fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniform4fvEXT(program uint32, location int32, count int32, value *float32) {
	C.glowProgramUniform4fvEXT(gpProgramUniform4fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4i(program uint32, location int32, v0 int32, v1 int32, v2 int32, v3 int32) {
	C.glowProgramUniform4i(gpProgramUniform4i, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2), (C.GLint)(v3))
}
func ProgramUniform4i64NV(program uint32, location int32, x int64, y int64, z int64, w int64) {
	C.glowProgramUniform4i64NV(gpProgramUniform4i64NV, (C.GLuint)(program), (C.GLint)(location), (C.GLint64EXT)(x), (C.GLint64EXT)(y), (C.GLint64EXT)(z), (C.GLint64EXT)(w))
}
func ProgramUniform4i64vNV(program uint32, location int32, count int32, value *int64) {
	C.glowProgramUniform4i64vNV(gpProgramUniform4i64vNV, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint64EXT)(unsafe.Pointer(value)))
}
func ProgramUniform4iEXT(program uint32, location int32, v0 int32, v1 int32, v2 int32, v3 int32) {
	C.glowProgramUniform4iEXT(gpProgramUniform4iEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2), (C.GLint)(v3))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4iv(program uint32, location int32, count int32, value *int32) {
	C.glowProgramUniform4iv(gpProgramUniform4iv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}
func ProgramUniform4ivEXT(program uint32, location int32, count int32, value *int32) {
	C.glowProgramUniform4ivEXT(gpProgramUniform4ivEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4ui(program uint32, location int32, v0 uint32, v1 uint32, v2 uint32, v3 uint32) {
	C.glowProgramUniform4ui(gpProgramUniform4ui, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2), (C.GLuint)(v3))
}
func ProgramUniform4ui64NV(program uint32, location int32, x uint64, y uint64, z uint64, w uint64) {
	C.glowProgramUniform4ui64NV(gpProgramUniform4ui64NV, (C.GLuint)(program), (C.GLint)(location), (C.GLuint64EXT)(x), (C.GLuint64EXT)(y), (C.GLuint64EXT)(z), (C.GLuint64EXT)(w))
}
func ProgramUniform4ui64vNV(program uint32, location int32, count int32, value *uint64) {
	C.glowProgramUniform4ui64vNV(gpProgramUniform4ui64vNV, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64EXT)(unsafe.Pointer(value)))
}
func ProgramUniform4uiEXT(program uint32, location int32, v0 uint32, v1 uint32, v2 uint32, v3 uint32) {
	C.glowProgramUniform4uiEXT(gpProgramUniform4uiEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2), (C.GLuint)(v3))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4uiv(program uint32, location int32, count int32, value *uint32) {
	C.glowProgramUniform4uiv(gpProgramUniform4uiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func ProgramUniform4uivEXT(program uint32, location int32, count int32, value *uint32) {
	C.glowProgramUniform4uivEXT(gpProgramUniform4uivEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func ProgramUniformHandleui64IMG(program uint32, location int32, value uint64) {
	C.glowProgramUniformHandleui64IMG(gpProgramUniformHandleui64IMG, (C.GLuint)(program), (C.GLint)(location), (C.GLuint64)(value))
}
func ProgramUniformHandleui64NV(program uint32, location int32, value uint64) {
	C.glowProgramUniformHandleui64NV(gpProgramUniformHandleui64NV, (C.GLuint)(program), (C.GLint)(location), (C.GLuint64)(value))
}
func ProgramUniformHandleui64vIMG(program uint32, location int32, count int32, values *uint64) {
	C.glowProgramUniformHandleui64vIMG(gpProgramUniformHandleui64vIMG, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64)(unsafe.Pointer(values)))
}
func ProgramUniformHandleui64vNV(program uint32, location int32, count int32, values *uint64) {
	C.glowProgramUniformHandleui64vNV(gpProgramUniformHandleui64vNV, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64)(unsafe.Pointer(values)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix2fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix2fv(gpProgramUniformMatrix2fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix2fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix2fvEXT(gpProgramUniformMatrix2fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix2x3fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix2x3fv(gpProgramUniformMatrix2x3fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix2x3fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix2x3fvEXT(gpProgramUniformMatrix2x3fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix2x4fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix2x4fv(gpProgramUniformMatrix2x4fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix2x4fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix2x4fvEXT(gpProgramUniformMatrix2x4fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix3fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix3fv(gpProgramUniformMatrix3fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix3fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix3fvEXT(gpProgramUniformMatrix3fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix3x2fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix3x2fv(gpProgramUniformMatrix3x2fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix3x2fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix3x2fvEXT(gpProgramUniformMatrix3x2fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix3x4fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix3x4fv(gpProgramUniformMatrix3x4fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix3x4fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix3x4fvEXT(gpProgramUniformMatrix3x4fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix4fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix4fv(gpProgramUniformMatrix4fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix4fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix4fvEXT(gpProgramUniformMatrix4fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix4x2fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix4x2fv(gpProgramUniformMatrix4x2fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix4x2fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix4x2fvEXT(gpProgramUniformMatrix4x2fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix4x3fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix4x3fv(gpProgramUniformMatrix4x3fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix4x3fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix4x3fvEXT(gpProgramUniformMatrix4x3fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

// push a named debug group into the command stream
func PushDebugGroup(source uint32, id uint32, length int32, message *uint8) {
	C.glowPushDebugGroup(gpPushDebugGroup, (C.GLenum)(source), (C.GLuint)(id), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(message)))
}
func PushDebugGroupKHR(source uint32, id uint32, length int32, message *uint8) {
	C.glowPushDebugGroupKHR(gpPushDebugGroupKHR, (C.GLenum)(source), (C.GLuint)(id), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(message)))
}
func PushGroupMarkerEXT(length int32, marker *uint8) {
	C.glowPushGroupMarkerEXT(gpPushGroupMarkerEXT, (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(marker)))
}
func QueryCounterEXT(id uint32, target uint32) {
	C.glowQueryCounterEXT(gpQueryCounterEXT, (C.GLuint)(id), (C.GLenum)(target))
}
func RasterSamplesEXT(samples uint32, fixedsamplelocations bool) {
	C.glowRasterSamplesEXT(gpRasterSamplesEXT, (C.GLuint)(samples), (C.GLboolean)(boolToInt(fixedsamplelocations)))
}

// select a color buffer source for pixels
func ReadBuffer(src uint32) {
	C.glowReadBuffer(gpReadBuffer, (C.GLenum)(src))
}
func ReadBufferIndexedEXT(src uint32, index int32) {
	C.glowReadBufferIndexedEXT(gpReadBufferIndexedEXT, (C.GLenum)(src), (C.GLint)(index))
}
func ReadBufferNV(mode uint32) {
	C.glowReadBufferNV(gpReadBufferNV, (C.GLenum)(mode))
}

// read a block of pixels from the frame buffer
func ReadPixels(x int32, y int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowReadPixels(gpReadPixels, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// read a block of pixels from the frame buffer
func ReadnPixels(x int32, y int32, width int32, height int32, format uint32, xtype uint32, bufSize int32, data unsafe.Pointer) {
	C.glowReadnPixels(gpReadnPixels, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(xtype), (C.GLsizei)(bufSize), data)
}
func ReadnPixelsEXT(x int32, y int32, width int32, height int32, format uint32, xtype uint32, bufSize int32, data unsafe.Pointer) {
	C.glowReadnPixelsEXT(gpReadnPixelsEXT, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(xtype), (C.GLsizei)(bufSize), data)
}
func ReadnPixelsKHR(x int32, y int32, width int32, height int32, format uint32, xtype uint32, bufSize int32, data unsafe.Pointer) {
	C.glowReadnPixelsKHR(gpReadnPixelsKHR, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(xtype), (C.GLsizei)(bufSize), data)
}
func ReleaseKeyedMutexWin32EXT(memory uint32, key uint64) bool {
	ret := C.glowReleaseKeyedMutexWin32EXT(gpReleaseKeyedMutexWin32EXT, (C.GLuint)(memory), (C.GLuint64)(key))
	return ret == TRUE
}

// release resources consumed by the implementation's shader compiler
func ReleaseShaderCompiler() {
	C.glowReleaseShaderCompiler(gpReleaseShaderCompiler)
}

// establish data storage, format and dimensions of a     renderbuffer object's image
func RenderbufferStorage(target uint32, internalformat uint32, width int32, height int32) {
	C.glowRenderbufferStorage(gpRenderbufferStorage, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}

// establish data storage, format, dimensions and sample count of     a renderbuffer object's image
func RenderbufferStorageMultisample(target uint32, samples int32, internalformat uint32, width int32, height int32) {
	C.glowRenderbufferStorageMultisample(gpRenderbufferStorageMultisample, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
func RenderbufferStorageMultisampleANGLE(target uint32, samples int32, internalformat uint32, width int32, height int32) {
	C.glowRenderbufferStorageMultisampleANGLE(gpRenderbufferStorageMultisampleANGLE, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
func RenderbufferStorageMultisampleAPPLE(target uint32, samples int32, internalformat uint32, width int32, height int32) {
	C.glowRenderbufferStorageMultisampleAPPLE(gpRenderbufferStorageMultisampleAPPLE, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
func RenderbufferStorageMultisampleEXT(target uint32, samples int32, internalformat uint32, width int32, height int32) {
	C.glowRenderbufferStorageMultisampleEXT(gpRenderbufferStorageMultisampleEXT, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
func RenderbufferStorageMultisampleIMG(target uint32, samples int32, internalformat uint32, width int32, height int32) {
	C.glowRenderbufferStorageMultisampleIMG(gpRenderbufferStorageMultisampleIMG, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
func RenderbufferStorageMultisampleNV(target uint32, samples int32, internalformat uint32, width int32, height int32) {
	C.glowRenderbufferStorageMultisampleNV(gpRenderbufferStorageMultisampleNV, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
func ResolveDepthValuesNV() {
	C.glowResolveDepthValuesNV(gpResolveDepthValuesNV)
}
func ResolveMultisampleFramebufferAPPLE() {
	C.glowResolveMultisampleFramebufferAPPLE(gpResolveMultisampleFramebufferAPPLE)
}

// resume transform feedback operations
func ResumeTransformFeedback() {
	C.glowResumeTransformFeedback(gpResumeTransformFeedback)
}

// specify multisample coverage parameters
func SampleCoverage(value float32, invert bool) {
	C.glowSampleCoverage(gpSampleCoverage, (C.GLfloat)(value), (C.GLboolean)(boolToInt(invert)))
}

// set the value of a sub-word of the sample mask
func SampleMaski(maskNumber uint32, mask uint32) {
	C.glowSampleMaski(gpSampleMaski, (C.GLuint)(maskNumber), (C.GLbitfield)(mask))
}
func SamplerParameterIivEXT(sampler uint32, pname uint32, param *int32) {
	C.glowSamplerParameterIivEXT(gpSamplerParameterIivEXT, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}
func SamplerParameterIivOES(sampler uint32, pname uint32, param *int32) {
	C.glowSamplerParameterIivOES(gpSamplerParameterIivOES, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}
func SamplerParameterIuivEXT(sampler uint32, pname uint32, param *uint32) {
	C.glowSamplerParameterIuivEXT(gpSamplerParameterIuivEXT, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(param)))
}
func SamplerParameterIuivOES(sampler uint32, pname uint32, param *uint32) {
	C.glowSamplerParameterIuivOES(gpSamplerParameterIuivOES, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(param)))
}
func SamplerParameterf(sampler uint32, pname uint32, param float32) {
	C.glowSamplerParameterf(gpSamplerParameterf, (C.GLuint)(sampler), (C.GLenum)(pname), (C.GLfloat)(param))
}
func SamplerParameterfv(sampler uint32, pname uint32, param *float32) {
	C.glowSamplerParameterfv(gpSamplerParameterfv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(param)))
}
func SamplerParameteri(sampler uint32, pname uint32, param int32) {
	C.glowSamplerParameteri(gpSamplerParameteri, (C.GLuint)(sampler), (C.GLenum)(pname), (C.GLint)(param))
}
func SamplerParameteriv(sampler uint32, pname uint32, param *int32) {
	C.glowSamplerParameteriv(gpSamplerParameteriv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}

// define the scissor box
func Scissor(x int32, y int32, width int32, height int32) {
	C.glowScissor(gpScissor, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
func ScissorArrayvNV(first uint32, count int32, v *int32) {
	C.glowScissorArrayvNV(gpScissorArrayvNV, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(v)))
}
func ScissorArrayvOES(first uint32, count int32, v *int32) {
	C.glowScissorArrayvOES(gpScissorArrayvOES, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(v)))
}
func ScissorIndexedNV(index uint32, left int32, bottom int32, width int32, height int32) {
	C.glowScissorIndexedNV(gpScissorIndexedNV, (C.GLuint)(index), (C.GLint)(left), (C.GLint)(bottom), (C.GLsizei)(width), (C.GLsizei)(height))
}
func ScissorIndexedOES(index uint32, left int32, bottom int32, width int32, height int32) {
	C.glowScissorIndexedOES(gpScissorIndexedOES, (C.GLuint)(index), (C.GLint)(left), (C.GLint)(bottom), (C.GLsizei)(width), (C.GLsizei)(height))
}
func ScissorIndexedvNV(index uint32, v *int32) {
	C.glowScissorIndexedvNV(gpScissorIndexedvNV, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func ScissorIndexedvOES(index uint32, v *int32) {
	C.glowScissorIndexedvOES(gpScissorIndexedvOES, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func SelectPerfMonitorCountersAMD(monitor uint32, enable bool, group uint32, numCounters int32, counterList *uint32) {
	C.glowSelectPerfMonitorCountersAMD(gpSelectPerfMonitorCountersAMD, (C.GLuint)(monitor), (C.GLboolean)(boolToInt(enable)), (C.GLuint)(group), (C.GLint)(numCounters), (*C.GLuint)(unsafe.Pointer(counterList)))
}
func SemaphoreParameterui64vEXT(semaphore uint32, pname uint32, params *uint64) {
	C.glowSemaphoreParameterui64vEXT(gpSemaphoreParameterui64vEXT, (C.GLuint)(semaphore), (C.GLenum)(pname), (*C.GLuint64)(unsafe.Pointer(params)))
}
func SetFenceNV(fence uint32, condition uint32) {
	C.glowSetFenceNV(gpSetFenceNV, (C.GLuint)(fence), (C.GLenum)(condition))
}

// load pre-compiled shader binaries
func ShaderBinary(count int32, shaders *uint32, binaryformat uint32, binary unsafe.Pointer, length int32) {
	C.glowShaderBinary(gpShaderBinary, (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(shaders)), (C.GLenum)(binaryformat), binary, (C.GLsizei)(length))
}

// Replaces the source code in a shader object
func ShaderSource(shader uint32, count int32, xstring **uint8, length *int32) {
	C.glowShaderSource(gpShaderSource, (C.GLuint)(shader), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(xstring)), (*C.GLint)(unsafe.Pointer(length)))
}
func SignalSemaphoreEXT(semaphore uint32, numBufferBarriers uint32, buffers *uint32, numTextureBarriers uint32, textures *uint32, dstLayouts *uint32) {
	C.glowSignalSemaphoreEXT(gpSignalSemaphoreEXT, (C.GLuint)(semaphore), (C.GLuint)(numBufferBarriers), (*C.GLuint)(unsafe.Pointer(buffers)), (C.GLuint)(numTextureBarriers), (*C.GLuint)(unsafe.Pointer(textures)), (*C.GLenum)(unsafe.Pointer(dstLayouts)))
}
func SignalVkFenceNV(vkFence uint64) {
	C.glowSignalVkFenceNV(gpSignalVkFenceNV, (C.GLuint64)(vkFence))
}
func SignalVkSemaphoreNV(vkSemaphore uint64) {
	C.glowSignalVkSemaphoreNV(gpSignalVkSemaphoreNV, (C.GLuint64)(vkSemaphore))
}
func StartTilingQCOM(x uint32, y uint32, width uint32, height uint32, preserveMask uint32) {
	C.glowStartTilingQCOM(gpStartTilingQCOM, (C.GLuint)(x), (C.GLuint)(y), (C.GLuint)(width), (C.GLuint)(height), (C.GLbitfield)(preserveMask))
}
func StencilFillPathInstancedNV(numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, fillMode uint32, mask uint32, transformType uint32, transformValues *float32) {
	C.glowStencilFillPathInstancedNV(gpStencilFillPathInstancedNV, (C.GLsizei)(numPaths), (C.GLenum)(pathNameType), paths, (C.GLuint)(pathBase), (C.GLenum)(fillMode), (C.GLuint)(mask), (C.GLenum)(transformType), (*C.GLfloat)(unsafe.Pointer(transformValues)))
}
func StencilFillPathNV(path uint32, fillMode uint32, mask uint32) {
	C.glowStencilFillPathNV(gpStencilFillPathNV, (C.GLuint)(path), (C.GLenum)(fillMode), (C.GLuint)(mask))
}

// set front and back function and reference value for stencil testing
func StencilFunc(xfunc uint32, ref int32, mask uint32) {
	C.glowStencilFunc(gpStencilFunc, (C.GLenum)(xfunc), (C.GLint)(ref), (C.GLuint)(mask))
}

// set front and/or back function and reference value for stencil testing
func StencilFuncSeparate(face uint32, xfunc uint32, ref int32, mask uint32) {
	C.glowStencilFuncSeparate(gpStencilFuncSeparate, (C.GLenum)(face), (C.GLenum)(xfunc), (C.GLint)(ref), (C.GLuint)(mask))
}

// control the front and back writing of individual bits in the stencil planes
func StencilMask(mask uint32) {
	C.glowStencilMask(gpStencilMask, (C.GLuint)(mask))
}

// control the front and/or back writing of individual bits in the stencil planes
func StencilMaskSeparate(face uint32, mask uint32) {
	C.glowStencilMaskSeparate(gpStencilMaskSeparate, (C.GLenum)(face), (C.GLuint)(mask))
}

// set front and back stencil test actions
func StencilOp(fail uint32, zfail uint32, zpass uint32) {
	C.glowStencilOp(gpStencilOp, (C.GLenum)(fail), (C.GLenum)(zfail), (C.GLenum)(zpass))
}

// set front and/or back stencil test actions
func StencilOpSeparate(face uint32, sfail uint32, dpfail uint32, dppass uint32) {
	C.glowStencilOpSeparate(gpStencilOpSeparate, (C.GLenum)(face), (C.GLenum)(sfail), (C.GLenum)(dpfail), (C.GLenum)(dppass))
}
func StencilStrokePathInstancedNV(numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, reference int32, mask uint32, transformType uint32, transformValues *float32) {
	C.glowStencilStrokePathInstancedNV(gpStencilStrokePathInstancedNV, (C.GLsizei)(numPaths), (C.GLenum)(pathNameType), paths, (C.GLuint)(pathBase), (C.GLint)(reference), (C.GLuint)(mask), (C.GLenum)(transformType), (*C.GLfloat)(unsafe.Pointer(transformValues)))
}
func StencilStrokePathNV(path uint32, reference int32, mask uint32) {
	C.glowStencilStrokePathNV(gpStencilStrokePathNV, (C.GLuint)(path), (C.GLint)(reference), (C.GLuint)(mask))
}
func StencilThenCoverFillPathInstancedNV(numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, fillMode uint32, mask uint32, coverMode uint32, transformType uint32, transformValues *float32) {
	C.glowStencilThenCoverFillPathInstancedNV(gpStencilThenCoverFillPathInstancedNV, (C.GLsizei)(numPaths), (C.GLenum)(pathNameType), paths, (C.GLuint)(pathBase), (C.GLenum)(fillMode), (C.GLuint)(mask), (C.GLenum)(coverMode), (C.GLenum)(transformType), (*C.GLfloat)(unsafe.Pointer(transformValues)))
}
func StencilThenCoverFillPathNV(path uint32, fillMode uint32, mask uint32, coverMode uint32) {
	C.glowStencilThenCoverFillPathNV(gpStencilThenCoverFillPathNV, (C.GLuint)(path), (C.GLenum)(fillMode), (C.GLuint)(mask), (C.GLenum)(coverMode))
}
func StencilThenCoverStrokePathInstancedNV(numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, reference int32, mask uint32, coverMode uint32, transformType uint32, transformValues *float32) {
	C.glowStencilThenCoverStrokePathInstancedNV(gpStencilThenCoverStrokePathInstancedNV, (C.GLsizei)(numPaths), (C.GLenum)(pathNameType), paths, (C.GLuint)(pathBase), (C.GLint)(reference), (C.GLuint)(mask), (C.GLenum)(coverMode), (C.GLenum)(transformType), (*C.GLfloat)(unsafe.Pointer(transformValues)))
}
func StencilThenCoverStrokePathNV(path uint32, reference int32, mask uint32, coverMode uint32) {
	C.glowStencilThenCoverStrokePathNV(gpStencilThenCoverStrokePathNV, (C.GLuint)(path), (C.GLint)(reference), (C.GLuint)(mask), (C.GLenum)(coverMode))
}
func SubpixelPrecisionBiasNV(xbits uint32, ybits uint32) {
	C.glowSubpixelPrecisionBiasNV(gpSubpixelPrecisionBiasNV, (C.GLuint)(xbits), (C.GLuint)(ybits))
}
func TestFenceNV(fence uint32) bool {
	ret := C.glowTestFenceNV(gpTestFenceNV, (C.GLuint)(fence))
	return ret == TRUE
}
func TexBufferEXT(target uint32, internalformat uint32, buffer uint32) {
	C.glowTexBufferEXT(gpTexBufferEXT, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLuint)(buffer))
}
func TexBufferOES(target uint32, internalformat uint32, buffer uint32) {
	C.glowTexBufferOES(gpTexBufferOES, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLuint)(buffer))
}
func TexBufferRangeEXT(target uint32, internalformat uint32, buffer uint32, offset int, size int) {
	C.glowTexBufferRangeEXT(gpTexBufferRangeEXT, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(size))
}
func TexBufferRangeOES(target uint32, internalformat uint32, buffer uint32, offset int, size int) {
	C.glowTexBufferRangeOES(gpTexBufferRangeOES, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(size))
}

// specify a two-dimensional texture image
func TexImage2D(target uint32, level int32, internalformat int32, width int32, height int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTexImage2D(gpTexImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// specify a three-dimensional texture image
func TexImage3D(target uint32, level int32, internalformat int32, width int32, height int32, depth int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTexImage3D(gpTexImage3D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}
func TexImage3DOES(target uint32, level int32, internalformat uint32, width int32, height int32, depth int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTexImage3DOES(gpTexImage3DOES, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}
func TexPageCommitmentEXT(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, commit bool) {
	C.glowTexPageCommitmentEXT(gpTexPageCommitmentEXT, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(boolToInt(commit)))
}
func TexParameterIivEXT(target uint32, pname uint32, params *int32) {
	C.glowTexParameterIivEXT(gpTexParameterIivEXT, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func TexParameterIivOES(target uint32, pname uint32, params *int32) {
	C.glowTexParameterIivOES(gpTexParameterIivOES, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func TexParameterIuivEXT(target uint32, pname uint32, params *uint32) {
	C.glowTexParameterIuivEXT(gpTexParameterIuivEXT, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func TexParameterIuivOES(target uint32, pname uint32, params *uint32) {
	C.glowTexParameterIuivOES(gpTexParameterIuivOES, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func TexParameterf(target uint32, pname uint32, param float32) {
	C.glowTexParameterf(gpTexParameterf, (C.GLenum)(target), (C.GLenum)(pname), (C.GLfloat)(param))
}
func TexParameterfv(target uint32, pname uint32, params *float32) {
	C.glowTexParameterfv(gpTexParameterfv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func TexParameteri(target uint32, pname uint32, param int32) {
	C.glowTexParameteri(gpTexParameteri, (C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}
func TexParameteriv(target uint32, pname uint32, params *int32) {
	C.glowTexParameteriv(gpTexParameteriv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func TexStorage1DEXT(target uint32, levels int32, internalformat uint32, width int32) {
	C.glowTexStorage1DEXT(gpTexStorage1DEXT, (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width))
}

// simultaneously specify storage for all levels of a two-dimensional or one-dimensional array texture
func TexStorage2D(target uint32, levels int32, internalformat uint32, width int32, height int32) {
	C.glowTexStorage2D(gpTexStorage2D, (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
func TexStorage2DEXT(target uint32, levels int32, internalformat uint32, width int32, height int32) {
	C.glowTexStorage2DEXT(gpTexStorage2DEXT, (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}

// specify storage for a two-dimensional multisample texture
func TexStorage2DMultisample(target uint32, samples int32, internalformat uint32, width int32, height int32, fixedsamplelocations bool) {
	C.glowTexStorage2DMultisample(gpTexStorage2DMultisample, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLboolean)(boolToInt(fixedsamplelocations)))
}

// simultaneously specify storage for all levels of a three-dimensional, two-dimensional array or cube-map array texture
func TexStorage3D(target uint32, levels int32, internalformat uint32, width int32, height int32, depth int32) {
	C.glowTexStorage3D(gpTexStorage3D, (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth))
}
func TexStorage3DEXT(target uint32, levels int32, internalformat uint32, width int32, height int32, depth int32) {
	C.glowTexStorage3DEXT(gpTexStorage3DEXT, (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth))
}
func TexStorage3DMultisampleOES(target uint32, samples int32, internalformat uint32, width int32, height int32, depth int32, fixedsamplelocations bool) {
	C.glowTexStorage3DMultisampleOES(gpTexStorage3DMultisampleOES, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(boolToInt(fixedsamplelocations)))
}
func TexStorageMem1DEXT(target uint32, levels int32, internalFormat uint32, width int32, memory uint32, offset uint64) {
	C.glowTexStorageMem1DEXT(gpTexStorageMem1DEXT, (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalFormat), (C.GLsizei)(width), (C.GLuint)(memory), (C.GLuint64)(offset))
}
func TexStorageMem2DEXT(target uint32, levels int32, internalFormat uint32, width int32, height int32, memory uint32, offset uint64) {
	C.glowTexStorageMem2DEXT(gpTexStorageMem2DEXT, (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalFormat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLuint)(memory), (C.GLuint64)(offset))
}
func TexStorageMem2DMultisampleEXT(target uint32, samples int32, internalFormat uint32, width int32, height int32, fixedSampleLocations bool, memory uint32, offset uint64) {
	C.glowTexStorageMem2DMultisampleEXT(gpTexStorageMem2DMultisampleEXT, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalFormat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLboolean)(boolToInt(fixedSampleLocations)), (C.GLuint)(memory), (C.GLuint64)(offset))
}
func TexStorageMem3DEXT(target uint32, levels int32, internalFormat uint32, width int32, height int32, depth int32, memory uint32, offset uint64) {
	C.glowTexStorageMem3DEXT(gpTexStorageMem3DEXT, (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalFormat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLuint)(memory), (C.GLuint64)(offset))
}
func TexStorageMem3DMultisampleEXT(target uint32, samples int32, internalFormat uint32, width int32, height int32, depth int32, fixedSampleLocations bool, memory uint32, offset uint64) {
	C.glowTexStorageMem3DMultisampleEXT(gpTexStorageMem3DMultisampleEXT, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalFormat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(boolToInt(fixedSampleLocations)), (C.GLuint)(memory), (C.GLuint64)(offset))
}

// specify a two-dimensional texture subimage
func TexSubImage2D(target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTexSubImage2D(gpTexSubImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// specify a three-dimensional texture subimage
func TexSubImage3D(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTexSubImage3D(gpTexSubImage3D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}
func TexSubImage3DOES(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTexSubImage3DOES(gpTexSubImage3DOES, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}
func TextureFoveationParametersQCOM(texture uint32, layer uint32, focalPoint uint32, focalX float32, focalY float32, gainX float32, gainY float32, foveaArea float32) {
	C.glowTextureFoveationParametersQCOM(gpTextureFoveationParametersQCOM, (C.GLuint)(texture), (C.GLuint)(layer), (C.GLuint)(focalPoint), (C.GLfloat)(focalX), (C.GLfloat)(focalY), (C.GLfloat)(gainX), (C.GLfloat)(gainY), (C.GLfloat)(foveaArea))
}
func TextureStorage1DEXT(texture uint32, target uint32, levels int32, internalformat uint32, width int32) {
	C.glowTextureStorage1DEXT(gpTextureStorage1DEXT, (C.GLuint)(texture), (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width))
}
func TextureStorage2DEXT(texture uint32, target uint32, levels int32, internalformat uint32, width int32, height int32) {
	C.glowTextureStorage2DEXT(gpTextureStorage2DEXT, (C.GLuint)(texture), (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
func TextureStorage3DEXT(texture uint32, target uint32, levels int32, internalformat uint32, width int32, height int32, depth int32) {
	C.glowTextureStorage3DEXT(gpTextureStorage3DEXT, (C.GLuint)(texture), (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth))
}
func TextureStorageMem1DEXT(texture uint32, levels int32, internalFormat uint32, width int32, memory uint32, offset uint64) {
	C.glowTextureStorageMem1DEXT(gpTextureStorageMem1DEXT, (C.GLuint)(texture), (C.GLsizei)(levels), (C.GLenum)(internalFormat), (C.GLsizei)(width), (C.GLuint)(memory), (C.GLuint64)(offset))
}
func TextureStorageMem2DEXT(texture uint32, levels int32, internalFormat uint32, width int32, height int32, memory uint32, offset uint64) {
	C.glowTextureStorageMem2DEXT(gpTextureStorageMem2DEXT, (C.GLuint)(texture), (C.GLsizei)(levels), (C.GLenum)(internalFormat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLuint)(memory), (C.GLuint64)(offset))
}
func TextureStorageMem2DMultisampleEXT(texture uint32, samples int32, internalFormat uint32, width int32, height int32, fixedSampleLocations bool, memory uint32, offset uint64) {
	C.glowTextureStorageMem2DMultisampleEXT(gpTextureStorageMem2DMultisampleEXT, (C.GLuint)(texture), (C.GLsizei)(samples), (C.GLenum)(internalFormat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLboolean)(boolToInt(fixedSampleLocations)), (C.GLuint)(memory), (C.GLuint64)(offset))
}
func TextureStorageMem3DEXT(texture uint32, levels int32, internalFormat uint32, width int32, height int32, depth int32, memory uint32, offset uint64) {
	C.glowTextureStorageMem3DEXT(gpTextureStorageMem3DEXT, (C.GLuint)(texture), (C.GLsizei)(levels), (C.GLenum)(internalFormat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLuint)(memory), (C.GLuint64)(offset))
}
func TextureStorageMem3DMultisampleEXT(texture uint32, samples int32, internalFormat uint32, width int32, height int32, depth int32, fixedSampleLocations bool, memory uint32, offset uint64) {
	C.glowTextureStorageMem3DMultisampleEXT(gpTextureStorageMem3DMultisampleEXT, (C.GLuint)(texture), (C.GLsizei)(samples), (C.GLenum)(internalFormat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(boolToInt(fixedSampleLocations)), (C.GLuint)(memory), (C.GLuint64)(offset))
}
func TextureViewEXT(texture uint32, target uint32, origtexture uint32, internalformat uint32, minlevel uint32, numlevels uint32, minlayer uint32, numlayers uint32) {
	C.glowTextureViewEXT(gpTextureViewEXT, (C.GLuint)(texture), (C.GLenum)(target), (C.GLuint)(origtexture), (C.GLenum)(internalformat), (C.GLuint)(minlevel), (C.GLuint)(numlevels), (C.GLuint)(minlayer), (C.GLuint)(numlayers))
}
func TextureViewOES(texture uint32, target uint32, origtexture uint32, internalformat uint32, minlevel uint32, numlevels uint32, minlayer uint32, numlayers uint32) {
	C.glowTextureViewOES(gpTextureViewOES, (C.GLuint)(texture), (C.GLenum)(target), (C.GLuint)(origtexture), (C.GLenum)(internalformat), (C.GLuint)(minlevel), (C.GLuint)(numlevels), (C.GLuint)(minlayer), (C.GLuint)(numlayers))
}

// specify values to record in transform feedback buffers
func TransformFeedbackVaryings(program uint32, count int32, varyings **uint8, bufferMode uint32) {
	C.glowTransformFeedbackVaryings(gpTransformFeedbackVaryings, (C.GLuint)(program), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(varyings)), (C.GLenum)(bufferMode))
}
func TransformPathNV(resultPath uint32, srcPath uint32, transformType uint32, transformValues *float32) {
	C.glowTransformPathNV(gpTransformPathNV, (C.GLuint)(resultPath), (C.GLuint)(srcPath), (C.GLenum)(transformType), (*C.GLfloat)(unsafe.Pointer(transformValues)))
}

// Specify the value of a uniform variable for the current program object
func Uniform1f(location int32, v0 float32) {
	C.glowUniform1f(gpUniform1f, (C.GLint)(location), (C.GLfloat)(v0))
}

// Specify the value of a uniform variable for the current program object
func Uniform1fv(location int32, count int32, value *float32) {
	C.glowUniform1fv(gpUniform1fv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform1i(location int32, v0 int32) {
	C.glowUniform1i(gpUniform1i, (C.GLint)(location), (C.GLint)(v0))
}
func Uniform1i64NV(location int32, x int64) {
	C.glowUniform1i64NV(gpUniform1i64NV, (C.GLint)(location), (C.GLint64EXT)(x))
}
func Uniform1i64vNV(location int32, count int32, value *int64) {
	C.glowUniform1i64vNV(gpUniform1i64vNV, (C.GLint)(location), (C.GLsizei)(count), (*C.GLint64EXT)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform1iv(location int32, count int32, value *int32) {
	C.glowUniform1iv(gpUniform1iv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform1ui(location int32, v0 uint32) {
	C.glowUniform1ui(gpUniform1ui, (C.GLint)(location), (C.GLuint)(v0))
}
func Uniform1ui64NV(location int32, x uint64) {
	C.glowUniform1ui64NV(gpUniform1ui64NV, (C.GLint)(location), (C.GLuint64EXT)(x))
}
func Uniform1ui64vNV(location int32, count int32, value *uint64) {
	C.glowUniform1ui64vNV(gpUniform1ui64vNV, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64EXT)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform1uiv(location int32, count int32, value *uint32) {
	C.glowUniform1uiv(gpUniform1uiv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform2f(location int32, v0 float32, v1 float32) {
	C.glowUniform2f(gpUniform2f, (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1))
}

// Specify the value of a uniform variable for the current program object
func Uniform2fv(location int32, count int32, value *float32) {
	C.glowUniform2fv(gpUniform2fv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform2i(location int32, v0 int32, v1 int32) {
	C.glowUniform2i(gpUniform2i, (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1))
}
func Uniform2i64NV(location int32, x int64, y int64) {
	C.glowUniform2i64NV(gpUniform2i64NV, (C.GLint)(location), (C.GLint64EXT)(x), (C.GLint64EXT)(y))
}
func Uniform2i64vNV(location int32, count int32, value *int64) {
	C.glowUniform2i64vNV(gpUniform2i64vNV, (C.GLint)(location), (C.GLsizei)(count), (*C.GLint64EXT)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform2iv(location int32, count int32, value *int32) {
	C.glowUniform2iv(gpUniform2iv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform2ui(location int32, v0 uint32, v1 uint32) {
	C.glowUniform2ui(gpUniform2ui, (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1))
}
func Uniform2ui64NV(location int32, x uint64, y uint64) {
	C.glowUniform2ui64NV(gpUniform2ui64NV, (C.GLint)(location), (C.GLuint64EXT)(x), (C.GLuint64EXT)(y))
}
func Uniform2ui64vNV(location int32, count int32, value *uint64) {
	C.glowUniform2ui64vNV(gpUniform2ui64vNV, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64EXT)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform2uiv(location int32, count int32, value *uint32) {
	C.glowUniform2uiv(gpUniform2uiv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform3f(location int32, v0 float32, v1 float32, v2 float32) {
	C.glowUniform3f(gpUniform3f, (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2))
}

// Specify the value of a uniform variable for the current program object
func Uniform3fv(location int32, count int32, value *float32) {
	C.glowUniform3fv(gpUniform3fv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform3i(location int32, v0 int32, v1 int32, v2 int32) {
	C.glowUniform3i(gpUniform3i, (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2))
}
func Uniform3i64NV(location int32, x int64, y int64, z int64) {
	C.glowUniform3i64NV(gpUniform3i64NV, (C.GLint)(location), (C.GLint64EXT)(x), (C.GLint64EXT)(y), (C.GLint64EXT)(z))
}
func Uniform3i64vNV(location int32, count int32, value *int64) {
	C.glowUniform3i64vNV(gpUniform3i64vNV, (C.GLint)(location), (C.GLsizei)(count), (*C.GLint64EXT)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform3iv(location int32, count int32, value *int32) {
	C.glowUniform3iv(gpUniform3iv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform3ui(location int32, v0 uint32, v1 uint32, v2 uint32) {
	C.glowUniform3ui(gpUniform3ui, (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2))
}
func Uniform3ui64NV(location int32, x uint64, y uint64, z uint64) {
	C.glowUniform3ui64NV(gpUniform3ui64NV, (C.GLint)(location), (C.GLuint64EXT)(x), (C.GLuint64EXT)(y), (C.GLuint64EXT)(z))
}
func Uniform3ui64vNV(location int32, count int32, value *uint64) {
	C.glowUniform3ui64vNV(gpUniform3ui64vNV, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64EXT)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform3uiv(location int32, count int32, value *uint32) {
	C.glowUniform3uiv(gpUniform3uiv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform4f(location int32, v0 float32, v1 float32, v2 float32, v3 float32) {
	C.glowUniform4f(gpUniform4f, (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLfloat)(v3))
}

// Specify the value of a uniform variable for the current program object
func Uniform4fv(location int32, count int32, value *float32) {
	C.glowUniform4fv(gpUniform4fv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform4i(location int32, v0 int32, v1 int32, v2 int32, v3 int32) {
	C.glowUniform4i(gpUniform4i, (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2), (C.GLint)(v3))
}
func Uniform4i64NV(location int32, x int64, y int64, z int64, w int64) {
	C.glowUniform4i64NV(gpUniform4i64NV, (C.GLint)(location), (C.GLint64EXT)(x), (C.GLint64EXT)(y), (C.GLint64EXT)(z), (C.GLint64EXT)(w))
}
func Uniform4i64vNV(location int32, count int32, value *int64) {
	C.glowUniform4i64vNV(gpUniform4i64vNV, (C.GLint)(location), (C.GLsizei)(count), (*C.GLint64EXT)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform4iv(location int32, count int32, value *int32) {
	C.glowUniform4iv(gpUniform4iv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform4ui(location int32, v0 uint32, v1 uint32, v2 uint32, v3 uint32) {
	C.glowUniform4ui(gpUniform4ui, (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2), (C.GLuint)(v3))
}
func Uniform4ui64NV(location int32, x uint64, y uint64, z uint64, w uint64) {
	C.glowUniform4ui64NV(gpUniform4ui64NV, (C.GLint)(location), (C.GLuint64EXT)(x), (C.GLuint64EXT)(y), (C.GLuint64EXT)(z), (C.GLuint64EXT)(w))
}
func Uniform4ui64vNV(location int32, count int32, value *uint64) {
	C.glowUniform4ui64vNV(gpUniform4ui64vNV, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64EXT)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform4uiv(location int32, count int32, value *uint32) {
	C.glowUniform4uiv(gpUniform4uiv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}

// assign a binding point to an active uniform block
func UniformBlockBinding(program uint32, uniformBlockIndex uint32, uniformBlockBinding uint32) {
	C.glowUniformBlockBinding(gpUniformBlockBinding, (C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLuint)(uniformBlockBinding))
}
func UniformHandleui64IMG(location int32, value uint64) {
	C.glowUniformHandleui64IMG(gpUniformHandleui64IMG, (C.GLint)(location), (C.GLuint64)(value))
}
func UniformHandleui64NV(location int32, value uint64) {
	C.glowUniformHandleui64NV(gpUniformHandleui64NV, (C.GLint)(location), (C.GLuint64)(value))
}
func UniformHandleui64vIMG(location int32, count int32, value *uint64) {
	C.glowUniformHandleui64vIMG(gpUniformHandleui64vIMG, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64)(unsafe.Pointer(value)))
}
func UniformHandleui64vNV(location int32, count int32, value *uint64) {
	C.glowUniformHandleui64vNV(gpUniformHandleui64vNV, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix2fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix2fv(gpUniformMatrix2fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix2x3fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix2x3fv(gpUniformMatrix2x3fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix2x3fvNV(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix2x3fvNV(gpUniformMatrix2x3fvNV, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix2x4fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix2x4fv(gpUniformMatrix2x4fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix2x4fvNV(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix2x4fvNV(gpUniformMatrix2x4fvNV, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix3fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix3fv(gpUniformMatrix3fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix3x2fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix3x2fv(gpUniformMatrix3x2fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix3x2fvNV(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix3x2fvNV(gpUniformMatrix3x2fvNV, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix3x4fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix3x4fv(gpUniformMatrix3x4fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix3x4fvNV(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix3x4fvNV(gpUniformMatrix3x4fvNV, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix4fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix4fv(gpUniformMatrix4fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix4x2fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix4x2fv(gpUniformMatrix4x2fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix4x2fvNV(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix4x2fvNV(gpUniformMatrix4x2fvNV, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix4x3fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix4x3fv(gpUniformMatrix4x3fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix4x3fvNV(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix4x3fvNV(gpUniformMatrix4x3fvNV, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

// release the mapping of a buffer object's data store into the client's address space
func UnmapBuffer(target uint32) bool {
	ret := C.glowUnmapBuffer(gpUnmapBuffer, (C.GLenum)(target))
	return ret == TRUE
}
func UnmapBufferOES(target uint32) bool {
	ret := C.glowUnmapBufferOES(gpUnmapBufferOES, (C.GLenum)(target))
	return ret == TRUE
}

// Installs a program object as part of current rendering state
func UseProgram(program uint32) {
	C.glowUseProgram(gpUseProgram, (C.GLuint)(program))
}

// bind stages of a program object to a program pipeline
func UseProgramStages(pipeline uint32, stages uint32, program uint32) {
	C.glowUseProgramStages(gpUseProgramStages, (C.GLuint)(pipeline), (C.GLbitfield)(stages), (C.GLuint)(program))
}
func UseProgramStagesEXT(pipeline uint32, stages uint32, program uint32) {
	C.glowUseProgramStagesEXT(gpUseProgramStagesEXT, (C.GLuint)(pipeline), (C.GLbitfield)(stages), (C.GLuint)(program))
}
func UseShaderProgramEXT(xtype uint32, program uint32) {
	C.glowUseShaderProgramEXT(gpUseShaderProgramEXT, (C.GLenum)(xtype), (C.GLuint)(program))
}

// Validates a program object
func ValidateProgram(program uint32) {
	C.glowValidateProgram(gpValidateProgram, (C.GLuint)(program))
}

// validate a program pipeline object against current GL state
func ValidateProgramPipeline(pipeline uint32) {
	C.glowValidateProgramPipeline(gpValidateProgramPipeline, (C.GLuint)(pipeline))
}
func ValidateProgramPipelineEXT(pipeline uint32) {
	C.glowValidateProgramPipelineEXT(gpValidateProgramPipelineEXT, (C.GLuint)(pipeline))
}
func VertexAttrib1f(index uint32, x float32) {
	C.glowVertexAttrib1f(gpVertexAttrib1f, (C.GLuint)(index), (C.GLfloat)(x))
}
func VertexAttrib1fv(index uint32, v *float32) {
	C.glowVertexAttrib1fv(gpVertexAttrib1fv, (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(v)))
}
func VertexAttrib2f(index uint32, x float32, y float32) {
	C.glowVertexAttrib2f(gpVertexAttrib2f, (C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y))
}
func VertexAttrib2fv(index uint32, v *float32) {
	C.glowVertexAttrib2fv(gpVertexAttrib2fv, (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(v)))
}
func VertexAttrib3f(index uint32, x float32, y float32, z float32) {
	C.glowVertexAttrib3f(gpVertexAttrib3f, (C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func VertexAttrib3fv(index uint32, v *float32) {
	C.glowVertexAttrib3fv(gpVertexAttrib3fv, (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(v)))
}
func VertexAttrib4f(index uint32, x float32, y float32, z float32, w float32) {
	C.glowVertexAttrib4f(gpVertexAttrib4f, (C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
func VertexAttrib4fv(index uint32, v *float32) {
	C.glowVertexAttrib4fv(gpVertexAttrib4fv, (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(v)))
}

// associate a vertex attribute and a vertex buffer binding for a vertex array object
func VertexAttribBinding(attribindex uint32, bindingindex uint32) {
	C.glowVertexAttribBinding(gpVertexAttribBinding, (C.GLuint)(attribindex), (C.GLuint)(bindingindex))
}

// modify the rate at which generic vertex attributes advance during instanced rendering
func VertexAttribDivisor(index uint32, divisor uint32) {
	C.glowVertexAttribDivisor(gpVertexAttribDivisor, (C.GLuint)(index), (C.GLuint)(divisor))
}
func VertexAttribDivisorANGLE(index uint32, divisor uint32) {
	C.glowVertexAttribDivisorANGLE(gpVertexAttribDivisorANGLE, (C.GLuint)(index), (C.GLuint)(divisor))
}
func VertexAttribDivisorEXT(index uint32, divisor uint32) {
	C.glowVertexAttribDivisorEXT(gpVertexAttribDivisorEXT, (C.GLuint)(index), (C.GLuint)(divisor))
}
func VertexAttribDivisorNV(index uint32, divisor uint32) {
	C.glowVertexAttribDivisorNV(gpVertexAttribDivisorNV, (C.GLuint)(index), (C.GLuint)(divisor))
}

// specify the organization of vertex arrays
func VertexAttribFormat(attribindex uint32, size int32, xtype uint32, normalized bool, relativeoffset uint32) {
	C.glowVertexAttribFormat(gpVertexAttribFormat, (C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (C.GLuint)(relativeoffset))
}
func VertexAttribI4i(index uint32, x int32, y int32, z int32, w int32) {
	C.glowVertexAttribI4i(gpVertexAttribI4i, (C.GLuint)(index), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
}
func VertexAttribI4iv(index uint32, v *int32) {
	C.glowVertexAttribI4iv(gpVertexAttribI4iv, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func VertexAttribI4ui(index uint32, x uint32, y uint32, z uint32, w uint32) {
	C.glowVertexAttribI4ui(gpVertexAttribI4ui, (C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y), (C.GLuint)(z), (C.GLuint)(w))
}
func VertexAttribI4uiv(index uint32, v *uint32) {
	C.glowVertexAttribI4uiv(gpVertexAttribI4uiv, (C.GLuint)(index), (*C.GLuint)(unsafe.Pointer(v)))
}
func VertexAttribIFormat(attribindex uint32, size int32, xtype uint32, relativeoffset uint32) {
	C.glowVertexAttribIFormat(gpVertexAttribIFormat, (C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(xtype), (C.GLuint)(relativeoffset))
}
func VertexAttribIPointer(index uint32, size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	C.glowVertexAttribIPointer(gpVertexAttribIPointer, (C.GLuint)(index), (C.GLint)(size), (C.GLenum)(xtype), (C.GLsizei)(stride), pointer)
}

// define an array of generic vertex attribute data
func VertexAttribPointer(index uint32, size int32, xtype uint32, normalized bool, stride int32, pointer unsafe.Pointer) {
	C.glowVertexAttribPointer(gpVertexAttribPointer, (C.GLuint)(index), (C.GLint)(size), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (C.GLsizei)(stride), pointer)
}

// modify the rate at which generic vertex attributes     advance
func VertexBindingDivisor(bindingindex uint32, divisor uint32) {
	C.glowVertexBindingDivisor(gpVertexBindingDivisor, (C.GLuint)(bindingindex), (C.GLuint)(divisor))
}

// set the viewport
func Viewport(x int32, y int32, width int32, height int32) {
	C.glowViewport(gpViewport, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
func ViewportArrayvNV(first uint32, count int32, v *float32) {
	C.glowViewportArrayvNV(gpViewportArrayvNV, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(v)))
}
func ViewportArrayvOES(first uint32, count int32, v *float32) {
	C.glowViewportArrayvOES(gpViewportArrayvOES, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(v)))
}
func ViewportIndexedfNV(index uint32, x float32, y float32, w float32, h float32) {
	C.glowViewportIndexedfNV(gpViewportIndexedfNV, (C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(w), (C.GLfloat)(h))
}
func ViewportIndexedfOES(index uint32, x float32, y float32, w float32, h float32) {
	C.glowViewportIndexedfOES(gpViewportIndexedfOES, (C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(w), (C.GLfloat)(h))
}
func ViewportIndexedfvNV(index uint32, v *float32) {
	C.glowViewportIndexedfvNV(gpViewportIndexedfvNV, (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(v)))
}
func ViewportIndexedfvOES(index uint32, v *float32) {
	C.glowViewportIndexedfvOES(gpViewportIndexedfvOES, (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(v)))
}
func ViewportPositionWScaleNV(index uint32, xcoeff float32, ycoeff float32) {
	C.glowViewportPositionWScaleNV(gpViewportPositionWScaleNV, (C.GLuint)(index), (C.GLfloat)(xcoeff), (C.GLfloat)(ycoeff))
}
func ViewportSwizzleNV(index uint32, swizzlex uint32, swizzley uint32, swizzlez uint32, swizzlew uint32) {
	C.glowViewportSwizzleNV(gpViewportSwizzleNV, (C.GLuint)(index), (C.GLenum)(swizzlex), (C.GLenum)(swizzley), (C.GLenum)(swizzlez), (C.GLenum)(swizzlew))
}
func WaitSemaphoreEXT(semaphore uint32, numBufferBarriers uint32, buffers *uint32, numTextureBarriers uint32, textures *uint32, srcLayouts *uint32) {
	C.glowWaitSemaphoreEXT(gpWaitSemaphoreEXT, (C.GLuint)(semaphore), (C.GLuint)(numBufferBarriers), (*C.GLuint)(unsafe.Pointer(buffers)), (C.GLuint)(numTextureBarriers), (*C.GLuint)(unsafe.Pointer(textures)), (*C.GLenum)(unsafe.Pointer(srcLayouts)))
}

// instruct the GL server to block until the specified sync object becomes signaled
func WaitSync(sync uintptr, flags uint32, timeout uint64) {
	C.glowWaitSync(gpWaitSync, (C.GLsync)(sync), (C.GLbitfield)(flags), (C.GLuint64)(timeout))
}
func WaitSyncAPPLE(sync uintptr, flags uint32, timeout uint64) {
	C.glowWaitSyncAPPLE(gpWaitSyncAPPLE, (C.GLsync)(sync), (C.GLbitfield)(flags), (C.GLuint64)(timeout))
}
func WaitVkSemaphoreNV(vkSemaphore uint64) {
	C.glowWaitVkSemaphoreNV(gpWaitVkSemaphoreNV, (C.GLuint64)(vkSemaphore))
}
func WeightPathsNV(resultPath uint32, numPaths int32, paths *uint32, weights *float32) {
	C.glowWeightPathsNV(gpWeightPathsNV, (C.GLuint)(resultPath), (C.GLsizei)(numPaths), (*C.GLuint)(unsafe.Pointer(paths)), (*C.GLfloat)(unsafe.Pointer(weights)))
}
func WindowRectanglesEXT(mode uint32, count int32, box *int32) {
	C.glowWindowRectanglesEXT(gpWindowRectanglesEXT, (C.GLenum)(mode), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(box)))
}

// Init initializes the OpenGL bindings by loading the function pointers (for
// each OpenGL function) from the active OpenGL context.
//
// It must be called under the presence of an active OpenGL context, e.g.,
// always after calling window.MakeContextCurrent() and always before calling
// any OpenGL functions exported by this package.
//
// On Windows, Init loads pointers that are context-specific (and hence you
// must re-init if switching between OpenGL contexts, although not calling Init
// again after switching between OpenGL contexts may work if the contexts belong
// to the same graphics driver/device).
//
// On macOS and the other POSIX systems, the behavior is different, but code
// written compatible with the Windows behavior is compatible with macOS and the
// other POSIX systems. That is, always Init under an active OpenGL context, and
// always re-init after switching graphics contexts.
//
// For information about caveats of Init, you should read the "Platform Specific
// Function Retrieval" section of https://www.opengl.org/wiki/Load_OpenGL_Functions.
func Init() error {
	return InitWithProcAddrFunc(getProcAddress)
}

// InitWithProcAddrFunc intializes the package using the specified OpenGL
// function pointer loading function. For more cases Init should be used
// instead.
func InitWithProcAddrFunc(getProcAddr func(name string) unsafe.Pointer) error {
	gpAcquireKeyedMutexWin32EXT = (C.GPACQUIREKEYEDMUTEXWIN32EXT)(getProcAddr("glAcquireKeyedMutexWin32EXT"))
	gpActiveProgramEXT = (C.GPACTIVEPROGRAMEXT)(getProcAddr("glActiveProgramEXT"))
	gpActiveShaderProgram = (C.GPACTIVESHADERPROGRAM)(getProcAddr("glActiveShaderProgram"))
	if gpActiveShaderProgram == nil {
		return errors.New("glActiveShaderProgram")
	}
	gpActiveShaderProgramEXT = (C.GPACTIVESHADERPROGRAMEXT)(getProcAddr("glActiveShaderProgramEXT"))
	gpActiveTexture = (C.GPACTIVETEXTURE)(getProcAddr("glActiveTexture"))
	if gpActiveTexture == nil {
		return errors.New("glActiveTexture")
	}
	gpAlphaFuncQCOM = (C.GPALPHAFUNCQCOM)(getProcAddr("glAlphaFuncQCOM"))
	gpApplyFramebufferAttachmentCMAAINTEL = (C.GPAPPLYFRAMEBUFFERATTACHMENTCMAAINTEL)(getProcAddr("glApplyFramebufferAttachmentCMAAINTEL"))
	gpAttachShader = (C.GPATTACHSHADER)(getProcAddr("glAttachShader"))
	if gpAttachShader == nil {
		return errors.New("glAttachShader")
	}
	gpBeginConditionalRenderNV = (C.GPBEGINCONDITIONALRENDERNV)(getProcAddr("glBeginConditionalRenderNV"))
	gpBeginPerfMonitorAMD = (C.GPBEGINPERFMONITORAMD)(getProcAddr("glBeginPerfMonitorAMD"))
	gpBeginPerfQueryINTEL = (C.GPBEGINPERFQUERYINTEL)(getProcAddr("glBeginPerfQueryINTEL"))
	gpBeginQuery = (C.GPBEGINQUERY)(getProcAddr("glBeginQuery"))
	if gpBeginQuery == nil {
		return errors.New("glBeginQuery")
	}
	gpBeginQueryEXT = (C.GPBEGINQUERYEXT)(getProcAddr("glBeginQueryEXT"))
	gpBeginTransformFeedback = (C.GPBEGINTRANSFORMFEEDBACK)(getProcAddr("glBeginTransformFeedback"))
	if gpBeginTransformFeedback == nil {
		return errors.New("glBeginTransformFeedback")
	}
	gpBindAttribLocation = (C.GPBINDATTRIBLOCATION)(getProcAddr("glBindAttribLocation"))
	if gpBindAttribLocation == nil {
		return errors.New("glBindAttribLocation")
	}
	gpBindBuffer = (C.GPBINDBUFFER)(getProcAddr("glBindBuffer"))
	if gpBindBuffer == nil {
		return errors.New("glBindBuffer")
	}
	gpBindBufferBase = (C.GPBINDBUFFERBASE)(getProcAddr("glBindBufferBase"))
	if gpBindBufferBase == nil {
		return errors.New("glBindBufferBase")
	}
	gpBindBufferRange = (C.GPBINDBUFFERRANGE)(getProcAddr("glBindBufferRange"))
	if gpBindBufferRange == nil {
		return errors.New("glBindBufferRange")
	}
	gpBindFragDataLocationEXT = (C.GPBINDFRAGDATALOCATIONEXT)(getProcAddr("glBindFragDataLocationEXT"))
	gpBindFragDataLocationIndexedEXT = (C.GPBINDFRAGDATALOCATIONINDEXEDEXT)(getProcAddr("glBindFragDataLocationIndexedEXT"))
	gpBindFramebuffer = (C.GPBINDFRAMEBUFFER)(getProcAddr("glBindFramebuffer"))
	if gpBindFramebuffer == nil {
		return errors.New("glBindFramebuffer")
	}
	gpBindImageTexture = (C.GPBINDIMAGETEXTURE)(getProcAddr("glBindImageTexture"))
	if gpBindImageTexture == nil {
		return errors.New("glBindImageTexture")
	}
	gpBindProgramPipeline = (C.GPBINDPROGRAMPIPELINE)(getProcAddr("glBindProgramPipeline"))
	if gpBindProgramPipeline == nil {
		return errors.New("glBindProgramPipeline")
	}
	gpBindProgramPipelineEXT = (C.GPBINDPROGRAMPIPELINEEXT)(getProcAddr("glBindProgramPipelineEXT"))
	gpBindRenderbuffer = (C.GPBINDRENDERBUFFER)(getProcAddr("glBindRenderbuffer"))
	if gpBindRenderbuffer == nil {
		return errors.New("glBindRenderbuffer")
	}
	gpBindSampler = (C.GPBINDSAMPLER)(getProcAddr("glBindSampler"))
	if gpBindSampler == nil {
		return errors.New("glBindSampler")
	}
	gpBindTexture = (C.GPBINDTEXTURE)(getProcAddr("glBindTexture"))
	if gpBindTexture == nil {
		return errors.New("glBindTexture")
	}
	gpBindTransformFeedback = (C.GPBINDTRANSFORMFEEDBACK)(getProcAddr("glBindTransformFeedback"))
	if gpBindTransformFeedback == nil {
		return errors.New("glBindTransformFeedback")
	}
	gpBindVertexArray = (C.GPBINDVERTEXARRAY)(getProcAddr("glBindVertexArray"))
	if gpBindVertexArray == nil {
		return errors.New("glBindVertexArray")
	}
	gpBindVertexArrayOES = (C.GPBINDVERTEXARRAYOES)(getProcAddr("glBindVertexArrayOES"))
	gpBindVertexBuffer = (C.GPBINDVERTEXBUFFER)(getProcAddr("glBindVertexBuffer"))
	if gpBindVertexBuffer == nil {
		return errors.New("glBindVertexBuffer")
	}
	gpBlendBarrierKHR = (C.GPBLENDBARRIERKHR)(getProcAddr("glBlendBarrierKHR"))
	gpBlendBarrierNV = (C.GPBLENDBARRIERNV)(getProcAddr("glBlendBarrierNV"))
	gpBlendColor = (C.GPBLENDCOLOR)(getProcAddr("glBlendColor"))
	if gpBlendColor == nil {
		return errors.New("glBlendColor")
	}
	gpBlendEquation = (C.GPBLENDEQUATION)(getProcAddr("glBlendEquation"))
	if gpBlendEquation == nil {
		return errors.New("glBlendEquation")
	}
	gpBlendEquationEXT = (C.GPBLENDEQUATIONEXT)(getProcAddr("glBlendEquationEXT"))
	gpBlendEquationSeparate = (C.GPBLENDEQUATIONSEPARATE)(getProcAddr("glBlendEquationSeparate"))
	if gpBlendEquationSeparate == nil {
		return errors.New("glBlendEquationSeparate")
	}
	gpBlendEquationSeparateiEXT = (C.GPBLENDEQUATIONSEPARATEIEXT)(getProcAddr("glBlendEquationSeparateiEXT"))
	gpBlendEquationSeparateiOES = (C.GPBLENDEQUATIONSEPARATEIOES)(getProcAddr("glBlendEquationSeparateiOES"))
	gpBlendEquationiEXT = (C.GPBLENDEQUATIONIEXT)(getProcAddr("glBlendEquationiEXT"))
	gpBlendEquationiOES = (C.GPBLENDEQUATIONIOES)(getProcAddr("glBlendEquationiOES"))
	gpBlendFunc = (C.GPBLENDFUNC)(getProcAddr("glBlendFunc"))
	if gpBlendFunc == nil {
		return errors.New("glBlendFunc")
	}
	gpBlendFuncSeparate = (C.GPBLENDFUNCSEPARATE)(getProcAddr("glBlendFuncSeparate"))
	if gpBlendFuncSeparate == nil {
		return errors.New("glBlendFuncSeparate")
	}
	gpBlendFuncSeparateiEXT = (C.GPBLENDFUNCSEPARATEIEXT)(getProcAddr("glBlendFuncSeparateiEXT"))
	gpBlendFuncSeparateiOES = (C.GPBLENDFUNCSEPARATEIOES)(getProcAddr("glBlendFuncSeparateiOES"))
	gpBlendFunciEXT = (C.GPBLENDFUNCIEXT)(getProcAddr("glBlendFunciEXT"))
	gpBlendFunciOES = (C.GPBLENDFUNCIOES)(getProcAddr("glBlendFunciOES"))
	gpBlendParameteriNV = (C.GPBLENDPARAMETERINV)(getProcAddr("glBlendParameteriNV"))
	gpBlitFramebuffer = (C.GPBLITFRAMEBUFFER)(getProcAddr("glBlitFramebuffer"))
	if gpBlitFramebuffer == nil {
		return errors.New("glBlitFramebuffer")
	}
	gpBlitFramebufferANGLE = (C.GPBLITFRAMEBUFFERANGLE)(getProcAddr("glBlitFramebufferANGLE"))
	gpBlitFramebufferNV = (C.GPBLITFRAMEBUFFERNV)(getProcAddr("glBlitFramebufferNV"))
	gpBufferData = (C.GPBUFFERDATA)(getProcAddr("glBufferData"))
	if gpBufferData == nil {
		return errors.New("glBufferData")
	}
	gpBufferStorageEXT = (C.GPBUFFERSTORAGEEXT)(getProcAddr("glBufferStorageEXT"))
	gpBufferStorageExternalEXT = (C.GPBUFFERSTORAGEEXTERNALEXT)(getProcAddr("glBufferStorageExternalEXT"))
	gpBufferStorageMemEXT = (C.GPBUFFERSTORAGEMEMEXT)(getProcAddr("glBufferStorageMemEXT"))
	gpBufferSubData = (C.GPBUFFERSUBDATA)(getProcAddr("glBufferSubData"))
	if gpBufferSubData == nil {
		return errors.New("glBufferSubData")
	}
	gpCheckFramebufferStatus = (C.GPCHECKFRAMEBUFFERSTATUS)(getProcAddr("glCheckFramebufferStatus"))
	if gpCheckFramebufferStatus == nil {
		return errors.New("glCheckFramebufferStatus")
	}
	gpClear = (C.GPCLEAR)(getProcAddr("glClear"))
	if gpClear == nil {
		return errors.New("glClear")
	}
	gpClearBufferfi = (C.GPCLEARBUFFERFI)(getProcAddr("glClearBufferfi"))
	if gpClearBufferfi == nil {
		return errors.New("glClearBufferfi")
	}
	gpClearBufferfv = (C.GPCLEARBUFFERFV)(getProcAddr("glClearBufferfv"))
	if gpClearBufferfv == nil {
		return errors.New("glClearBufferfv")
	}
	gpClearBufferiv = (C.GPCLEARBUFFERIV)(getProcAddr("glClearBufferiv"))
	if gpClearBufferiv == nil {
		return errors.New("glClearBufferiv")
	}
	gpClearBufferuiv = (C.GPCLEARBUFFERUIV)(getProcAddr("glClearBufferuiv"))
	if gpClearBufferuiv == nil {
		return errors.New("glClearBufferuiv")
	}
	gpClearColor = (C.GPCLEARCOLOR)(getProcAddr("glClearColor"))
	if gpClearColor == nil {
		return errors.New("glClearColor")
	}
	gpClearDepthf = (C.GPCLEARDEPTHF)(getProcAddr("glClearDepthf"))
	if gpClearDepthf == nil {
		return errors.New("glClearDepthf")
	}
	gpClearPixelLocalStorageuiEXT = (C.GPCLEARPIXELLOCALSTORAGEUIEXT)(getProcAddr("glClearPixelLocalStorageuiEXT"))
	gpClearStencil = (C.GPCLEARSTENCIL)(getProcAddr("glClearStencil"))
	if gpClearStencil == nil {
		return errors.New("glClearStencil")
	}
	gpClearTexImageEXT = (C.GPCLEARTEXIMAGEEXT)(getProcAddr("glClearTexImageEXT"))
	gpClearTexSubImageEXT = (C.GPCLEARTEXSUBIMAGEEXT)(getProcAddr("glClearTexSubImageEXT"))
	gpClientWaitSync = (C.GPCLIENTWAITSYNC)(getProcAddr("glClientWaitSync"))
	if gpClientWaitSync == nil {
		return errors.New("glClientWaitSync")
	}
	gpClientWaitSyncAPPLE = (C.GPCLIENTWAITSYNCAPPLE)(getProcAddr("glClientWaitSyncAPPLE"))
	gpClipControlEXT = (C.GPCLIPCONTROLEXT)(getProcAddr("glClipControlEXT"))
	gpColorMask = (C.GPCOLORMASK)(getProcAddr("glColorMask"))
	if gpColorMask == nil {
		return errors.New("glColorMask")
	}
	gpColorMaskiEXT = (C.GPCOLORMASKIEXT)(getProcAddr("glColorMaskiEXT"))
	gpColorMaskiOES = (C.GPCOLORMASKIOES)(getProcAddr("glColorMaskiOES"))
	gpCompileShader = (C.GPCOMPILESHADER)(getProcAddr("glCompileShader"))
	if gpCompileShader == nil {
		return errors.New("glCompileShader")
	}
	gpCompressedTexImage2D = (C.GPCOMPRESSEDTEXIMAGE2D)(getProcAddr("glCompressedTexImage2D"))
	if gpCompressedTexImage2D == nil {
		return errors.New("glCompressedTexImage2D")
	}
	gpCompressedTexImage3D = (C.GPCOMPRESSEDTEXIMAGE3D)(getProcAddr("glCompressedTexImage3D"))
	if gpCompressedTexImage3D == nil {
		return errors.New("glCompressedTexImage3D")
	}
	gpCompressedTexImage3DOES = (C.GPCOMPRESSEDTEXIMAGE3DOES)(getProcAddr("glCompressedTexImage3DOES"))
	gpCompressedTexSubImage2D = (C.GPCOMPRESSEDTEXSUBIMAGE2D)(getProcAddr("glCompressedTexSubImage2D"))
	if gpCompressedTexSubImage2D == nil {
		return errors.New("glCompressedTexSubImage2D")
	}
	gpCompressedTexSubImage3D = (C.GPCOMPRESSEDTEXSUBIMAGE3D)(getProcAddr("glCompressedTexSubImage3D"))
	if gpCompressedTexSubImage3D == nil {
		return errors.New("glCompressedTexSubImage3D")
	}
	gpCompressedTexSubImage3DOES = (C.GPCOMPRESSEDTEXSUBIMAGE3DOES)(getProcAddr("glCompressedTexSubImage3DOES"))
	gpConservativeRasterParameteriNV = (C.GPCONSERVATIVERASTERPARAMETERINV)(getProcAddr("glConservativeRasterParameteriNV"))
	gpCopyBufferSubData = (C.GPCOPYBUFFERSUBDATA)(getProcAddr("glCopyBufferSubData"))
	if gpCopyBufferSubData == nil {
		return errors.New("glCopyBufferSubData")
	}
	gpCopyBufferSubDataNV = (C.GPCOPYBUFFERSUBDATANV)(getProcAddr("glCopyBufferSubDataNV"))
	gpCopyImageSubDataEXT = (C.GPCOPYIMAGESUBDATAEXT)(getProcAddr("glCopyImageSubDataEXT"))
	gpCopyImageSubDataOES = (C.GPCOPYIMAGESUBDATAOES)(getProcAddr("glCopyImageSubDataOES"))
	gpCopyPathNV = (C.GPCOPYPATHNV)(getProcAddr("glCopyPathNV"))
	gpCopyTexImage2D = (C.GPCOPYTEXIMAGE2D)(getProcAddr("glCopyTexImage2D"))
	if gpCopyTexImage2D == nil {
		return errors.New("glCopyTexImage2D")
	}
	gpCopyTexSubImage2D = (C.GPCOPYTEXSUBIMAGE2D)(getProcAddr("glCopyTexSubImage2D"))
	if gpCopyTexSubImage2D == nil {
		return errors.New("glCopyTexSubImage2D")
	}
	gpCopyTexSubImage3D = (C.GPCOPYTEXSUBIMAGE3D)(getProcAddr("glCopyTexSubImage3D"))
	if gpCopyTexSubImage3D == nil {
		return errors.New("glCopyTexSubImage3D")
	}
	gpCopyTexSubImage3DOES = (C.GPCOPYTEXSUBIMAGE3DOES)(getProcAddr("glCopyTexSubImage3DOES"))
	gpCopyTextureLevelsAPPLE = (C.GPCOPYTEXTURELEVELSAPPLE)(getProcAddr("glCopyTextureLevelsAPPLE"))
	gpCoverFillPathInstancedNV = (C.GPCOVERFILLPATHINSTANCEDNV)(getProcAddr("glCoverFillPathInstancedNV"))
	gpCoverFillPathNV = (C.GPCOVERFILLPATHNV)(getProcAddr("glCoverFillPathNV"))
	gpCoverStrokePathInstancedNV = (C.GPCOVERSTROKEPATHINSTANCEDNV)(getProcAddr("glCoverStrokePathInstancedNV"))
	gpCoverStrokePathNV = (C.GPCOVERSTROKEPATHNV)(getProcAddr("glCoverStrokePathNV"))
	gpCoverageMaskNV = (C.GPCOVERAGEMASKNV)(getProcAddr("glCoverageMaskNV"))
	gpCoverageModulationNV = (C.GPCOVERAGEMODULATIONNV)(getProcAddr("glCoverageModulationNV"))
	gpCoverageModulationTableNV = (C.GPCOVERAGEMODULATIONTABLENV)(getProcAddr("glCoverageModulationTableNV"))
	gpCoverageOperationNV = (C.GPCOVERAGEOPERATIONNV)(getProcAddr("glCoverageOperationNV"))
	gpCreateMemoryObjectsEXT = (C.GPCREATEMEMORYOBJECTSEXT)(getProcAddr("glCreateMemoryObjectsEXT"))
	gpCreatePerfQueryINTEL = (C.GPCREATEPERFQUERYINTEL)(getProcAddr("glCreatePerfQueryINTEL"))
	gpCreateProgram = (C.GPCREATEPROGRAM)(getProcAddr("glCreateProgram"))
	if gpCreateProgram == nil {
		return errors.New("glCreateProgram")
	}
	gpCreateShader = (C.GPCREATESHADER)(getProcAddr("glCreateShader"))
	if gpCreateShader == nil {
		return errors.New("glCreateShader")
	}
	gpCreateShaderProgramEXT = (C.GPCREATESHADERPROGRAMEXT)(getProcAddr("glCreateShaderProgramEXT"))
	gpCreateShaderProgramv = (C.GPCREATESHADERPROGRAMV)(getProcAddr("glCreateShaderProgramv"))
	if gpCreateShaderProgramv == nil {
		return errors.New("glCreateShaderProgramv")
	}
	gpCreateShaderProgramvEXT = (C.GPCREATESHADERPROGRAMVEXT)(getProcAddr("glCreateShaderProgramvEXT"))
	gpCullFace = (C.GPCULLFACE)(getProcAddr("glCullFace"))
	if gpCullFace == nil {
		return errors.New("glCullFace")
	}
	gpDebugMessageCallback = (C.GPDEBUGMESSAGECALLBACK)(getProcAddr("glDebugMessageCallback"))
	gpDebugMessageCallbackKHR = (C.GPDEBUGMESSAGECALLBACKKHR)(getProcAddr("glDebugMessageCallbackKHR"))
	gpDebugMessageControl = (C.GPDEBUGMESSAGECONTROL)(getProcAddr("glDebugMessageControl"))
	gpDebugMessageControlKHR = (C.GPDEBUGMESSAGECONTROLKHR)(getProcAddr("glDebugMessageControlKHR"))
	gpDebugMessageInsert = (C.GPDEBUGMESSAGEINSERT)(getProcAddr("glDebugMessageInsert"))
	gpDebugMessageInsertKHR = (C.GPDEBUGMESSAGEINSERTKHR)(getProcAddr("glDebugMessageInsertKHR"))
	gpDeleteBuffers = (C.GPDELETEBUFFERS)(getProcAddr("glDeleteBuffers"))
	if gpDeleteBuffers == nil {
		return errors.New("glDeleteBuffers")
	}
	gpDeleteFencesNV = (C.GPDELETEFENCESNV)(getProcAddr("glDeleteFencesNV"))
	gpDeleteFramebuffers = (C.GPDELETEFRAMEBUFFERS)(getProcAddr("glDeleteFramebuffers"))
	if gpDeleteFramebuffers == nil {
		return errors.New("glDeleteFramebuffers")
	}
	gpDeleteMemoryObjectsEXT = (C.GPDELETEMEMORYOBJECTSEXT)(getProcAddr("glDeleteMemoryObjectsEXT"))
	gpDeletePathsNV = (C.GPDELETEPATHSNV)(getProcAddr("glDeletePathsNV"))
	gpDeletePerfMonitorsAMD = (C.GPDELETEPERFMONITORSAMD)(getProcAddr("glDeletePerfMonitorsAMD"))
	gpDeletePerfQueryINTEL = (C.GPDELETEPERFQUERYINTEL)(getProcAddr("glDeletePerfQueryINTEL"))
	gpDeleteProgram = (C.GPDELETEPROGRAM)(getProcAddr("glDeleteProgram"))
	if gpDeleteProgram == nil {
		return errors.New("glDeleteProgram")
	}
	gpDeleteProgramPipelines = (C.GPDELETEPROGRAMPIPELINES)(getProcAddr("glDeleteProgramPipelines"))
	if gpDeleteProgramPipelines == nil {
		return errors.New("glDeleteProgramPipelines")
	}
	gpDeleteProgramPipelinesEXT = (C.GPDELETEPROGRAMPIPELINESEXT)(getProcAddr("glDeleteProgramPipelinesEXT"))
	gpDeleteQueries = (C.GPDELETEQUERIES)(getProcAddr("glDeleteQueries"))
	if gpDeleteQueries == nil {
		return errors.New("glDeleteQueries")
	}
	gpDeleteQueriesEXT = (C.GPDELETEQUERIESEXT)(getProcAddr("glDeleteQueriesEXT"))
	gpDeleteRenderbuffers = (C.GPDELETERENDERBUFFERS)(getProcAddr("glDeleteRenderbuffers"))
	if gpDeleteRenderbuffers == nil {
		return errors.New("glDeleteRenderbuffers")
	}
	gpDeleteSamplers = (C.GPDELETESAMPLERS)(getProcAddr("glDeleteSamplers"))
	if gpDeleteSamplers == nil {
		return errors.New("glDeleteSamplers")
	}
	gpDeleteSemaphoresEXT = (C.GPDELETESEMAPHORESEXT)(getProcAddr("glDeleteSemaphoresEXT"))
	gpDeleteShader = (C.GPDELETESHADER)(getProcAddr("glDeleteShader"))
	if gpDeleteShader == nil {
		return errors.New("glDeleteShader")
	}
	gpDeleteSync = (C.GPDELETESYNC)(getProcAddr("glDeleteSync"))
	if gpDeleteSync == nil {
		return errors.New("glDeleteSync")
	}
	gpDeleteSyncAPPLE = (C.GPDELETESYNCAPPLE)(getProcAddr("glDeleteSyncAPPLE"))
	gpDeleteTextures = (C.GPDELETETEXTURES)(getProcAddr("glDeleteTextures"))
	if gpDeleteTextures == nil {
		return errors.New("glDeleteTextures")
	}
	gpDeleteTransformFeedbacks = (C.GPDELETETRANSFORMFEEDBACKS)(getProcAddr("glDeleteTransformFeedbacks"))
	if gpDeleteTransformFeedbacks == nil {
		return errors.New("glDeleteTransformFeedbacks")
	}
	gpDeleteVertexArrays = (C.GPDELETEVERTEXARRAYS)(getProcAddr("glDeleteVertexArrays"))
	if gpDeleteVertexArrays == nil {
		return errors.New("glDeleteVertexArrays")
	}
	gpDeleteVertexArraysOES = (C.GPDELETEVERTEXARRAYSOES)(getProcAddr("glDeleteVertexArraysOES"))
	gpDepthFunc = (C.GPDEPTHFUNC)(getProcAddr("glDepthFunc"))
	if gpDepthFunc == nil {
		return errors.New("glDepthFunc")
	}
	gpDepthMask = (C.GPDEPTHMASK)(getProcAddr("glDepthMask"))
	if gpDepthMask == nil {
		return errors.New("glDepthMask")
	}
	gpDepthRangeArrayfvNV = (C.GPDEPTHRANGEARRAYFVNV)(getProcAddr("glDepthRangeArrayfvNV"))
	gpDepthRangeArrayfvOES = (C.GPDEPTHRANGEARRAYFVOES)(getProcAddr("glDepthRangeArrayfvOES"))
	gpDepthRangeIndexedfNV = (C.GPDEPTHRANGEINDEXEDFNV)(getProcAddr("glDepthRangeIndexedfNV"))
	gpDepthRangeIndexedfOES = (C.GPDEPTHRANGEINDEXEDFOES)(getProcAddr("glDepthRangeIndexedfOES"))
	gpDepthRangef = (C.GPDEPTHRANGEF)(getProcAddr("glDepthRangef"))
	if gpDepthRangef == nil {
		return errors.New("glDepthRangef")
	}
	gpDetachShader = (C.GPDETACHSHADER)(getProcAddr("glDetachShader"))
	if gpDetachShader == nil {
		return errors.New("glDetachShader")
	}
	gpDisable = (C.GPDISABLE)(getProcAddr("glDisable"))
	if gpDisable == nil {
		return errors.New("glDisable")
	}
	gpDisableDriverControlQCOM = (C.GPDISABLEDRIVERCONTROLQCOM)(getProcAddr("glDisableDriverControlQCOM"))
	gpDisableVertexAttribArray = (C.GPDISABLEVERTEXATTRIBARRAY)(getProcAddr("glDisableVertexAttribArray"))
	if gpDisableVertexAttribArray == nil {
		return errors.New("glDisableVertexAttribArray")
	}
	gpDisableiEXT = (C.GPDISABLEIEXT)(getProcAddr("glDisableiEXT"))
	gpDisableiNV = (C.GPDISABLEINV)(getProcAddr("glDisableiNV"))
	gpDisableiOES = (C.GPDISABLEIOES)(getProcAddr("glDisableiOES"))
	gpDiscardFramebufferEXT = (C.GPDISCARDFRAMEBUFFEREXT)(getProcAddr("glDiscardFramebufferEXT"))
	gpDispatchCompute = (C.GPDISPATCHCOMPUTE)(getProcAddr("glDispatchCompute"))
	if gpDispatchCompute == nil {
		return errors.New("glDispatchCompute")
	}
	gpDispatchComputeIndirect = (C.GPDISPATCHCOMPUTEINDIRECT)(getProcAddr("glDispatchComputeIndirect"))
	if gpDispatchComputeIndirect == nil {
		return errors.New("glDispatchComputeIndirect")
	}
	gpDrawArrays = (C.GPDRAWARRAYS)(getProcAddr("glDrawArrays"))
	if gpDrawArrays == nil {
		return errors.New("glDrawArrays")
	}
	gpDrawArraysIndirect = (C.GPDRAWARRAYSINDIRECT)(getProcAddr("glDrawArraysIndirect"))
	if gpDrawArraysIndirect == nil {
		return errors.New("glDrawArraysIndirect")
	}
	gpDrawArraysInstanced = (C.GPDRAWARRAYSINSTANCED)(getProcAddr("glDrawArraysInstanced"))
	if gpDrawArraysInstanced == nil {
		return errors.New("glDrawArraysInstanced")
	}
	gpDrawArraysInstancedANGLE = (C.GPDRAWARRAYSINSTANCEDANGLE)(getProcAddr("glDrawArraysInstancedANGLE"))
	gpDrawArraysInstancedBaseInstanceEXT = (C.GPDRAWARRAYSINSTANCEDBASEINSTANCEEXT)(getProcAddr("glDrawArraysInstancedBaseInstanceEXT"))
	gpDrawArraysInstancedEXT = (C.GPDRAWARRAYSINSTANCEDEXT)(getProcAddr("glDrawArraysInstancedEXT"))
	gpDrawArraysInstancedNV = (C.GPDRAWARRAYSINSTANCEDNV)(getProcAddr("glDrawArraysInstancedNV"))
	gpDrawBuffers = (C.GPDRAWBUFFERS)(getProcAddr("glDrawBuffers"))
	if gpDrawBuffers == nil {
		return errors.New("glDrawBuffers")
	}
	gpDrawBuffersEXT = (C.GPDRAWBUFFERSEXT)(getProcAddr("glDrawBuffersEXT"))
	gpDrawBuffersIndexedEXT = (C.GPDRAWBUFFERSINDEXEDEXT)(getProcAddr("glDrawBuffersIndexedEXT"))
	gpDrawBuffersNV = (C.GPDRAWBUFFERSNV)(getProcAddr("glDrawBuffersNV"))
	gpDrawElements = (C.GPDRAWELEMENTS)(getProcAddr("glDrawElements"))
	if gpDrawElements == nil {
		return errors.New("glDrawElements")
	}
	gpDrawElementsBaseVertexEXT = (C.GPDRAWELEMENTSBASEVERTEXEXT)(getProcAddr("glDrawElementsBaseVertexEXT"))
	gpDrawElementsBaseVertexOES = (C.GPDRAWELEMENTSBASEVERTEXOES)(getProcAddr("glDrawElementsBaseVertexOES"))
	gpDrawElementsIndirect = (C.GPDRAWELEMENTSINDIRECT)(getProcAddr("glDrawElementsIndirect"))
	if gpDrawElementsIndirect == nil {
		return errors.New("glDrawElementsIndirect")
	}
	gpDrawElementsInstanced = (C.GPDRAWELEMENTSINSTANCED)(getProcAddr("glDrawElementsInstanced"))
	if gpDrawElementsInstanced == nil {
		return errors.New("glDrawElementsInstanced")
	}
	gpDrawElementsInstancedANGLE = (C.GPDRAWELEMENTSINSTANCEDANGLE)(getProcAddr("glDrawElementsInstancedANGLE"))
	gpDrawElementsInstancedBaseInstanceEXT = (C.GPDRAWELEMENTSINSTANCEDBASEINSTANCEEXT)(getProcAddr("glDrawElementsInstancedBaseInstanceEXT"))
	gpDrawElementsInstancedBaseVertexBaseInstanceEXT = (C.GPDRAWELEMENTSINSTANCEDBASEVERTEXBASEINSTANCEEXT)(getProcAddr("glDrawElementsInstancedBaseVertexBaseInstanceEXT"))
	gpDrawElementsInstancedBaseVertexEXT = (C.GPDRAWELEMENTSINSTANCEDBASEVERTEXEXT)(getProcAddr("glDrawElementsInstancedBaseVertexEXT"))
	gpDrawElementsInstancedBaseVertexOES = (C.GPDRAWELEMENTSINSTANCEDBASEVERTEXOES)(getProcAddr("glDrawElementsInstancedBaseVertexOES"))
	gpDrawElementsInstancedEXT = (C.GPDRAWELEMENTSINSTANCEDEXT)(getProcAddr("glDrawElementsInstancedEXT"))
	gpDrawElementsInstancedNV = (C.GPDRAWELEMENTSINSTANCEDNV)(getProcAddr("glDrawElementsInstancedNV"))
	gpDrawRangeElements = (C.GPDRAWRANGEELEMENTS)(getProcAddr("glDrawRangeElements"))
	if gpDrawRangeElements == nil {
		return errors.New("glDrawRangeElements")
	}
	gpDrawRangeElementsBaseVertexEXT = (C.GPDRAWRANGEELEMENTSBASEVERTEXEXT)(getProcAddr("glDrawRangeElementsBaseVertexEXT"))
	gpDrawRangeElementsBaseVertexOES = (C.GPDRAWRANGEELEMENTSBASEVERTEXOES)(getProcAddr("glDrawRangeElementsBaseVertexOES"))
	gpDrawTransformFeedbackEXT = (C.GPDRAWTRANSFORMFEEDBACKEXT)(getProcAddr("glDrawTransformFeedbackEXT"))
	gpDrawTransformFeedbackInstancedEXT = (C.GPDRAWTRANSFORMFEEDBACKINSTANCEDEXT)(getProcAddr("glDrawTransformFeedbackInstancedEXT"))
	gpDrawVkImageNV = (C.GPDRAWVKIMAGENV)(getProcAddr("glDrawVkImageNV"))
	gpEGLImageTargetRenderbufferStorageOES = (C.GPEGLIMAGETARGETRENDERBUFFERSTORAGEOES)(getProcAddr("glEGLImageTargetRenderbufferStorageOES"))
	gpEGLImageTargetTexStorageEXT = (C.GPEGLIMAGETARGETTEXSTORAGEEXT)(getProcAddr("glEGLImageTargetTexStorageEXT"))
	gpEGLImageTargetTexture2DOES = (C.GPEGLIMAGETARGETTEXTURE2DOES)(getProcAddr("glEGLImageTargetTexture2DOES"))
	gpEGLImageTargetTextureStorageEXT = (C.GPEGLIMAGETARGETTEXTURESTORAGEEXT)(getProcAddr("glEGLImageTargetTextureStorageEXT"))
	gpEnable = (C.GPENABLE)(getProcAddr("glEnable"))
	if gpEnable == nil {
		return errors.New("glEnable")
	}
	gpEnableDriverControlQCOM = (C.GPENABLEDRIVERCONTROLQCOM)(getProcAddr("glEnableDriverControlQCOM"))
	gpEnableVertexAttribArray = (C.GPENABLEVERTEXATTRIBARRAY)(getProcAddr("glEnableVertexAttribArray"))
	if gpEnableVertexAttribArray == nil {
		return errors.New("glEnableVertexAttribArray")
	}
	gpEnableiEXT = (C.GPENABLEIEXT)(getProcAddr("glEnableiEXT"))
	gpEnableiNV = (C.GPENABLEINV)(getProcAddr("glEnableiNV"))
	gpEnableiOES = (C.GPENABLEIOES)(getProcAddr("glEnableiOES"))
	gpEndConditionalRenderNV = (C.GPENDCONDITIONALRENDERNV)(getProcAddr("glEndConditionalRenderNV"))
	gpEndPerfMonitorAMD = (C.GPENDPERFMONITORAMD)(getProcAddr("glEndPerfMonitorAMD"))
	gpEndPerfQueryINTEL = (C.GPENDPERFQUERYINTEL)(getProcAddr("glEndPerfQueryINTEL"))
	gpEndQuery = (C.GPENDQUERY)(getProcAddr("glEndQuery"))
	if gpEndQuery == nil {
		return errors.New("glEndQuery")
	}
	gpEndQueryEXT = (C.GPENDQUERYEXT)(getProcAddr("glEndQueryEXT"))
	gpEndTilingQCOM = (C.GPENDTILINGQCOM)(getProcAddr("glEndTilingQCOM"))
	gpEndTransformFeedback = (C.GPENDTRANSFORMFEEDBACK)(getProcAddr("glEndTransformFeedback"))
	if gpEndTransformFeedback == nil {
		return errors.New("glEndTransformFeedback")
	}
	gpExtGetBufferPointervQCOM = (C.GPEXTGETBUFFERPOINTERVQCOM)(getProcAddr("glExtGetBufferPointervQCOM"))
	gpExtGetBuffersQCOM = (C.GPEXTGETBUFFERSQCOM)(getProcAddr("glExtGetBuffersQCOM"))
	gpExtGetFramebuffersQCOM = (C.GPEXTGETFRAMEBUFFERSQCOM)(getProcAddr("glExtGetFramebuffersQCOM"))
	gpExtGetProgramBinarySourceQCOM = (C.GPEXTGETPROGRAMBINARYSOURCEQCOM)(getProcAddr("glExtGetProgramBinarySourceQCOM"))
	gpExtGetProgramsQCOM = (C.GPEXTGETPROGRAMSQCOM)(getProcAddr("glExtGetProgramsQCOM"))
	gpExtGetRenderbuffersQCOM = (C.GPEXTGETRENDERBUFFERSQCOM)(getProcAddr("glExtGetRenderbuffersQCOM"))
	gpExtGetShadersQCOM = (C.GPEXTGETSHADERSQCOM)(getProcAddr("glExtGetShadersQCOM"))
	gpExtGetTexLevelParameterivQCOM = (C.GPEXTGETTEXLEVELPARAMETERIVQCOM)(getProcAddr("glExtGetTexLevelParameterivQCOM"))
	gpExtGetTexSubImageQCOM = (C.GPEXTGETTEXSUBIMAGEQCOM)(getProcAddr("glExtGetTexSubImageQCOM"))
	gpExtGetTexturesQCOM = (C.GPEXTGETTEXTURESQCOM)(getProcAddr("glExtGetTexturesQCOM"))
	gpExtIsProgramBinaryQCOM = (C.GPEXTISPROGRAMBINARYQCOM)(getProcAddr("glExtIsProgramBinaryQCOM"))
	gpExtTexObjectStateOverrideiQCOM = (C.GPEXTTEXOBJECTSTATEOVERRIDEIQCOM)(getProcAddr("glExtTexObjectStateOverrideiQCOM"))
	gpFenceSync = (C.GPFENCESYNC)(getProcAddr("glFenceSync"))
	if gpFenceSync == nil {
		return errors.New("glFenceSync")
	}
	gpFenceSyncAPPLE = (C.GPFENCESYNCAPPLE)(getProcAddr("glFenceSyncAPPLE"))
	gpFinish = (C.GPFINISH)(getProcAddr("glFinish"))
	if gpFinish == nil {
		return errors.New("glFinish")
	}
	gpFinishFenceNV = (C.GPFINISHFENCENV)(getProcAddr("glFinishFenceNV"))
	gpFlush = (C.GPFLUSH)(getProcAddr("glFlush"))
	if gpFlush == nil {
		return errors.New("glFlush")
	}
	gpFlushMappedBufferRange = (C.GPFLUSHMAPPEDBUFFERRANGE)(getProcAddr("glFlushMappedBufferRange"))
	if gpFlushMappedBufferRange == nil {
		return errors.New("glFlushMappedBufferRange")
	}
	gpFlushMappedBufferRangeEXT = (C.GPFLUSHMAPPEDBUFFERRANGEEXT)(getProcAddr("glFlushMappedBufferRangeEXT"))
	gpFragmentCoverageColorNV = (C.GPFRAGMENTCOVERAGECOLORNV)(getProcAddr("glFragmentCoverageColorNV"))
	gpFramebufferFetchBarrierEXT = (C.GPFRAMEBUFFERFETCHBARRIEREXT)(getProcAddr("glFramebufferFetchBarrierEXT"))
	gpFramebufferFetchBarrierQCOM = (C.GPFRAMEBUFFERFETCHBARRIERQCOM)(getProcAddr("glFramebufferFetchBarrierQCOM"))
	gpFramebufferFoveationConfigQCOM = (C.GPFRAMEBUFFERFOVEATIONCONFIGQCOM)(getProcAddr("glFramebufferFoveationConfigQCOM"))
	gpFramebufferFoveationParametersQCOM = (C.GPFRAMEBUFFERFOVEATIONPARAMETERSQCOM)(getProcAddr("glFramebufferFoveationParametersQCOM"))
	gpFramebufferParameteri = (C.GPFRAMEBUFFERPARAMETERI)(getProcAddr("glFramebufferParameteri"))
	if gpFramebufferParameteri == nil {
		return errors.New("glFramebufferParameteri")
	}
	gpFramebufferPixelLocalStorageSizeEXT = (C.GPFRAMEBUFFERPIXELLOCALSTORAGESIZEEXT)(getProcAddr("glFramebufferPixelLocalStorageSizeEXT"))
	gpFramebufferRenderbuffer = (C.GPFRAMEBUFFERRENDERBUFFER)(getProcAddr("glFramebufferRenderbuffer"))
	if gpFramebufferRenderbuffer == nil {
		return errors.New("glFramebufferRenderbuffer")
	}
	gpFramebufferSampleLocationsfvNV = (C.GPFRAMEBUFFERSAMPLELOCATIONSFVNV)(getProcAddr("glFramebufferSampleLocationsfvNV"))
	gpFramebufferTexture2D = (C.GPFRAMEBUFFERTEXTURE2D)(getProcAddr("glFramebufferTexture2D"))
	if gpFramebufferTexture2D == nil {
		return errors.New("glFramebufferTexture2D")
	}
	gpFramebufferTexture2DDownsampleIMG = (C.GPFRAMEBUFFERTEXTURE2DDOWNSAMPLEIMG)(getProcAddr("glFramebufferTexture2DDownsampleIMG"))
	gpFramebufferTexture2DMultisampleEXT = (C.GPFRAMEBUFFERTEXTURE2DMULTISAMPLEEXT)(getProcAddr("glFramebufferTexture2DMultisampleEXT"))
	gpFramebufferTexture2DMultisampleIMG = (C.GPFRAMEBUFFERTEXTURE2DMULTISAMPLEIMG)(getProcAddr("glFramebufferTexture2DMultisampleIMG"))
	gpFramebufferTexture3DOES = (C.GPFRAMEBUFFERTEXTURE3DOES)(getProcAddr("glFramebufferTexture3DOES"))
	gpFramebufferTextureEXT = (C.GPFRAMEBUFFERTEXTUREEXT)(getProcAddr("glFramebufferTextureEXT"))
	gpFramebufferTextureLayer = (C.GPFRAMEBUFFERTEXTURELAYER)(getProcAddr("glFramebufferTextureLayer"))
	if gpFramebufferTextureLayer == nil {
		return errors.New("glFramebufferTextureLayer")
	}
	gpFramebufferTextureLayerDownsampleIMG = (C.GPFRAMEBUFFERTEXTURELAYERDOWNSAMPLEIMG)(getProcAddr("glFramebufferTextureLayerDownsampleIMG"))
	gpFramebufferTextureMultisampleMultiviewOVR = (C.GPFRAMEBUFFERTEXTUREMULTISAMPLEMULTIVIEWOVR)(getProcAddr("glFramebufferTextureMultisampleMultiviewOVR"))
	gpFramebufferTextureMultiviewOVR = (C.GPFRAMEBUFFERTEXTUREMULTIVIEWOVR)(getProcAddr("glFramebufferTextureMultiviewOVR"))
	gpFramebufferTextureOES = (C.GPFRAMEBUFFERTEXTUREOES)(getProcAddr("glFramebufferTextureOES"))
	gpFrontFace = (C.GPFRONTFACE)(getProcAddr("glFrontFace"))
	if gpFrontFace == nil {
		return errors.New("glFrontFace")
	}
	gpGenBuffers = (C.GPGENBUFFERS)(getProcAddr("glGenBuffers"))
	if gpGenBuffers == nil {
		return errors.New("glGenBuffers")
	}
	gpGenFencesNV = (C.GPGENFENCESNV)(getProcAddr("glGenFencesNV"))
	gpGenFramebuffers = (C.GPGENFRAMEBUFFERS)(getProcAddr("glGenFramebuffers"))
	if gpGenFramebuffers == nil {
		return errors.New("glGenFramebuffers")
	}
	gpGenPathsNV = (C.GPGENPATHSNV)(getProcAddr("glGenPathsNV"))
	gpGenPerfMonitorsAMD = (C.GPGENPERFMONITORSAMD)(getProcAddr("glGenPerfMonitorsAMD"))
	gpGenProgramPipelines = (C.GPGENPROGRAMPIPELINES)(getProcAddr("glGenProgramPipelines"))
	if gpGenProgramPipelines == nil {
		return errors.New("glGenProgramPipelines")
	}
	gpGenProgramPipelinesEXT = (C.GPGENPROGRAMPIPELINESEXT)(getProcAddr("glGenProgramPipelinesEXT"))
	gpGenQueries = (C.GPGENQUERIES)(getProcAddr("glGenQueries"))
	if gpGenQueries == nil {
		return errors.New("glGenQueries")
	}
	gpGenQueriesEXT = (C.GPGENQUERIESEXT)(getProcAddr("glGenQueriesEXT"))
	gpGenRenderbuffers = (C.GPGENRENDERBUFFERS)(getProcAddr("glGenRenderbuffers"))
	if gpGenRenderbuffers == nil {
		return errors.New("glGenRenderbuffers")
	}
	gpGenSamplers = (C.GPGENSAMPLERS)(getProcAddr("glGenSamplers"))
	if gpGenSamplers == nil {
		return errors.New("glGenSamplers")
	}
	gpGenSemaphoresEXT = (C.GPGENSEMAPHORESEXT)(getProcAddr("glGenSemaphoresEXT"))
	gpGenTextures = (C.GPGENTEXTURES)(getProcAddr("glGenTextures"))
	if gpGenTextures == nil {
		return errors.New("glGenTextures")
	}
	gpGenTransformFeedbacks = (C.GPGENTRANSFORMFEEDBACKS)(getProcAddr("glGenTransformFeedbacks"))
	if gpGenTransformFeedbacks == nil {
		return errors.New("glGenTransformFeedbacks")
	}
	gpGenVertexArrays = (C.GPGENVERTEXARRAYS)(getProcAddr("glGenVertexArrays"))
	if gpGenVertexArrays == nil {
		return errors.New("glGenVertexArrays")
	}
	gpGenVertexArraysOES = (C.GPGENVERTEXARRAYSOES)(getProcAddr("glGenVertexArraysOES"))
	gpGenerateMipmap = (C.GPGENERATEMIPMAP)(getProcAddr("glGenerateMipmap"))
	if gpGenerateMipmap == nil {
		return errors.New("glGenerateMipmap")
	}
	gpGetActiveAttrib = (C.GPGETACTIVEATTRIB)(getProcAddr("glGetActiveAttrib"))
	if gpGetActiveAttrib == nil {
		return errors.New("glGetActiveAttrib")
	}
	gpGetActiveUniform = (C.GPGETACTIVEUNIFORM)(getProcAddr("glGetActiveUniform"))
	if gpGetActiveUniform == nil {
		return errors.New("glGetActiveUniform")
	}
	gpGetActiveUniformBlockName = (C.GPGETACTIVEUNIFORMBLOCKNAME)(getProcAddr("glGetActiveUniformBlockName"))
	if gpGetActiveUniformBlockName == nil {
		return errors.New("glGetActiveUniformBlockName")
	}
	gpGetActiveUniformBlockiv = (C.GPGETACTIVEUNIFORMBLOCKIV)(getProcAddr("glGetActiveUniformBlockiv"))
	if gpGetActiveUniformBlockiv == nil {
		return errors.New("glGetActiveUniformBlockiv")
	}
	gpGetActiveUniformsiv = (C.GPGETACTIVEUNIFORMSIV)(getProcAddr("glGetActiveUniformsiv"))
	if gpGetActiveUniformsiv == nil {
		return errors.New("glGetActiveUniformsiv")
	}
	gpGetAttachedShaders = (C.GPGETATTACHEDSHADERS)(getProcAddr("glGetAttachedShaders"))
	if gpGetAttachedShaders == nil {
		return errors.New("glGetAttachedShaders")
	}
	gpGetAttribLocation = (C.GPGETATTRIBLOCATION)(getProcAddr("glGetAttribLocation"))
	if gpGetAttribLocation == nil {
		return errors.New("glGetAttribLocation")
	}
	gpGetBooleani_v = (C.GPGETBOOLEANI_V)(getProcAddr("glGetBooleani_v"))
	if gpGetBooleani_v == nil {
		return errors.New("glGetBooleani_v")
	}
	gpGetBooleanv = (C.GPGETBOOLEANV)(getProcAddr("glGetBooleanv"))
	if gpGetBooleanv == nil {
		return errors.New("glGetBooleanv")
	}
	gpGetBufferParameteri64v = (C.GPGETBUFFERPARAMETERI64V)(getProcAddr("glGetBufferParameteri64v"))
	if gpGetBufferParameteri64v == nil {
		return errors.New("glGetBufferParameteri64v")
	}
	gpGetBufferParameteriv = (C.GPGETBUFFERPARAMETERIV)(getProcAddr("glGetBufferParameteriv"))
	if gpGetBufferParameteriv == nil {
		return errors.New("glGetBufferParameteriv")
	}
	gpGetBufferPointerv = (C.GPGETBUFFERPOINTERV)(getProcAddr("glGetBufferPointerv"))
	if gpGetBufferPointerv == nil {
		return errors.New("glGetBufferPointerv")
	}
	gpGetBufferPointervOES = (C.GPGETBUFFERPOINTERVOES)(getProcAddr("glGetBufferPointervOES"))
	gpGetCoverageModulationTableNV = (C.GPGETCOVERAGEMODULATIONTABLENV)(getProcAddr("glGetCoverageModulationTableNV"))
	gpGetDebugMessageLog = (C.GPGETDEBUGMESSAGELOG)(getProcAddr("glGetDebugMessageLog"))
	gpGetDebugMessageLogKHR = (C.GPGETDEBUGMESSAGELOGKHR)(getProcAddr("glGetDebugMessageLogKHR"))
	gpGetDriverControlStringQCOM = (C.GPGETDRIVERCONTROLSTRINGQCOM)(getProcAddr("glGetDriverControlStringQCOM"))
	gpGetDriverControlsQCOM = (C.GPGETDRIVERCONTROLSQCOM)(getProcAddr("glGetDriverControlsQCOM"))
	gpGetError = (C.GPGETERROR)(getProcAddr("glGetError"))
	if gpGetError == nil {
		return errors.New("glGetError")
	}
	gpGetFenceivNV = (C.GPGETFENCEIVNV)(getProcAddr("glGetFenceivNV"))
	gpGetFirstPerfQueryIdINTEL = (C.GPGETFIRSTPERFQUERYIDINTEL)(getProcAddr("glGetFirstPerfQueryIdINTEL"))
	gpGetFloati_vNV = (C.GPGETFLOATI_VNV)(getProcAddr("glGetFloati_vNV"))
	gpGetFloati_vOES = (C.GPGETFLOATI_VOES)(getProcAddr("glGetFloati_vOES"))
	gpGetFloatv = (C.GPGETFLOATV)(getProcAddr("glGetFloatv"))
	if gpGetFloatv == nil {
		return errors.New("glGetFloatv")
	}
	gpGetFragDataIndexEXT = (C.GPGETFRAGDATAINDEXEXT)(getProcAddr("glGetFragDataIndexEXT"))
	gpGetFragDataLocation = (C.GPGETFRAGDATALOCATION)(getProcAddr("glGetFragDataLocation"))
	if gpGetFragDataLocation == nil {
		return errors.New("glGetFragDataLocation")
	}
	gpGetFramebufferAttachmentParameteriv = (C.GPGETFRAMEBUFFERATTACHMENTPARAMETERIV)(getProcAddr("glGetFramebufferAttachmentParameteriv"))
	if gpGetFramebufferAttachmentParameteriv == nil {
		return errors.New("glGetFramebufferAttachmentParameteriv")
	}
	gpGetFramebufferParameteriv = (C.GPGETFRAMEBUFFERPARAMETERIV)(getProcAddr("glGetFramebufferParameteriv"))
	if gpGetFramebufferParameteriv == nil {
		return errors.New("glGetFramebufferParameteriv")
	}
	gpGetFramebufferPixelLocalStorageSizeEXT = (C.GPGETFRAMEBUFFERPIXELLOCALSTORAGESIZEEXT)(getProcAddr("glGetFramebufferPixelLocalStorageSizeEXT"))
	gpGetGraphicsResetStatus = (C.GPGETGRAPHICSRESETSTATUS)(getProcAddr("glGetGraphicsResetStatus"))
	gpGetGraphicsResetStatusEXT = (C.GPGETGRAPHICSRESETSTATUSEXT)(getProcAddr("glGetGraphicsResetStatusEXT"))
	gpGetGraphicsResetStatusKHR = (C.GPGETGRAPHICSRESETSTATUSKHR)(getProcAddr("glGetGraphicsResetStatusKHR"))
	gpGetImageHandleNV = (C.GPGETIMAGEHANDLENV)(getProcAddr("glGetImageHandleNV"))
	gpGetInteger64i_v = (C.GPGETINTEGER64I_V)(getProcAddr("glGetInteger64i_v"))
	if gpGetInteger64i_v == nil {
		return errors.New("glGetInteger64i_v")
	}
	gpGetInteger64v = (C.GPGETINTEGER64V)(getProcAddr("glGetInteger64v"))
	if gpGetInteger64v == nil {
		return errors.New("glGetInteger64v")
	}
	gpGetInteger64vAPPLE = (C.GPGETINTEGER64VAPPLE)(getProcAddr("glGetInteger64vAPPLE"))
	gpGetIntegeri_v = (C.GPGETINTEGERI_V)(getProcAddr("glGetIntegeri_v"))
	if gpGetIntegeri_v == nil {
		return errors.New("glGetIntegeri_v")
	}
	gpGetIntegeri_vEXT = (C.GPGETINTEGERI_VEXT)(getProcAddr("glGetIntegeri_vEXT"))
	gpGetIntegerv = (C.GPGETINTEGERV)(getProcAddr("glGetIntegerv"))
	if gpGetIntegerv == nil {
		return errors.New("glGetIntegerv")
	}
	gpGetInternalformatSampleivNV = (C.GPGETINTERNALFORMATSAMPLEIVNV)(getProcAddr("glGetInternalformatSampleivNV"))
	gpGetInternalformativ = (C.GPGETINTERNALFORMATIV)(getProcAddr("glGetInternalformativ"))
	if gpGetInternalformativ == nil {
		return errors.New("glGetInternalformativ")
	}
	gpGetMemoryObjectParameterivEXT = (C.GPGETMEMORYOBJECTPARAMETERIVEXT)(getProcAddr("glGetMemoryObjectParameterivEXT"))
	gpGetMultisamplefv = (C.GPGETMULTISAMPLEFV)(getProcAddr("glGetMultisamplefv"))
	if gpGetMultisamplefv == nil {
		return errors.New("glGetMultisamplefv")
	}
	gpGetNextPerfQueryIdINTEL = (C.GPGETNEXTPERFQUERYIDINTEL)(getProcAddr("glGetNextPerfQueryIdINTEL"))
	gpGetObjectLabel = (C.GPGETOBJECTLABEL)(getProcAddr("glGetObjectLabel"))
	gpGetObjectLabelEXT = (C.GPGETOBJECTLABELEXT)(getProcAddr("glGetObjectLabelEXT"))
	gpGetObjectLabelKHR = (C.GPGETOBJECTLABELKHR)(getProcAddr("glGetObjectLabelKHR"))
	gpGetObjectPtrLabel = (C.GPGETOBJECTPTRLABEL)(getProcAddr("glGetObjectPtrLabel"))
	gpGetObjectPtrLabelKHR = (C.GPGETOBJECTPTRLABELKHR)(getProcAddr("glGetObjectPtrLabelKHR"))
	gpGetPathCommandsNV = (C.GPGETPATHCOMMANDSNV)(getProcAddr("glGetPathCommandsNV"))
	gpGetPathCoordsNV = (C.GPGETPATHCOORDSNV)(getProcAddr("glGetPathCoordsNV"))
	gpGetPathDashArrayNV = (C.GPGETPATHDASHARRAYNV)(getProcAddr("glGetPathDashArrayNV"))
	gpGetPathLengthNV = (C.GPGETPATHLENGTHNV)(getProcAddr("glGetPathLengthNV"))
	gpGetPathMetricRangeNV = (C.GPGETPATHMETRICRANGENV)(getProcAddr("glGetPathMetricRangeNV"))
	gpGetPathMetricsNV = (C.GPGETPATHMETRICSNV)(getProcAddr("glGetPathMetricsNV"))
	gpGetPathParameterfvNV = (C.GPGETPATHPARAMETERFVNV)(getProcAddr("glGetPathParameterfvNV"))
	gpGetPathParameterivNV = (C.GPGETPATHPARAMETERIVNV)(getProcAddr("glGetPathParameterivNV"))
	gpGetPathSpacingNV = (C.GPGETPATHSPACINGNV)(getProcAddr("glGetPathSpacingNV"))
	gpGetPerfCounterInfoINTEL = (C.GPGETPERFCOUNTERINFOINTEL)(getProcAddr("glGetPerfCounterInfoINTEL"))
	gpGetPerfMonitorCounterDataAMD = (C.GPGETPERFMONITORCOUNTERDATAAMD)(getProcAddr("glGetPerfMonitorCounterDataAMD"))
	gpGetPerfMonitorCounterInfoAMD = (C.GPGETPERFMONITORCOUNTERINFOAMD)(getProcAddr("glGetPerfMonitorCounterInfoAMD"))
	gpGetPerfMonitorCounterStringAMD = (C.GPGETPERFMONITORCOUNTERSTRINGAMD)(getProcAddr("glGetPerfMonitorCounterStringAMD"))
	gpGetPerfMonitorCountersAMD = (C.GPGETPERFMONITORCOUNTERSAMD)(getProcAddr("glGetPerfMonitorCountersAMD"))
	gpGetPerfMonitorGroupStringAMD = (C.GPGETPERFMONITORGROUPSTRINGAMD)(getProcAddr("glGetPerfMonitorGroupStringAMD"))
	gpGetPerfMonitorGroupsAMD = (C.GPGETPERFMONITORGROUPSAMD)(getProcAddr("glGetPerfMonitorGroupsAMD"))
	gpGetPerfQueryDataINTEL = (C.GPGETPERFQUERYDATAINTEL)(getProcAddr("glGetPerfQueryDataINTEL"))
	gpGetPerfQueryIdByNameINTEL = (C.GPGETPERFQUERYIDBYNAMEINTEL)(getProcAddr("glGetPerfQueryIdByNameINTEL"))
	gpGetPerfQueryInfoINTEL = (C.GPGETPERFQUERYINFOINTEL)(getProcAddr("glGetPerfQueryInfoINTEL"))
	gpGetPointerv = (C.GPGETPOINTERV)(getProcAddr("glGetPointerv"))
	gpGetPointervKHR = (C.GPGETPOINTERVKHR)(getProcAddr("glGetPointervKHR"))
	gpGetProgramBinary = (C.GPGETPROGRAMBINARY)(getProcAddr("glGetProgramBinary"))
	if gpGetProgramBinary == nil {
		return errors.New("glGetProgramBinary")
	}
	gpGetProgramBinaryOES = (C.GPGETPROGRAMBINARYOES)(getProcAddr("glGetProgramBinaryOES"))
	gpGetProgramInfoLog = (C.GPGETPROGRAMINFOLOG)(getProcAddr("glGetProgramInfoLog"))
	if gpGetProgramInfoLog == nil {
		return errors.New("glGetProgramInfoLog")
	}
	gpGetProgramInterfaceiv = (C.GPGETPROGRAMINTERFACEIV)(getProcAddr("glGetProgramInterfaceiv"))
	if gpGetProgramInterfaceiv == nil {
		return errors.New("glGetProgramInterfaceiv")
	}
	gpGetProgramPipelineInfoLog = (C.GPGETPROGRAMPIPELINEINFOLOG)(getProcAddr("glGetProgramPipelineInfoLog"))
	if gpGetProgramPipelineInfoLog == nil {
		return errors.New("glGetProgramPipelineInfoLog")
	}
	gpGetProgramPipelineInfoLogEXT = (C.GPGETPROGRAMPIPELINEINFOLOGEXT)(getProcAddr("glGetProgramPipelineInfoLogEXT"))
	gpGetProgramPipelineiv = (C.GPGETPROGRAMPIPELINEIV)(getProcAddr("glGetProgramPipelineiv"))
	if gpGetProgramPipelineiv == nil {
		return errors.New("glGetProgramPipelineiv")
	}
	gpGetProgramPipelineivEXT = (C.GPGETPROGRAMPIPELINEIVEXT)(getProcAddr("glGetProgramPipelineivEXT"))
	gpGetProgramResourceIndex = (C.GPGETPROGRAMRESOURCEINDEX)(getProcAddr("glGetProgramResourceIndex"))
	if gpGetProgramResourceIndex == nil {
		return errors.New("glGetProgramResourceIndex")
	}
	gpGetProgramResourceLocation = (C.GPGETPROGRAMRESOURCELOCATION)(getProcAddr("glGetProgramResourceLocation"))
	if gpGetProgramResourceLocation == nil {
		return errors.New("glGetProgramResourceLocation")
	}
	gpGetProgramResourceLocationIndexEXT = (C.GPGETPROGRAMRESOURCELOCATIONINDEXEXT)(getProcAddr("glGetProgramResourceLocationIndexEXT"))
	gpGetProgramResourceName = (C.GPGETPROGRAMRESOURCENAME)(getProcAddr("glGetProgramResourceName"))
	if gpGetProgramResourceName == nil {
		return errors.New("glGetProgramResourceName")
	}
	gpGetProgramResourcefvNV = (C.GPGETPROGRAMRESOURCEFVNV)(getProcAddr("glGetProgramResourcefvNV"))
	gpGetProgramResourceiv = (C.GPGETPROGRAMRESOURCEIV)(getProcAddr("glGetProgramResourceiv"))
	if gpGetProgramResourceiv == nil {
		return errors.New("glGetProgramResourceiv")
	}
	gpGetProgramiv = (C.GPGETPROGRAMIV)(getProcAddr("glGetProgramiv"))
	if gpGetProgramiv == nil {
		return errors.New("glGetProgramiv")
	}
	gpGetQueryObjecti64vEXT = (C.GPGETQUERYOBJECTI64VEXT)(getProcAddr("glGetQueryObjecti64vEXT"))
	gpGetQueryObjectivEXT = (C.GPGETQUERYOBJECTIVEXT)(getProcAddr("glGetQueryObjectivEXT"))
	gpGetQueryObjectui64vEXT = (C.GPGETQUERYOBJECTUI64VEXT)(getProcAddr("glGetQueryObjectui64vEXT"))
	gpGetQueryObjectuiv = (C.GPGETQUERYOBJECTUIV)(getProcAddr("glGetQueryObjectuiv"))
	if gpGetQueryObjectuiv == nil {
		return errors.New("glGetQueryObjectuiv")
	}
	gpGetQueryObjectuivEXT = (C.GPGETQUERYOBJECTUIVEXT)(getProcAddr("glGetQueryObjectuivEXT"))
	gpGetQueryiv = (C.GPGETQUERYIV)(getProcAddr("glGetQueryiv"))
	if gpGetQueryiv == nil {
		return errors.New("glGetQueryiv")
	}
	gpGetQueryivEXT = (C.GPGETQUERYIVEXT)(getProcAddr("glGetQueryivEXT"))
	gpGetRenderbufferParameteriv = (C.GPGETRENDERBUFFERPARAMETERIV)(getProcAddr("glGetRenderbufferParameteriv"))
	if gpGetRenderbufferParameteriv == nil {
		return errors.New("glGetRenderbufferParameteriv")
	}
	gpGetSamplerParameterIivEXT = (C.GPGETSAMPLERPARAMETERIIVEXT)(getProcAddr("glGetSamplerParameterIivEXT"))
	gpGetSamplerParameterIivOES = (C.GPGETSAMPLERPARAMETERIIVOES)(getProcAddr("glGetSamplerParameterIivOES"))
	gpGetSamplerParameterIuivEXT = (C.GPGETSAMPLERPARAMETERIUIVEXT)(getProcAddr("glGetSamplerParameterIuivEXT"))
	gpGetSamplerParameterIuivOES = (C.GPGETSAMPLERPARAMETERIUIVOES)(getProcAddr("glGetSamplerParameterIuivOES"))
	gpGetSamplerParameterfv = (C.GPGETSAMPLERPARAMETERFV)(getProcAddr("glGetSamplerParameterfv"))
	if gpGetSamplerParameterfv == nil {
		return errors.New("glGetSamplerParameterfv")
	}
	gpGetSamplerParameteriv = (C.GPGETSAMPLERPARAMETERIV)(getProcAddr("glGetSamplerParameteriv"))
	if gpGetSamplerParameteriv == nil {
		return errors.New("glGetSamplerParameteriv")
	}
	gpGetSemaphoreParameterui64vEXT = (C.GPGETSEMAPHOREPARAMETERUI64VEXT)(getProcAddr("glGetSemaphoreParameterui64vEXT"))
	gpGetShaderInfoLog = (C.GPGETSHADERINFOLOG)(getProcAddr("glGetShaderInfoLog"))
	if gpGetShaderInfoLog == nil {
		return errors.New("glGetShaderInfoLog")
	}
	gpGetShaderPrecisionFormat = (C.GPGETSHADERPRECISIONFORMAT)(getProcAddr("glGetShaderPrecisionFormat"))
	if gpGetShaderPrecisionFormat == nil {
		return errors.New("glGetShaderPrecisionFormat")
	}
	gpGetShaderSource = (C.GPGETSHADERSOURCE)(getProcAddr("glGetShaderSource"))
	if gpGetShaderSource == nil {
		return errors.New("glGetShaderSource")
	}
	gpGetShaderiv = (C.GPGETSHADERIV)(getProcAddr("glGetShaderiv"))
	if gpGetShaderiv == nil {
		return errors.New("glGetShaderiv")
	}
	gpGetString = (C.GPGETSTRING)(getProcAddr("glGetString"))
	if gpGetString == nil {
		return errors.New("glGetString")
	}
	gpGetStringi = (C.GPGETSTRINGI)(getProcAddr("glGetStringi"))
	if gpGetStringi == nil {
		return errors.New("glGetStringi")
	}
	gpGetSynciv = (C.GPGETSYNCIV)(getProcAddr("glGetSynciv"))
	if gpGetSynciv == nil {
		return errors.New("glGetSynciv")
	}
	gpGetSyncivAPPLE = (C.GPGETSYNCIVAPPLE)(getProcAddr("glGetSyncivAPPLE"))
	gpGetTexLevelParameterfv = (C.GPGETTEXLEVELPARAMETERFV)(getProcAddr("glGetTexLevelParameterfv"))
	if gpGetTexLevelParameterfv == nil {
		return errors.New("glGetTexLevelParameterfv")
	}
	gpGetTexLevelParameteriv = (C.GPGETTEXLEVELPARAMETERIV)(getProcAddr("glGetTexLevelParameteriv"))
	if gpGetTexLevelParameteriv == nil {
		return errors.New("glGetTexLevelParameteriv")
	}
	gpGetTexParameterIivEXT = (C.GPGETTEXPARAMETERIIVEXT)(getProcAddr("glGetTexParameterIivEXT"))
	gpGetTexParameterIivOES = (C.GPGETTEXPARAMETERIIVOES)(getProcAddr("glGetTexParameterIivOES"))
	gpGetTexParameterIuivEXT = (C.GPGETTEXPARAMETERIUIVEXT)(getProcAddr("glGetTexParameterIuivEXT"))
	gpGetTexParameterIuivOES = (C.GPGETTEXPARAMETERIUIVOES)(getProcAddr("glGetTexParameterIuivOES"))
	gpGetTexParameterfv = (C.GPGETTEXPARAMETERFV)(getProcAddr("glGetTexParameterfv"))
	if gpGetTexParameterfv == nil {
		return errors.New("glGetTexParameterfv")
	}
	gpGetTexParameteriv = (C.GPGETTEXPARAMETERIV)(getProcAddr("glGetTexParameteriv"))
	if gpGetTexParameteriv == nil {
		return errors.New("glGetTexParameteriv")
	}
	gpGetTextureHandleIMG = (C.GPGETTEXTUREHANDLEIMG)(getProcAddr("glGetTextureHandleIMG"))
	gpGetTextureHandleNV = (C.GPGETTEXTUREHANDLENV)(getProcAddr("glGetTextureHandleNV"))
	gpGetTextureSamplerHandleIMG = (C.GPGETTEXTURESAMPLERHANDLEIMG)(getProcAddr("glGetTextureSamplerHandleIMG"))
	gpGetTextureSamplerHandleNV = (C.GPGETTEXTURESAMPLERHANDLENV)(getProcAddr("glGetTextureSamplerHandleNV"))
	gpGetTransformFeedbackVarying = (C.GPGETTRANSFORMFEEDBACKVARYING)(getProcAddr("glGetTransformFeedbackVarying"))
	if gpGetTransformFeedbackVarying == nil {
		return errors.New("glGetTransformFeedbackVarying")
	}
	gpGetTranslatedShaderSourceANGLE = (C.GPGETTRANSLATEDSHADERSOURCEANGLE)(getProcAddr("glGetTranslatedShaderSourceANGLE"))
	gpGetUniformBlockIndex = (C.GPGETUNIFORMBLOCKINDEX)(getProcAddr("glGetUniformBlockIndex"))
	if gpGetUniformBlockIndex == nil {
		return errors.New("glGetUniformBlockIndex")
	}
	gpGetUniformIndices = (C.GPGETUNIFORMINDICES)(getProcAddr("glGetUniformIndices"))
	if gpGetUniformIndices == nil {
		return errors.New("glGetUniformIndices")
	}
	gpGetUniformLocation = (C.GPGETUNIFORMLOCATION)(getProcAddr("glGetUniformLocation"))
	if gpGetUniformLocation == nil {
		return errors.New("glGetUniformLocation")
	}
	gpGetUniformfv = (C.GPGETUNIFORMFV)(getProcAddr("glGetUniformfv"))
	if gpGetUniformfv == nil {
		return errors.New("glGetUniformfv")
	}
	gpGetUniformi64vNV = (C.GPGETUNIFORMI64VNV)(getProcAddr("glGetUniformi64vNV"))
	gpGetUniformiv = (C.GPGETUNIFORMIV)(getProcAddr("glGetUniformiv"))
	if gpGetUniformiv == nil {
		return errors.New("glGetUniformiv")
	}
	gpGetUniformuiv = (C.GPGETUNIFORMUIV)(getProcAddr("glGetUniformuiv"))
	if gpGetUniformuiv == nil {
		return errors.New("glGetUniformuiv")
	}
	gpGetUnsignedBytei_vEXT = (C.GPGETUNSIGNEDBYTEI_VEXT)(getProcAddr("glGetUnsignedBytei_vEXT"))
	gpGetUnsignedBytevEXT = (C.GPGETUNSIGNEDBYTEVEXT)(getProcAddr("glGetUnsignedBytevEXT"))
	gpGetVertexAttribIiv = (C.GPGETVERTEXATTRIBIIV)(getProcAddr("glGetVertexAttribIiv"))
	if gpGetVertexAttribIiv == nil {
		return errors.New("glGetVertexAttribIiv")
	}
	gpGetVertexAttribIuiv = (C.GPGETVERTEXATTRIBIUIV)(getProcAddr("glGetVertexAttribIuiv"))
	if gpGetVertexAttribIuiv == nil {
		return errors.New("glGetVertexAttribIuiv")
	}
	gpGetVertexAttribPointerv = (C.GPGETVERTEXATTRIBPOINTERV)(getProcAddr("glGetVertexAttribPointerv"))
	if gpGetVertexAttribPointerv == nil {
		return errors.New("glGetVertexAttribPointerv")
	}
	gpGetVertexAttribfv = (C.GPGETVERTEXATTRIBFV)(getProcAddr("glGetVertexAttribfv"))
	if gpGetVertexAttribfv == nil {
		return errors.New("glGetVertexAttribfv")
	}
	gpGetVertexAttribiv = (C.GPGETVERTEXATTRIBIV)(getProcAddr("glGetVertexAttribiv"))
	if gpGetVertexAttribiv == nil {
		return errors.New("glGetVertexAttribiv")
	}
	gpGetVkProcAddrNV = (C.GPGETVKPROCADDRNV)(getProcAddr("glGetVkProcAddrNV"))
	gpGetnUniformfv = (C.GPGETNUNIFORMFV)(getProcAddr("glGetnUniformfv"))
	gpGetnUniformfvEXT = (C.GPGETNUNIFORMFVEXT)(getProcAddr("glGetnUniformfvEXT"))
	gpGetnUniformfvKHR = (C.GPGETNUNIFORMFVKHR)(getProcAddr("glGetnUniformfvKHR"))
	gpGetnUniformiv = (C.GPGETNUNIFORMIV)(getProcAddr("glGetnUniformiv"))
	gpGetnUniformivEXT = (C.GPGETNUNIFORMIVEXT)(getProcAddr("glGetnUniformivEXT"))
	gpGetnUniformivKHR = (C.GPGETNUNIFORMIVKHR)(getProcAddr("glGetnUniformivKHR"))
	gpGetnUniformuiv = (C.GPGETNUNIFORMUIV)(getProcAddr("glGetnUniformuiv"))
	gpGetnUniformuivKHR = (C.GPGETNUNIFORMUIVKHR)(getProcAddr("glGetnUniformuivKHR"))
	gpHint = (C.GPHINT)(getProcAddr("glHint"))
	if gpHint == nil {
		return errors.New("glHint")
	}
	gpImportMemoryFdEXT = (C.GPIMPORTMEMORYFDEXT)(getProcAddr("glImportMemoryFdEXT"))
	gpImportMemoryWin32HandleEXT = (C.GPIMPORTMEMORYWIN32HANDLEEXT)(getProcAddr("glImportMemoryWin32HandleEXT"))
	gpImportMemoryWin32NameEXT = (C.GPIMPORTMEMORYWIN32NAMEEXT)(getProcAddr("glImportMemoryWin32NameEXT"))
	gpImportSemaphoreFdEXT = (C.GPIMPORTSEMAPHOREFDEXT)(getProcAddr("glImportSemaphoreFdEXT"))
	gpImportSemaphoreWin32HandleEXT = (C.GPIMPORTSEMAPHOREWIN32HANDLEEXT)(getProcAddr("glImportSemaphoreWin32HandleEXT"))
	gpImportSemaphoreWin32NameEXT = (C.GPIMPORTSEMAPHOREWIN32NAMEEXT)(getProcAddr("glImportSemaphoreWin32NameEXT"))
	gpInsertEventMarkerEXT = (C.GPINSERTEVENTMARKEREXT)(getProcAddr("glInsertEventMarkerEXT"))
	gpInterpolatePathsNV = (C.GPINTERPOLATEPATHSNV)(getProcAddr("glInterpolatePathsNV"))
	gpInvalidateFramebuffer = (C.GPINVALIDATEFRAMEBUFFER)(getProcAddr("glInvalidateFramebuffer"))
	if gpInvalidateFramebuffer == nil {
		return errors.New("glInvalidateFramebuffer")
	}
	gpInvalidateSubFramebuffer = (C.GPINVALIDATESUBFRAMEBUFFER)(getProcAddr("glInvalidateSubFramebuffer"))
	if gpInvalidateSubFramebuffer == nil {
		return errors.New("glInvalidateSubFramebuffer")
	}
	gpIsBuffer = (C.GPISBUFFER)(getProcAddr("glIsBuffer"))
	if gpIsBuffer == nil {
		return errors.New("glIsBuffer")
	}
	gpIsEnabled = (C.GPISENABLED)(getProcAddr("glIsEnabled"))
	if gpIsEnabled == nil {
		return errors.New("glIsEnabled")
	}
	gpIsEnablediEXT = (C.GPISENABLEDIEXT)(getProcAddr("glIsEnablediEXT"))
	gpIsEnablediNV = (C.GPISENABLEDINV)(getProcAddr("glIsEnablediNV"))
	gpIsEnablediOES = (C.GPISENABLEDIOES)(getProcAddr("glIsEnablediOES"))
	gpIsFenceNV = (C.GPISFENCENV)(getProcAddr("glIsFenceNV"))
	gpIsFramebuffer = (C.GPISFRAMEBUFFER)(getProcAddr("glIsFramebuffer"))
	if gpIsFramebuffer == nil {
		return errors.New("glIsFramebuffer")
	}
	gpIsImageHandleResidentNV = (C.GPISIMAGEHANDLERESIDENTNV)(getProcAddr("glIsImageHandleResidentNV"))
	gpIsMemoryObjectEXT = (C.GPISMEMORYOBJECTEXT)(getProcAddr("glIsMemoryObjectEXT"))
	gpIsPathNV = (C.GPISPATHNV)(getProcAddr("glIsPathNV"))
	gpIsPointInFillPathNV = (C.GPISPOINTINFILLPATHNV)(getProcAddr("glIsPointInFillPathNV"))
	gpIsPointInStrokePathNV = (C.GPISPOINTINSTROKEPATHNV)(getProcAddr("glIsPointInStrokePathNV"))
	gpIsProgram = (C.GPISPROGRAM)(getProcAddr("glIsProgram"))
	if gpIsProgram == nil {
		return errors.New("glIsProgram")
	}
	gpIsProgramPipeline = (C.GPISPROGRAMPIPELINE)(getProcAddr("glIsProgramPipeline"))
	if gpIsProgramPipeline == nil {
		return errors.New("glIsProgramPipeline")
	}
	gpIsProgramPipelineEXT = (C.GPISPROGRAMPIPELINEEXT)(getProcAddr("glIsProgramPipelineEXT"))
	gpIsQuery = (C.GPISQUERY)(getProcAddr("glIsQuery"))
	if gpIsQuery == nil {
		return errors.New("glIsQuery")
	}
	gpIsQueryEXT = (C.GPISQUERYEXT)(getProcAddr("glIsQueryEXT"))
	gpIsRenderbuffer = (C.GPISRENDERBUFFER)(getProcAddr("glIsRenderbuffer"))
	if gpIsRenderbuffer == nil {
		return errors.New("glIsRenderbuffer")
	}
	gpIsSampler = (C.GPISSAMPLER)(getProcAddr("glIsSampler"))
	if gpIsSampler == nil {
		return errors.New("glIsSampler")
	}
	gpIsSemaphoreEXT = (C.GPISSEMAPHOREEXT)(getProcAddr("glIsSemaphoreEXT"))
	gpIsShader = (C.GPISSHADER)(getProcAddr("glIsShader"))
	if gpIsShader == nil {
		return errors.New("glIsShader")
	}
	gpIsSync = (C.GPISSYNC)(getProcAddr("glIsSync"))
	if gpIsSync == nil {
		return errors.New("glIsSync")
	}
	gpIsSyncAPPLE = (C.GPISSYNCAPPLE)(getProcAddr("glIsSyncAPPLE"))
	gpIsTexture = (C.GPISTEXTURE)(getProcAddr("glIsTexture"))
	if gpIsTexture == nil {
		return errors.New("glIsTexture")
	}
	gpIsTextureHandleResidentNV = (C.GPISTEXTUREHANDLERESIDENTNV)(getProcAddr("glIsTextureHandleResidentNV"))
	gpIsTransformFeedback = (C.GPISTRANSFORMFEEDBACK)(getProcAddr("glIsTransformFeedback"))
	if gpIsTransformFeedback == nil {
		return errors.New("glIsTransformFeedback")
	}
	gpIsVertexArray = (C.GPISVERTEXARRAY)(getProcAddr("glIsVertexArray"))
	if gpIsVertexArray == nil {
		return errors.New("glIsVertexArray")
	}
	gpIsVertexArrayOES = (C.GPISVERTEXARRAYOES)(getProcAddr("glIsVertexArrayOES"))
	gpLabelObjectEXT = (C.GPLABELOBJECTEXT)(getProcAddr("glLabelObjectEXT"))
	gpLineWidth = (C.GPLINEWIDTH)(getProcAddr("glLineWidth"))
	if gpLineWidth == nil {
		return errors.New("glLineWidth")
	}
	gpLinkProgram = (C.GPLINKPROGRAM)(getProcAddr("glLinkProgram"))
	if gpLinkProgram == nil {
		return errors.New("glLinkProgram")
	}
	gpMakeImageHandleNonResidentNV = (C.GPMAKEIMAGEHANDLENONRESIDENTNV)(getProcAddr("glMakeImageHandleNonResidentNV"))
	gpMakeImageHandleResidentNV = (C.GPMAKEIMAGEHANDLERESIDENTNV)(getProcAddr("glMakeImageHandleResidentNV"))
	gpMakeTextureHandleNonResidentNV = (C.GPMAKETEXTUREHANDLENONRESIDENTNV)(getProcAddr("glMakeTextureHandleNonResidentNV"))
	gpMakeTextureHandleResidentNV = (C.GPMAKETEXTUREHANDLERESIDENTNV)(getProcAddr("glMakeTextureHandleResidentNV"))
	gpMapBufferOES = (C.GPMAPBUFFEROES)(getProcAddr("glMapBufferOES"))
	gpMapBufferRange = (C.GPMAPBUFFERRANGE)(getProcAddr("glMapBufferRange"))
	if gpMapBufferRange == nil {
		return errors.New("glMapBufferRange")
	}
	gpMapBufferRangeEXT = (C.GPMAPBUFFERRANGEEXT)(getProcAddr("glMapBufferRangeEXT"))
	gpMatrixFrustumEXT = (C.GPMATRIXFRUSTUMEXT)(getProcAddr("glMatrixFrustumEXT"))
	gpMatrixLoad3x2fNV = (C.GPMATRIXLOAD3X2FNV)(getProcAddr("glMatrixLoad3x2fNV"))
	gpMatrixLoad3x3fNV = (C.GPMATRIXLOAD3X3FNV)(getProcAddr("glMatrixLoad3x3fNV"))
	gpMatrixLoadIdentityEXT = (C.GPMATRIXLOADIDENTITYEXT)(getProcAddr("glMatrixLoadIdentityEXT"))
	gpMatrixLoadTranspose3x3fNV = (C.GPMATRIXLOADTRANSPOSE3X3FNV)(getProcAddr("glMatrixLoadTranspose3x3fNV"))
	gpMatrixLoadTransposedEXT = (C.GPMATRIXLOADTRANSPOSEDEXT)(getProcAddr("glMatrixLoadTransposedEXT"))
	gpMatrixLoadTransposefEXT = (C.GPMATRIXLOADTRANSPOSEFEXT)(getProcAddr("glMatrixLoadTransposefEXT"))
	gpMatrixLoaddEXT = (C.GPMATRIXLOADDEXT)(getProcAddr("glMatrixLoaddEXT"))
	gpMatrixLoadfEXT = (C.GPMATRIXLOADFEXT)(getProcAddr("glMatrixLoadfEXT"))
	gpMatrixMult3x2fNV = (C.GPMATRIXMULT3X2FNV)(getProcAddr("glMatrixMult3x2fNV"))
	gpMatrixMult3x3fNV = (C.GPMATRIXMULT3X3FNV)(getProcAddr("glMatrixMult3x3fNV"))
	gpMatrixMultTranspose3x3fNV = (C.GPMATRIXMULTTRANSPOSE3X3FNV)(getProcAddr("glMatrixMultTranspose3x3fNV"))
	gpMatrixMultTransposedEXT = (C.GPMATRIXMULTTRANSPOSEDEXT)(getProcAddr("glMatrixMultTransposedEXT"))
	gpMatrixMultTransposefEXT = (C.GPMATRIXMULTTRANSPOSEFEXT)(getProcAddr("glMatrixMultTransposefEXT"))
	gpMatrixMultdEXT = (C.GPMATRIXMULTDEXT)(getProcAddr("glMatrixMultdEXT"))
	gpMatrixMultfEXT = (C.GPMATRIXMULTFEXT)(getProcAddr("glMatrixMultfEXT"))
	gpMatrixOrthoEXT = (C.GPMATRIXORTHOEXT)(getProcAddr("glMatrixOrthoEXT"))
	gpMatrixPopEXT = (C.GPMATRIXPOPEXT)(getProcAddr("glMatrixPopEXT"))
	gpMatrixPushEXT = (C.GPMATRIXPUSHEXT)(getProcAddr("glMatrixPushEXT"))
	gpMatrixRotatedEXT = (C.GPMATRIXROTATEDEXT)(getProcAddr("glMatrixRotatedEXT"))
	gpMatrixRotatefEXT = (C.GPMATRIXROTATEFEXT)(getProcAddr("glMatrixRotatefEXT"))
	gpMatrixScaledEXT = (C.GPMATRIXSCALEDEXT)(getProcAddr("glMatrixScaledEXT"))
	gpMatrixScalefEXT = (C.GPMATRIXSCALEFEXT)(getProcAddr("glMatrixScalefEXT"))
	gpMatrixTranslatedEXT = (C.GPMATRIXTRANSLATEDEXT)(getProcAddr("glMatrixTranslatedEXT"))
	gpMatrixTranslatefEXT = (C.GPMATRIXTRANSLATEFEXT)(getProcAddr("glMatrixTranslatefEXT"))
	gpMaxShaderCompilerThreadsKHR = (C.GPMAXSHADERCOMPILERTHREADSKHR)(getProcAddr("glMaxShaderCompilerThreadsKHR"))
	gpMemoryBarrier = (C.GPMEMORYBARRIER)(getProcAddr("glMemoryBarrier"))
	if gpMemoryBarrier == nil {
		return errors.New("glMemoryBarrier")
	}
	gpMemoryBarrierByRegion = (C.GPMEMORYBARRIERBYREGION)(getProcAddr("glMemoryBarrierByRegion"))
	if gpMemoryBarrierByRegion == nil {
		return errors.New("glMemoryBarrierByRegion")
	}
	gpMemoryObjectParameterivEXT = (C.GPMEMORYOBJECTPARAMETERIVEXT)(getProcAddr("glMemoryObjectParameterivEXT"))
	gpMinSampleShadingOES = (C.GPMINSAMPLESHADINGOES)(getProcAddr("glMinSampleShadingOES"))
	gpMultiDrawArraysEXT = (C.GPMULTIDRAWARRAYSEXT)(getProcAddr("glMultiDrawArraysEXT"))
	gpMultiDrawArraysIndirectEXT = (C.GPMULTIDRAWARRAYSINDIRECTEXT)(getProcAddr("glMultiDrawArraysIndirectEXT"))
	gpMultiDrawElementsBaseVertexEXT = (C.GPMULTIDRAWELEMENTSBASEVERTEXEXT)(getProcAddr("glMultiDrawElementsBaseVertexEXT"))
	gpMultiDrawElementsEXT = (C.GPMULTIDRAWELEMENTSEXT)(getProcAddr("glMultiDrawElementsEXT"))
	gpMultiDrawElementsIndirectEXT = (C.GPMULTIDRAWELEMENTSINDIRECTEXT)(getProcAddr("glMultiDrawElementsIndirectEXT"))
	gpNamedBufferStorageExternalEXT = (C.GPNAMEDBUFFERSTORAGEEXTERNALEXT)(getProcAddr("glNamedBufferStorageExternalEXT"))
	gpNamedBufferStorageMemEXT = (C.GPNAMEDBUFFERSTORAGEMEMEXT)(getProcAddr("glNamedBufferStorageMemEXT"))
	gpNamedFramebufferSampleLocationsfvNV = (C.GPNAMEDFRAMEBUFFERSAMPLELOCATIONSFVNV)(getProcAddr("glNamedFramebufferSampleLocationsfvNV"))
	gpObjectLabel = (C.GPOBJECTLABEL)(getProcAddr("glObjectLabel"))
	gpObjectLabelKHR = (C.GPOBJECTLABELKHR)(getProcAddr("glObjectLabelKHR"))
	gpObjectPtrLabel = (C.GPOBJECTPTRLABEL)(getProcAddr("glObjectPtrLabel"))
	gpObjectPtrLabelKHR = (C.GPOBJECTPTRLABELKHR)(getProcAddr("glObjectPtrLabelKHR"))
	gpPatchParameteriEXT = (C.GPPATCHPARAMETERIEXT)(getProcAddr("glPatchParameteriEXT"))
	gpPatchParameteriOES = (C.GPPATCHPARAMETERIOES)(getProcAddr("glPatchParameteriOES"))
	gpPathCommandsNV = (C.GPPATHCOMMANDSNV)(getProcAddr("glPathCommandsNV"))
	gpPathCoordsNV = (C.GPPATHCOORDSNV)(getProcAddr("glPathCoordsNV"))
	gpPathCoverDepthFuncNV = (C.GPPATHCOVERDEPTHFUNCNV)(getProcAddr("glPathCoverDepthFuncNV"))
	gpPathDashArrayNV = (C.GPPATHDASHARRAYNV)(getProcAddr("glPathDashArrayNV"))
	gpPathGlyphIndexArrayNV = (C.GPPATHGLYPHINDEXARRAYNV)(getProcAddr("glPathGlyphIndexArrayNV"))
	gpPathGlyphIndexRangeNV = (C.GPPATHGLYPHINDEXRANGENV)(getProcAddr("glPathGlyphIndexRangeNV"))
	gpPathGlyphRangeNV = (C.GPPATHGLYPHRANGENV)(getProcAddr("glPathGlyphRangeNV"))
	gpPathGlyphsNV = (C.GPPATHGLYPHSNV)(getProcAddr("glPathGlyphsNV"))
	gpPathMemoryGlyphIndexArrayNV = (C.GPPATHMEMORYGLYPHINDEXARRAYNV)(getProcAddr("glPathMemoryGlyphIndexArrayNV"))
	gpPathParameterfNV = (C.GPPATHPARAMETERFNV)(getProcAddr("glPathParameterfNV"))
	gpPathParameterfvNV = (C.GPPATHPARAMETERFVNV)(getProcAddr("glPathParameterfvNV"))
	gpPathParameteriNV = (C.GPPATHPARAMETERINV)(getProcAddr("glPathParameteriNV"))
	gpPathParameterivNV = (C.GPPATHPARAMETERIVNV)(getProcAddr("glPathParameterivNV"))
	gpPathStencilDepthOffsetNV = (C.GPPATHSTENCILDEPTHOFFSETNV)(getProcAddr("glPathStencilDepthOffsetNV"))
	gpPathStencilFuncNV = (C.GPPATHSTENCILFUNCNV)(getProcAddr("glPathStencilFuncNV"))
	gpPathStringNV = (C.GPPATHSTRINGNV)(getProcAddr("glPathStringNV"))
	gpPathSubCommandsNV = (C.GPPATHSUBCOMMANDSNV)(getProcAddr("glPathSubCommandsNV"))
	gpPathSubCoordsNV = (C.GPPATHSUBCOORDSNV)(getProcAddr("glPathSubCoordsNV"))
	gpPauseTransformFeedback = (C.GPPAUSETRANSFORMFEEDBACK)(getProcAddr("glPauseTransformFeedback"))
	if gpPauseTransformFeedback == nil {
		return errors.New("glPauseTransformFeedback")
	}
	gpPixelStorei = (C.GPPIXELSTOREI)(getProcAddr("glPixelStorei"))
	if gpPixelStorei == nil {
		return errors.New("glPixelStorei")
	}
	gpPointAlongPathNV = (C.GPPOINTALONGPATHNV)(getProcAddr("glPointAlongPathNV"))
	gpPolygonModeNV = (C.GPPOLYGONMODENV)(getProcAddr("glPolygonModeNV"))
	gpPolygonOffset = (C.GPPOLYGONOFFSET)(getProcAddr("glPolygonOffset"))
	if gpPolygonOffset == nil {
		return errors.New("glPolygonOffset")
	}
	gpPolygonOffsetClampEXT = (C.GPPOLYGONOFFSETCLAMPEXT)(getProcAddr("glPolygonOffsetClampEXT"))
	gpPopDebugGroup = (C.GPPOPDEBUGGROUP)(getProcAddr("glPopDebugGroup"))
	gpPopDebugGroupKHR = (C.GPPOPDEBUGGROUPKHR)(getProcAddr("glPopDebugGroupKHR"))
	gpPopGroupMarkerEXT = (C.GPPOPGROUPMARKEREXT)(getProcAddr("glPopGroupMarkerEXT"))
	gpPrimitiveBoundingBoxEXT = (C.GPPRIMITIVEBOUNDINGBOXEXT)(getProcAddr("glPrimitiveBoundingBoxEXT"))
	gpPrimitiveBoundingBoxOES = (C.GPPRIMITIVEBOUNDINGBOXOES)(getProcAddr("glPrimitiveBoundingBoxOES"))
	gpProgramBinary = (C.GPPROGRAMBINARY)(getProcAddr("glProgramBinary"))
	if gpProgramBinary == nil {
		return errors.New("glProgramBinary")
	}
	gpProgramBinaryOES = (C.GPPROGRAMBINARYOES)(getProcAddr("glProgramBinaryOES"))
	gpProgramParameteri = (C.GPPROGRAMPARAMETERI)(getProcAddr("glProgramParameteri"))
	if gpProgramParameteri == nil {
		return errors.New("glProgramParameteri")
	}
	gpProgramParameteriEXT = (C.GPPROGRAMPARAMETERIEXT)(getProcAddr("glProgramParameteriEXT"))
	gpProgramPathFragmentInputGenNV = (C.GPPROGRAMPATHFRAGMENTINPUTGENNV)(getProcAddr("glProgramPathFragmentInputGenNV"))
	gpProgramUniform1f = (C.GPPROGRAMUNIFORM1F)(getProcAddr("glProgramUniform1f"))
	if gpProgramUniform1f == nil {
		return errors.New("glProgramUniform1f")
	}
	gpProgramUniform1fEXT = (C.GPPROGRAMUNIFORM1FEXT)(getProcAddr("glProgramUniform1fEXT"))
	gpProgramUniform1fv = (C.GPPROGRAMUNIFORM1FV)(getProcAddr("glProgramUniform1fv"))
	if gpProgramUniform1fv == nil {
		return errors.New("glProgramUniform1fv")
	}
	gpProgramUniform1fvEXT = (C.GPPROGRAMUNIFORM1FVEXT)(getProcAddr("glProgramUniform1fvEXT"))
	gpProgramUniform1i = (C.GPPROGRAMUNIFORM1I)(getProcAddr("glProgramUniform1i"))
	if gpProgramUniform1i == nil {
		return errors.New("glProgramUniform1i")
	}
	gpProgramUniform1i64NV = (C.GPPROGRAMUNIFORM1I64NV)(getProcAddr("glProgramUniform1i64NV"))
	gpProgramUniform1i64vNV = (C.GPPROGRAMUNIFORM1I64VNV)(getProcAddr("glProgramUniform1i64vNV"))
	gpProgramUniform1iEXT = (C.GPPROGRAMUNIFORM1IEXT)(getProcAddr("glProgramUniform1iEXT"))
	gpProgramUniform1iv = (C.GPPROGRAMUNIFORM1IV)(getProcAddr("glProgramUniform1iv"))
	if gpProgramUniform1iv == nil {
		return errors.New("glProgramUniform1iv")
	}
	gpProgramUniform1ivEXT = (C.GPPROGRAMUNIFORM1IVEXT)(getProcAddr("glProgramUniform1ivEXT"))
	gpProgramUniform1ui = (C.GPPROGRAMUNIFORM1UI)(getProcAddr("glProgramUniform1ui"))
	if gpProgramUniform1ui == nil {
		return errors.New("glProgramUniform1ui")
	}
	gpProgramUniform1ui64NV = (C.GPPROGRAMUNIFORM1UI64NV)(getProcAddr("glProgramUniform1ui64NV"))
	gpProgramUniform1ui64vNV = (C.GPPROGRAMUNIFORM1UI64VNV)(getProcAddr("glProgramUniform1ui64vNV"))
	gpProgramUniform1uiEXT = (C.GPPROGRAMUNIFORM1UIEXT)(getProcAddr("glProgramUniform1uiEXT"))
	gpProgramUniform1uiv = (C.GPPROGRAMUNIFORM1UIV)(getProcAddr("glProgramUniform1uiv"))
	if gpProgramUniform1uiv == nil {
		return errors.New("glProgramUniform1uiv")
	}
	gpProgramUniform1uivEXT = (C.GPPROGRAMUNIFORM1UIVEXT)(getProcAddr("glProgramUniform1uivEXT"))
	gpProgramUniform2f = (C.GPPROGRAMUNIFORM2F)(getProcAddr("glProgramUniform2f"))
	if gpProgramUniform2f == nil {
		return errors.New("glProgramUniform2f")
	}
	gpProgramUniform2fEXT = (C.GPPROGRAMUNIFORM2FEXT)(getProcAddr("glProgramUniform2fEXT"))
	gpProgramUniform2fv = (C.GPPROGRAMUNIFORM2FV)(getProcAddr("glProgramUniform2fv"))
	if gpProgramUniform2fv == nil {
		return errors.New("glProgramUniform2fv")
	}
	gpProgramUniform2fvEXT = (C.GPPROGRAMUNIFORM2FVEXT)(getProcAddr("glProgramUniform2fvEXT"))
	gpProgramUniform2i = (C.GPPROGRAMUNIFORM2I)(getProcAddr("glProgramUniform2i"))
	if gpProgramUniform2i == nil {
		return errors.New("glProgramUniform2i")
	}
	gpProgramUniform2i64NV = (C.GPPROGRAMUNIFORM2I64NV)(getProcAddr("glProgramUniform2i64NV"))
	gpProgramUniform2i64vNV = (C.GPPROGRAMUNIFORM2I64VNV)(getProcAddr("glProgramUniform2i64vNV"))
	gpProgramUniform2iEXT = (C.GPPROGRAMUNIFORM2IEXT)(getProcAddr("glProgramUniform2iEXT"))
	gpProgramUniform2iv = (C.GPPROGRAMUNIFORM2IV)(getProcAddr("glProgramUniform2iv"))
	if gpProgramUniform2iv == nil {
		return errors.New("glProgramUniform2iv")
	}
	gpProgramUniform2ivEXT = (C.GPPROGRAMUNIFORM2IVEXT)(getProcAddr("glProgramUniform2ivEXT"))
	gpProgramUniform2ui = (C.GPPROGRAMUNIFORM2UI)(getProcAddr("glProgramUniform2ui"))
	if gpProgramUniform2ui == nil {
		return errors.New("glProgramUniform2ui")
	}
	gpProgramUniform2ui64NV = (C.GPPROGRAMUNIFORM2UI64NV)(getProcAddr("glProgramUniform2ui64NV"))
	gpProgramUniform2ui64vNV = (C.GPPROGRAMUNIFORM2UI64VNV)(getProcAddr("glProgramUniform2ui64vNV"))
	gpProgramUniform2uiEXT = (C.GPPROGRAMUNIFORM2UIEXT)(getProcAddr("glProgramUniform2uiEXT"))
	gpProgramUniform2uiv = (C.GPPROGRAMUNIFORM2UIV)(getProcAddr("glProgramUniform2uiv"))
	if gpProgramUniform2uiv == nil {
		return errors.New("glProgramUniform2uiv")
	}
	gpProgramUniform2uivEXT = (C.GPPROGRAMUNIFORM2UIVEXT)(getProcAddr("glProgramUniform2uivEXT"))
	gpProgramUniform3f = (C.GPPROGRAMUNIFORM3F)(getProcAddr("glProgramUniform3f"))
	if gpProgramUniform3f == nil {
		return errors.New("glProgramUniform3f")
	}
	gpProgramUniform3fEXT = (C.GPPROGRAMUNIFORM3FEXT)(getProcAddr("glProgramUniform3fEXT"))
	gpProgramUniform3fv = (C.GPPROGRAMUNIFORM3FV)(getProcAddr("glProgramUniform3fv"))
	if gpProgramUniform3fv == nil {
		return errors.New("glProgramUniform3fv")
	}
	gpProgramUniform3fvEXT = (C.GPPROGRAMUNIFORM3FVEXT)(getProcAddr("glProgramUniform3fvEXT"))
	gpProgramUniform3i = (C.GPPROGRAMUNIFORM3I)(getProcAddr("glProgramUniform3i"))
	if gpProgramUniform3i == nil {
		return errors.New("glProgramUniform3i")
	}
	gpProgramUniform3i64NV = (C.GPPROGRAMUNIFORM3I64NV)(getProcAddr("glProgramUniform3i64NV"))
	gpProgramUniform3i64vNV = (C.GPPROGRAMUNIFORM3I64VNV)(getProcAddr("glProgramUniform3i64vNV"))
	gpProgramUniform3iEXT = (C.GPPROGRAMUNIFORM3IEXT)(getProcAddr("glProgramUniform3iEXT"))
	gpProgramUniform3iv = (C.GPPROGRAMUNIFORM3IV)(getProcAddr("glProgramUniform3iv"))
	if gpProgramUniform3iv == nil {
		return errors.New("glProgramUniform3iv")
	}
	gpProgramUniform3ivEXT = (C.GPPROGRAMUNIFORM3IVEXT)(getProcAddr("glProgramUniform3ivEXT"))
	gpProgramUniform3ui = (C.GPPROGRAMUNIFORM3UI)(getProcAddr("glProgramUniform3ui"))
	if gpProgramUniform3ui == nil {
		return errors.New("glProgramUniform3ui")
	}
	gpProgramUniform3ui64NV = (C.GPPROGRAMUNIFORM3UI64NV)(getProcAddr("glProgramUniform3ui64NV"))
	gpProgramUniform3ui64vNV = (C.GPPROGRAMUNIFORM3UI64VNV)(getProcAddr("glProgramUniform3ui64vNV"))
	gpProgramUniform3uiEXT = (C.GPPROGRAMUNIFORM3UIEXT)(getProcAddr("glProgramUniform3uiEXT"))
	gpProgramUniform3uiv = (C.GPPROGRAMUNIFORM3UIV)(getProcAddr("glProgramUniform3uiv"))
	if gpProgramUniform3uiv == nil {
		return errors.New("glProgramUniform3uiv")
	}
	gpProgramUniform3uivEXT = (C.GPPROGRAMUNIFORM3UIVEXT)(getProcAddr("glProgramUniform3uivEXT"))
	gpProgramUniform4f = (C.GPPROGRAMUNIFORM4F)(getProcAddr("glProgramUniform4f"))
	if gpProgramUniform4f == nil {
		return errors.New("glProgramUniform4f")
	}
	gpProgramUniform4fEXT = (C.GPPROGRAMUNIFORM4FEXT)(getProcAddr("glProgramUniform4fEXT"))
	gpProgramUniform4fv = (C.GPPROGRAMUNIFORM4FV)(getProcAddr("glProgramUniform4fv"))
	if gpProgramUniform4fv == nil {
		return errors.New("glProgramUniform4fv")
	}
	gpProgramUniform4fvEXT = (C.GPPROGRAMUNIFORM4FVEXT)(getProcAddr("glProgramUniform4fvEXT"))
	gpProgramUniform4i = (C.GPPROGRAMUNIFORM4I)(getProcAddr("glProgramUniform4i"))
	if gpProgramUniform4i == nil {
		return errors.New("glProgramUniform4i")
	}
	gpProgramUniform4i64NV = (C.GPPROGRAMUNIFORM4I64NV)(getProcAddr("glProgramUniform4i64NV"))
	gpProgramUniform4i64vNV = (C.GPPROGRAMUNIFORM4I64VNV)(getProcAddr("glProgramUniform4i64vNV"))
	gpProgramUniform4iEXT = (C.GPPROGRAMUNIFORM4IEXT)(getProcAddr("glProgramUniform4iEXT"))
	gpProgramUniform4iv = (C.GPPROGRAMUNIFORM4IV)(getProcAddr("glProgramUniform4iv"))
	if gpProgramUniform4iv == nil {
		return errors.New("glProgramUniform4iv")
	}
	gpProgramUniform4ivEXT = (C.GPPROGRAMUNIFORM4IVEXT)(getProcAddr("glProgramUniform4ivEXT"))
	gpProgramUniform4ui = (C.GPPROGRAMUNIFORM4UI)(getProcAddr("glProgramUniform4ui"))
	if gpProgramUniform4ui == nil {
		return errors.New("glProgramUniform4ui")
	}
	gpProgramUniform4ui64NV = (C.GPPROGRAMUNIFORM4UI64NV)(getProcAddr("glProgramUniform4ui64NV"))
	gpProgramUniform4ui64vNV = (C.GPPROGRAMUNIFORM4UI64VNV)(getProcAddr("glProgramUniform4ui64vNV"))
	gpProgramUniform4uiEXT = (C.GPPROGRAMUNIFORM4UIEXT)(getProcAddr("glProgramUniform4uiEXT"))
	gpProgramUniform4uiv = (C.GPPROGRAMUNIFORM4UIV)(getProcAddr("glProgramUniform4uiv"))
	if gpProgramUniform4uiv == nil {
		return errors.New("glProgramUniform4uiv")
	}
	gpProgramUniform4uivEXT = (C.GPPROGRAMUNIFORM4UIVEXT)(getProcAddr("glProgramUniform4uivEXT"))
	gpProgramUniformHandleui64IMG = (C.GPPROGRAMUNIFORMHANDLEUI64IMG)(getProcAddr("glProgramUniformHandleui64IMG"))
	gpProgramUniformHandleui64NV = (C.GPPROGRAMUNIFORMHANDLEUI64NV)(getProcAddr("glProgramUniformHandleui64NV"))
	gpProgramUniformHandleui64vIMG = (C.GPPROGRAMUNIFORMHANDLEUI64VIMG)(getProcAddr("glProgramUniformHandleui64vIMG"))
	gpProgramUniformHandleui64vNV = (C.GPPROGRAMUNIFORMHANDLEUI64VNV)(getProcAddr("glProgramUniformHandleui64vNV"))
	gpProgramUniformMatrix2fv = (C.GPPROGRAMUNIFORMMATRIX2FV)(getProcAddr("glProgramUniformMatrix2fv"))
	if gpProgramUniformMatrix2fv == nil {
		return errors.New("glProgramUniformMatrix2fv")
	}
	gpProgramUniformMatrix2fvEXT = (C.GPPROGRAMUNIFORMMATRIX2FVEXT)(getProcAddr("glProgramUniformMatrix2fvEXT"))
	gpProgramUniformMatrix2x3fv = (C.GPPROGRAMUNIFORMMATRIX2X3FV)(getProcAddr("glProgramUniformMatrix2x3fv"))
	if gpProgramUniformMatrix2x3fv == nil {
		return errors.New("glProgramUniformMatrix2x3fv")
	}
	gpProgramUniformMatrix2x3fvEXT = (C.GPPROGRAMUNIFORMMATRIX2X3FVEXT)(getProcAddr("glProgramUniformMatrix2x3fvEXT"))
	gpProgramUniformMatrix2x4fv = (C.GPPROGRAMUNIFORMMATRIX2X4FV)(getProcAddr("glProgramUniformMatrix2x4fv"))
	if gpProgramUniformMatrix2x4fv == nil {
		return errors.New("glProgramUniformMatrix2x4fv")
	}
	gpProgramUniformMatrix2x4fvEXT = (C.GPPROGRAMUNIFORMMATRIX2X4FVEXT)(getProcAddr("glProgramUniformMatrix2x4fvEXT"))
	gpProgramUniformMatrix3fv = (C.GPPROGRAMUNIFORMMATRIX3FV)(getProcAddr("glProgramUniformMatrix3fv"))
	if gpProgramUniformMatrix3fv == nil {
		return errors.New("glProgramUniformMatrix3fv")
	}
	gpProgramUniformMatrix3fvEXT = (C.GPPROGRAMUNIFORMMATRIX3FVEXT)(getProcAddr("glProgramUniformMatrix3fvEXT"))
	gpProgramUniformMatrix3x2fv = (C.GPPROGRAMUNIFORMMATRIX3X2FV)(getProcAddr("glProgramUniformMatrix3x2fv"))
	if gpProgramUniformMatrix3x2fv == nil {
		return errors.New("glProgramUniformMatrix3x2fv")
	}
	gpProgramUniformMatrix3x2fvEXT = (C.GPPROGRAMUNIFORMMATRIX3X2FVEXT)(getProcAddr("glProgramUniformMatrix3x2fvEXT"))
	gpProgramUniformMatrix3x4fv = (C.GPPROGRAMUNIFORMMATRIX3X4FV)(getProcAddr("glProgramUniformMatrix3x4fv"))
	if gpProgramUniformMatrix3x4fv == nil {
		return errors.New("glProgramUniformMatrix3x4fv")
	}
	gpProgramUniformMatrix3x4fvEXT = (C.GPPROGRAMUNIFORMMATRIX3X4FVEXT)(getProcAddr("glProgramUniformMatrix3x4fvEXT"))
	gpProgramUniformMatrix4fv = (C.GPPROGRAMUNIFORMMATRIX4FV)(getProcAddr("glProgramUniformMatrix4fv"))
	if gpProgramUniformMatrix4fv == nil {
		return errors.New("glProgramUniformMatrix4fv")
	}
	gpProgramUniformMatrix4fvEXT = (C.GPPROGRAMUNIFORMMATRIX4FVEXT)(getProcAddr("glProgramUniformMatrix4fvEXT"))
	gpProgramUniformMatrix4x2fv = (C.GPPROGRAMUNIFORMMATRIX4X2FV)(getProcAddr("glProgramUniformMatrix4x2fv"))
	if gpProgramUniformMatrix4x2fv == nil {
		return errors.New("glProgramUniformMatrix4x2fv")
	}
	gpProgramUniformMatrix4x2fvEXT = (C.GPPROGRAMUNIFORMMATRIX4X2FVEXT)(getProcAddr("glProgramUniformMatrix4x2fvEXT"))
	gpProgramUniformMatrix4x3fv = (C.GPPROGRAMUNIFORMMATRIX4X3FV)(getProcAddr("glProgramUniformMatrix4x3fv"))
	if gpProgramUniformMatrix4x3fv == nil {
		return errors.New("glProgramUniformMatrix4x3fv")
	}
	gpProgramUniformMatrix4x3fvEXT = (C.GPPROGRAMUNIFORMMATRIX4X3FVEXT)(getProcAddr("glProgramUniformMatrix4x3fvEXT"))
	gpPushDebugGroup = (C.GPPUSHDEBUGGROUP)(getProcAddr("glPushDebugGroup"))
	gpPushDebugGroupKHR = (C.GPPUSHDEBUGGROUPKHR)(getProcAddr("glPushDebugGroupKHR"))
	gpPushGroupMarkerEXT = (C.GPPUSHGROUPMARKEREXT)(getProcAddr("glPushGroupMarkerEXT"))
	gpQueryCounterEXT = (C.GPQUERYCOUNTEREXT)(getProcAddr("glQueryCounterEXT"))
	gpRasterSamplesEXT = (C.GPRASTERSAMPLESEXT)(getProcAddr("glRasterSamplesEXT"))
	gpReadBuffer = (C.GPREADBUFFER)(getProcAddr("glReadBuffer"))
	if gpReadBuffer == nil {
		return errors.New("glReadBuffer")
	}
	gpReadBufferIndexedEXT = (C.GPREADBUFFERINDEXEDEXT)(getProcAddr("glReadBufferIndexedEXT"))
	gpReadBufferNV = (C.GPREADBUFFERNV)(getProcAddr("glReadBufferNV"))
	gpReadPixels = (C.GPREADPIXELS)(getProcAddr("glReadPixels"))
	if gpReadPixels == nil {
		return errors.New("glReadPixels")
	}
	gpReadnPixels = (C.GPREADNPIXELS)(getProcAddr("glReadnPixels"))
	gpReadnPixelsEXT = (C.GPREADNPIXELSEXT)(getProcAddr("glReadnPixelsEXT"))
	gpReadnPixelsKHR = (C.GPREADNPIXELSKHR)(getProcAddr("glReadnPixelsKHR"))
	gpReleaseKeyedMutexWin32EXT = (C.GPRELEASEKEYEDMUTEXWIN32EXT)(getProcAddr("glReleaseKeyedMutexWin32EXT"))
	gpReleaseShaderCompiler = (C.GPRELEASESHADERCOMPILER)(getProcAddr("glReleaseShaderCompiler"))
	if gpReleaseShaderCompiler == nil {
		return errors.New("glReleaseShaderCompiler")
	}
	gpRenderbufferStorage = (C.GPRENDERBUFFERSTORAGE)(getProcAddr("glRenderbufferStorage"))
	if gpRenderbufferStorage == nil {
		return errors.New("glRenderbufferStorage")
	}
	gpRenderbufferStorageMultisample = (C.GPRENDERBUFFERSTORAGEMULTISAMPLE)(getProcAddr("glRenderbufferStorageMultisample"))
	if gpRenderbufferStorageMultisample == nil {
		return errors.New("glRenderbufferStorageMultisample")
	}
	gpRenderbufferStorageMultisampleANGLE = (C.GPRENDERBUFFERSTORAGEMULTISAMPLEANGLE)(getProcAddr("glRenderbufferStorageMultisampleANGLE"))
	gpRenderbufferStorageMultisampleAPPLE = (C.GPRENDERBUFFERSTORAGEMULTISAMPLEAPPLE)(getProcAddr("glRenderbufferStorageMultisampleAPPLE"))
	gpRenderbufferStorageMultisampleEXT = (C.GPRENDERBUFFERSTORAGEMULTISAMPLEEXT)(getProcAddr("glRenderbufferStorageMultisampleEXT"))
	gpRenderbufferStorageMultisampleIMG = (C.GPRENDERBUFFERSTORAGEMULTISAMPLEIMG)(getProcAddr("glRenderbufferStorageMultisampleIMG"))
	gpRenderbufferStorageMultisampleNV = (C.GPRENDERBUFFERSTORAGEMULTISAMPLENV)(getProcAddr("glRenderbufferStorageMultisampleNV"))
	gpResolveDepthValuesNV = (C.GPRESOLVEDEPTHVALUESNV)(getProcAddr("glResolveDepthValuesNV"))
	gpResolveMultisampleFramebufferAPPLE = (C.GPRESOLVEMULTISAMPLEFRAMEBUFFERAPPLE)(getProcAddr("glResolveMultisampleFramebufferAPPLE"))
	gpResumeTransformFeedback = (C.GPRESUMETRANSFORMFEEDBACK)(getProcAddr("glResumeTransformFeedback"))
	if gpResumeTransformFeedback == nil {
		return errors.New("glResumeTransformFeedback")
	}
	gpSampleCoverage = (C.GPSAMPLECOVERAGE)(getProcAddr("glSampleCoverage"))
	if gpSampleCoverage == nil {
		return errors.New("glSampleCoverage")
	}
	gpSampleMaski = (C.GPSAMPLEMASKI)(getProcAddr("glSampleMaski"))
	if gpSampleMaski == nil {
		return errors.New("glSampleMaski")
	}
	gpSamplerParameterIivEXT = (C.GPSAMPLERPARAMETERIIVEXT)(getProcAddr("glSamplerParameterIivEXT"))
	gpSamplerParameterIivOES = (C.GPSAMPLERPARAMETERIIVOES)(getProcAddr("glSamplerParameterIivOES"))
	gpSamplerParameterIuivEXT = (C.GPSAMPLERPARAMETERIUIVEXT)(getProcAddr("glSamplerParameterIuivEXT"))
	gpSamplerParameterIuivOES = (C.GPSAMPLERPARAMETERIUIVOES)(getProcAddr("glSamplerParameterIuivOES"))
	gpSamplerParameterf = (C.GPSAMPLERPARAMETERF)(getProcAddr("glSamplerParameterf"))
	if gpSamplerParameterf == nil {
		return errors.New("glSamplerParameterf")
	}
	gpSamplerParameterfv = (C.GPSAMPLERPARAMETERFV)(getProcAddr("glSamplerParameterfv"))
	if gpSamplerParameterfv == nil {
		return errors.New("glSamplerParameterfv")
	}
	gpSamplerParameteri = (C.GPSAMPLERPARAMETERI)(getProcAddr("glSamplerParameteri"))
	if gpSamplerParameteri == nil {
		return errors.New("glSamplerParameteri")
	}
	gpSamplerParameteriv = (C.GPSAMPLERPARAMETERIV)(getProcAddr("glSamplerParameteriv"))
	if gpSamplerParameteriv == nil {
		return errors.New("glSamplerParameteriv")
	}
	gpScissor = (C.GPSCISSOR)(getProcAddr("glScissor"))
	if gpScissor == nil {
		return errors.New("glScissor")
	}
	gpScissorArrayvNV = (C.GPSCISSORARRAYVNV)(getProcAddr("glScissorArrayvNV"))
	gpScissorArrayvOES = (C.GPSCISSORARRAYVOES)(getProcAddr("glScissorArrayvOES"))
	gpScissorIndexedNV = (C.GPSCISSORINDEXEDNV)(getProcAddr("glScissorIndexedNV"))
	gpScissorIndexedOES = (C.GPSCISSORINDEXEDOES)(getProcAddr("glScissorIndexedOES"))
	gpScissorIndexedvNV = (C.GPSCISSORINDEXEDVNV)(getProcAddr("glScissorIndexedvNV"))
	gpScissorIndexedvOES = (C.GPSCISSORINDEXEDVOES)(getProcAddr("glScissorIndexedvOES"))
	gpSelectPerfMonitorCountersAMD = (C.GPSELECTPERFMONITORCOUNTERSAMD)(getProcAddr("glSelectPerfMonitorCountersAMD"))
	gpSemaphoreParameterui64vEXT = (C.GPSEMAPHOREPARAMETERUI64VEXT)(getProcAddr("glSemaphoreParameterui64vEXT"))
	gpSetFenceNV = (C.GPSETFENCENV)(getProcAddr("glSetFenceNV"))
	gpShaderBinary = (C.GPSHADERBINARY)(getProcAddr("glShaderBinary"))
	if gpShaderBinary == nil {
		return errors.New("glShaderBinary")
	}
	gpShaderSource = (C.GPSHADERSOURCE)(getProcAddr("glShaderSource"))
	if gpShaderSource == nil {
		return errors.New("glShaderSource")
	}
	gpSignalSemaphoreEXT = (C.GPSIGNALSEMAPHOREEXT)(getProcAddr("glSignalSemaphoreEXT"))
	gpSignalVkFenceNV = (C.GPSIGNALVKFENCENV)(getProcAddr("glSignalVkFenceNV"))
	gpSignalVkSemaphoreNV = (C.GPSIGNALVKSEMAPHORENV)(getProcAddr("glSignalVkSemaphoreNV"))
	gpStartTilingQCOM = (C.GPSTARTTILINGQCOM)(getProcAddr("glStartTilingQCOM"))
	gpStencilFillPathInstancedNV = (C.GPSTENCILFILLPATHINSTANCEDNV)(getProcAddr("glStencilFillPathInstancedNV"))
	gpStencilFillPathNV = (C.GPSTENCILFILLPATHNV)(getProcAddr("glStencilFillPathNV"))
	gpStencilFunc = (C.GPSTENCILFUNC)(getProcAddr("glStencilFunc"))
	if gpStencilFunc == nil {
		return errors.New("glStencilFunc")
	}
	gpStencilFuncSeparate = (C.GPSTENCILFUNCSEPARATE)(getProcAddr("glStencilFuncSeparate"))
	if gpStencilFuncSeparate == nil {
		return errors.New("glStencilFuncSeparate")
	}
	gpStencilMask = (C.GPSTENCILMASK)(getProcAddr("glStencilMask"))
	if gpStencilMask == nil {
		return errors.New("glStencilMask")
	}
	gpStencilMaskSeparate = (C.GPSTENCILMASKSEPARATE)(getProcAddr("glStencilMaskSeparate"))
	if gpStencilMaskSeparate == nil {
		return errors.New("glStencilMaskSeparate")
	}
	gpStencilOp = (C.GPSTENCILOP)(getProcAddr("glStencilOp"))
	if gpStencilOp == nil {
		return errors.New("glStencilOp")
	}
	gpStencilOpSeparate = (C.GPSTENCILOPSEPARATE)(getProcAddr("glStencilOpSeparate"))
	if gpStencilOpSeparate == nil {
		return errors.New("glStencilOpSeparate")
	}
	gpStencilStrokePathInstancedNV = (C.GPSTENCILSTROKEPATHINSTANCEDNV)(getProcAddr("glStencilStrokePathInstancedNV"))
	gpStencilStrokePathNV = (C.GPSTENCILSTROKEPATHNV)(getProcAddr("glStencilStrokePathNV"))
	gpStencilThenCoverFillPathInstancedNV = (C.GPSTENCILTHENCOVERFILLPATHINSTANCEDNV)(getProcAddr("glStencilThenCoverFillPathInstancedNV"))
	gpStencilThenCoverFillPathNV = (C.GPSTENCILTHENCOVERFILLPATHNV)(getProcAddr("glStencilThenCoverFillPathNV"))
	gpStencilThenCoverStrokePathInstancedNV = (C.GPSTENCILTHENCOVERSTROKEPATHINSTANCEDNV)(getProcAddr("glStencilThenCoverStrokePathInstancedNV"))
	gpStencilThenCoverStrokePathNV = (C.GPSTENCILTHENCOVERSTROKEPATHNV)(getProcAddr("glStencilThenCoverStrokePathNV"))
	gpSubpixelPrecisionBiasNV = (C.GPSUBPIXELPRECISIONBIASNV)(getProcAddr("glSubpixelPrecisionBiasNV"))
	gpTestFenceNV = (C.GPTESTFENCENV)(getProcAddr("glTestFenceNV"))
	gpTexBufferEXT = (C.GPTEXBUFFEREXT)(getProcAddr("glTexBufferEXT"))
	gpTexBufferOES = (C.GPTEXBUFFEROES)(getProcAddr("glTexBufferOES"))
	gpTexBufferRangeEXT = (C.GPTEXBUFFERRANGEEXT)(getProcAddr("glTexBufferRangeEXT"))
	gpTexBufferRangeOES = (C.GPTEXBUFFERRANGEOES)(getProcAddr("glTexBufferRangeOES"))
	gpTexImage2D = (C.GPTEXIMAGE2D)(getProcAddr("glTexImage2D"))
	if gpTexImage2D == nil {
		return errors.New("glTexImage2D")
	}
	gpTexImage3D = (C.GPTEXIMAGE3D)(getProcAddr("glTexImage3D"))
	if gpTexImage3D == nil {
		return errors.New("glTexImage3D")
	}
	gpTexImage3DOES = (C.GPTEXIMAGE3DOES)(getProcAddr("glTexImage3DOES"))
	gpTexPageCommitmentEXT = (C.GPTEXPAGECOMMITMENTEXT)(getProcAddr("glTexPageCommitmentEXT"))
	gpTexParameterIivEXT = (C.GPTEXPARAMETERIIVEXT)(getProcAddr("glTexParameterIivEXT"))
	gpTexParameterIivOES = (C.GPTEXPARAMETERIIVOES)(getProcAddr("glTexParameterIivOES"))
	gpTexParameterIuivEXT = (C.GPTEXPARAMETERIUIVEXT)(getProcAddr("glTexParameterIuivEXT"))
	gpTexParameterIuivOES = (C.GPTEXPARAMETERIUIVOES)(getProcAddr("glTexParameterIuivOES"))
	gpTexParameterf = (C.GPTEXPARAMETERF)(getProcAddr("glTexParameterf"))
	if gpTexParameterf == nil {
		return errors.New("glTexParameterf")
	}
	gpTexParameterfv = (C.GPTEXPARAMETERFV)(getProcAddr("glTexParameterfv"))
	if gpTexParameterfv == nil {
		return errors.New("glTexParameterfv")
	}
	gpTexParameteri = (C.GPTEXPARAMETERI)(getProcAddr("glTexParameteri"))
	if gpTexParameteri == nil {
		return errors.New("glTexParameteri")
	}
	gpTexParameteriv = (C.GPTEXPARAMETERIV)(getProcAddr("glTexParameteriv"))
	if gpTexParameteriv == nil {
		return errors.New("glTexParameteriv")
	}
	gpTexStorage1DEXT = (C.GPTEXSTORAGE1DEXT)(getProcAddr("glTexStorage1DEXT"))
	gpTexStorage2D = (C.GPTEXSTORAGE2D)(getProcAddr("glTexStorage2D"))
	if gpTexStorage2D == nil {
		return errors.New("glTexStorage2D")
	}
	gpTexStorage2DEXT = (C.GPTEXSTORAGE2DEXT)(getProcAddr("glTexStorage2DEXT"))
	gpTexStorage2DMultisample = (C.GPTEXSTORAGE2DMULTISAMPLE)(getProcAddr("glTexStorage2DMultisample"))
	if gpTexStorage2DMultisample == nil {
		return errors.New("glTexStorage2DMultisample")
	}
	gpTexStorage3D = (C.GPTEXSTORAGE3D)(getProcAddr("glTexStorage3D"))
	if gpTexStorage3D == nil {
		return errors.New("glTexStorage3D")
	}
	gpTexStorage3DEXT = (C.GPTEXSTORAGE3DEXT)(getProcAddr("glTexStorage3DEXT"))
	gpTexStorage3DMultisampleOES = (C.GPTEXSTORAGE3DMULTISAMPLEOES)(getProcAddr("glTexStorage3DMultisampleOES"))
	gpTexStorageMem1DEXT = (C.GPTEXSTORAGEMEM1DEXT)(getProcAddr("glTexStorageMem1DEXT"))
	gpTexStorageMem2DEXT = (C.GPTEXSTORAGEMEM2DEXT)(getProcAddr("glTexStorageMem2DEXT"))
	gpTexStorageMem2DMultisampleEXT = (C.GPTEXSTORAGEMEM2DMULTISAMPLEEXT)(getProcAddr("glTexStorageMem2DMultisampleEXT"))
	gpTexStorageMem3DEXT = (C.GPTEXSTORAGEMEM3DEXT)(getProcAddr("glTexStorageMem3DEXT"))
	gpTexStorageMem3DMultisampleEXT = (C.GPTEXSTORAGEMEM3DMULTISAMPLEEXT)(getProcAddr("glTexStorageMem3DMultisampleEXT"))
	gpTexSubImage2D = (C.GPTEXSUBIMAGE2D)(getProcAddr("glTexSubImage2D"))
	if gpTexSubImage2D == nil {
		return errors.New("glTexSubImage2D")
	}
	gpTexSubImage3D = (C.GPTEXSUBIMAGE3D)(getProcAddr("glTexSubImage3D"))
	if gpTexSubImage3D == nil {
		return errors.New("glTexSubImage3D")
	}
	gpTexSubImage3DOES = (C.GPTEXSUBIMAGE3DOES)(getProcAddr("glTexSubImage3DOES"))
	gpTextureFoveationParametersQCOM = (C.GPTEXTUREFOVEATIONPARAMETERSQCOM)(getProcAddr("glTextureFoveationParametersQCOM"))
	gpTextureStorage1DEXT = (C.GPTEXTURESTORAGE1DEXT)(getProcAddr("glTextureStorage1DEXT"))
	gpTextureStorage2DEXT = (C.GPTEXTURESTORAGE2DEXT)(getProcAddr("glTextureStorage2DEXT"))
	gpTextureStorage3DEXT = (C.GPTEXTURESTORAGE3DEXT)(getProcAddr("glTextureStorage3DEXT"))
	gpTextureStorageMem1DEXT = (C.GPTEXTURESTORAGEMEM1DEXT)(getProcAddr("glTextureStorageMem1DEXT"))
	gpTextureStorageMem2DEXT = (C.GPTEXTURESTORAGEMEM2DEXT)(getProcAddr("glTextureStorageMem2DEXT"))
	gpTextureStorageMem2DMultisampleEXT = (C.GPTEXTURESTORAGEMEM2DMULTISAMPLEEXT)(getProcAddr("glTextureStorageMem2DMultisampleEXT"))
	gpTextureStorageMem3DEXT = (C.GPTEXTURESTORAGEMEM3DEXT)(getProcAddr("glTextureStorageMem3DEXT"))
	gpTextureStorageMem3DMultisampleEXT = (C.GPTEXTURESTORAGEMEM3DMULTISAMPLEEXT)(getProcAddr("glTextureStorageMem3DMultisampleEXT"))
	gpTextureViewEXT = (C.GPTEXTUREVIEWEXT)(getProcAddr("glTextureViewEXT"))
	gpTextureViewOES = (C.GPTEXTUREVIEWOES)(getProcAddr("glTextureViewOES"))
	gpTransformFeedbackVaryings = (C.GPTRANSFORMFEEDBACKVARYINGS)(getProcAddr("glTransformFeedbackVaryings"))
	if gpTransformFeedbackVaryings == nil {
		return errors.New("glTransformFeedbackVaryings")
	}
	gpTransformPathNV = (C.GPTRANSFORMPATHNV)(getProcAddr("glTransformPathNV"))
	gpUniform1f = (C.GPUNIFORM1F)(getProcAddr("glUniform1f"))
	if gpUniform1f == nil {
		return errors.New("glUniform1f")
	}
	gpUniform1fv = (C.GPUNIFORM1FV)(getProcAddr("glUniform1fv"))
	if gpUniform1fv == nil {
		return errors.New("glUniform1fv")
	}
	gpUniform1i = (C.GPUNIFORM1I)(getProcAddr("glUniform1i"))
	if gpUniform1i == nil {
		return errors.New("glUniform1i")
	}
	gpUniform1i64NV = (C.GPUNIFORM1I64NV)(getProcAddr("glUniform1i64NV"))
	gpUniform1i64vNV = (C.GPUNIFORM1I64VNV)(getProcAddr("glUniform1i64vNV"))
	gpUniform1iv = (C.GPUNIFORM1IV)(getProcAddr("glUniform1iv"))
	if gpUniform1iv == nil {
		return errors.New("glUniform1iv")
	}
	gpUniform1ui = (C.GPUNIFORM1UI)(getProcAddr("glUniform1ui"))
	if gpUniform1ui == nil {
		return errors.New("glUniform1ui")
	}
	gpUniform1ui64NV = (C.GPUNIFORM1UI64NV)(getProcAddr("glUniform1ui64NV"))
	gpUniform1ui64vNV = (C.GPUNIFORM1UI64VNV)(getProcAddr("glUniform1ui64vNV"))
	gpUniform1uiv = (C.GPUNIFORM1UIV)(getProcAddr("glUniform1uiv"))
	if gpUniform1uiv == nil {
		return errors.New("glUniform1uiv")
	}
	gpUniform2f = (C.GPUNIFORM2F)(getProcAddr("glUniform2f"))
	if gpUniform2f == nil {
		return errors.New("glUniform2f")
	}
	gpUniform2fv = (C.GPUNIFORM2FV)(getProcAddr("glUniform2fv"))
	if gpUniform2fv == nil {
		return errors.New("glUniform2fv")
	}
	gpUniform2i = (C.GPUNIFORM2I)(getProcAddr("glUniform2i"))
	if gpUniform2i == nil {
		return errors.New("glUniform2i")
	}
	gpUniform2i64NV = (C.GPUNIFORM2I64NV)(getProcAddr("glUniform2i64NV"))
	gpUniform2i64vNV = (C.GPUNIFORM2I64VNV)(getProcAddr("glUniform2i64vNV"))
	gpUniform2iv = (C.GPUNIFORM2IV)(getProcAddr("glUniform2iv"))
	if gpUniform2iv == nil {
		return errors.New("glUniform2iv")
	}
	gpUniform2ui = (C.GPUNIFORM2UI)(getProcAddr("glUniform2ui"))
	if gpUniform2ui == nil {
		return errors.New("glUniform2ui")
	}
	gpUniform2ui64NV = (C.GPUNIFORM2UI64NV)(getProcAddr("glUniform2ui64NV"))
	gpUniform2ui64vNV = (C.GPUNIFORM2UI64VNV)(getProcAddr("glUniform2ui64vNV"))
	gpUniform2uiv = (C.GPUNIFORM2UIV)(getProcAddr("glUniform2uiv"))
	if gpUniform2uiv == nil {
		return errors.New("glUniform2uiv")
	}
	gpUniform3f = (C.GPUNIFORM3F)(getProcAddr("glUniform3f"))
	if gpUniform3f == nil {
		return errors.New("glUniform3f")
	}
	gpUniform3fv = (C.GPUNIFORM3FV)(getProcAddr("glUniform3fv"))
	if gpUniform3fv == nil {
		return errors.New("glUniform3fv")
	}
	gpUniform3i = (C.GPUNIFORM3I)(getProcAddr("glUniform3i"))
	if gpUniform3i == nil {
		return errors.New("glUniform3i")
	}
	gpUniform3i64NV = (C.GPUNIFORM3I64NV)(getProcAddr("glUniform3i64NV"))
	gpUniform3i64vNV = (C.GPUNIFORM3I64VNV)(getProcAddr("glUniform3i64vNV"))
	gpUniform3iv = (C.GPUNIFORM3IV)(getProcAddr("glUniform3iv"))
	if gpUniform3iv == nil {
		return errors.New("glUniform3iv")
	}
	gpUniform3ui = (C.GPUNIFORM3UI)(getProcAddr("glUniform3ui"))
	if gpUniform3ui == nil {
		return errors.New("glUniform3ui")
	}
	gpUniform3ui64NV = (C.GPUNIFORM3UI64NV)(getProcAddr("glUniform3ui64NV"))
	gpUniform3ui64vNV = (C.GPUNIFORM3UI64VNV)(getProcAddr("glUniform3ui64vNV"))
	gpUniform3uiv = (C.GPUNIFORM3UIV)(getProcAddr("glUniform3uiv"))
	if gpUniform3uiv == nil {
		return errors.New("glUniform3uiv")
	}
	gpUniform4f = (C.GPUNIFORM4F)(getProcAddr("glUniform4f"))
	if gpUniform4f == nil {
		return errors.New("glUniform4f")
	}
	gpUniform4fv = (C.GPUNIFORM4FV)(getProcAddr("glUniform4fv"))
	if gpUniform4fv == nil {
		return errors.New("glUniform4fv")
	}
	gpUniform4i = (C.GPUNIFORM4I)(getProcAddr("glUniform4i"))
	if gpUniform4i == nil {
		return errors.New("glUniform4i")
	}
	gpUniform4i64NV = (C.GPUNIFORM4I64NV)(getProcAddr("glUniform4i64NV"))
	gpUniform4i64vNV = (C.GPUNIFORM4I64VNV)(getProcAddr("glUniform4i64vNV"))
	gpUniform4iv = (C.GPUNIFORM4IV)(getProcAddr("glUniform4iv"))
	if gpUniform4iv == nil {
		return errors.New("glUniform4iv")
	}
	gpUniform4ui = (C.GPUNIFORM4UI)(getProcAddr("glUniform4ui"))
	if gpUniform4ui == nil {
		return errors.New("glUniform4ui")
	}
	gpUniform4ui64NV = (C.GPUNIFORM4UI64NV)(getProcAddr("glUniform4ui64NV"))
	gpUniform4ui64vNV = (C.GPUNIFORM4UI64VNV)(getProcAddr("glUniform4ui64vNV"))
	gpUniform4uiv = (C.GPUNIFORM4UIV)(getProcAddr("glUniform4uiv"))
	if gpUniform4uiv == nil {
		return errors.New("glUniform4uiv")
	}
	gpUniformBlockBinding = (C.GPUNIFORMBLOCKBINDING)(getProcAddr("glUniformBlockBinding"))
	if gpUniformBlockBinding == nil {
		return errors.New("glUniformBlockBinding")
	}
	gpUniformHandleui64IMG = (C.GPUNIFORMHANDLEUI64IMG)(getProcAddr("glUniformHandleui64IMG"))
	gpUniformHandleui64NV = (C.GPUNIFORMHANDLEUI64NV)(getProcAddr("glUniformHandleui64NV"))
	gpUniformHandleui64vIMG = (C.GPUNIFORMHANDLEUI64VIMG)(getProcAddr("glUniformHandleui64vIMG"))
	gpUniformHandleui64vNV = (C.GPUNIFORMHANDLEUI64VNV)(getProcAddr("glUniformHandleui64vNV"))
	gpUniformMatrix2fv = (C.GPUNIFORMMATRIX2FV)(getProcAddr("glUniformMatrix2fv"))
	if gpUniformMatrix2fv == nil {
		return errors.New("glUniformMatrix2fv")
	}
	gpUniformMatrix2x3fv = (C.GPUNIFORMMATRIX2X3FV)(getProcAddr("glUniformMatrix2x3fv"))
	if gpUniformMatrix2x3fv == nil {
		return errors.New("glUniformMatrix2x3fv")
	}
	gpUniformMatrix2x3fvNV = (C.GPUNIFORMMATRIX2X3FVNV)(getProcAddr("glUniformMatrix2x3fvNV"))
	gpUniformMatrix2x4fv = (C.GPUNIFORMMATRIX2X4FV)(getProcAddr("glUniformMatrix2x4fv"))
	if gpUniformMatrix2x4fv == nil {
		return errors.New("glUniformMatrix2x4fv")
	}
	gpUniformMatrix2x4fvNV = (C.GPUNIFORMMATRIX2X4FVNV)(getProcAddr("glUniformMatrix2x4fvNV"))
	gpUniformMatrix3fv = (C.GPUNIFORMMATRIX3FV)(getProcAddr("glUniformMatrix3fv"))
	if gpUniformMatrix3fv == nil {
		return errors.New("glUniformMatrix3fv")
	}
	gpUniformMatrix3x2fv = (C.GPUNIFORMMATRIX3X2FV)(getProcAddr("glUniformMatrix3x2fv"))
	if gpUniformMatrix3x2fv == nil {
		return errors.New("glUniformMatrix3x2fv")
	}
	gpUniformMatrix3x2fvNV = (C.GPUNIFORMMATRIX3X2FVNV)(getProcAddr("glUniformMatrix3x2fvNV"))
	gpUniformMatrix3x4fv = (C.GPUNIFORMMATRIX3X4FV)(getProcAddr("glUniformMatrix3x4fv"))
	if gpUniformMatrix3x4fv == nil {
		return errors.New("glUniformMatrix3x4fv")
	}
	gpUniformMatrix3x4fvNV = (C.GPUNIFORMMATRIX3X4FVNV)(getProcAddr("glUniformMatrix3x4fvNV"))
	gpUniformMatrix4fv = (C.GPUNIFORMMATRIX4FV)(getProcAddr("glUniformMatrix4fv"))
	if gpUniformMatrix4fv == nil {
		return errors.New("glUniformMatrix4fv")
	}
	gpUniformMatrix4x2fv = (C.GPUNIFORMMATRIX4X2FV)(getProcAddr("glUniformMatrix4x2fv"))
	if gpUniformMatrix4x2fv == nil {
		return errors.New("glUniformMatrix4x2fv")
	}
	gpUniformMatrix4x2fvNV = (C.GPUNIFORMMATRIX4X2FVNV)(getProcAddr("glUniformMatrix4x2fvNV"))
	gpUniformMatrix4x3fv = (C.GPUNIFORMMATRIX4X3FV)(getProcAddr("glUniformMatrix4x3fv"))
	if gpUniformMatrix4x3fv == nil {
		return errors.New("glUniformMatrix4x3fv")
	}
	gpUniformMatrix4x3fvNV = (C.GPUNIFORMMATRIX4X3FVNV)(getProcAddr("glUniformMatrix4x3fvNV"))
	gpUnmapBuffer = (C.GPUNMAPBUFFER)(getProcAddr("glUnmapBuffer"))
	if gpUnmapBuffer == nil {
		return errors.New("glUnmapBuffer")
	}
	gpUnmapBufferOES = (C.GPUNMAPBUFFEROES)(getProcAddr("glUnmapBufferOES"))
	gpUseProgram = (C.GPUSEPROGRAM)(getProcAddr("glUseProgram"))
	if gpUseProgram == nil {
		return errors.New("glUseProgram")
	}
	gpUseProgramStages = (C.GPUSEPROGRAMSTAGES)(getProcAddr("glUseProgramStages"))
	if gpUseProgramStages == nil {
		return errors.New("glUseProgramStages")
	}
	gpUseProgramStagesEXT = (C.GPUSEPROGRAMSTAGESEXT)(getProcAddr("glUseProgramStagesEXT"))
	gpUseShaderProgramEXT = (C.GPUSESHADERPROGRAMEXT)(getProcAddr("glUseShaderProgramEXT"))
	gpValidateProgram = (C.GPVALIDATEPROGRAM)(getProcAddr("glValidateProgram"))
	if gpValidateProgram == nil {
		return errors.New("glValidateProgram")
	}
	gpValidateProgramPipeline = (C.GPVALIDATEPROGRAMPIPELINE)(getProcAddr("glValidateProgramPipeline"))
	if gpValidateProgramPipeline == nil {
		return errors.New("glValidateProgramPipeline")
	}
	gpValidateProgramPipelineEXT = (C.GPVALIDATEPROGRAMPIPELINEEXT)(getProcAddr("glValidateProgramPipelineEXT"))
	gpVertexAttrib1f = (C.GPVERTEXATTRIB1F)(getProcAddr("glVertexAttrib1f"))
	if gpVertexAttrib1f == nil {
		return errors.New("glVertexAttrib1f")
	}
	gpVertexAttrib1fv = (C.GPVERTEXATTRIB1FV)(getProcAddr("glVertexAttrib1fv"))
	if gpVertexAttrib1fv == nil {
		return errors.New("glVertexAttrib1fv")
	}
	gpVertexAttrib2f = (C.GPVERTEXATTRIB2F)(getProcAddr("glVertexAttrib2f"))
	if gpVertexAttrib2f == nil {
		return errors.New("glVertexAttrib2f")
	}
	gpVertexAttrib2fv = (C.GPVERTEXATTRIB2FV)(getProcAddr("glVertexAttrib2fv"))
	if gpVertexAttrib2fv == nil {
		return errors.New("glVertexAttrib2fv")
	}
	gpVertexAttrib3f = (C.GPVERTEXATTRIB3F)(getProcAddr("glVertexAttrib3f"))
	if gpVertexAttrib3f == nil {
		return errors.New("glVertexAttrib3f")
	}
	gpVertexAttrib3fv = (C.GPVERTEXATTRIB3FV)(getProcAddr("glVertexAttrib3fv"))
	if gpVertexAttrib3fv == nil {
		return errors.New("glVertexAttrib3fv")
	}
	gpVertexAttrib4f = (C.GPVERTEXATTRIB4F)(getProcAddr("glVertexAttrib4f"))
	if gpVertexAttrib4f == nil {
		return errors.New("glVertexAttrib4f")
	}
	gpVertexAttrib4fv = (C.GPVERTEXATTRIB4FV)(getProcAddr("glVertexAttrib4fv"))
	if gpVertexAttrib4fv == nil {
		return errors.New("glVertexAttrib4fv")
	}
	gpVertexAttribBinding = (C.GPVERTEXATTRIBBINDING)(getProcAddr("glVertexAttribBinding"))
	if gpVertexAttribBinding == nil {
		return errors.New("glVertexAttribBinding")
	}
	gpVertexAttribDivisor = (C.GPVERTEXATTRIBDIVISOR)(getProcAddr("glVertexAttribDivisor"))
	if gpVertexAttribDivisor == nil {
		return errors.New("glVertexAttribDivisor")
	}
	gpVertexAttribDivisorANGLE = (C.GPVERTEXATTRIBDIVISORANGLE)(getProcAddr("glVertexAttribDivisorANGLE"))
	gpVertexAttribDivisorEXT = (C.GPVERTEXATTRIBDIVISOREXT)(getProcAddr("glVertexAttribDivisorEXT"))
	gpVertexAttribDivisorNV = (C.GPVERTEXATTRIBDIVISORNV)(getProcAddr("glVertexAttribDivisorNV"))
	gpVertexAttribFormat = (C.GPVERTEXATTRIBFORMAT)(getProcAddr("glVertexAttribFormat"))
	if gpVertexAttribFormat == nil {
		return errors.New("glVertexAttribFormat")
	}
	gpVertexAttribI4i = (C.GPVERTEXATTRIBI4I)(getProcAddr("glVertexAttribI4i"))
	if gpVertexAttribI4i == nil {
		return errors.New("glVertexAttribI4i")
	}
	gpVertexAttribI4iv = (C.GPVERTEXATTRIBI4IV)(getProcAddr("glVertexAttribI4iv"))
	if gpVertexAttribI4iv == nil {
		return errors.New("glVertexAttribI4iv")
	}
	gpVertexAttribI4ui = (C.GPVERTEXATTRIBI4UI)(getProcAddr("glVertexAttribI4ui"))
	if gpVertexAttribI4ui == nil {
		return errors.New("glVertexAttribI4ui")
	}
	gpVertexAttribI4uiv = (C.GPVERTEXATTRIBI4UIV)(getProcAddr("glVertexAttribI4uiv"))
	if gpVertexAttribI4uiv == nil {
		return errors.New("glVertexAttribI4uiv")
	}
	gpVertexAttribIFormat = (C.GPVERTEXATTRIBIFORMAT)(getProcAddr("glVertexAttribIFormat"))
	if gpVertexAttribIFormat == nil {
		return errors.New("glVertexAttribIFormat")
	}
	gpVertexAttribIPointer = (C.GPVERTEXATTRIBIPOINTER)(getProcAddr("glVertexAttribIPointer"))
	if gpVertexAttribIPointer == nil {
		return errors.New("glVertexAttribIPointer")
	}
	gpVertexAttribPointer = (C.GPVERTEXATTRIBPOINTER)(getProcAddr("glVertexAttribPointer"))
	if gpVertexAttribPointer == nil {
		return errors.New("glVertexAttribPointer")
	}
	gpVertexBindingDivisor = (C.GPVERTEXBINDINGDIVISOR)(getProcAddr("glVertexBindingDivisor"))
	if gpVertexBindingDivisor == nil {
		return errors.New("glVertexBindingDivisor")
	}
	gpViewport = (C.GPVIEWPORT)(getProcAddr("glViewport"))
	if gpViewport == nil {
		return errors.New("glViewport")
	}
	gpViewportArrayvNV = (C.GPVIEWPORTARRAYVNV)(getProcAddr("glViewportArrayvNV"))
	gpViewportArrayvOES = (C.GPVIEWPORTARRAYVOES)(getProcAddr("glViewportArrayvOES"))
	gpViewportIndexedfNV = (C.GPVIEWPORTINDEXEDFNV)(getProcAddr("glViewportIndexedfNV"))
	gpViewportIndexedfOES = (C.GPVIEWPORTINDEXEDFOES)(getProcAddr("glViewportIndexedfOES"))
	gpViewportIndexedfvNV = (C.GPVIEWPORTINDEXEDFVNV)(getProcAddr("glViewportIndexedfvNV"))
	gpViewportIndexedfvOES = (C.GPVIEWPORTINDEXEDFVOES)(getProcAddr("glViewportIndexedfvOES"))
	gpViewportPositionWScaleNV = (C.GPVIEWPORTPOSITIONWSCALENV)(getProcAddr("glViewportPositionWScaleNV"))
	gpViewportSwizzleNV = (C.GPVIEWPORTSWIZZLENV)(getProcAddr("glViewportSwizzleNV"))
	gpWaitSemaphoreEXT = (C.GPWAITSEMAPHOREEXT)(getProcAddr("glWaitSemaphoreEXT"))
	gpWaitSync = (C.GPWAITSYNC)(getProcAddr("glWaitSync"))
	if gpWaitSync == nil {
		return errors.New("glWaitSync")
	}
	gpWaitSyncAPPLE = (C.GPWAITSYNCAPPLE)(getProcAddr("glWaitSyncAPPLE"))
	gpWaitVkSemaphoreNV = (C.GPWAITVKSEMAPHORENV)(getProcAddr("glWaitVkSemaphoreNV"))
	gpWeightPathsNV = (C.GPWEIGHTPATHSNV)(getProcAddr("glWeightPathsNV"))
	gpWindowRectanglesEXT = (C.GPWINDOWRECTANGLESEXT)(getProcAddr("glWindowRectanglesEXT"))
	return nil
}