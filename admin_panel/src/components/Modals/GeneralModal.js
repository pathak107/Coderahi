import { useContext } from "react";
import { ModalContext } from "../../context/modalContext";

const GeneralModal = ({children}) => {
    const {state, actions}= useContext(ModalContext)
    return (
        <div className={`modal ${state.isOpen?"modal-open":""}`} >
            <div className="modal-box relative">
                <a className="btn btn-sm btn-circle absolute right-2 top-2"
                    onClick={()=>{
                        actions.closeModal()
                    }}
                >
                    ✕
                </a>
                {children}
            </div>
        </div>
    );
}

export default GeneralModal;