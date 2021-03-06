package sdl

/*
#include "gosdl.h"
*/
import "C"

type BlendMode C.SDL_BlendMode

const (
	BLENDMODE_NONE  BlendMode = C.SDL_BLENDMODE_NONE
	BLENDMODE_BLEND BlendMode = C.SDL_BLENDMODE_BLEND
	BLENDMODE_ADD   BlendMode = C.SDL_BLENDMODE_ADD
	BLENDMODE_MOD   BlendMode = C.SDL_BLENDMODE_MOD
)
