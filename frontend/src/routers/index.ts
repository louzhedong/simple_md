import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router';
import MainPage from '../pages/MainPage/Index.vue';
import EditorPage from '../pages/EditorPage/Index.vue';
import DetailPage from '../pages/DetailPage/Index.vue';

const routes = [
	{
		path: '/',
		component: MainPage
	},
	{
		path: '/editor',
		component: EditorPage
	},
	{
		path: '/detail',
		component: DetailPage
	}
];

const router = createRouter({
	history: createWebHashHistory(),
	routes: routes as RouteRecordRaw[]
})

export default router;