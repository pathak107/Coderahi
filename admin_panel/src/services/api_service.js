import axios from 'axios'

const instance = axios.create({
    baseURL: 'http://localhost:8080/api/v1/',
    timeout: 1000,
    // headers: {'X-Custom-Header': 'foobar'}
});
export const getCourseByIDWithSectionsAndPosts= async(course_id)=>{
    const res= await instance.get(`course/${course_id}?section=true&post=true`)
    console.log(res)
    return res
}
export const createCourse = async ({ title, desc, cost }) => {
        const res= await instance.post('course', {
            title: title,
            desc_short: desc,
            cost: parseInt(cost)
        })
        return res
}
export const editCourse = () => { }
export const deleteCourse = () => { }

export const createPost = () => { }
export const editPost = () => { }
export const deletePost = () => { }

export const createSection = async ({course_id, title, desc}) => {
    const res= await instance.post('course/section', {
        title: title,
        desc,
        course_id: parseInt(course_id)
    })
    return res
}
export const editSection = async ({section_id, title, desc}) => {
    const res= await instance.patch(`course/section/${section_id}`, {
        title: title,
        desc,
    })
    return res
}
export const deleteSection = async (section_id) => {
    const res= await instance.delete(`course/section/${section_id}`)
    return res
}