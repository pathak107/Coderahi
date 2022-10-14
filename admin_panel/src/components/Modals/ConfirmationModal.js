import { useContext } from "react";
import { ConfirmModalCtx } from "../../context/confirmModalCtx";

const ConfirmationModal = () => {
    const { state, actions } = useContext(ConfirmModalCtx)
    return (
        <div className={`modal ${state.isOpen ? "modal-open" : ""}`} >
            <div className="modal-box relative">
                <h3 className="font-bold text-lg">Are you sure?</h3>
                <div className="modal-action">
                    <a className="btn"
                        onClick={() => {
                            state.onYesAction.action()
                            actions.closeModal()
                        }}
                    >
                        Yes
                    </a>
                    <a className="btn"
                        onClick={()=>{
                            actions.closeModal()
                        }}
                    >
                        No
                    </a>
                </div>
            </div>
        </div>
    );
}

export default ConfirmationModal;