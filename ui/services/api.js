import axios from 'axios'

const instance = axios.create({
    baseURL: 'http://localhost:8080/api/v1/',
    timeout: 1000,
    // headers: {'X-Custom-Header': 'foobar'}
});

export const getAllCourses = async () => {
    const res = await instance.get(`course?section=true&post=true`)
    return res.data.data.courses
}

export const getCourseBySlugWithSectionsAndPosts = async (slug) => {
    const res = await instance.get(`course/slug/${slug}?section=true&post=true`)
    return res.data.data.course
}

export const getPostBySlug = async (slug) => {
    const res = await instance.get(`post/slug/${slug}`)
    return res.data.data.post
}