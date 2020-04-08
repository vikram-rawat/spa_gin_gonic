-- FUNCTION: public.get_address()

-- DROP FUNCTION public.get_address();

CREATE OR REPLACE FUNCTION public.get_address(
	)
    RETURNS TABLE(states text) 
    LANGUAGE 'plpgsql'

    COST 100
    VOLATILE 
    ROWS 1000
AS $BODY$
BEGIN

return query 
select 
s.states_name
from
statics.states as s;

END; 

$BODY$;

ALTER FUNCTION public.get_address()
    OWNER TO postgres;

