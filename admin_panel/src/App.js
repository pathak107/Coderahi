import Home from './pages/Home'
import {
  createBrowserRouter,
  RouterProvider,
  Route,
} from "react-router-dom";
import {
  useQuery,
  useMutation,
  useQueryClient,
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'

import CourseForm from './components/Course/CourseForm';
import Course from './pages/Course';
import CourseList from './pages/CourseList';
import ModalContextProvider from './context/modalContext';
import ConfirmModalCtxProvider, { ConfirmModalCtx } from './context/confirmModalCtx';
import Header from './components/Layout/Header';

const queryClient = new QueryClient()

const router = createBrowserRouter([
  {
    path: "/",
    element: <><Header/><Home /></>,
  },
  {
    path: "/course/:course_id",
    element: <><Header/><Course /></>,
  },
  {
    path: "/course/",
    element: <><Header/><CourseList /></>,
  },
  {
    path: "/course/:course_id/post/:post_id",
    element: <><Header/><Course /></>,
  },
]);

function App() {

  return (
    <>
      <QueryClientProvider client={queryClient}>
        <ReactQueryDevtools initialIsOpen={true} />
        <ModalContextProvider>
          <ConfirmModalCtxProvider>
            <RouterProvider router={router} />
          </ConfirmModalCtxProvider>
        </ModalContextProvider>
      </QueryClientProvider>
    </>
  );
}

export default App;
