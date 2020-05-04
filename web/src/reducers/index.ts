import { testGetPostReducer as testGetPost } from './testGetPostReducer';
import { imageListReducer as imageList } from './imageListReducer';
import { slideShowReducer as slideShow } from './slideShowReducer';
import { reducer as loading } from './loadingReducer';

export const reducers = { testGetPost, slideShow, imageList, loading };
