
CREATE TABLE public.tb_ticket
(
    id_ticket bigserial NOT NULL,
	nu_cpf_cnpj character varying(19) NOT NULL,
    st_private character varying(12) NOT NULL,
    st_incompleto character varying(12) NOT NULL,
    dt_ultima_compra character varying(22),
    vl_ticket_medio character varying(22),
    vl_ticket_ultima_compra character varying(24),
    nu_cnpj_loja_frequente character varying(20),
    nu_cnpj_loja_ultima_compra character varying(21),
    st_valido_cpf_cnpj boolean,
    st_valido_cnpj_loja_frequente boolean,
    st_valido_cnpj_ultima_compra boolean
)

TABLESPACE pg_default;

ALTER TABLE public.tb_ticket
    OWNER to postgres;
	
create or replace function fc_valida_cpf
(
    p_cpf in character varying, 
    p_valida_nulo in boolean default false
)
returns boolean as 
$$
declare
    
    v_cpf_invalidos character varying[10] 
    default array['00000000000', '11111111111',
                  '22222222222', '33333333333',
                  '44444444444', '55555555555',
                  '66666666666', '77777777777',
                  '88888888888', '99999999999'];
                  
    v_cpf_quebrado smallint[];
    
    c_posicao_dv1 constant smallint default 10;    
    v_arranjo_dv1 smallint[9] default array[10,9,8,7,6,5,4,3,2];
    v_soma_dv1 smallint default 0;
    v_resto_dv1 double precision default 0;
    
    c_posicao_dv2 constant smallint default 11;
    v_arranjo_dv2 smallint[10] default array[11,10,9,8,7,6,5,4,3,2];
    v_soma_dv2 smallint default 0;
    v_resto_dv2 double precision default 0;
    
begin
    if p_valida_nulo and nullif(p_cpf, '') is null then
        return true;
    end if;
    if (not (p_cpf ~* '^([0-9]{11})$' or 
             p_cpf ~* '^([0-9]{3}\.[0-9]{3}\.[0-9]{3}\-[0-9]{2})$')
        ) or
        p_cpf = any (v_cpf_invalidos) or
        p_cpf is null
    then
        return false;    
    end if;
    
v_cpf_quebrado := regexp_split_to_array(
                    regexp_replace(p_cpf, '[^0-9]', '', 'g'), '');
    -------------------------------- Digito Verificador 1
    for t in 1..9 loop
        v_soma_dv1 := v_soma_dv1 + 
                     (v_cpf_quebrado[t] * v_arranjo_dv1[t]);
    end loop;
    v_resto_dv1 := ((10 * v_soma_dv1) % 11) % 10;
        
    if (v_resto_dv1 != v_cpf_quebrado[c_posicao_dv1]) 
    then
        return false;
    end if;
    
    -------------------------------- Digito Verificador 2
    for t in 1..10 loop
        v_soma_dv2 := v_soma_dv2 + 
                     (v_cpf_quebrado[t] * v_arranjo_dv2[t]);
    end loop;
    v_resto_dv2 := ((10 * v_soma_dv2) % 11) % 10;
    
    return (v_resto_dv2 = v_cpf_quebrado[c_posicao_dv2]);    
    
end;
$$ language plpgsql;	



create or replace function fc_valida_cnpj
(
    in p_cnpj character varying, 
    in p_fg_permite_nulo boolean default false
)
returns boolean as
$$
declare
    
    v_cnpj_invalidos character varying[10] 
    default array['00000000000000', '11111111111111',
                  '22222222222222', '33333333333333',
                  '44444444444444', '55555555555555',
                  '66666666666666', '77777777777777',
                  '88888888888888', '99999999999999'];
                  
    v_cnpj_quebrado smallint[];
    
    c_posicao_dv1 constant smallint default 13;
    v_arranjo_dv1 smallint[12] default array[5,4,3,2,9,8,7,6,5,4,3,2];
    v_soma_dv1 smallint default 0;
    v_resto_dv1 double precision default 0;
    
    c_posicao_dv2 constant smallint default 14;
    v_arranjo_dv2 smallint[13] default array[6,5,4,3,2,9,8,7,6,5,4,3,2];
    v_soma_dv2 smallint default 0;
    v_resto_dv2 double precision default 0;
    
begin
    
    if p_fg_permite_nulo and nullif(p_cnpj, '') is null then
        return true;
    end if;
    
    if (not (p_cnpj ~* '^([0-9]{14})$' or 
             p_cnpj ~* '^([0-9]{2}\.[0-9]{3}\.[0-9]{3}\/[0-9]{4}\-[0-9]{2})$')) or
        p_cnpj = any (v_cnpj_invalidos) or
        p_cnpj is null
    then
        return false;    
    end if;
    
    v_cnpj_quebrado := regexp_split_to_array(
      regexp_replace(p_cnpj, '[^0-9]', '', 'g'), '');
        
    -- Realiza o calculo do primeiro digito
    for t in 1..12 loop
        v_soma_dv1 := v_soma_dv1 + 
      (v_cnpj_quebrado[t] * v_arranjo_dv1[t]);
    end loop;
    v_resto_dv1 := ((10 * v_soma_dv1) % 11) % 10;
       
    if (v_resto_dv1 != v_cnpj_quebrado[13]) 
    then
        return false;
    end if;
    
    -- Realiza o calculo do segundo digito    
    for t in 1..13 loop
        v_soma_dv2 := v_soma_dv2 + 
      (v_cnpj_quebrado[t] * v_arranjo_dv2[t]);
    end loop;
    v_resto_dv2 := ((10 * v_soma_dv2) % 11) % 10;
    
    return (v_resto_dv2 = v_cnpj_quebrado[c_posicao_dv2]);    
    
end;
$$ language plpgsql;