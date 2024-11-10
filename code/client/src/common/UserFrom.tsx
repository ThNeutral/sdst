export interface UserFromField {
  name: string;
  type: React.HTMLInputTypeAttribute;
  placeholder?: string;
  isError: boolean;
  errorMessage: string;
}

interface UserFromParams {
  fields: UserFromField[];
  buttonName: string;
  textBig: string;
  textSmall: string;
  submitHandler: (e: React.FormEvent<HTMLFormElement>) => void;
}

export function UserFrom(params: UserFromParams) {
  return (
    <div className="userform">
      <p className="userform-text-big">{params.textBig}</p>
      <p className="userform-text-small">{params.textSmall}</p>
      <form className="userform-form" onSubmit={params.submitHandler}>
        {params.fields.map((field) => {
          return (
            <div key={field.name}>
              <input
                className="userform-form-input"
                name={field.name}
                type={field.type}
                placeholder={field.placeholder}
              />
              {field.isError ? (
                <p className="userform-form-input-error">{field.errorMessage}</p>
              ) : (
                <></>
              )}
            </div>
          );
        })}
        <button className="userfrom-form-button" type="submit">
          {params.buttonName}
        </button>
      </form>
    </div>
  );
}
